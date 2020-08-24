const express = require('express');
const { graphqlHTTP } = require('express-graphql');
const { addPath } = require('graphql/jsutils/Path');
const {
  GraphQLSchema,
  GraphQLObjectType,
  GraphQLString,
  buildSchema
} = require('graphql');

const localGraphQL = require('./graphql/index');

const path = require('path');
const logger = require('morgan');

const app = express()

app.use(logger('dev'));

/*
  Types: String, Int, Float, Boolean, and ID.
  `!` Means it can't be nullable
  make lists with brackets: [Int]
*/

class Obj {
  constructor(name, value) {
    this.name = name;
    this.value = value;
  }
}

const schema = buildSchema(`
  type Obj {
    name: String,
    value: Int
  }

  type Query {
    id: ID,
    hello: String,
    random: Float!,
    ints: [Int],
    addWithMessage(num1: Int!, num2: Int!, msg: String): String,
    getObj(name: String, value: Int): Obj
  }

  type Mutation {
    setHello(value: String): String
  }
`);

let hello = 'Hello World...'

var root = {
  hello: () => {
    return hello;
  },
  random: () => {
    return Math.random();
  },
  ints: () => {
    return [1, 2, 3]
  },
  addWithMessage: ({ num1, num2, msg }) => {
    return msg + ": " + (num1 + num2);
  },
  getObj: ({ name, value }) => {
    return new Obj(name, value)
    // return 2;
  },
  setHello: ({ value }) => {
    hello = value;
    return hello;
  }
}

app.use('/graphql', graphqlHTTP({
  // schema,
  // schema: buildSchema(localGraphQL.schema),
  schema: localGraphQL.objSchema,
  // rootValue: root,
  // rootValue: localGraphQL.root,
  graphiql: process.env.NODE_ENV === 'development'
}))

app.get('/test', (req, res, next) => {
  res.send('tested');
})

let port = process.env.PORT || 5000;

app.listen(port, () => {
  console.log('app listening on ', port);
})