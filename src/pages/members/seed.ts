import { Member } from '@pm/pages/members/meta.ts';
import { Category, Level, Position, Status } from '@pm/common/constants';

export const seedMembers: Member[] = [
  {
    name: 'Nguyen Van A',
    level: Level.LV1,
    positions: [Position.PM],
    kpi: 100,
    category: Category.OFFICIAL,
    totalEffort: 100,
    joinedDate: '2021-01-21',
    status: Status.ACTIVE,
  },
  {
    name: 'Nguyen Van B',
    level: Level.LV2,
    positions: [Position.FE],
    kpi: 100,
    category: Category.BUFFER,
    totalEffort: 100,
    joinedDate: '2021-01-01',
    status: Status.INACTIVE,
  },
  {
    name: 'Nguyen Van C',
    level: Level.LV3,
    positions: [Position.DESIGNER, Position.QC],
    kpi: 100,
    category: Category.OJD,
    totalEffort: 50,
    joinedDate: '2021-01-01',
    status: Status.PENDING,
  },
];
