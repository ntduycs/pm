import { StyledEaWeeklyReportTable } from './styles.ts';
import { EaWeeklyMember } from '@pm/models';
import { ColumnProps } from 'antd/es/table';
import { EffortAllocation } from '@pm/common/constants';
import { Table } from 'antd';
import { uniqueId } from 'lodash';
import { useRecoilValue, useResetRecoilState } from 'recoil';
import { effortAllocationState } from '@pm/atoms';
import { useEffect, useMemo } from 'react';
import { Members } from '@pm/common/utils';

export const EaWeeklyReportTable = () => {
  const eaState = useRecoilValue(effortAllocationState);
  const resetEaState = useResetRecoilState(effortAllocationState);

  // Reset effort allocation state when unmount
  useEffect(() => {
    return () => resetEaState();
  }, [resetEaState]);

  const eaColumns: ColumnProps<EaWeeklyMember>[] = useMemo(
    () =>
      Object.values(EffortAllocation.Category).map((item: { label: string; type: number }) => ({
        title: item.label,
        dataIndex: item.label,
        key: item.label,
        sorter: (a, b) => {
          const effortA = a.efforts.find((effort) => effort.category === item.label);
          const effortB = b.efforts.find((effort) => effort.category === item.label);
          return (effortA?.time ?? 0) - (effortB?.time ?? 0);
        },
        width: 180,
        align: 'center',
        render: (_, rc: EaWeeklyMember) => {
          const effort = rc.efforts.find((effort) => effort.category === item.label);
          return effort ? effort.time : 0;
        },
        className: `ea-weekly-report-table__column ea-weekly-report-table__column--${
          item.type === 1 ? 'billable' : 'non-billable'
        }`,
      })),
    [],
  );

  const columns: ColumnProps<EaWeeklyMember>[] = [
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
      sorter: (a, b) => {
        return a.member.name.localeCompare(b.member.name);
      },
      width: 220,
      fixed: 'left',
      defaultSortOrder: 'ascend',
      render: (_, rc: EaWeeklyMember) => {
        return `${rc.member.name} (${Members.shortenPosition(rc.member.positions[0])})`;
      },
    },
    ...eaColumns,
    {
      title: 'Total',
      dataIndex: 'total',
      key: 'total',
      sorter: (a, b) => {
        return (
          a.efforts.reduce((acc, effort) => acc + effort.time, 0) -
          b.efforts.reduce((acc, effort) => acc + effort.time, 0)
        );
      },
      width: 140,
      align: 'center',
      fixed: 'right',
      render: (_, rc: EaWeeklyMember) => {
        return rc.efforts.reduce((acc, effort) => acc + effort.time, 0).toFixed(2);
      },
    },
    {
      title: 'Busy Rate',
      dataIndex: 'busyRate',
      key: 'busyRate',
      width: 150,
      align: 'center',
      fixed: 'right',
      render: (_, rc: EaWeeklyMember) => {
        const total = rc.efforts.reduce((acc, effort) => acc + effort.time, 0);
        const productTime = rc.efforts.reduce(
          (acc, effort) => acc + (effort.is_product_time ? effort.time : 0),
          0,
        );
        return `${Math.round((productTime / total) * 100).toFixed(2)}%`;
      },
    },
  ];

  return (
    <StyledEaWeeklyReportTable>
      <Table
        columns={columns}
        dataSource={eaState || []}
        rowKey={() => uniqueId()}
        pagination={false}
        loading={false}
        scroll={{ x: 1500, y: '40%' }}
        footer={() => <div>Total: {eaState?.length ?? 0}</div>}
      />
    </StyledEaWeeklyReportTable>
  );
};
