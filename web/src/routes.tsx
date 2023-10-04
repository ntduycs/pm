import { RouteObject, useRoutes } from 'react-router-dom';
import { DefaultLayout } from '@pm/layouts';
import { Home, Members, PaPc } from '@pm/pages';

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
      {
        path: '/pa-pc',
        element: <PaPc />,
      },
    ],
  },
];

const Routes = () => {
  return useRoutes(routes);
};

export default Routes;
