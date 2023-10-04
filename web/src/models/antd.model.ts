import { TablePaginationConfig } from 'antd';
import { FilterValue, SorterResult } from 'antd/es/table/interface';

export interface TableParams<T extends object> {
  pagination?: TablePaginationConfig;
  sorter?: SorterResult<T>;
  filters?: Record<string, FilterValue>;
}
