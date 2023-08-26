import { useRecoilValue } from 'recoil';
import { langState } from '@pm/atoms/lang.tsx';
import { useCallback } from 'react';
import i18next from 'i18next';
import { StyledDefaultLayout } from '@pm/layouts/styles.ts';
import { Outlet } from 'react-router-dom';

export const DefaultLayout = () => {
  const lang = useRecoilValue(langState);

  useCallback(() => {
    i18next.changeLanguage(lang).then(() => {
      console.log(`lang: ${lang}`);
    });
  }, [lang]);

  return (
    <StyledDefaultLayout>
      <Outlet />
    </StyledDefaultLayout>
  );
};
