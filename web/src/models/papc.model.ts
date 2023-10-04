import { Member } from '@pm/models/member.model.ts';
import { PageRequest, PageResponse } from '@pm/models/common.model.ts';

export interface PaPcResult {
  id: number;
  member: Member;
  technical_score: number;
  productivity_score: number;
  collaboration_score: number;
  development_score: number;
  period: string;
}

export interface PaPcTechnicalScore {
  id: number;
  pa_pc_id: number;
  type: 'soft-skills' | 'hard-skills';
  skill: string;
  self_score: number;
  peer_score?: number;
  manager_score: number;
}

export interface ListPaPcResultsRequest extends PageRequest {
  member_id?: number;
  period?: string;
}

export interface UpsertPaPcResultRequest {
  id?: number;
  member_id: number;
  technical_score: number;
  productivity_score: number;
  collaboration_score: number;
  development_score: number;
  period: string;
}

export interface ListPaPcResultsResponse extends PageResponse<PaPcResult> {}
