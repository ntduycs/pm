import { Button, Table, TablePaginationConfig, Tag } from 'antd';
import { useQuery } from '@tanstack/react-query';
import { listMembersAPI } from '@pm/services';
import {
  MemberConstant,
  TCategory,
  TPosition,
  TStatus,
  ApiConstant,
  UiConstant,
} from '@pm/common/constants';
import { Status } from '@pm/pages/members/components';
import { Member } from '@pm/models';
import { DeleteOutlined, EditOutlined } from '@ant-design/icons';
import { useState } from 'react';
import { FilterValue, SorterResult } from 'antd/es/table/interface';
import { capitalize, toString } from 'lodash';
import { ColumnProps } from 'antd/es/table';
import dayjs from 'dayjs';
import { StyledListMembers } from '@pm/pages/members/components/list/styles.ts';
import { color } from '@pm/styles';
import { TableFooter } from '@pm/components';

type ListMembersProps = {
  toggleUpsertModal: (rc?: Member) => void;
  toggleDeleteModal: (rc?: Member) => void;
};

interface TableParams {
  pagination?: TablePaginationConfig;
  sorter?: SorterResult<Member>;
  filters?: Record<string, FilterValue>;
}

export const ListMembersTable = ({ toggleUpsertModal, toggleDeleteModal }: ListMembersProps) => {
  const [tableParams, setTableParams] = useState<TableParams>({
    pagination: {
      defaultCurrent: 1,
      defaultPageSize: 10,
      showSizeChanger: true,
      position: ['bottomRight', 'topRight'],
    },
  });

  const { isFetching, data: members } = useQuery({
    queryKey: ['members', tableParams],
    queryFn: async () => {
      const response = await listMembersAPI({
        page: tableParams.pagination?.current || ApiConstant.DefaultPage,
        size: tableParams.pagination?.pageSize || ApiConstant.DefaultPageSize,
        sort: toString(tableParams.sorter?.field) || 'name',
        direction: tableParams.sorter?.order || ApiConstant.DefaultSortDirection,
        category: tableParams.filters?.category
          ? (tableParams.filters?.category as unknown as TCategory)
          : undefined,
        positions: tableParams.filters?.positions
          ? (tableParams.filters?.positions as unknown as TPosition[])
          : undefined,
        status: tableParams.filters?.status
          ? (tableParams.filters?.status as unknown as TStatus)
          : undefined,
      });
      setTableParams({
        ...tableParams,
        pagination: {
          ...tableParams.pagination,
          total: response.total,
          current: response.page,
          pageSize: response.size,
        },
      });
      return response;
    },
    enabled: true,
    keepPreviousData: true,
  });

  const columns: ColumnProps<Member>[] = [
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
      sorter: true,
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
      render: (level: number) => `Level ${level}`,
    },
    {
      title: 'Positions',
      dataIndex: 'positions',
      key: 'positions',
      render: (positions: TPosition[]) => positions.join(', '),
      filters: Object.entries(MemberConstant.Position).map(([, value]) => ({
        text: value,
        value: value,
      })),
    },
    {
      title: 'KPI (%)',
      dataIndex: 'kpi',
      key: 'kpi',
      render: (kpi: number) => `${kpi}%`,
    },
    {
      title: 'Category',
      dataIndex: 'category',
      key: 'category',
      filters: Object.entries(MemberConstant.Category).map(([, value]) => ({
        text: capitalize(value),
        value: value,
      })),
      filterMultiple: false,
      render: (category: TCategory) => capitalize(category),
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
      render: (start_date: string) =>
        start_date
          ? dayjs(start_date).format(UiConstant.DefaultDateFormat)
          : UiConstant.DefaultEmptyValue,
    },
    {
      title: 'End Date',
      dataIndex: 'end_date',
      key: 'end_date',
      render: (end_date: string) =>
        end_date
          ? dayjs(end_date).format(UiConstant.DefaultDateFormat)
          : UiConstant.DefaultEmptyValue,
    },
    {
      title: 'Status',
      dataIndex: 'status',
      key: 'status',
      render: (status: TStatus) => <Status status={status} />,
      filters: Object.entries(MemberConstant.Status).map(([, value]) => ({
        text: capitalize(value),
        value: value,
      })),
      filterMultiple: false,
    },
    {
      title: 'Action(s)',
      dataIndex: 'actions',
      key: 'actions',
      render: (_: unknown, rc: Member) => {
        return (
          <div className='actions'>
            <Tag
              className='action'
              color={color.blue}
              onClick={() => toggleUpsertModal(rc)}
            >
              <EditOutlined />
            </Tag>
            <Tag
              className='action'
              color={color.red}
              onClick={() => toggleDeleteModal(rc)}
            >
              <DeleteOutlined />
            </Tag>
          </div>
        );
      },
    },
  ];

  const handleTableChange = (
    pagination: TablePaginationConfig,
    filters: Record<string, FilterValue | null>,
    sorter: SorterResult<Member>,
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

  return (
    <StyledListMembers>
      <Button
        className='new-member-btn'
        type='primary'
        onClick={() => toggleUpsertModal()}
      >
        New Member
      </Button>
      <Table
        columns={columns}
        rowKey={(record: Member) => record.id}
        rowClassName={(record: Member) => `${record.status}-row`}
        pagination={tableParams.pagination}
        loading={isFetching}
        onChange={(pagination, filters, sorter) =>
          handleTableChange(pagination, filters, sorter as SorterResult<Member>)
        }
        dataSource={members?.items || []}
        footer={() => <TableFooter data={members} />}
      />
    </StyledListMembers>
  );
};
