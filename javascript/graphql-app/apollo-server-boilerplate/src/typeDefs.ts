import { gql } from 'apollo-server-express';

export const typeDefs = gql`
  type Query {
    getUser(id: Int!): User
  }
  type Mutation {
    addUser(username: String!, password: String!): Boolean!
  }

  type User {
    id: Int!
    username: String!,
    password: String!
  }
`;