export type TCategory = string;
export type TLevel = string;
export type TStatus = string;
export type TPosition = string;

const Category = {
  OFFICIAL: 'Official' as TCategory,
  BUFFER: 'Buffer' as TCategory,
  OJD: 'OJD' as TCategory,
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
  LV11: 'LV11' as TLevel,
  LV12: 'LV12' as TLevel,
  LV13: 'LV13' as TLevel,
  LV14: 'LV14' as TLevel,
  LV15: 'LV15' as TLevel,
};

const Status = {
  ACTIVE: 'Active' as TStatus,
  INACTIVE: 'Inactive' as TStatus,
  PENDING: 'Pending' as TStatus,
};

const Position = {
  PM: 'Project Manager' as TPosition,
  BE: 'Back-end' as TPosition,
  FE: 'Front-end' as TPosition,
  QC: 'QA/Tester' as TPosition,
  BA: 'Business Analyst' as TPosition,
  DESIGNER: 'Designer' as TPosition,
  MOBILE: 'Mobile' as TPosition,
};

export const EMember = {
  Category,
  Level,
  Status,
  Position,
};
