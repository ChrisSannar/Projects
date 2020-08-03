const {
  GraphQLSchema,
  GraphQLObjectType,
  GraphQLString
} = require('graphql');

let lorem = ["Lorem", "ipsum", "dolor", "sit", "amet,", "consectetur", "adipiscing", "elit,", "sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore", "magna", "aliqua.", "Ut", "enim", "ad", "minim", "veniam,", "quis", "nostrud", "exercitation", "ullamco", "laboris", "nisi", "ut", "aliquip", "ex", "ea", "commodo", "consequat.", "Duis", "aute", "irure", "dolor", "in", "reprehenderit", "in", "voluptate", "velit", "esse", "cillum", "dolore", "eu", "fugiat", "nulla", "pariatur.", "Excepteur", "sint", "occaecat", "cupidatat", "non", "proident,", "sunt", "in", "culpa", "qui", "officia", "deserunt", "mollit", "anim", "id", "est", "laborum."];

module.exports = {
  schema: `
    type Query {
      hello: String,
      message(value: String!): [String]
    }

    type Mutation {
      addLib(value: String!): [String],
      removeLib(value: String!): [String]
    }
  `,
  root: {
    hello: () => {
      return "Howdy";
    },
    message: ({ value }) => {
      let result = lorem.filter(wurd => wurd.includes(value));
      return result;
    },
    addLib: ({ value }) => {
      lorem.push(value);
      return lorem;
    },
    removeLib: async ({ value }) => {
      return await new Promise((res, rej) => {
        setTimeout(() => {
          lorem = lorem.filter(wurd => wurd !== value);
          res(lorem)
        }, 1000)
      })
    }
  }
}