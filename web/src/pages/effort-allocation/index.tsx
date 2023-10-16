import { StyledEffortAllocation } from '@pm/pages/effort-allocation/styles.ts';
import {
  EaWeeklyReportCharts,
  EaWeeklyReportImportModal,
  EaWeeklyReportTable,
} from '@pm/pages/effort-allocation/components';

export const EffortAllocation = () => {
  return (
    <StyledEffortAllocation>
      <EaWeeklyReportImportModal />
      <EaWeeklyReportCharts />
      <EaWeeklyReportTable />
    </StyledEffortAllocation>
  );
};
