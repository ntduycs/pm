import { TCategory, TLevel, TPosition, TStatus } from '@pm/common/constants';

export interface Member {
  id: string;
  name: string;
  level: TLevel;
  positions: TPosition[];
  kpi: number;
  category: TCategory;
  totalEffort: number;
  startDate: string; // YYYY-MM-DD
  endDate: string; // YYYY-MM-DD
  status: TStatus;
}

export const columns = [
  {
    title: 'Name',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: 'Level',
    dataIndex: 'level',
    key: 'level',
  },
  {
    title: 'Positions',
    dataIndex: 'positions',
    key: 'positions',
  },
  {
    title: 'KPI (%)',
    dataIndex: 'kpi',
    key: 'kpi',
  },
  {
    title: 'Category',
    dataIndex: 'category',
    key: 'category',
  },
  {
    title: 'Total Effort (%)',
    dataIndex: 'totalEffort',
    key: 'totalEffort',
  },
  {
    title: 'Joined Date',
    dataIndex: 'joinedDate',
    key: 'joinedDate',
  },
  {
    title: 'Status',
    dataIndex: 'status',
    key: 'status',
  },
];
