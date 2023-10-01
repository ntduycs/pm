export interface PageResponse<T> {
  page: number;
  size: number;
  total: number;
  pages: number;
  count: number;
  items: T[];
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
  direction: string;
}
