import { Member } from '@pm/pages/members/model.ts';
import { EMember } from '@pm/common/constants';

export const seedMembers: Member[] = [
  {
    name: 'Nguyen Van A',
    email: 'a.nguyen@gmail.com',
    level: EMember.Level.LV1,
    positions: [EMember.Position.PM],
    kpi: 100,
    category: EMember.Category.OFFICIAL,
    total_effort: 100,
    start_date: '2021-01-21',
    status: EMember.Status.ACTIVE,
  },
  {
    name: 'Nguyen Van B',
    email: 'b.nguyen@gmail.com',
    level: EMember.Level.LV2,
    positions: [EMember.Position.FE],
    kpi: 100,
    category: EMember.Category.BUFFER,
    total_effort: 100,
    start_date: '2021-01-01',
    status: EMember.Status.INACTIVE,
  },
  {
    name: 'Nguyen Van C',
    email: 'c.nguyen@gmail.com',
    level: EMember.Level.LV3,
    positions: [EMember.Position.DESIGNER, EMember.Position.QC],
    kpi: 100,
    category: EMember.Category.OJD,
    total_effort: 50,
    start_date: '2021-01-01',
    status: EMember.Status.PENDING,
  },
];
