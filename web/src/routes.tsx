import { RouteObject, useRoutes } from 'react-router-dom';
import { DefaultLayout } from '@pm/layouts';
import { Home, Members } from '@pm/pages';

const routes: RouteObject[] = [
  {
    path: '/',
    element: <DefaultLayout />,
    children: [
      {
        path: '/',
        element: <Home />,
      },
      {
        path: '/members',
        element: <Members />,
      },
    ],
  },
];

const Routes = () => {
  return useRoutes(routes);
};

export default Routes;
