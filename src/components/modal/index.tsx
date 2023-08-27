import { StyledModal } from '@pm/components/modal/styles.ts';
import { ReactNode, MouseEvent } from 'react';

interface ModalProps {
  title: ReactNode;
  children: ReactNode;
  isOpen: boolean;
  onOk?: (e?: MouseEvent<HTMLButtonElement>) => void;
  onCancel?: (e?: MouseEvent<HTMLButtonElement>) => void;
  footer?: ReactNode;
}

export const Modal = (props: ModalProps) => {
  return (
    <StyledModal
      title={props.title}
      open={props.isOpen}
      destroyOnClose={true}
      closable={true}
      footer={props.footer ?? null}
      onOk={props.onOk}
      onCancel={props.onCancel}
    >
      {props.children}
    </StyledModal>
  );
};
