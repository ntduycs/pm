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

export const boxShadow = () => css`
  -webkit-box-shadow: rgba(149, 157, 165, 0.2) 0 0.5em 1.5em;
  -moz-box-shadow: rgba(149, 157, 165, 0.2) 0 0.5em 1.5em;
  box-shadow: rgba(149, 157, 165, 0.2) 0 0.5em 1.5em;
`;
