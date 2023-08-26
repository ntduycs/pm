import { Suspense } from 'react';
import { BaseuiProvider, RecoilProvider, StyletronProvider, TanstackProvider } from '@pm/providers';
import { BrowserRouter } from 'react-router-dom';
import Routes from '@pm/routes.tsx';

const App = () => {
  return (
    <TanstackProvider>
      <RecoilProvider>
        <StyletronProvider>
          <BaseuiProvider>
            <Suspense fallback={<div>Loading...</div>}>
              <BrowserRouter>
                <Routes />
              </BrowserRouter>
            </Suspense>
          </BaseuiProvider>
        </StyletronProvider>
      </RecoilProvider>
    </TanstackProvider>
  );
};

export default App;
