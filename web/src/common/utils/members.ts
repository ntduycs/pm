import { MemberConstant, TMemberPosition } from '@pm/common/constants';

const shortenPosition = (role: TMemberPosition) => {
  switch (role) {
    case MemberConstant.Position.PM:
      return 'PM';
    case MemberConstant.Position.BE:
      return 'BE';
    case MemberConstant.Position.FE:
      return 'FE';
    case MemberConstant.Position.QC:
      return 'QC';
    case MemberConstant.Position.BA:
      return 'BA';
    case MemberConstant.Position.DESIGNER:
      return 'Designer';
    case MemberConstant.Position.MOBILE:
      return 'Mobile';
    default:
      return '';
  }
};

export const Members = {
  shortenPosition,
};
