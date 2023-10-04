import { Member, PaPcResult } from '@pm/models';
import { useQuery } from '@tanstack/react-query';
import { listPaPcResultsAPI } from '@pm/services/papc.service.ts';
import { ApiConstant } from '@pm/common/constants';
import { toString } from 'lodash';
import { ColumnProps } from 'antd/es/table';
import { useTableParams } from '@pm/hooks';
import { StyledListMembers } from '@pm/pages/papc/components/list/styles.ts';
import { Button, Table, Tag } from 'antd';
import { SorterResult } from 'antd/es/table/interface';
import { TableFooter } from '@pm/components';
import { color } from '@pm/styles';
import { EditOutlined } from '@ant-design/icons';

type ListMembersProps = {
  toggleUpsertModal: (rc?: PaPcResult) => void;
};

export const ListPaPcResultsTable = ({ toggleUpsertModal }: ListMembersProps) => {
  const { tableParams, setTableParams, onTableParamsChange } = useTableParams<PaPcResult>();

  const { isFetching, data: paPcResults } = useQuery({
    queryKey: ['pa-pc-results', tableParams],
    queryFn: async () => {
      const response = await listPaPcResultsAPI({
        page: tableParams.pagination?.current || ApiConstant.DefaultPage,
        size: tableParams.pagination?.pageSize || ApiConstant.DefaultPageSize,
        sort: toString(tableParams.sorter?.field) || 'member_id',
        direction: tableParams.sorter?.order || ApiConstant.DefaultSortDirection,
        member_id: tableParams.filters?.member_id
          ? (tableParams.filters?.member_id as unknown as number)
          : undefined,
        period: tableParams.filters?.period
          ? (tableParams.filters?.period as unknown as string)
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

  const columns: ColumnProps<PaPcResult>[] = [
    {
      title: 'Member',
      dataIndex: 'member',
      key: 'member',
      render: (member: Member) => member.name,
    },
    {
      title: 'Period',
      dataIndex: 'period',
      key: 'period',
    },
    {
      title: 'Technical Score (20%)',
      dataIndex: 'technical_score',
      key: 'technical_score',
      sorter: true,
      render: (score: number) => score.toFixed(2),
    },
    {
      title: 'Productivity Score (45%)',
      dataIndex: 'productivity_score',
      key: 'productivity_score',
      sorter: true,
      render: (score: number) => score.toFixed(2),
    },
    {
      title: 'Collaboration Score (15%)',
      dataIndex: 'collaboration_score',
      key: 'collaboration_score',
      sorter: true,
      render: (score: number) => score.toFixed(2),
    },
    {
      title: 'Development Score (10%)',
      dataIndex: 'development_score',
      key: 'development_score',
      sorter: true,
      render: (score: number) => score.toFixed(2),
    },
    {
      title: 'Total Score',
      dataIndex: 'total_score',
      key: 'total_score',
      render: (_, rc) => {
        return (
          rc.technical_score * 0.2 +
          rc.productivity_score * 0.45 +
          rc.collaboration_score * 0.15 +
          rc.development_score * 0.1 +
          0.6
        ).toFixed(2);
      },
    },
    {
      title: 'Actions',
      dataIndex: 'actions',
      key: 'actions',
      render: (_, rc) => {
        return (
          <div className='actions'>
            <Tag
              className='action'
              color={color.blue}
              onClick={() => toggleUpsertModal(rc)}
            >
              <EditOutlined />
            </Tag>
          </div>
        );
      },
    },
  ];

  return (
    <StyledListMembers>
      <Button
        type='primary'
        onClick={() => toggleUpsertModal()}
        style={{ marginBottom: '1em' }}
        className='table-button'
      >
        New
      </Button>
      <Table
        columns={columns}
        rowKey={(record) => record.id}
        pagination={tableParams.pagination}
        loading={isFetching}
        onChange={(pagination, filters, sorter) => {
          onTableParamsChange(pagination, filters, sorter as SorterResult<PaPcResult>);
        }}
        dataSource={paPcResults?.items || []}
        footer={() => <TableFooter data={paPcResults} />}
      />
    </StyledListMembers>
  );
};
