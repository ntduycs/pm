import { Button, Modal, notification, Space } from 'antd';
import { StyledDeleteMember } from '@pm/pages/members/components/delete/styles.ts';
import { Member } from '@pm/models';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { deleteMemberAPI } from '@pm/services';
import React from 'react';
import { Apis } from '@pm/common/utils';

type DeleteMemberProps = {
  member?: Member;
  isOpen: boolean;
  setOpen: (isOpen: boolean) => void;
};

export const DeleteMemberModal = ({ member, isOpen, setOpen }: DeleteMemberProps) => {
  const [api, contextHolder] = notification.useNotification();
  const queryClient = useQueryClient();

  const { mutateAsync: deleteMember, isLoading } = useMutation({
    mutationKey: ['members'],
    mutationFn: deleteMemberAPI,
    onSuccess: async () => {
      await queryClient.refetchQueries(['members']);
      closeModal();
    },
  });

  const closeModal = () => {
    setOpen(false);
  };

  const onFormSubmit = async () => {
    try {
      const { message } = await deleteMember(member!.id);
      api.success({
        message,
        placement: 'topRight',
        duration: 2,
      });
    } catch (error) {
      api.error({
        message: 'Error!',
        description: Apis.getErrorDescription(error),
        placement: 'topRight',
        duration: 3,
      });
    }
  };

  if (!member) return null;

  return (
    <>
      {contextHolder}
      <StyledDeleteMember>
        <Modal
          open={isOpen}
          title={`Delete Member`}
          destroyOnClose={true}
          keyboard={true}
          closable={false}
          footer={null}
        >
          <p>Are you sure you want to delete this member?</p>
          <ul>
            <li>Member Name: {member!.name}</li>
            <li>Member Email: {member!.email}</li>
          </ul>
          <p>
            This action cannot be undone. Deleting the member will remove them from any associated
            data.
          </p>
          <Space>
            <Button
              type='primary'
              danger={true}
              htmlType='submit'
              disabled={isLoading}
              onClick={onFormSubmit}
            >
              {isLoading ? 'Loading...' : 'Confirm'}
            </Button>
            <Button
              type='default'
              htmlType='reset'
              onClick={closeModal}
            >
              Cancel
            </Button>
          </Space>
        </Modal>
      </StyledDeleteMember>
    </>
  );
};
