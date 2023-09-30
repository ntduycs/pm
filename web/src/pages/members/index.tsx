import { StyledMembers } from '@pm/pages/members/styles.ts';
import { UpsertMemberModal, ListMembersTable } from '@pm/pages/members/components';
import { useModal } from '@pm/hooks';
import { useCallback, useEffect, useState } from 'react';
import { Member } from '@pm/models';

export const Members = () => {
  const { isOpen: isUpsertModalOpened, setOpen: showHideUpsertModal } = useModal();
  const [selectedMember, setSelectedMember] = useState<Member | undefined>(undefined);

  const toggleUpsertModal = useCallback(
    (member?: Member) => {
      setSelectedMember(member);
      showHideUpsertModal(true);
    },
    [showHideUpsertModal],
  );

  return (
    <StyledMembers>
      <UpsertMemberModal
        isOpen={isUpsertModalOpened}
        setOpen={showHideUpsertModal}
        member={selectedMember}
      />
      <ListMembersTable toggleUpsertModal={toggleUpsertModal} />
    </StyledMembers>
  );
};
