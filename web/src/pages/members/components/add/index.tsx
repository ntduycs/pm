import { StyledAddMember } from '@pm/pages/members/components/add/styles.ts';
import React, { useState } from 'react';
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
import { Member } from '@pm/pages/members/model.ts';
import { InfoCircleOutlined } from '@ant-design/icons';
import JSONPretty from 'react-json-pretty';
import { JsonRenderer } from '@pm/components';
import { useMutation, useQuery } from '@tanstack/react-query';
import { listMembersAPI, upsertMemberAPI } from '@pm/services';
import { UpsertMemberRequest } from '@pm/models';
import { useQueryClient } from '@pm/hooks';

export const AddMember = () => {
  const [isOpen, setIsOpen] = useState(false);
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
      onSuccess: async (data) => {
        await queryClient.invalidateQueries(['members', 'list']);
        await listMembersQuery.refetch();
      },
    },
  );

  const [form] = Form.useForm();
  const formValues = Form.useWatch([], form);

  const openModal = () => {
    setIsOpen(true);
  };

  const closeModal = () => {
    setIsOpen(false);
  };

  const onFormSubmit = async (values: UpsertMemberRequest) => {
    try {
      console.log(values.end_date);
      values = {
        ...values,
        start_date: values.start_date ? dayjs(values.start_date).format('YYYY-MM-DD') : '',
        end_date: values.end_date ? dayjs(values.end_date).format('YYYY-MM-DD') : '',
      };
      const { message } = await upsertMember(values);
      api.success({
        message,
        description: <JsonRenderer json={values} />,
        placement: 'topRight',
        duration: 2,
      });
      setIsOpen(false);
    } catch (error: unknown) {
      api.error({
        message: 'Error!',
        description: error instanceof Error ? error.message : 'Unknown error',
        placement: 'topRight',
        duration: 3,
      });
    }
  };

  const onFormCancel = () => {
    form.resetFields();
    setIsOpen(false);
  };

  return (
    <>
      {contextHolder}
      <StyledAddMember>
        <Button onClick={openModal}>New Member</Button>
        <Modal
          open={isOpen}
          title='New Member'
          destroyOnClose
          keyboard
          closable={false}
          footer={null}
        >
          <Form
            form={form}
            initialValues={{
              name: '',
              level: EMember.Level.LV1,
              positions: [],
              kpi: 0.5,
              category: EMember.Category.OFFICIAL,
              total_effort: 100,
              status: EMember.Status.ACTIVE.toLowerCase(),
            }}
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
              normalize={(value) => value?.toLowerCase()}
            >
              <Select
                options={Object.values(EMember.Status).map((status) => ({
                  label: status,
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
                  onClick={onFormCancel}
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
