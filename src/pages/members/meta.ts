import { TCategory, TLevel, TPosition, TStatus } from '@pm/common/constants';

export interface Member {
  name: string;
  level: TLevel;
  positions: TPosition[];
  kpi: number;
  category: TCategory;
  totalEffort: number;
  joinedDate: string; // YYYY-MM-DD
  status: TStatus;
}

export const columns = [
  'Name',
  'Level',
  'Position',
  'KPI (%)',
  'Category',
  'Total Effort (%)',
  'Joined Date',
  'Status',
];
