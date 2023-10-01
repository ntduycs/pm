import { MemberConstant, ThemeConstant, TStatus } from '@pm/common/constants';
import { Tag } from 'antd';
import { capitalize } from 'lodash';
import { StyledStatus } from '@pm/pages/members/components/status/styles.ts';

export const Status = ({ status }: { status: TStatus }) => {
  return (
    <StyledStatus>
      {status === MemberConstant.Status.ACTIVE ? (
        <Tag
          className='status'
          color={ThemeConstant.Color.GREEN}
        >
          {capitalize(status)}
        </Tag>
      ) : status === MemberConstant.Status.INACTIVE ? (
        <Tag
          className='status'
          color={ThemeConstant.Color.ORANGE}
        >
          {capitalize(status)}
        </Tag>
      ) : (
        <Tag
          className='status'
          color={ThemeConstant.Color.YELLOW}
        >
          {capitalize(status)}
        </Tag>
      )}
    </StyledStatus>
  );
};
