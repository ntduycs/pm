import { Suspense } from 'react';
import { AntdProvider, RecoilProvider, TanstackProvider } from '@pm/providers';
import { BrowserRouter } from 'react-router-dom';
import Routes from '@pm/routes.tsx';

const App = () => {
  return (
    <TanstackProvider>
      <RecoilProvider>
        <AntdProvider>
          <Suspense fallback={<div>Loading...</div>}>
            <BrowserRouter>
              <Routes />
            </BrowserRouter>
          </Suspense>
        </AntdProvider>
      </RecoilProvider>
    </TanstackProvider>
  );
};

export default App;
