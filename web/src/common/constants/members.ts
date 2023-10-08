export type TMemberCategory = string;
export type TMemberLevel = number;
export type TMemberStatus = string;
export type TMemberPosition = string;

const Category = {
  OFFICIAL: 'official' as TMemberCategory,
  BUFFER: 'buffer' as TMemberCategory,
};

const Level = {
  LV1: 1 as TMemberLevel,
  LV2: 2 as TMemberLevel,
  LV3: 3 as TMemberLevel,
  LV4: 4 as TMemberLevel,
  LV5: 5 as TMemberLevel,
  LV6: 6 as TMemberLevel,
  LV7: 7 as TMemberLevel,
  LV8: 8 as TMemberLevel,
  LV9: 9 as TMemberLevel,
  LV10: 10 as TMemberLevel,
  LV11: 11 as TMemberLevel,
  LV12: 12 as TMemberLevel,
};

const Status = {
  ACTIVE: 'active' as TMemberStatus,
  INACTIVE: 'inactive' as TMemberStatus,
  PENDING: 'pending' as TMemberStatus,
};

const Position = {
  PM: 'Project Manager' as TMemberPosition,
  BE: 'Backend Engineer' as TMemberPosition,
  FE: 'Frontend Engineer' as TMemberPosition,
  QC: 'QA Engineer' as TMemberPosition,
  BA: 'Business Analyst' as TMemberPosition,
  DESIGNER: 'UI/UX Designer' as TMemberPosition,
  MOBILE: 'Mobile Developer' as TMemberPosition,
};

export const MemberConstant = {
  Category,
  Level,
  Status,
  Position,
};
