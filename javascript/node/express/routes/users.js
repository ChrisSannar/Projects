var express = require('express');
var router = express.Router();
var users = require('../schema/user.schema');

// Main login page, in the case the the user isn't already registered.
router.get('/login', function (req, res, next) {
  let error = '';
  if (req.session.user && req.cookies.user_sid) {
    res.redirect('/');
  } else {
    // req.session.user = 'asdf'
    res.render('login', { error });
  }
});

// Attempting to login
router.post('/login', function (req, res, next) {
  let body = req.body;

  // Check to see if the user exists
  users.findOne({ "username": body.username }, (err, user) => {

    // Any error we encounter, or any user we can't find, handle accordingly
    if (err) {
      return next();
    }
    else if (!user) {
      res.redirect('/users/login');
    }

    // If we're successful, match the password using a built in method
    else {
      user.comparePassword(body.password, function (err, passing) {
        if (err) return next();
        if (passing) {
          req.session.user = Object.assign({}, body); // Set the session with the user
          res.redirect('/');
        } else {
          return res.redirect('/users/login');
        }
      })
    }
  })
});

// Logging out, we just clear the cookie and redirect to the beginning.
router.get('/logout', function (req, res, next) {
  if (req.session.user && req.cookies.user_sid) {
    res.clearCookie('user_sid');
    res.redirect('/');
  } else {
    res.redirect('/users/login');
  }
});

// Reseting the password. Most of the logic is in the schema
router.post('/new-password', function (req, res, next) {
  let { username, newPassword } = req.body;
  users.findOneAndUpdate(
    { username: username },
    { password: newPassword },
    function (err, result) {
      if (err) return next(err);

      res.redirect('/users/login');
    });
});

module.exports = router;
