import { StyledMemberInfo } from '@pm/pages/papc/components/info/styles.ts';
import { Form, Input, Modal } from 'antd';
import { Member } from '@pm/models';
import { useEffect } from 'react';
import { capitalize } from 'lodash';

type MemberInfoModalProps = {
  member?: Member;
  isOpen: boolean;
  setOpen: (isOpen: boolean) => void;
};

export const MemberInfoModal = ({
  member,
  isOpen: isModalOpened,
  setOpen: showHideModal,
}: MemberInfoModalProps) => {
  const [form] = Form.useForm();

  useEffect(() => {
    if (member) {
      form.setFieldsValue({
        ...member,
        positions: member.positions.join(', '),
        category: capitalize(member.category),
      });
    }
  }, [form, member]);

  return (
    <StyledMemberInfo>
      <Modal
        open={isModalOpened}
        onCancel={() => showHideModal(false)}
        destroyOnClose={true}
        keyboard={true}
        closable={false}
        footer={null}
        title={member?.name}
      >
        <Form
          form={form}
          labelCol={{ span: 6 }}
          labelAlign='left'
          wrapperCol={{ span: 18 }}
          autoComplete='off'
          disabled={true}
        >
          <Form.Item
            name='level'
            label='Level'
          >
            <Input disabled />
          </Form.Item>
          <Form.Item
            name='positions'
            label='Positions'
          >
            <Input disabled />
          </Form.Item>
          <Form.Item
            name='category'
            label='Category'
          >
            <Input disabled />
          </Form.Item>
          <Form.Item
            name='kpi'
            label='KPI (%)'
          >
            <Input disabled />
          </Form.Item>
          <Form.Item
            name='total_effort'
            label='Total Effort (%)'
          >
            <Input disabled />
          </Form.Item>
        </Form>
      </Modal>
    </StyledMemberInfo>
  );
};
