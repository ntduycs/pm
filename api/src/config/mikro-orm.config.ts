import { MikroOrmModuleSyncOptions } from '@mikro-orm/nestjs/typings';

export default {
  entities: ['src/core/entities'],
  entitiesTs: ['dist/core/entities'],
  type: 'postgresql',
  host: process.env.DB_HOST,
  port: Number(process.env.DB_PORT),
  dbName: process.env.DB_NAME,
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  debug: process.env.PROFILE === 'local',
  migrations: {
    path: 'src/migrations',
    pattern: /^[\w-]+\d+\.ts$/,
    tableName: 'migrations',
    transactional: true,
    allOrNothing: true,
    emit: 'ts',
  },
} as MikroOrmModuleSyncOptions;
