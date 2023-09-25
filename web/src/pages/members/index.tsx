import { StyledMembers } from '@pm/pages/members/styles.ts';
import { UpsertMemberModal, ListMembersTable } from '@pm/pages/members/components';
import { useModal } from '@pm/hooks';
import { useCallback, useState } from 'react';
import { Member } from '@pm/models';

export const Members = () => {
  const { isOpen: isUpsertModalOpened, setOpen: showHideUpsertModal } = useModal();
  const [selectedMember, setSelectedMember] = useState<Member | undefined>(undefined);

  const toggleUpsertModal = useCallback(
    (member?: Member) => {
      showHideUpsertModal(true);
      setSelectedMember(member);
    },
    [showHideUpsertModal],
  );

  return (
    <StyledMembers>
      <UpsertMemberModal
        isOpen={isUpsertModalOpened}
        setOpen={showHideUpsertModal}
      />
      <ListMembersTable toggleUpsertModal={toggleUpsertModal} />
    </StyledMembers>
  );
};
