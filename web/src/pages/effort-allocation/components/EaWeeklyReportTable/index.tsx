import { StyledEaWeeklyReportTable } from './styles.ts';
import { useTableParams } from '@pm/hooks';
import { EaWeeklyMember } from '@pm/models';
import { ColumnProps } from 'antd/es/table';
import { ApiConstant, EffortAllocation } from '@pm/common/constants';
import { Button, Table } from 'antd';
import { uniqueId, toString } from 'lodash';
import { SorterResult } from 'antd/es/table/interface';
import { useQuery } from '@tanstack/react-query';
import { getEaWeeklyReportAPI } from '@pm/services';

export const EaWeeklyReportTable = () => {
  const { tableParams, setTableParams, onTableParamsChange } = useTableParams<EaWeeklyMember>();

  const { isFetching, data: weeklyReport } = useQuery({
    queryKey: ['effort-allocation', 'weekly', tableParams],
    queryFn: async () => {
      const response = await getEaWeeklyReportAPI({
        sort: toString(tableParams.sorter?.field) || 'name',
        direction: tableParams.sorter?.order || ApiConstant.DefaultSortDirection,
      });
      setTableParams({
        ...tableParams,
      });

      return response;
    },
    enabled: true,
    keepPreviousData: true,
  });

  const eaColumns: ColumnProps<EaWeeklyMember>[] = Object.values(EffortAllocation.Category).map(
    (item: { label: string }) => ({
      title: item.label,
      dataIndex: item.label,
      key: item.label,
      sorter: true,
      width: 180,
      render: (_, rc: EaWeeklyMember) => {
        const effort = rc.efforts.find((effort) => effort.category === item.label);
        return effort ? effort.time : 0;
      },
    }),
  );

  const columns: ColumnProps<EaWeeklyMember>[] = [
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
      sorter: true,
      width: 200,
      fixed: 'left',
      defaultSortOrder: 'ascend',
      render: (_, rc: EaWeeklyMember) => {
        return rc.member.name;
      },
    },
    ...eaColumns,
    {
      title: 'Total',
      dataIndex: 'total',
      key: 'total',
      sorter: true,
      width: 160,
      render: (_, rc: EaWeeklyMember) => {
        return rc.efforts.reduce((acc, effort) => acc + effort.time, 0);
      },
    },
    {
      title: 'Busy Rate',
      dataIndex: 'busyRate',
      key: 'busyRate',
      width: 160,
      render: (_, rc: EaWeeklyMember) => {
        return `${Math.round(
          (rc.efforts.reduce((acc, effort) => acc + effort.time, 0) / 40) * 100,
        ).toFixed(2)}%`;
      },
    },
  ];

  return (
    <StyledEaWeeklyReportTable>
      <Table
        columns={columns}
        dataSource={weeklyReport?.items || []}
        rowKey={() => uniqueId()}
        pagination={false}
        loading={isFetching}
        onChange={(pagination, filters, sorter) =>
          onTableParamsChange(pagination, filters, sorter as SorterResult<EaWeeklyMember>)
        }
        scroll={{ x: 1500 }}
        footer={() => <div>Total: {weeklyReport?.total}</div>}
      />
    </StyledEaWeeklyReportTable>
  );
};
