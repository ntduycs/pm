import { atom } from 'recoil';
import { Language } from '@pm/common/constants';

export const langState = atom({
  key: 'langState',
  default: Language.EN,
});
