import styled from 'styled-components';

export const StyledListMembers = styled.div`
  position: relative;

  .table-button {
    position: absolute;
    top: 18px;
    z-index: 9;
  }

  .member-info {
    display: flex;
    align-items: center;

    .member-tooltip {
      margin-left: 0.625em;
    }

    &:hover {
      cursor: pointer;
    }
  }
`;
