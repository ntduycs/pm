import { StyledDefaultLayout } from '@pm/layouts/styles.ts';
import { Outlet } from 'react-router-dom';

export const DefaultLayout = () => {
  return (
    <StyledDefaultLayout>
      <Outlet />
    </StyledDefaultLayout>
  );
};
