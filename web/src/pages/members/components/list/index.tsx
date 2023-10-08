import { Button, Table, Tag } from 'antd';
import { useQuery } from '@tanstack/react-query';
import { listMembersAPI } from '@pm/services';
import {
  MemberConstant,
  TMemberCategory,
  TMemberPosition,
  TMemberStatus,
  ApiConstant,
  UiConstant,
} from '@pm/common/constants';
import { Status } from '@pm/pages/members/components';
import { Member } from '@pm/models';
import { DeleteOutlined, EditOutlined } from '@ant-design/icons';
import { SorterResult } from 'antd/es/table/interface';
import { capitalize, toString } from 'lodash';
import { ColumnProps } from 'antd/es/table';
import dayjs from 'dayjs';
import { StyledListMembers } from '@pm/pages/members/components/list/styles.ts';
import { color } from '@pm/styles';
import { TableFooter } from '@pm/components';
import { useTableParams } from '@pm/hooks';

type ListMembersProps = {
  toggleUpsertModal: (rc?: Member) => void;
  toggleDeleteModal: (rc?: Member) => void;
};

export const ListMembersTable = ({ toggleUpsertModal, toggleDeleteModal }: ListMembersProps) => {
  const { tableParams, setTableParams, onTableParamsChange } = useTableParams<Member>();

  const { isFetching, data: members } = useQuery({
    queryKey: ['members', tableParams],
    queryFn: async () => {
      const response = await listMembersAPI({
        page: tableParams.pagination?.current || ApiConstant.DefaultPage,
        size: tableParams.pagination?.pageSize || ApiConstant.DefaultPageSize,
        sort: toString(tableParams.sorter?.field) || 'name',
        direction: tableParams.sorter?.order || ApiConstant.DefaultSortDirection,
        category: tableParams.filters?.category
          ? (tableParams.filters?.category as unknown as TMemberCategory)
          : undefined,
        positions: tableParams.filters?.positions
          ? (tableParams.filters?.positions as unknown as TMemberPosition[])
          : undefined,
        status: tableParams.filters?.status
          ? (tableParams.filters?.status as unknown as TMemberStatus)
          : MemberConstant.Status.ACTIVE,
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
      render: (positions: TMemberPosition[]) => positions.join(', '),
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
      render: (category: TMemberCategory) => capitalize(category),
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
      render: (status: TMemberStatus) => <Status status={status} />,
      filters: Object.entries(MemberConstant.Status).map(([, value]) => ({
        text: capitalize(value),
        value: value,
      })),
      filterMultiple: false,
      defaultFilteredValue: [MemberConstant.Status.ACTIVE],
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

  return (
    <StyledListMembers>
      <Button
        className='new-member-btn table-button'
        type='primary'
        onClick={() => toggleUpsertModal()}
      >
        New
      </Button>
      <Table
        columns={columns}
        rowKey={(record: Member) => record.id}
        rowClassName={(record: Member) => `${record.status}-row`}
        pagination={tableParams.pagination}
        loading={isFetching}
        onChange={(pagination, filters, sorter) =>
          onTableParamsChange(pagination, filters, sorter as SorterResult<Member>)
        }
        dataSource={members?.items || []}
        footer={() => <TableFooter data={members} />}
      />
    </StyledListMembers>
  );
};
