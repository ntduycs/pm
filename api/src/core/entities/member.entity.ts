import { Entity, Enum, Property } from '@mikro-orm/core';
import { ELevel, EPosition, EStatus } from '../enums/member.enum';
import { BaseEntity } from './base.entity';

@Entity({
  tableName: 'members',
})
export class MemberEntity extends BaseEntity {
  @Property()
  name!: string;

  @Enum(() => ELevel)
  level!: ELevel;

  @Enum({ items: () => EPosition, array: true })
  positions!: EPosition[];

  @Property()
  kpi!: number;

  @Property()
  category!: string;

  @Property()
  totalEffort!: number;

  @Property()
  joinedDate: Date = new Date();

  @Enum(() => EStatus)
  status!: EStatus;
}
