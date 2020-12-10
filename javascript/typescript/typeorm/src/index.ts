import { PostgresConnectionOptions } from 'typeorm/driver/postgres/PostgresConnectionOptions';
import { createConnection, Connection } from 'typeorm';
import models from './models';
const { Test, Test2 } = models;

const typeOrmConfig: PostgresConnectionOptions = {
  type: "postgres",
  host: "0.0.0.0",
  port: 5432,
  username: "root",
  password: "password",
  database: "typeormTest",
  synchronize: true,
  logging: false,
  entities: [
    Test, Test2
  ]
};

async function populateDatabase(connection: Connection) {
  await connection.query('DELETE FROM test;');
  await connection.query('DELETE FROM test2;');
  const test = new Test();
  test.value = "asdf"
  const test2 = new Test();
  test2.value = "zxcv";

  const test3 = new Test2();
  test3.value = 'qwer';

  test.others = [test3];
  test2.others = [test3];

  await connection.getRepository(Test).save(test);
  await connection.getRepository(Test).save(test2);
  await connection.getRepository(Test2).save(test3);

  // Because the 
  await connection.getRepository(Test).delete({ value: "asdf" });
}

(async () => {
  const connection: Connection = await createConnection(typeOrmConfig);
  await populateDatabase(connection);


  const repo: any = await connection.getRepository(Test);
  // const repo2: any = await connection.getRepository(Test2);
  console.log(...await repo.find());

  await connection.close();

})()
