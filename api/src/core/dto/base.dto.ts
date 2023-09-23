export abstract class PageRequestDto {
  readonly page: number = 1;
  readonly limit: number = 20;
  readonly sort: string = 'id';
  readonly order: string = 'asc';
}

export abstract class PageResponseDto<T> {
  public page: number;
  public size: number;
  public total: number;
  public pages: number;
  public items: T[];
}

export abstract class SingularResponseDto<T> {
  public item: T;
}
