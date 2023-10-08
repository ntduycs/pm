export type TPaPcCategory = string;

const Category = {
  EXCEEDS_EXPECTATIONS: 'Exceeds Expectations' as TPaPcCategory,
  MEETS_EXPECTATIONS: 'Meets Expectations' as TPaPcCategory,
  NEEDS_IMPROVEMENT: 'Needs Improvement' as TPaPcCategory,
  UNDER_EXPECTATIONS: 'Under Expectations' as TPaPcCategory,
};

export const PaPcConstants = {
  Category,
};
