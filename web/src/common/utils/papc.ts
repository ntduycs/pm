import { MemberConstant, TMemberLevel } from '@pm/common/constants';
import { PaPcConstants, TPaPcCategory } from '@pm/common/constants/papc.ts';
import { PaPcResult } from '@pm/models';
import { Tag } from 'antd';

const evalTotalScore = (paPcResult: PaPcResult): number => {
  return Number(
    (
      paPcResult.technical_score * 0.2 +
      paPcResult.productivity_score * 0.45 +
      paPcResult.collaboration_score * 0.15 +
      paPcResult.development_score * 0.1 +
      0.6
    ).toFixed(2),
  );
};

const evalCategory = (level: TMemberLevel, totalScore: number): TPaPcCategory => {
  switch (level) {
    case MemberConstant.Level.LV1:
      if (totalScore >= 5.4 && totalScore <= 5.6) {
        return PaPcConstants.Category.MEETS_EXPECTATIONS;
      } else if (totalScore > 5.6) {
        return PaPcConstants.Category.EXCEEDS_EXPECTATIONS;
      } else if (totalScore < 5.2) {
        return PaPcConstants.Category.UNDER_EXPECTATIONS;
      } else {
        return PaPcConstants.Category.NEEDS_IMPROVEMENT;
      }
    case MemberConstant.Level.LV2:
      if (totalScore >= 5.6 && totalScore <= 5.8) {
        return PaPcConstants.Category.MEETS_EXPECTATIONS;
      } else if (totalScore > 5.8) {
        return PaPcConstants.Category.EXCEEDS_EXPECTATIONS;
      } else if (totalScore < 5.4) {
        return PaPcConstants.Category.UNDER_EXPECTATIONS;
      } else {
        return PaPcConstants.Category.NEEDS_IMPROVEMENT;
      }
    case MemberConstant.Level.LV3:
      if (totalScore >= 5.8 && totalScore <= 6) {
        return PaPcConstants.Category.MEETS_EXPECTATIONS;
      } else if (totalScore > 6) {
        return PaPcConstants.Category.EXCEEDS_EXPECTATIONS;
      } else if (totalScore < 5.6) {
        return PaPcConstants.Category.UNDER_EXPECTATIONS;
      } else {
        return PaPcConstants.Category.NEEDS_IMPROVEMENT;
      }
    case MemberConstant.Level.LV4:
      if (totalScore >= 6 && totalScore <= 6.2) {
        return PaPcConstants.Category.MEETS_EXPECTATIONS;
      } else if (totalScore > 6.2) {
        return PaPcConstants.Category.EXCEEDS_EXPECTATIONS;
      } else if (totalScore < 5.8) {
        return PaPcConstants.Category.UNDER_EXPECTATIONS;
      } else {
        return PaPcConstants.Category.NEEDS_IMPROVEMENT;
      }
    case MemberConstant.Level.LV5:
      if (totalScore >= 6.2 && totalScore <= 6.4) {
        return PaPcConstants.Category.MEETS_EXPECTATIONS;
      } else if (totalScore > 6.4) {
        return PaPcConstants.Category.EXCEEDS_EXPECTATIONS;
      } else if (totalScore < 6) {
        return PaPcConstants.Category.UNDER_EXPECTATIONS;
      } else {
        return PaPcConstants.Category.NEEDS_IMPROVEMENT;
      }
    case MemberConstant.Level.LV6:
      if (totalScore >= 6.4 && totalScore <= 6.6) {
        return PaPcConstants.Category.MEETS_EXPECTATIONS;
      } else if (totalScore > 6.6) {
        return PaPcConstants.Category.EXCEEDS_EXPECTATIONS;
      } else if (totalScore < 6.2) {
        return PaPcConstants.Category.UNDER_EXPECTATIONS;
      } else {
        return PaPcConstants.Category.NEEDS_IMPROVEMENT;
      }
    default:
      throw new Error('Invalid level');
  }
};

export const PaPcs = {
  evalCategory,
  evalTotalScore,
};
