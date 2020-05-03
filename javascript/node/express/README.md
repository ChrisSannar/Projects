# Express Projects

## Sessions

First, I wanted to establish and understand sessions properly, so I followed some tutorials till I began to understand their nature and purpose. A session is the server keeping track of a connection to a particular client, so no matter how many times that client refreshes and connects they will still have access to that same information.

Session data is kept server-side while the means to connect to that session is store in a cookie. Once a connection is established, the server uses the cookie to determine if there is a sessions with this particular server. This session can also hold important content such as shopping cart info as well as login information. That way, when the user returns to the page, the server can simply log them in automatically because of the cookie tied to the browser.

The purpose of why I built this project was to understand how user authentication works with the backend and have a much more secure and fluid experience with my application. Also it is a nice playground to implement Mongoose before I apply it to code-explainer and its dashboard.

Technologies used:

- Express.js - Server and application
- Mongodb - Store the users and session information
- Mongoose - Have an organized schema to validate data with the database.

Bugs/Things learned:

- Express will run the first path listed that matches the pattern, so be sure to have those paths lower on the list such as the root path ('/'). That goes at the very last.
- It's important to seperate components from each other so they can be easily interchanged with other projects. The mongoose schemas should be in their own file and let the app connect to the database so they can remained loosely coupled while maintaining high cohesion.
