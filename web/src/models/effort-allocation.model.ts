import { Member } from '@pm/models/member.model.ts';
import { ListRequest, ListResponse } from '@pm/models/common.model.ts';

export interface EaWeeklyMember {
  member: Member;
  efforts: EaWeeklyEffort[];
}

export interface EaWeeklyEffort {
  category: string;
  time: number;
}

export interface GetEaWeeklyReportRequest extends ListRequest {}

export interface GetEaWeeklyReportResponse extends ListResponse<EaWeeklyMember> {}
