import { ECategory, ELevel, EPosition, EStatus } from '../enums/member.enum';
import { PageResponseDto, SingularResponseDto } from './base.dto';

export class ListMemberDto extends PageResponseDto<MemberDto> {}
export class GetMemberDto extends SingularResponseDto<MemberDto> {}
export class UpsertMemberDto {
  id?: string;
  name: string;
  level: ELevel;
  positions: EPosition[];
  kpi: number;
  category: ECategory;
  totalEffort: number;
  startDate: string; // YYYY-MM-DD
  status: EStatus;
}
export class DeleteMemberDto {
  id: string;
}

export class MemberDto {
  id: string;
  name: string;
  level: ELevel;
  positions: EPosition[];
  kpi: number;
  category: ECategory;
  totalEffort: number;
  startDate: string; // YYYY-MM-DD
  status: EStatus;
}
