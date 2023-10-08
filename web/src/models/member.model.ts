import { PageRequest, PageResponse, SingularResponse } from '@pm/models/common.model.ts';
import {
  TMemberCategory,
  TMemberLevel,
  TMemberPosition,
  TMemberStatus,
} from '@pm/common/constants';

export interface Member {
  id: number;
  name: string;
  email: string;
  level: TMemberLevel;
  positions: TMemberPosition[];
  category: TMemberCategory;
  start_date?: string;
  end_date?: string;
  status: TMemberStatus;
  total_effort: number;
  kpi: number;
}

export interface ListMembersRequest extends PageRequest {
  category?: TMemberCategory;
  positions?: TMemberPosition[];
  status?: TMemberStatus;
}

export interface UpsertMemberRequest {
  id?: number;
  name: string;
  email: string;
  level: TMemberLevel;
  position: TMemberPosition[];
  category: TMemberCategory;
  start_date?: string;
  end_date?: string;
  status: TMemberStatus;
  total_effort: number;
  kpi: number;
}

export interface ListMembersResponse extends PageResponse<Member> {}
export interface GetMemberResponse extends SingularResponse<Member> {}
