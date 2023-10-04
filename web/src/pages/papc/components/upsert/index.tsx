import { PaPcResult, UpsertPaPcResultRequest } from '@pm/models';
import { Button, Form, Input, InputNumber, Modal, notification, Select, Space } from 'antd';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { upsertPaPcResultAPI } from '@pm/services/papc.service.ts';
import { useEffect } from 'react';
import { StyledUpsertPaPcResult } from '@pm/pages/papc/components/upsert/styles.ts';
import { listMembersAPI } from '@pm/services';
import { TSortDirection } from '@pm/common/constants';
import { capitalize, deburr, toString } from 'lodash';
import { Strings } from '@pm/common/utils';

type UpsertPaPcResultModalProps = {
  paPcResult?: PaPcResult;
  isOpen: boolean;
  setOpen: (isOpen: boolean) => void;
};

export const UpsertPaPcResultModal = ({
  paPcResult,
  isOpen: isModalOpened,
  setOpen: showHideUpsert,
}: UpsertPaPcResultModalProps) => {
  const [api, contextHolder] = notification.useNotification();
  const queryClient = useQueryClient();

  const [form] = Form.useForm();

  const { mutateAsync: upsertPaPcResult, isLoading: isSubmitting } = useMutation({
    mutationKey: ['pa-pc-results'],
    mutationFn: upsertPaPcResultAPI,
    onSuccess: async () => {
      await queryClient.refetchQueries(['pa-pc-results']);
      form.resetFields();
      closeModal();
    },
  });

  const { isFetching, data: members } = useQuery({
    queryKey: ['members'],
    queryFn: async () => {
      return await listMembersAPI({
        page: 1,
        size: 100,
        sort: 'name',
        direction: 'asc' as TSortDirection,
      });
    },
    enabled: true,
  });

  const closeModal = () => {
    showHideUpsert(false);
  };

  const onFormSubmit = async (values: UpsertPaPcResultRequest) => {
    try {
      const { message } = await upsertPaPcResult(values);
      api.success({
        message: 'Success!',
        description: message,
        placement: 'topRight',
        duration: 2,
      });
    } catch (error: unknown) {
      api.error({
        message: 'Error!',
        description: error instanceof Error ? error.message : 'Something went wrong!',
        placement: 'topRight',
        duration: 3,
      });
    }
  };

  useEffect(() => {
    if (paPcResult) {
      form.setFieldsValue({
        ...paPcResult,
        member_id: paPcResult.member.id,
      });
    } else {
      form.resetFields();
      form.setFieldsValue({
        period: '2023-Q3',
      });
    }
  }, [form, paPcResult]);

  return (
    <>
      {contextHolder}
      <StyledUpsertPaPcResult>
        <Modal
          open={isModalOpened}
          title={paPcResult ? 'Edit PA/PC Result' : 'Add PA/PC Result'}
          destroyOnClose
          keyboard
          closable={false}
          footer={null}
        >
          <Form
            form={form}
            labelCol={{ span: 8 }}
            labelAlign='left'
            wrapperCol={{ span: 16 }}
            autoComplete='off'
            onFinish={onFormSubmit}
            scrollToFirstError
          >
            <Form.Item
              name='id'
              hidden
            >
              <Input />
            </Form.Item>
            <Form.Item
              name='member_id'
              label='Member'
              hasFeedback
              rules={[
                {
                  required: true,
                  message: 'Please select member!',
                },
              ]}
            >
              <Select
                options={members?.items.map((member) => ({
                  label: `${member.name} (${capitalize(member.category)})`,
                  value: member.id,
                }))}
                loading={isFetching}
                allowClear={false}
                showSearch
                filterOption={(input, option) =>
                  toString(Strings.removeDiacritics(option?.label))
                    .toLowerCase()
                    .indexOf(input.toLowerCase()) >= 0
                }
              />
            </Form.Item>
            <Form.Item
              name='period'
              label='Period'
              hasFeedback
            >
              <Input disabled />
            </Form.Item>
            <Form.Item
              name='technical_score'
              label='Technical Score'
              hasFeedback
              rules={[
                {
                  required: true,
                  message: 'Please input technical score!',
                },
                {
                  type: 'number',
                  min: 1,
                  max: 10,
                  message: 'Please input number from 1 to 10!',
                },
              ]}
            >
              <InputNumber />
            </Form.Item>
            <Form.Item
              name='productivity_score'
              label='Productivity Score'
              hasFeedback
              rules={[
                {
                  required: true,
                  message: 'Please input productivity score!',
                },
                {
                  type: 'number',
                  min: 1,
                  max: 10,
                  message: 'Please input number from 1 to 10!',
                },
              ]}
              validateDebounce={400}
            >
              <InputNumber />
            </Form.Item>
            <Form.Item
              name='collaboration_score'
              label='Collaboration Score'
              hasFeedback
              rules={[
                {
                  required: true,
                  message: 'Please input collaboration score!',
                },
                {
                  type: 'number',
                  min: 1,
                  max: 10,
                  message: 'Please input number from 1 to 10!',
                },
              ]}
            >
              <InputNumber />
            </Form.Item>
            <Form.Item
              name='development_score'
              label='Development Score'
              hasFeedback
              rules={[
                {
                  required: true,
                  message: 'Please input development score!',
                },
                {
                  type: 'number',
                  min: 1,
                  max: 10,
                  message: 'Please input number from 1 to 10!',
                },
              ]}
            >
              <InputNumber />
            </Form.Item>
            <Form.Item wrapperCol={{ offset: 8, span: 12 }}>
              <Space>
                <Button
                  type='primary'
                  htmlType='submit'
                  disabled={isSubmitting}
                >
                  {isSubmitting ? 'Loading...' : 'Submit'}
                </Button>
                <Button
                  type='default'
                  htmlType='button'
                  onClick={closeModal}
                >
                  Cancel
                </Button>
              </Space>
            </Form.Item>
          </Form>
        </Modal>
      </StyledUpsertPaPcResult>
    </>
  );
};
