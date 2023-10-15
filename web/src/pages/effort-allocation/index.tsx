import { StyledEffortAllocation } from '@pm/pages/effort-allocation/styles.ts';
import {
  EaWeeklyReportImportModal,
  EaWeeklyReportTable,
} from '@pm/pages/effort-allocation/components';

export const EffortAllocation = () => {
  return (
    <StyledEffortAllocation>
      <EaWeeklyReportImportModal />
      <EaWeeklyReportTable />
    </StyledEffortAllocation>
  );
};
