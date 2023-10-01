import { StyledMembers } from '@pm/pages/members/styles.ts';
import {
  UpsertMemberModal,
  ListMembersTable,
  DeleteMemberModal,
} from '@pm/pages/members/components';
import { useModal } from '@pm/hooks';
import { useCallback, useState } from 'react';
import { Member } from '@pm/models';

export const Members = () => {
  const { isOpen: isUpsertModalOpened, setOpen: showHideUpsertModal } = useModal();
  const { isOpen: isDeleteModalOpened, setOpen: showHideDeleteModal } = useModal();
  const [selectedUpsertMember, setSelectedUpsertMember] = useState<Member | undefined>(undefined);
  const [selectedDeleteMember, setSelectedDeleteMember] = useState<Member | undefined>(undefined);

  const toggleUpsertModal = useCallback(
    (member?: Member) => {
      setSelectedUpsertMember(member);
      showHideUpsertModal(true);
    },
    [showHideUpsertModal],
  );

  const toggleDeleteModal = useCallback(
    (member?: Member) => {
      setSelectedDeleteMember(member);
      showHideDeleteModal(true);
    },
    [showHideDeleteModal],
  );

  return (
    <StyledMembers>
      <ListMembersTable
        toggleUpsertModal={toggleUpsertModal}
        toggleDeleteModal={toggleDeleteModal}
      />
      <UpsertMemberModal
        isOpen={isUpsertModalOpened}
        setOpen={showHideUpsertModal}
        member={selectedUpsertMember}
      />
      <DeleteMemberModal
        member={selectedDeleteMember}
        isOpen={isDeleteModalOpened}
        setOpen={showHideDeleteModal}
      />
    </StyledMembers>
  );
};
