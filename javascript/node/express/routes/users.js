var express = require('express');
var router = express.Router();
var mongoose = require('mongoose');
mongoose.connect('mongodb://localhost/testing', { useNewUrlParser: true, useUnifiedTopology: true });

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
  if (body.username === 'asdf' && body.password === 'qwer') {
    req.session.user = Object.assign({}, body);
    res.redirect('/');
  } else {
    res.redirect('/users/login');
  }
});

router.get('/logout', function (req, res, next) {
  if (req.session.user && req.cookies.user_sid) {
    res.clearCookie('user_sid');
    res.redirect('/');
  } else {
    res.redirect('/users/login');
  }
})

module.exports = router;
