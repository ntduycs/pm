import { StyledPaPc } from '@pm/pages/papc/styles.ts';
import { ListPaPcResultsTable, UpsertPaPcResultModal } from '@pm/pages/papc/components';
import { useModal } from '@pm/hooks';
import { useCallback, useState } from 'react';
import { PaPcResult } from '@pm/models';

export const PaPc = () => {
  const { isOpen: isUpsertModalOpened, setOpen: showHideUpsertModal } = useModal();
  const [selectedUpsertPaPcResult, setSelectedUpsertPaPcResult] = useState<PaPcResult | undefined>(
    undefined,
  );

  const toggleUpsertModal = useCallback(
    (paPcResult?: PaPcResult) => {
      setSelectedUpsertPaPcResult(paPcResult);
      showHideUpsertModal(true);
    },
    [showHideUpsertModal],
  );

  return (
    <StyledPaPc>
      <ListPaPcResultsTable toggleUpsertModal={toggleUpsertModal} />
      <UpsertPaPcResultModal
        isOpen={isUpsertModalOpened}
        setOpen={showHideUpsertModal}
        paPcResult={selectedUpsertPaPcResult}
      />
    </StyledPaPc>
  );
};
