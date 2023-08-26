import { Client } from 'styletron-engine-atomic';
import { PropsWithChildren } from 'react';
import { Provider } from 'styletron-react';

const engine = new Client();

export const StyletronProvider = ({ children }: PropsWithChildren) => (
  <Provider value={engine}>{children}</Provider>
);
