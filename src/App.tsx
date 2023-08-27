import { Suspense } from 'react';
import { RecoilProvider, TanstackProvider } from '@pm/providers';
import { BrowserRouter } from 'react-router-dom';
import Routes from '@pm/routes.tsx';

const App = () => {
  return (
    <TanstackProvider>
      <RecoilProvider>
        <Suspense fallback={<div>Loading...</div>}>
          <BrowserRouter>
            <Routes />
          </BrowserRouter>
        </Suspense>
      </RecoilProvider>
    </TanstackProvider>
  );
};

export default App;
