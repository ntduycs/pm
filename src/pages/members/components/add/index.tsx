import { StyledAddMember } from '@pm/pages/members/components/add/styles.ts';
import { Button, KIND } from 'baseui/button';
import { SyntheticEvent, useMemo, useState } from 'react';
import { ModalBody, ModalButton, ModalFooter, ModalHeader } from 'baseui/modal';
import { Modal } from '@pm/components';
import { Controller, SubmitHandler, useForm } from 'react-hook-form';
import { Member } from '@pm/pages/members/meta.ts';
import { Input } from 'baseui/input';
import { Combobox } from 'baseui/combobox';
import { Category, Level, Position, Status } from '@pm/common/constants';
import dayjs from 'dayjs';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';

const schema = yup.object({
  name: yup.string().required().trim(),
  level: yup.string().required().oneOf(Object.values(Level)),
  kpi: yup.number().required(),
  category: yup.string().required().oneOf(Object.values(Category)),
  positions: yup
    .array()
    .of(yup.string().required().oneOf(Object.values(Position)))
    .required(),
  joinedDate: yup
    .string()
    .required()
    .matches(/\d{4}-\d{2}-\d{2}/), // YYYY-MM-DD
  status: yup.string().required().oneOf(Object.values(Status)),
  totalEffort: yup.number().required().min(0).max(100),
});

export const AddMember = () => {
  const [isOpen, setIsOpen] = useState(false);

  const [typedLevel, setTypedLevel] = useState('');

  const { handleSubmit, control } = useForm<Member>({
    defaultValues: {
      name: '',
      level: Level.LV1,
      kpi: 0,
      category: Category.OFFICIAL,
      positions: [],
      joinedDate: dayjs().format('YYYY-MM-DD'),
      status: Status.PENDING,
      totalEffort: 0,
    },
    resolver: yupResolver(schema),
  });

  const openModal = (e: SyntheticEvent<HTMLButtonElement>) => {
    setIsOpen(true);
  };

  const closeModal = () => {
    setIsOpen(false);
  };

  const onSubmit: SubmitHandler<Member> = (data: Member) => {
    console.log(data);
  };

  const onClickSubmit = () => {
    console.log('submit');
  };

  const filteredLevelOptions = useMemo(() => {
    const options = Object.values(Level).map((level) => ({ id: level, label: level }));
    return options.filter((level) => level.label.includes(typedLevel));
  }, [typedLevel]);

  return (
    <StyledAddMember>
      <Button onClick={openModal}>New Member</Button>
      <Modal isOpen={isOpen}>
        <ModalHeader>Add Member</ModalHeader>
        <ModalBody>
          <form onSubmit={handleSubmit(onSubmit)}>
            <Controller
              name='name'
              control={control}
              render={({ field: { ref, ...field }, formState: { errors } }) => {
                return (
                  <Input
                    clearable
                    required
                    {...field}
                  />
                );
              }}
            />
            <Controller
              name='level'
              control={control}
              render={({ field: { ref, ...field }, formState: { errors } }) => {
                return (
                  <Combobox
                    options={filteredLevelOptions}
                    mapOptionToString={(option: { id: string; label: string }) => option.label}
                    {...field}
                  />
                );
              }}
            />
          </form>
        </ModalBody>
        <ModalFooter>
          <ModalButton
            kind={KIND.tertiary}
            onClick={closeModal}
          >
            Cancel
          </ModalButton>
          <ModalButton
            onClick={onClickSubmit}
            kind={KIND.primary}
          >
            Okay
          </ModalButton>
        </ModalFooter>
      </Modal>
    </StyledAddMember>
  );
};
