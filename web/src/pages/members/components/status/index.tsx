import { EMember, ETheme, TStatus } from '@pm/common/constants';
import { Tag } from 'antd';
import { capitalize } from 'lodash';
import { StyledStatus } from '@pm/pages/members/components/status/styles.ts';

export const Status = ({ status }: { status: TStatus }) => {
  status = status.toLowerCase();

  return (
    <StyledStatus>
      {status === EMember.Status.ACTIVE.toLowerCase() ? (
        <Tag
          className='status'
          color={ETheme.Color.GREEN}
        >
          {capitalize(status)}
        </Tag>
      ) : status === EMember.Status.INACTIVE.toLowerCase() ? (
        <Tag
          className='status'
          color={ETheme.Color.ORANGE}
        >
          {capitalize(status)}
        </Tag>
      ) : (
        <Tag
          className='status'
          color={ETheme.Color.YELLOW}
        >
          {capitalize(status)}
        </Tag>
      )}
    </StyledStatus>
  );
};
