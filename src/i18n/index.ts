import i18next from 'i18next';
import { initReactI18next } from 'react-i18next';
import vi from './locales/vi.json';
import en from './locales/en.json';

export const defaultNS = 'en';
export const resources = {
  en: { translation: en },
  vi: { translation: vi },
} as const;

i18next
  .use(initReactI18next)
  .init({
    fallbackLng: 'en',
    interpolation: { escapeValue: false },
    resources,
  })
  .then(() => {
    console.log('i18next is ready');
  });

export default i18next;
