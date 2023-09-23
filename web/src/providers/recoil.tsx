import { RecoilRoot } from 'recoil';
import { PropsWithChildren } from 'react';

export const RecoilProvider = ({ children }: PropsWithChildren) => {
  return <RecoilRoot>{children}</RecoilRoot>;
};
