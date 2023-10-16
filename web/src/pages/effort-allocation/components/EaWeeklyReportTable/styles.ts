import styled from 'styled-components';
import { color } from '@pm/styles';

export const StyledEaWeeklyReportTable = styled.div`
  position: relative;

  .table-button {
    margin: 1.125em 0;
  }

  th.ea-weekly-report-table__column {
    &--billable {
      background-color: ${color.greenBg} !important;
    }

    &--non-billable {
      background-color: ${color.orangeBg} !important;
    }
  }
`;
