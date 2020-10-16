const express = require('express')
const { graphqlHTTP } = require('express-graphql')
const {
  GraphQLSchema,
  GraphQLObjectType,
  GraphQLString,
  GraphQLList,
  GraphQLInt,
  GraphQLNonNull,
} = require('graphql')
const app = express()
const cors = require('cors')

const PORT = process.env.PORT || 4000

app.use(
  cors({
    origin: 'http://localhost:5000',
  }),
)

// app.use(function (req, res, next) {
//   console.log('mark')
//   res.setHeader('Access-Control-Allow-Origin', '*')
//   res.setHeader('Access-Control-Allow-Methods', '*')
//   res.setHeader('Access-Control-Allow-Headers', '*')
//   next()
// })

const authors = [
  { id: 1, name: 'J. K. Rowling' },
  { id: 2, name: 'J. R. R. Tolkien' },
  { id: 3, name: 'Brent Weeks' },
]

const books = [
  { id: 1, name: 'Harry Potter and the Chamber of Secrets', authorId: 1 },
  { id: 2, name: 'Harry Potter and the Prisoner of Azkaban', authorId: 1 },
  { id: 3, name: 'Harry Potter and the Goblet of Fire', authorId: 1 },
  { id: 4, name: 'The Fellowship of the Ring', authorId: 2 },
  { id: 5, name: 'The Two Towers', authorId: 2 },
  { id: 6, name: 'The Return of the King', authorId: 2 },
  { id: 7, name: 'The Way of Shadows', authorId: 3 },
  { id: 8, name: 'Beyond the Shadows', authorId: 3 },
]

// const schema = new GraphQLSchema({
//   query: new GraphQLObjectType({
//     name: 'HelloWorld',
//     fields: () => ({
//       message: {
//         type: GraphQLString,
//         resolve: () => 'Hello World',
//       },
//       test: {
//         type: GraphQLString,
//         resolve: () => 'Yowdy',
//       },
//     }),
//   }),
// })

const BookType = new GraphQLObjectType({
  name: 'book',
  description: 'This represents a book writen by an author',
  fields: () => ({
    // These need to match the object properties
    id: { type: GraphQLNonNull(GraphQLInt) },
    name: { type: GraphQLNonNull(GraphQLString) },
    authorId: { type: GraphQLNonNull(GraphQLInt) },

    // Any extra types that don't exist on the object will be null
    test: { type: GraphQLString },

    // We can also represent other types
    author: {
      type: AuthorType,

      // When we resolve, the parameter this the type we're currently in
      resolve: (book) => {
        return authors.find((author) => author.id === book.authorId)
      },
    },
  }),
})

const AuthorType = new GraphQLObjectType({
  name: 'author',
  description: 'This represents an author',
  fields: () => ({
    // These need to match the object properties
    id: { type: GraphQLNonNull(GraphQLInt) },
    name: { type: GraphQLNonNull(GraphQLString) },

    // test: {
    //   type: GraphQLString,
    //   resolve: (author) => author.id + author.name,
    // },
    books: {
      type: new GraphQLList(BookType),
      resolve: (author) => {
        return books.filter((book) => book.authorId === author.id)
      },
    },
  }),
})

const RootQueryType = new GraphQLObjectType({
  name: 'Query',
  diescription: 'Root Query',

  // This needs to be a function in the case that it references other types
  fields: () => ({
    // New Field to query by
    hello: {
      type: GraphQLString, // What kind of data it is
      description: "It's a test", // A description of it for documentation
      resolve: () => "I'm a test", // What it actually returns
    },

    book: {
      // If we want a unique javascript object inside it, we need to create a type
      type: BookType,
      description: 'A Single Book',

      // We can get arguments passed in to query our results
      args: {
        id: { type: GraphQLInt },
      },

      // Use those args to find the given book
      resolve: (parent, args) => books.find((book) => book.id === args.id),
    },
    books: {
      type: new GraphQLList(BookType),
      description: 'List of Books',
      resolve: () => books,
    },
    author: {
      type: AuthorType,
      description: 'A Single Author',
      args: {
        id: { type: GraphQLInt },
      },
      resolve: (parent, args) =>
        authors.find((author) => author.id === args.id),
    },
    authors: {
      type: new GraphQLList(AuthorType),
      description: 'List of All Authors',
      resolve: () => authors,
    },
  }),
})

// For mutation,s we need a root type
const RootMutationType = new GraphQLObjectType({
  name: 'Mutation',
  description: 'Root Mutation',

  // The list of mutations we can do
  fields: () => ({
    addBook: {
      type: BookType, // What schema we're saving it at
      description: 'Adding a Book',

      // How graphQL can relate to the type
      args: {
        name: { type: GraphQLNonNull(GraphQLString) },
        authorId: { type: GraphQLNonNull(GraphQLInt) },
      },

      // What we're doing to cause the mutation and what we're sending back
      resolve: (parent, args) => {
        const book = {
          id: books.length + 1,
          name: args.name,
          authorId: args.authorId,
        }
        books.push(book)
        return book
      },
    },

    addAuthor: {
      type: AuthorType,
      description: 'Adding an Author',
      args: {
        name: { type: GraphQLNonNull(GraphQLString) },
      },
      resolve: (parent, args) => {
        const author = {
          id: authors.length + 1,
          name: args.name,
        }
        authors.push(author)
        return author
      },
    },
  }),
})

const schema = new GraphQLSchema({
  query: RootQueryType,
  mutation: RootMutationType,
})

app.use(
  '/graphql',
  graphqlHTTP({
    // graphiql: true,
    schema,
  }),
)

app.get('/', (req, res) => res.send('ok'))

app.listen(PORT, () => console.log('App running on', PORT))
