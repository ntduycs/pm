export type TCategory = string;
export type TLevel = number;
export type TStatus = string;
export type TPosition = string;

const Category = {
  OFFICIAL: 'official' as TCategory,
  BUFFER: 'buffer' as TCategory,
};

const Level = {
  LV1: 1 as TLevel,
  LV2: 2 as TLevel,
  LV3: 3 as TLevel,
  LV4: 4 as TLevel,
  LV5: 5 as TLevel,
  LV6: 6 as TLevel,
  LV7: 7 as TLevel,
  LV8: 8 as TLevel,
  LV9: 9 as TLevel,
  LV10: 10 as TLevel,
  LV11: 11 as TLevel,
  LV12: 12 as TLevel,
};

const Status = {
  ACTIVE: 'active' as TStatus,
  INACTIVE: 'inactive' as TStatus,
  PENDING: 'pending' as TStatus,
};

const Position = {
  PM: 'Project Manager' as TPosition,
  BE: 'Backend Engineer' as TPosition,
  FE: 'Frontend Engineer' as TPosition,
  QC: 'QA Engineer' as TPosition,
  BA: 'Business Analyst' as TPosition,
  DESIGNER: 'UI/UX Designer' as TPosition,
  MOBILE: 'Mobile Developer' as TPosition,
};

export const MemberConstant = {
  Category,
  Level,
  Status,
  Position,
};
