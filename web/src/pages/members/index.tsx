import { StyledMembers } from '@pm/pages/members/styles.ts';
import { columns } from '@pm/pages/members/model.ts';
import { AddMember } from '@pm/pages/members/components';
import { Spin, Table } from 'antd';
import { useQuery } from '@tanstack/react-query';
import { listMembersAPI } from '@pm/services';

export const Members = () => {
  const { isFetching, isError, data, error } = useQuery({
    queryKey: ['members', 'list'],
    queryFn: listMembersAPI,
    meta: {
      page: 1,
      size: 10,
    },
  });

  return (
    <StyledMembers>
      <AddMember />
      <Spin spinning={isFetching}>
        <Table
          columns={columns}
          dataSource={data?.items.map((item) => ({
            ...item,
            key: item.id,
          }))}
        />
      </Spin>
    </StyledMembers>
  );
};
