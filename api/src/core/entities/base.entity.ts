import { PrimaryKey, Utils } from '@mikro-orm/core';
import { v4 } from 'uuid';
import { EntityName } from '@mikro-orm/nestjs';

export abstract class BaseEntity {
  @PrimaryKey()
  id: string = v4();
}

export const getRepositoryToken = <T>(entity: EntityName<T>) => {
  return `${Utils.className(entity)}Repository`;
};
