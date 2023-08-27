import { Member } from '@pm/pages/members/meta.ts';
import dayjs from 'dayjs';

export const adapt = (member: Member) => ({
  key: member.name,
  name: member.name,
  level: member.level,
  positions: member.positions.join(', '),
  kpi: member.kpi,
  category: member.category,
  totalEffort: member.totalEffort,
  joinedDate: dayjs(member.joinedDate).format('MMM DD, YYYY'),
  status: member.status,
});
