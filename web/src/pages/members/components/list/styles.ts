import styled from 'styled-components';
import { color } from '@pm/styles';

export const StyledListMembers = styled.div`
  position: relative;

  .table-button {
    position: absolute;
    top: 18px;
    z-index: 9;
  }

  .active-row {
  }

  .inactive-row {
    background-color: ${color.grayBg};
  }

  .action {
    &:hover {
      cursor: pointer;
    }
  }
`;
