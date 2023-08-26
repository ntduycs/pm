import { BaseProvider, LightTheme } from 'baseui';
import { PropsWithChildren } from 'react';

export const BaseuiProvider = ({ children }: PropsWithChildren) => {
  return <BaseProvider theme={LightTheme}>{children}</BaseProvider>;
};
