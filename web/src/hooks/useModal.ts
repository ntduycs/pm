import { useCallback, useState } from 'react';

export const useModal = (isOpenDefault?: boolean) => {
  const [isOpen, setIsOpen] = useState(() => !!isOpenDefault);

  const setOpen = useCallback((state: boolean) => {
    setIsOpen(state);
  }, []);

  return { isOpen, setOpen };
};
