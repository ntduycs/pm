import { StyledAddMember } from '@pm/pages/members/components/upsert/styles.ts';
import React, { useEffect } from 'react';
import { ApiConstant, MemberConstant, UiConstant } from '@pm/common/constants';
import dayjs from 'dayjs';
import { Button, DatePicker, Form, Input, Modal, notification, Select, Space } from 'antd';
import { InfoCircleOutlined } from '@ant-design/icons';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { upsertMemberAPI } from '@pm/services';
import { Member, UpsertMemberRequest } from '@pm/models';
import { capitalize } from 'lodash';

type UpsertMemberProps = {
  member?: Member;
  isOpen: boolean;
  setOpen: (isOpen: boolean) => void;
};

const calcStatus = (start_date?: string, end_date?: string) => {
  const now = dayjs();

  if (!start_date && !end_date) {
    return MemberConstant.Status.ACTIVE;
  } else if (start_date && !end_date) {
    return now.isBefore(dayjs(start_date))
      ? MemberConstant.Status.PENDING
      : MemberConstant.Status.ACTIVE;
  } else if (!start_date && end_date) {
    return now.isAfter(dayjs(end_date))
      ? MemberConstant.Status.INACTIVE
      : MemberConstant.Status.ACTIVE;
  } else {
    return now.isBefore(dayjs(start_date))
      ? MemberConstant.Status.PENDING
      : now.isAfter(dayjs(end_date))
      ? MemberConstant.Status.INACTIVE
      : MemberConstant.Status.ACTIVE;
  }
};

