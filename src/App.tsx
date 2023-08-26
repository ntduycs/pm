import { Suspense, useState } from 'react';
import reactLogo from './assets/react.svg';
import viteLogo from '/vite.svg';
import './App.css';
import { BaseuiProvider, RecoilProvider, StyletronProvider, TanstackProvider } from '@pm/providers';
import { I18nextProvider } from 'react-i18next';
import i18n from '@pm/i18n';
import { BrowserRouter } from 'react-router-dom';
import Routes from '@pm/routes.tsx';

const App = () => {
  return (
    <TanstackProvider>
      <RecoilProvider>
        <StyletronProvider>
          <BaseuiProvider>
            <I18nextProvider i18n={i18n}>
              <Suspense fallback={<div>Loading...</div>}>
                <BrowserRouter>
                  <Routes />
                </BrowserRouter>
              </Suspense>
            </I18nextProvider>
          </BaseuiProvider>
        </StyletronProvider>
      </RecoilProvider>
    </TanstackProvider>
  );
};

export default App;
