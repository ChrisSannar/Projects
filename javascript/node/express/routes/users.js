var express = require('express');
var router = express.Router();
var users = require('../schema/user.schema');

// *** TESTING
router.get('/', function (req, res, next) {
  users.find().then(result => {
    res.json(result);
  })
});
// ***

/* GET home page. */
router.get('/login', function (req, res, next) {
  let error = '';
  if (req.session.user && req.cookies.user_sid) {
    res.redirect('/');
  } else {
    // req.session.user = 'asdf'
    res.render('login', { error });
  }
});

router.post('/login', function (req, res, next) {
  let body = req.body;
  users.findOne({ "username": body.username }, (err, user) => {
    if (err) {
      return next();
    }
    else if (!user) {
      res.redirect('/users/login');
    }
    else {
      user.comparePassword(body.password, function (err, match) {
        console.log('MATCH', match)
        if (err) return next();
        if (match) {
          req.session.user = Object.assign({}, body);
          res.redirect('/');
        } else {
          return res.redirect('/users/login');
        }
      })
    }
  })
});

router.get('/logout', function (req, res, next) {
  if (req.session.user && req.cookies.user_sid) {
    res.clearCookie('user_sid');
    res.redirect('/');
  } else {
    res.redirect('/users/login');
  }
});

router.post('/new-password', function (req, res, next) {
  let { username, newPassword } = req.body;
  users.findOneAndUpdate(
    { username: username },
    { password: newPassword },
    function (err, result) {
      if (err) return next(err);

      console.log('RESULT', result);
      res.redirect('/users/login');
    });
});

module.exports = router;
