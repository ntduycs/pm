import { PaPcResult, UpsertPaPcResultRequest } from '@pm/models';
import { Button, Form, Input, Modal, notification, Select, Space } from 'antd';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { upsertPaPcResultAPI } from '@pm/services';
import { useEffect } from 'react';
import { StyledUpsertPaPcResult } from '@pm/pages/papc/components/upsert/styles.ts';
import { listMembersAPI } from '@pm/services';
import { TSortDirection, UiConstant } from '@pm/common/constants';
import { capitalize, toString } from 'lodash';
import { Apis, Strings } from '@pm/common/utils';

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
      form.setFieldValue('period', import.meta.env.VITE_DEFAULT_PA_PC_PERIOD);
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
      const { message } = await upsertPaPcResult({
        ...values,
        technical_score: Number(values.technical_score),
        productivity_score: Number(values.productivity_score),
        collaboration_score: Number(values.collaboration_score),
        development_score: Number(values.development_score),
      });
      api.success({
        message: UiConstant.SuccessMessage,
        description: message,
        placement: 'topRight',
        duration: 2,
      });
    } catch (error: unknown) {
      api.error({
        message: UiConstant.ErrorMessage,
        description: Apis.getErrorDescription(error),
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
        period: import.meta.env.VITE_DEFAULT_PA_PC_PERIOD,
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
              name='period'
              label='Period'
              hasFeedback
            >
              <Input disabled />
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
              name='technical_score'
              label='Technical Score'
              hasFeedback
              rules={[
                {
                  required: true,
                  message: 'Please input technical score!',
                },
                {
                  pattern: new RegExp(/^[0-9](\.[0-9]{1,2})?$/),
                  message: 'Please input number with 2 decimal places!',
                },
              ]}
            >
              <Input />
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
                  pattern: new RegExp(/^[0-9](\.[0-9]{1,2})?$/),
                  message: 'Please input number with 2 decimal places!',
                },
              ]}
            >
              <Input />
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
                  pattern: new RegExp(/^[0-9](\.[0-9]{1,2})?$/),
                  message: 'Please input number with 2 decimal places!',
                },
              ]}
            >
              <Input />
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
                  pattern: new RegExp(/^[0-9](\.[0-9]{1,2})?$/),
                  message: 'Please input number with 2 decimal places!',
                },
              ]}
            >
              <Input />
            </Form.Item>
            <Form.Item
              name='note'
              label='Note'
            >
              <Input.TextArea
                autoSize={{
                  minRows: 3,
                  maxRows: 10,
                }}
              />
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
