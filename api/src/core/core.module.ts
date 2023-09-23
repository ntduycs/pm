import { Module } from '@nestjs/common';
import { MemberRepository } from './repositories/member.repository';
import { MikroOrmModule } from '@mikro-orm/nestjs';
import { MemberEntity } from './entities/member.entity';

@Module({
  providers: [MemberRepository],
  exports: [MemberRepository],
  imports: [MikroOrmModule.forFeature([MemberEntity])],
})
export class CoreModule {}
