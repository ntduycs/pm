import { PageResponse, SingularResponse } from '@pm/models/common.model.ts';
import { TCategory, TLevel, TPosition, TStatus } from '@pm/common/constants';

export interface Member {
  id: number;
  name: string;
  email: string;
  level: TLevel;
  position: TPosition[];
  category: TCategory;
  start_date?: string;
  end_date?: string;
  status: TStatus;
  total_effort: number;
  kpi: number;
}

export interface UpsertMemberRequest {
  id?: number;
  name: string;
  email: string;
  level: TLevel;
  position: TPosition[];
  category: TCategory;
  start_date?: string;
  end_date?: string;
  status: TStatus;
  total_effort: number;
  kpi: number;
}

export interface ListMembersResponse extends PageResponse<Member> {}
export interface GetMemberResponse extends SingularResponse<Member> {}
