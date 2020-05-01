var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function (req, res, next) {
  if (req.session.user && req.cookies.user_sid) {
    res.redirect('/');
  } else {
    req.session.user = 'asdf'
    res.send(`${req.session.user}`);
  }
});

module.exports = router;
