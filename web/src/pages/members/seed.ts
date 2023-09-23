import { Member } from '@pm/pages/members/model.ts';
import { EMember } from '@pm/common/constants';

export const seedMembers: Member[] = [
  {
    name: 'Nguyen Van A',
    level: EMember.Level.LV1,
    positions: [EMember.Position.PM],
    kpi: 100,
    category: EMember.Category.OFFICIAL,
    totalEffort: 100,
    startDate: '2021-01-21',
    status: EMember.Status.ACTIVE,
  },
  {
    name: 'Nguyen Van B',
    level: EMember.Level.LV2,
    positions: [EMember.Position.FE],
    kpi: 100,
    category: EMember.Category.BUFFER,
    totalEffort: 100,
    startDate: '2021-01-01',
    status: EMember.Status.INACTIVE,
  },
  {
    name: 'Nguyen Van C',
    level: EMember.Level.LV3,
    positions: [EMember.Position.DESIGNER, EMember.Position.QC],
    kpi: 100,
    category: EMember.Category.OJD,
    totalEffort: 50,
    startDate: '2021-01-01',
    status: EMember.Status.PENDING,
  },
];
