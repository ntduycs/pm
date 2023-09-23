import { Injectable } from '@nestjs/common';
import { UpsertMemberDto } from '../core/dto/member.dto';
import { MemberRepository } from '../core/repositories/member.repository';

@Injectable()
export class MemberService {
  constructor(private readonly memberRepository: MemberRepository) {}

  create(dto: UpsertMemberDto) {
    return 'This action adds a new member';
  }

  findAll() {
    return `This action returns all member`;
  }

  findOne(id: string) {
    try {
      return this.memberRepository.findById(id);
    } catch (error) {
      console.log(error);
      throw new Error('Member not found');
    }
  }

  update(id: number, dto: UpsertMemberDto) {
    return `This action updates a #${id} member`;
  }

  remove(id: number) {
    return `This action removes a #${id} member`;
  }
}
