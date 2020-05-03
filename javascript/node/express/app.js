var createError = require('http-errors');
var express = require('express');
var path = require('path');
var cookieParser = require('cookie-parser');
var logger = require('morgan');
var sessions = require('express-session');
require('dotenv').config();

// Set up all mongoose settings and connect to the database
var mongoose = require('mongoose');
mongoose.set('useCreateIndex', true);
mongoose.set('useFindAndModify', false);
mongoose.connect('mongodb://localhost/testing', { useNewUrlParser: true, useUnifiedTopology: true });
var db = mongoose.connection;
db.on('error', console.error.bind(console, 'connection error:'));
const MongoStore = require('connect-mongo')(sessions);  // Used to save sessions in the db

// Routes
var indexRouter = require('./routes/index');
var usersRouter = require('./routes/users');

var app = express();

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'ejs');

// Middleware
app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));
app.use(sessions({
  key: 'user_sid',
  secret: process.env.SECRET,
  resave: false,
  saveUninitialized: false,
  cookie: {
    maxAge: 1000 * 60 * 60 * 24 * 30  // Set for 60 days
  },
  store: new MongoStore({ mongooseConnection: db })
}));

// Make sure there is a user tied to each session. If not, reset.
app.use((req, res, next) => {
  if (req.cookies.user_sid && !req.session.user) {
    res.clearCookie('user_sid');
  }
  next();
})

// This may seem a little redundant from above, but it's an extra check specifically for root
//   Otherwise we get stuck in a cookie loop. Can't really think of another word for it beside that...
const sessionChecker = (req, res, next) => {
  if (req.session.user && req.cookies.user_sid) {
    next();
  } else {
    res.redirect('/users/login');
  }
}

app.use('/users', usersRouter);   // Handle user stuff
app.use('/', sessionChecker, indexRouter);  // Main page

// catch 404 and forward to error handler
app.use(function (req, res, next) {
  next(createError(404));
});

// error handler
app.use(function (err, req, res, next) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.render('error');
});

module.exports = app;
