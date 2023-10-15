import { TSortDirection } from '@pm/common/constants';

export interface PageResponse<T> {
  page: number;
  size: number;
  total: number;
  pages: number;
  count: number;
  items: T[];
}

export interface ListResponse<T> {
  items: T[];
  total: number;
}

export interface SingularResponse<T> {
  item: T;
}

export interface EmptyResponse {
  message: string;
}

export interface ErrorResponse {
  error_message: string;
}

export interface PageRequest {
  page: number;
  size: number;
  sort: string;
  direction: TSortDirection;
}

export interface ListRequest {
  sort: string;
  direction: TSortDirection;
}
