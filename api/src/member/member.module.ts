import { Module } from '@nestjs/common';
import { MemberService } from './member.service';
import { MemberController } from './member.controller';
import { CoreModule } from '../core/core.module';

@Module({
  controllers: [MemberController],
  providers: [MemberService],
  imports: [CoreModule],
})
export class MemberModule {}
