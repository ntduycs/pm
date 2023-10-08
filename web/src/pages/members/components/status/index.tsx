import { MemberConstant, TMemberStatus } from '@pm/common/constants';
import { Tag } from 'antd';
import { capitalize } from 'lodash';
import { StyledStatus } from '@pm/pages/members/components/status/styles.ts';
import { color } from '@pm/styles';

export const Status = ({ status }: { status: TMemberStatus }) => {
  return (
    <StyledStatus>
      {status === MemberConstant.Status.ACTIVE ? (
        <Tag
          className='status'
          color={color.blue}
        >
          {capitalize(status)}
        </Tag>
      ) : status === MemberConstant.Status.INACTIVE ? (
        <Tag
          className='status'
          color={color.gray}
        >
          {capitalize(status)}
        </Tag>
      ) : (
        <Tag
          className='status'
          color={color.yellow}
        >
          {capitalize(status)}
        </Tag>
      )}
    </StyledStatus>
  );
};
