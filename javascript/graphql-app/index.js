const express = require('express');
const { graphqlHTTP } = require('express-graphql');
const { buildSchema } = require('graphql');

var schema = buildSchema(`
    type Query {
        hello: String
    }
`);

var root = {
    hello: () => {
        return 'Hello World!';
    }
}

const app = express();
const PORT = process.env.PORT || 5000;
app.use('/graphql', graphqlHTTP({
    schema: schema,
    rootValue: root,
    graphiql: true,
}));
app.listen(PORT, () => console.log('Server Listening On ', PORT))