export const UpsertMemberModal = ({
  member,
  isOpen: isModalOpened,
  setOpen: showHideModal,
}: UpsertMemberProps) => {
  const [api, contextHolder] = notification.useNotification();
  const queryClient = useQueryClient();

  const { mutateAsync: upsertMember, isLoading } = useMutation({
    mutationKey: ['members', 'upsert'],
    mutationFn: upsertMemberAPI,
    onSuccess: async () => {
      await queryClient.refetchQueries(['members']);
      form.resetFields();
      closeModal();
    },
  });

  const [form] = Form.useForm();

  const closeModal = () => {
    showHideModal(false);
  };

  const onFormSubmit = async (values: UpsertMemberRequest) => {
    try {
      values = {
        ...values,
        email: `${values.email}${ApiConstant.DefaultEmailSuffix}`,
        kpi: Number(values.kpi),
        total_effort: Number(values.total_effort),
        status: calcStatus(values.start_date, values.end_date),
        start_date: values.start_date
          ? dayjs(values.start_date).format(ApiConstant.DefaultDateFormat)
          : '',
        end_date: values.end_date
          ? dayjs(values.end_date).format(ApiConstant.DefaultDateFormat)
          : '',
      };
      const { message } = await upsertMember(values);
      api.success({
        message: UiConstant.SuccessMessage,
        description: message,
        placement: 'topRight',
        duration: 2,
      });
    } catch (error: unknown) {
      api.error({
        message: UiConstant.ErrorMessage,
        description: error instanceof Error ? error.message : 'Unknown error',
        placement: 'topRight',
        duration: 3,
      });
    }
  };

  useEffect(() => {
    if (member) {
      form.setFieldsValue({
        id: member.id,
        name: member.name,
        email: member.email.substring(0, member.email.indexOf('@')),
        jira_name: member.jira_name,
        level: member.level,
        positions: member.positions,
        kpi: member.kpi,
        category: member.category,
        total_effort: member.total_effort,
        status: member.status,
        start_date: member.start_date ? dayjs(member.start_date) : undefined,
        end_date: member.end_date ? dayjs(member.end_date) : undefined,
      });
    } else {
      form.resetFields();
      form.setFieldsValue({
        level: MemberConstant.Level.LV1,
        category: MemberConstant.Category.BUFFER,
        total_effort: 100,
        status: MemberConstant.Status.ACTIVE,
      });
    }
  }, [form, member]);

  return (
    <>
      {contextHolder}
      <StyledAddMember>
        <Modal
          open={isModalOpened}
          title={member ? 'Edit Member' : 'Add Member'}
          destroyOnClose={true}
          keyboard={true}
          closable={false}
          footer={null}
        >
          <Form
            form={form}
            labelCol={{ span: 6 }}
            labelAlign='left'
            wrapperCol={{ span: 18 }}
            autoComplete='off'
            onFinish={onFormSubmit}
            scrollToFirstError
          >
            <Form.Item
              name='id'
              label='ID'
              hidden
            >
              <Input />
            </Form.Item>
            <Form.Item
              name='name'
              label='Name'
              hasFeedback
              rules={[
                {
                  required: true,
                  message: 'Please input name',
                },
              ]}
            >
              <Input />
            </Form.Item>
            <Form.Item
              name='email'
              label='Email'
              hasFeedback
              rules={[
                {
                  required: true,
                  whitespace: false,
                  message: 'Please input valid email',
                },
              ]}
            >
              <Input addonAfter={<span>@banvien.com.vn</span>} />
            </Form.Item>
            <Form.Item
              name='jira_name'
              label='Jira Name'
              hasFeedback
              rules={[
                {
                  required: true,
                  message: 'Please input Jira name',
                },
              ]}
            >
              <Input />
            </Form.Item>
            <Form.Item
              name='level'
              label='Level'
              hasFeedback
              rules={[
                {
                  required: true,
                  message: 'Please select level',
                },
              ]}
            >
              <Select
                options={Object.values(MemberConstant.Level).map((level) => ({
                  label: `Level ${level}`,
                  value: level,
                }))}
                allowClear={false}
              />
            </Form.Item>
            <Form.Item
              name='positions'
              label='Positions'
              hasFeedback
              rules={[
                {
                  required: true,
                  message: 'Please select at least one position',
                },
              ]}
              tooltip={{
                title: 'Please select at least one position',
                icon: <InfoCircleOutlined />,
              }}
            >
              <Select
                mode='multiple'
                options={Object.values(MemberConstant.Position).map((position) => ({
                  label: position,
                  value: position,
                }))}
                allowClear
                autoClearSearchValue
              />
            </Form.Item>
            <Form.Item
              name='kpi'
              label='KPI'
              hasFeedback
              rules={[
                {
                  required: true,
                  message: 'Please input KPI',
                },
                {
                  pattern: /^([0-9]|[1-9][0-9]|100)$/,
                  message: 'Please input number from 0 to 100',
                },
              ]}
              tooltip={{
                title: 'Please input KPI (0 - 100)',
                icon: <InfoCircleOutlined />,
              }}
            >
              <Input
                addonAfter={<span>%</span>}
                style={{ width: '100%' }}
              />
            </Form.Item>
            <Form.Item
              name='category'
              label='Category'
              hasFeedback
              rules={[
                {
                  required: true,
                  message: 'Please select category',
                },
              ]}
            >
              <Select
                options={Object.values(MemberConstant.Category).map((category) => ({
                  label: capitalize(category),
                  value: category,
                }))}
                allowClear={false}
              />
            </Form.Item>
            <Form.Item
              name='total_effort'
              label='Total Effort'
              hasFeedback
              validateDebounce={300}
              rules={[
                {
                  required: true,
                  message: 'Please input total effort',
                },
                {
                  pattern: /^([0-9]|[1-9][0-9]|100)$/,
                  message: 'Please input number from 0 to 100',
                },
              ]}
              tooltip={{
                title: 'Please input total effort (0 - 100)',
                icon: <InfoCircleOutlined />,
              }}
            >
              <Input
                style={{ width: '100%' }}
                addonAfter={<span>%</span>}
              />
            </Form.Item>
            <Form.Item
              name='start_date'
              label='Start Date'
              rules={[
                {
                  required: false,
                  transform: (value) => value?.format(ApiConstant.DefaultDateFormat),
                },
              ]}
            >
              <DatePicker
                style={{ width: '100%' }}
                disabled={!member}
              />
            </Form.Item>
            <Form.Item
              name='end_date'
              label='End Date'
              rules={[
                {
                  required: false,
                  transform: (value) => value?.format(ApiConstant.DefaultDateFormat),
                },
              ]}
            >
              <DatePicker
                style={{ width: '100%' }}
                disabled={!member}
              />
            </Form.Item>
            <Form.Item wrapperCol={{ span: 12, offset: 6 }}>
              <Space>
                <Button
                  type='primary'
                  htmlType='submit'
                  disabled={isLoading}
                >
                  {isLoading ? 'Loading...' : 'Submit'}
                </Button>
                <Button
                  htmlType='button'
                  onClick={closeModal}
                >
                  Cancel
                </Button>
              </Space>
            </Form.Item>
          </Form>
        </Modal>
      </StyledAddMember>
    </>
  );
};
