export type TCategory = string;
export type TLevel = string;
export type TStatus = string;
export type TPosition = string;

const Category = {
  OFFICIAL: 'Official' as TCategory,
  BUFFER: 'Buffer' as TCategory,
};

const Level = {
  LV1: 'LV1' as TLevel,
  LV2: 'LV2' as TLevel,
  LV3: 'LV3' as TLevel,
  LV4: 'LV4' as TLevel,
  LV5: 'LV5' as TLevel,
  LV6: 'LV6' as TLevel,
  LV7: 'LV7' as TLevel,
  LV8: 'LV8' as TLevel,
  LV9: 'LV9' as TLevel,
  LV10: 'LV10' as TLevel,
};

const Status = {
  ACTIVE: 'Active' as TStatus,
  INACTIVE: 'Inactive' as TStatus,
  PENDING: 'Pending' as TStatus,
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

export const EMember = {
  Category,
  Level,
  Status,
  Position,
};
