import { RouteObject, useRoutes } from 'react-router-dom';
import { DefaultLayout } from '@pm/layouts';
import { EffortAllocation, Home, Members, PaPcResults } from '@pm/pages';

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
        path: '/pa-pc-results',
        element: <PaPcResults />,
      },
      {
        path: '/effort-allocation',
        element: <EffortAllocation />,
      },
    ],
  },
];

const Routes = () => {
  return useRoutes(routes);
};

export default Routes;
