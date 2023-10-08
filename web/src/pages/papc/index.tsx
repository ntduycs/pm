import { StyledPaPcResults } from '@pm/pages/papc/styles.ts';
import {
  ListPaPcResultsTable,
  MemberInfoModal,
  UpsertPaPcResultModal,
} from '@pm/pages/papc/components';
import { useModal } from '@pm/hooks';
import { useCallback, useState } from 'react';
import { Member, PaPcResult } from '@pm/models';

export const PaPcResults = () => {
  const { isOpen: isUpsertModalOpened, setOpen: showHideUpsertModal } = useModal();
  const [selectedUpsertPaPcResult, setSelectedUpsertPaPcResult] = useState<PaPcResult | undefined>(
    undefined,
  );

  const { isOpen: isMemberInfoModalOpened, setOpen: showHideMemberInfoModal } = useModal();
  const [selectedMember, setSelectedMember] = useState<Member | undefined>(undefined);

  const toggleUpsertModal = useCallback(
    (paPcResult?: PaPcResult) => {
      setSelectedUpsertPaPcResult(paPcResult);
      showHideUpsertModal(true);
    },
    [showHideUpsertModal],
  );

  const toggleMemberInfoModal = useCallback(
    (member?: Member) => {
      setSelectedMember(member);
      showHideMemberInfoModal(true);
    },
    [showHideMemberInfoModal],
  );

  return (
    <StyledPaPcResults>
      <ListPaPcResultsTable
        toggleUpsertModal={toggleUpsertModal}
        toggleMemberInfoModal={toggleMemberInfoModal}
      />
      <UpsertPaPcResultModal
        isOpen={isUpsertModalOpened}
        setOpen={showHideUpsertModal}
        paPcResult={selectedUpsertPaPcResult}
      />
      <MemberInfoModal
        isOpen={isMemberInfoModalOpened}
        setOpen={showHideMemberInfoModal}
        member={selectedMember}
      />
    </StyledPaPcResults>
  );
};
