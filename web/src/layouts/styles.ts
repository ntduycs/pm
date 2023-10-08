import styled from 'styled-components';
import { boxShadow, color } from '@pm/styles';

export const StyledDefaultLayout = styled.div`
  .aside {
    height: 100%;

    .ant-menu {
      height: calc(100vh - 4em);
      border-right: none;
      border-radius: 8px;
      ${boxShadow()};
      padding: 0.75em;
    }
  }

  .content {
    ${boxShadow()};
    padding: 0 1.25em;
    overflow-y: scroll;
    height: calc(100vh - 4em);
    border-radius: 8px;
    background-color: ${color.white};
  }
`;
