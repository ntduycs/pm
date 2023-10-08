import { Member, PaPcResult } from '@pm/models';
import { useQuery } from '@tanstack/react-query';
import { listPaPcResultsAPI } from '@pm/services/papc.service.ts';
import { ApiConstant } from '@pm/common/constants';
import { toString } from 'lodash';
import { ColumnProps } from 'antd/es/table';
import { useTableParams } from '@pm/hooks';
import { StyledListMembers } from '@pm/pages/papc/components/list/styles.ts';
import { Button, Table, Tag, Tooltip } from 'antd';
import { SorterResult } from 'antd/es/table/interface';
import { TableFooter } from '@pm/components';
import { color } from '@pm/styles';
import { EditOutlined, InfoCircleOutlined } from '@ant-design/icons';
import { PaPcs } from '@pm/common/utils/papc.ts';
import { PaPcConstants, TPaPcCategory } from '@pm/common/constants/papc.ts';

type ListMembersProps = {
  toggleUpsertModal: (rc?: PaPcResult) => void;
  toggleMemberInfoModal: (rc?: Member) => void;
};

const renderCategoryTag = (category: TPaPcCategory) => {
  switch (category) {
    case PaPcConstants.Category.EXCEEDS_EXPECTATIONS:
      return <Tag color={color.blue}>{category}</Tag>;
    case PaPcConstants.Category.MEETS_EXPECTATIONS:
      return <Tag color={color.green}>{category}</Tag>;
    case PaPcConstants.Category.NEEDS_IMPROVEMENT:
      return <Tag color={color.orange}>{category}</Tag>;
    case PaPcConstants.Category.UNDER_EXPECTATIONS:
      return <Tag color={color.red}>{category}</Tag>;
    default:
      throw new Error('Invalid category');
  }
};

export const ListPaPcResultsTable = ({
  toggleUpsertModal,
  toggleMemberInfoModal,
}: ListMembersProps) => {
  const { tableParams, setTableParams, onTableParamsChange } = useTableParams<PaPcResult>();

  const { isFetching, data: paPcResults } = useQuery({
    queryKey: ['pa-pc-results', tableParams],
    queryFn: async () => {
      const response = await listPaPcResultsAPI({
        page: tableParams.pagination?.current || ApiConstant.DefaultPage,
        size: tableParams.pagination?.pageSize || ApiConstant.DefaultPageSize,
        sort: toString(tableParams.sorter?.field) || 'name',
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
      dataIndex: 'name',
      key: 'name',
      width: 220,
      render: (_, rc) => {
        return (
          <div
            className='member-info'
            onClick={() => toggleMemberInfoModal(rc.member)}
          >
            <span>{rc.member.name}</span>
            <span className='member-tooltip'>
              <InfoCircleOutlined />
            </span>
          </div>
        );
      },
      sorter: true,
      fixed: 'left',
      defaultSortOrder: 'ascend',
    },
    {
      title: 'Level',
      dataIndex: 'level',
      key: 'level',
      width: 80,
      render: (_, rc) => `Level ${rc.member.level}`,
      fixed: 'left',
    },
    {
      title: 'Period',
      dataIndex: 'period',
      key: 'period',
      width: 100,
      filters: ['2023-Q3'].map((value) => ({
        text: value,
        value: value,
      })),
    },
    {
      title: 'Technical (20%)',
      dataIndex: 'technical_score',
      key: 'technical_score',
      sorter: true,
      render: (score: number) => score.toFixed(2),
      width: 140,
      align: 'center',
    },
    {
      title: 'Productivity (45%)',
      dataIndex: 'productivity_score',
      key: 'productivity_score',
      sorter: true,
      render: (score: number) => score.toFixed(2),
      width: 140,
      align: 'center',
    },
    {
      title: 'Collaboration (15%)',
      dataIndex: 'collaboration_score',
      key: 'collaboration_score',
      sorter: true,
      render: (score: number) => score.toFixed(2),
      width: 140,
      align: 'center',
    },
    {
      title: 'Development (10%)',
      dataIndex: 'development_score',
      key: 'development_score',
      sorter: true,
      render: (score: number) => score.toFixed(2),
      width: 140,
      align: 'center',
    },
    {
      title: 'Average',
      dataIndex: 'total_score',
      key: 'total_score',
      render: (_, rc) => {
        return PaPcs.evalTotalScore(rc);
      },
      width: 100,
      align: 'center',
    },
    {
      title: 'Category',
      dataIndex: 'category',
      key: 'category',
      render: (_, rc) => {
        const category = PaPcs.evalCategory(rc.member.level, PaPcs.evalTotalScore(rc));
        return renderCategoryTag(category);
      },
      width: 180,
      align: 'center',
    },
    {
      title: 'Note',
      dataIndex: 'note',
      key: 'note',
      width: 180,
      ellipsis: true,
      render: (note: string) => {
        return (
          <Tooltip
            title={note}
            placement='topLeft'
            autoAdjustOverflow
          >
            <span>{note}</span>
          </Tooltip>
        );
      },
    },
    {
      title: 'Actions',
      dataIndex: 'actions',
      key: 'actions',
      width: 100,
      align: 'center',
      fixed: 'right',
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
        sticky={true}
        scroll={{ x: 1500 }}
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
