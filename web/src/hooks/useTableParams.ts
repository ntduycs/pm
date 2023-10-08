import { useMemo, useState } from 'react';
import { TableParams } from '@pm/models';
import { TablePaginationConfig } from 'antd';
import { FilterValue, SorterResult } from 'antd/es/table/interface';
import { ApiConstant } from '@pm/common/constants';

export const useTableParams = <T extends object>() => {
  const [tableParams, setTableParams] = useState<TableParams<T>>({
    pagination: {
      defaultCurrent: ApiConstant.DefaultPage,
      defaultPageSize: ApiConstant.DefaultPageSize,
      showSizeChanger: true,
      position: ['bottomRight', 'topRight'],
    },
  });

  const onTableParamsChange = (
    pagination: TablePaginationConfig,
    filters: Record<string, FilterValue | null>,
    sorter: SorterResult<T>,
  ) => {
    // remove null values from filters
    const nonNullFilters = Object.fromEntries(
      Object.entries(filters).filter(([_, v]) => v !== null),
    ) as Record<string, FilterValue>;

    setTableParams({
      pagination,
      filters: nonNullFilters,
      sorter: {
        field: sorter.field,
        order: sorter.order,
      },
    });
  };

  return { tableParams, setTableParams, onTableParamsChange };
};
