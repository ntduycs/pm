import { TCategory, TLevel, TPosition, TStatus } from '@pm/common/constants';
import { capitalize } from 'lodash';

export interface Member {
  id?: string;
  name: string;
  email: string;
  level: TLevel;
  positions: TPosition[];
  kpi: number;
  category: TCategory;
  total_effort: number;
  start_date?: string; // YYYY-MM-DD
  end_date?: string; // YYYY-MM-DD
  status: TStatus;
}

export const columns = [
  {
    title: 'Name',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: 'Email',
    dataIndex: 'email',
    key: 'email',
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
    render: (kpi: number) => `${kpi * 100}%`,
  },
  {
    title: 'Category',
    dataIndex: 'category',
    key: 'category',
  },
  {
    title: 'Total Effort (%)',
    dataIndex: 'total_effort',
    key: 'total_effort',
    render: (total_effort: number) => `${total_effort}%`,
  },
  {
    title: 'Start Date',
    dataIndex: 'start_date',
    key: 'start_date',
  },
  {
    title: 'End Date',
    dataIndex: 'end_date',
    key: 'end_date',
  },
  {
    title: 'Status',
    dataIndex: 'status',
    key: 'status',
    render: (status: TStatus) => capitalize(status),
  },
];
