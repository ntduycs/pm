import { Button, Spin, Table, Tag } from 'antd';
import { useQuery } from '@tanstack/react-query';
import { listMembersAPI } from '@pm/services';
import { ETheme, TPosition, TStatus } from '@pm/common/constants';
import { Status } from '@pm/pages/members/components';
import { Member } from '@pm/models';
import { DeleteOutlined, EditOutlined } from '@ant-design/icons';

type ListMembersProps = {
  toggleUpsertModal: (rc?: Member) => void;
};

export const ListMembersTable = ({ toggleUpsertModal }: ListMembersProps) => {
  const { isFetching, data, error } = useQuery({
    queryKey: ['members', 'list'],
    queryFn: listMembersAPI,
    meta: {
      page: 1,
      size: 10,
    },
  });

  const columns = [
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
      render: (positions: TPosition[]) => positions.join(', '),
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
      render: (status: TStatus) => <Status status={status} />,
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
              color={ETheme.Color.BLUE}
              onClick={() => toggleUpsertModal(rc)}
            >
              <EditOutlined />
            </Tag>
            <Tag
              className='action'
              color={ETheme.Color.ORANGE}
              onClick={() => console.log('delete')}
            >
              <DeleteOutlined />
            </Tag>
          </div>
        );
      },
    },
  ];

  return (
    <>
      <Button
        type='primary'
        onClick={() => toggleUpsertModal()}
      >
        New Member
      </Button>
      <Spin spinning={isFetching}>
        <Table
          columns={columns}
          dataSource={data?.items.map((item) => ({
            ...item,
            key: item.id,
          }))}
        />
      </Spin>
    </>
  );
};
