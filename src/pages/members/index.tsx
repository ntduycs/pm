import { StyledMembers } from '@pm/pages/members/styles.ts';
import { Table } from 'baseui/table-semantic';
import { columns } from '@pm/pages/members/meta.ts';
import { seedMembers } from '@pm/pages/members/seed.ts';
import { Arrays } from '@pm/common/utils';
import { AddMember } from '@pm/pages/members/components';
import { adapt } from '@pm/pages/members/util.ts';

export const Members = () => {
  return (
    <StyledMembers>
      <AddMember />
      <Table
        columns={columns}
        data={Arrays.fromObjects(seedMembers.map(adapt))}
      />
    </StyledMembers>
  );
};
