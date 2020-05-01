var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function (req, res, next) {
  console.log('mark');
  if (!req.session.count) {
    req.session.count = 1;
  } else {
    req.session.count += 1;
  }
  res.render('index', { title: 'Express', count: req.session.count });
});

module.exports = router;
