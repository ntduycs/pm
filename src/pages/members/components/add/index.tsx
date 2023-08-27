import { StyledAddMember } from '@pm/pages/members/components/add/styles.ts';
import { useEffect, useState } from 'react';
import { Modal } from '@pm/components';
import { Category, Level, Position, Status } from '@pm/common/constants';
import dayjs from 'dayjs';
import { Button, DatePicker, Form, Input, InputNumber, Select } from 'antd';
import { Member } from '@pm/pages/members/meta.ts';
import { InfoCircleOutlined } from '@ant-design/icons';

export const AddMember = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [isSubmittable, setIsSubmittable] = useState(false);

  const [form] = Form.useForm();
  const formValues = Form.useWatch([], form);

  const openModal = () => {
    setIsOpen(true);
  };

  const closeModal = () => {
    setIsOpen(false);
  };

  const onFormSubmit = (values: Member) => {
    console.log(values);
    setIsOpen(false);
  };

  const onFormCancel = () => {
    form.resetFields();
    setIsOpen(false);
  };

  useEffect(() => {
    form.validateFields({ validateOnly: true }).then(
      () => {
        setIsSubmittable(true);
      },
      () => {
        setIsSubmittable(false);
      },
    );
  }, [form, formValues]);

  return (
    <StyledAddMember>
      <Button onClick={openModal}>New Member</Button>
      <Modal
        isOpen={isOpen}
        title='Add Member'
      >
        <Form
          form={form}
          initialValues={{
            name: '',
            level: Level.LV1,
            positions: [],
            kpi: 0,
            category: Category.OFFICIAL,
            totalEffort: 0,
            joinDate: dayjs(),
            status: Status.ACTIVE,
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
            rules={[
              {
                required: true,
                message: 'Please select level',
              },
            ]}
          >
            <Select
              options={Object.values(Level).map((level) => ({
                label: level,
                value: level,
              }))}
              allowClear={false}
            />
          </Form.Item>
          <Form.Item
            name='positions'
            label='Positions'
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
              options={Object.values(Position).map((position) => ({
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
            rules={[
              {
                required: true,
                message: 'Please select category',
              },
            ]}
          >
            <Select
              options={Object.values(Category).map((category) => ({
                label: category,
                value: category,
              }))}
              allowClear={false}
            />
          </Form.Item>
          <Form.Item
            name='totalEffort'
            label='Total Effort'
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
            name='joinDate'
            label='Join Date'
            rules={[
              {
                required: true,
                message: 'Please input join date',
              },
            ]}
          >
            <DatePicker />
          </Form.Item>
          {/*======== Buttons========*/}
          <Form.Item>
            <Button
              type='primary'
              htmlType='submit'
              disabled={!isSubmittable}
            >
              Save
            </Button>
          </Form.Item>
          <Form.Item>
            <Button
              type='default'
              htmlType='button'
              onClick={onFormCancel}
            >
              Cancel
            </Button>
          </Form.Item>
        </Form>
      </Modal>
    </StyledAddMember>
  );
};
