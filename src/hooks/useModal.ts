import { useState } from 'react';

export const useModal = (isOpenDefault?: boolean) => {
  const [isOpen, setIsOpen] = useState(() => (isOpenDefault ? isOpenDefault : false));

  const toggle = () => setIsOpen(!isOpen);

  return { isOpen, toggle };
};
