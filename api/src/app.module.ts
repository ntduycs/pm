import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ConfigModule } from '@nestjs/config';
import { databaseConfig, serverConfig } from './config';
import { MemberModule } from './member/member.module';
import { MikroOrmModule } from '@mikro-orm/nestjs';
import { CoreModule } from './core/core.module';
import mikroConfig from './config/mikro.config';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      load: [serverConfig, databaseConfig],
    }),
    MikroOrmModule.forRoot(mikroConfig),
    CoreModule,
    MemberModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
