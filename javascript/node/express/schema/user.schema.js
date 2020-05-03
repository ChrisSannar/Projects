var mongoose = require('mongoose');
var Schema = mongoose.Schema;
var bcrypt = require('bcrypt');
const SALT_WORK_FACTOR = 10;

var UserSchema = Schema({
  username: { type: String, required: true, index: { unique: true } },
  password: { type: String, required: true },
});

// When we change the password
UserSchema.pre('findOneAndUpdate', function (next) {
  var user = this;

  // if (!user.isModified('password')) return next();

  // Get some salt
  bcrypt.genSalt(SALT_WORK_FACTOR, function (err, salt) {
    if (err) throw err;

    // hash the password
    bcrypt.hash(user._update.password, salt, function (err, hash) {
      if (err) throw err;

      user._update.password = hash;
      next();
    });
  });
});

// Compares the the incoming password with the hashed one in the database
UserSchema.methods.comparePassword = function (candidate, callback) {
  bcrypt.compare(candidate, this.password, function (err, isMatch) {
    if (err) return callback(err);
    callback(null, isMatch);
  });
}

module.exports = mongoose.model('User', UserSchema, 'users');