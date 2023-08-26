import { ROLE, SIZE } from 'baseui/modal';
import { StyledModal } from '@pm/components/modal/styles.ts';
import { ReactNode } from 'react';

interface ModalProps {
  children: ReactNode;
  isOpen?: boolean;
  onClose?: () => void;
  size?: 'default' | 'full' | 'auto';
}

export const Modal = (props: ModalProps) => {
  return (
    <StyledModal
      isOpen={props.isOpen ?? false}
      onClose={() => {
        props.onClose?.();
      }}
      closeable
      animate
      autoFocus
      size={props.size ?? SIZE.default}
      // overrides={{
      //   Root: {
      //     style: {
      //       zIndex: 9999,
      //     },
      //   },
      // }}
    >
      {props.children}
    </StyledModal>
  );
};
