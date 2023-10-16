import { atom } from 'recoil';
import { EaWeeklyMember } from '@pm/models';

export const effortAllocationState = atom({
  key: 'effortAllocationState',
  default: [] as EaWeeklyMember[],
});
