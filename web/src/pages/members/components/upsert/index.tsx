import { StyledAddMember } from '@pm/pages/members/components/upsert/styles.ts';
import React, { useEffect, useLayoutEffect, useState } from 'react';
import { EMember } from '@pm/common/constants';
import dayjs from 'dayjs';
import {
  Button,
  DatePicker,
  Form,
  Input,
  InputNumber,
  Modal,
  notification,
  Select,
  Space,
} from 'antd';
import { InfoCircleOutlined } from '@ant-design/icons';
import { useMutation, useQuery } from '@tanstack/react-query';
import { listMembersAPI, upsertMemberAPI } from '@pm/services';
import { Member, UpsertMemberRequest } from '@pm/models';
import { useQueryClient } from '@pm/hooks';
import { Status } from '@pm/pages/members/components';
import { capitalize } from 'lodash';

type UpsertMemberProps = {
  member?: Member;
  isOpen: boolean;
  setOpen: (isOpen: boolean) => void;
};

export const UpsertMemberModal = ({
  member,
  isOpen: isModalOpened,
  setOpen: showHideModal,
}: UpsertMemberProps) => {
  const [api, contextHolder] = notification.useNotification();
  const queryClient = useQueryClient();

  const listMembersQuery = useQuery({
    queryKey: ['members', 'list'],
    queryFn: listMembersAPI,
    enabled: false,
  });

  // eslint-disable-next-line @tanstack/query/prefer-query-object-syntax
  const { mutateAsync: upsertMember, isLoading } = useMutation(
    ['members', 'upsert'],
    upsertMemberAPI,
    {
      onSuccess: async () => {
        await queryClient.invalidateQueries(['members', 'list']);
        await listMembersQuery.refetch();
        closeModal();
      },
    },
  );

  const [form] = Form.useForm();

  const closeModal = () => {
    showHideModal(false);
    form.resetFields();
  };

  const onFormSubmit = async (values: UpsertMemberRequest) => {
    try {
      values = {
        ...values,
        kpi: values.kpi / 100,
        status: values.status?.toLowerCase(),
        start_date: values.start_date ? dayjs(values.start_date).format('YYYY-MM-DD') : '',
        end_date: values.end_date ? dayjs(values.end_date).format('YYYY-MM-DD') : '',
      };
      const { message } = await upsertMember(values);
      api.success({
        message,
        placement: 'topRight',
        duration: 2,
      });
      closeModal();
    } catch (error: unknown) {
      api.error({
        message: 'Error!',
        description: error instanceof Error ? error.message : 'Unknown error',
        placement: 'topRight',
        duration: 3,
      });
    }
  };

  useEffect(() => {
    if (member) {
      form.setFieldsValue({
        name: member.name,
        email: member.email,
        level: member.level,
        positions: member.positions,
        kpi: member.kpi,
        category: member.category,
        total_effort: member.total_effort,
        status: capitalize(member.status),
        start_date: member.start_date ? dayjs(member.start_date) : undefined,
        end_date: member.end_date ? dayjs(member.end_date) : undefined,
      });
    }
  }, [form, member]);

  console.log('render upsert member modal');

  return (
    <>
      {contextHolder}
      <StyledAddMember>
        <Modal
          open={isModalOpened}
          title='New Member'
          destroyOnClose
          keyboard
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
                  type: 'email',
                  message: 'Please input valid email',
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
                options={Object.values(EMember.Level).map((level) => ({
                  label: level,
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
                options={Object.values(EMember.Position).map((position) => ({
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
              ]}
              tooltip={{
                title: 'Please input KPI (0 - 100)',
                icon: <InfoCircleOutlined />,
              }}
            >
              <InputNumber
                min={0}
                max={100}
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
                options={Object.values(EMember.Category).map((category) => ({
                  label: category,
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
              ]}
              tooltip={{
                title: 'Please input total effort (0 - 100)',
                icon: <InfoCircleOutlined />,
              }}
            >
              <InputNumber
                min={0}
                max={100}
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
                  transform: (value) => value?.format('YYYY-MM-DD'),
                },
              ]}
            >
              <DatePicker style={{ width: '100%' }} />
            </Form.Item>
            <Form.Item
              name='end_date'
              label='End Date'
              rules={[
                {
                  required: false,
                  transform: (value) => value?.format('YYYY-MM-DD'),
                },
              ]}
            >
              <DatePicker style={{ width: '100%' }} />
            </Form.Item>
            <Form.Item
              name='status'
              label='Status'
              hasFeedback
              rules={[
                {
                  required: true,
                  message: 'Please select status',
                },
              ]}
            >
              <Select
                options={Object.values(EMember.Status).map((status) => ({
                  label: <Status status={status} />,
                  value: status,
                }))}
                allowClear={false}
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
                  htmlType='reset'
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
