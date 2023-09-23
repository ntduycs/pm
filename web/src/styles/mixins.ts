import { css } from 'styled-components';

export const flex = (justifyContent: string, alignItems: string) => css`
  display: flex;
  justify-content: ${justifyContent};
  align-items: ${alignItems};
`;

export const transition = (property = 'all', duration = '0.4s', timingFunction: 'ease') => css`
  -webkit-transition: ${property} ${duration} ${timingFunction};
  -o-transition: ${property} ${duration} ${timingFunction};
  transition: ${property} ${duration} ${timingFunction};
`;
