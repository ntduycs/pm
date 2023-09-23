import { MemberEntity } from '../entities/member.entity';
import { EntityRepository } from '@mikro-orm/postgresql';
import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@mikro-orm/nestjs';

@Injectable()
export class MemberRepository {
  constructor(
    @InjectRepository(MemberEntity)
    private readonly memberRepository: EntityRepository<MemberEntity>,
  ) {}

  async findById(id: string): Promise<MemberEntity> {
    return this.memberRepository.findOneOrFail({ id });
  }

  async findAll(): Promise<MemberEntity[]> {
    return this.memberRepository.findAll();
  }

  async save(dto: MemberEntity): Promise<MemberEntity> {
    return this.memberRepository.upsert(dto);
  }

  async delete(id: string): Promise<void> {
    await this.memberRepository.nativeDelete({ id });
  }
}
