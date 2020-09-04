let dburl = 'postgres://altcoinsensei:n!Bjw+tY6Q&e8U9X@51.159.24.51:17166/acn-crypto-trade';
// let dburl = 'postgres://chris:@localhost:5432/test';

let knex = require('knex')({
  client: "pg",
  connection: dburl,
  searchPath: ['knex', 'public']
  // connection: {
  //   user: 'chris',
  //   host: 'localhost',
  //   database: 'test',
  //   password: '',
  //   port: 5432
  // }
});

// knex.debug(true);

async function getTest() {
  try {
    let result = await knex.select().from('users');
    console.log(result);
    process.exit();
  } catch (e) {
    console.error(e);
    process.exit();
  }
}

getTest();



