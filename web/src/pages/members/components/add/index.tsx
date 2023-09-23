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

export const AddMember = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [api, contextHolder] = notification.useNotification();

  const [form] = Form.useForm();
  const formValues = Form.useWatch([], form);

  const openModal = () => {
    setIsOpen(true);
  };

  const closeModal = () => {
    setIsOpen(false);
  };

  const onFormSubmit = (values: Member) => {
    setIsOpen(false);
    api.info({
      message: 'Done!',
      description: (
        <JSONPretty
          data={values}
          theme={{
            main: 'line-height:1.3;color:#444444;background:inherit;overflow:auto;',
            error: 'line-height:1.3;color:#444444;background:#272822;overflow:auto;',
            key: 'color:#f92672;',
            string: 'color:#444444;',
            value: 'color:#1677ff;',
            boolean: 'color:#1677ff;',
          }}
        />
      ),
      placement: 'top',
      duration: 3,
    });
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
              kpi: 0,
              category: EMember.Category.OFFICIAL,
              total_effort: 100,
              start_date: dayjs(),
              status: EMember.Status.ACTIVE,
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
              />
            </Form.Item>
            <Form.Item
              name='start_date'
              label='Start Date'
              hasFeedback
              rules={[
                {
                  required: true,
                  message: 'Please input start date',
                },
              ]}
            >
              <DatePicker />
            </Form.Item>
            <Form.Item wrapperCol={{ span: 12, offset: 6 }}>
              <Space>
                <Button
                  type='primary'
                  htmlType='submit'
                >
                  Submit
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
