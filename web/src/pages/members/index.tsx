import { StyledMembers } from '@pm/pages/members/styles.ts';
import { columns } from '@pm/pages/members/model.ts';
import { seedMembers } from '@pm/pages/members/seed.ts';
import { AddMember } from '@pm/pages/members/components';
import { adapt } from '@pm/pages/members/util.ts';
import { Table } from 'antd';

export const Members = () => {
  return (
    <StyledMembers>
      <AddMember />
      <Table
        columns={columns}
        dataSource={seedMembers.map(adapt)}
      />
    </StyledMembers>
  );
};
