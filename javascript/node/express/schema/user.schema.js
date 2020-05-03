var mongoose = require('mongoose');
var Schema = mongoose.Schema;
var bcrypt = require('bcrypt');
const SALT_WORK_FACTOR = 10;

var UserSchema = Schema({
  username: { type: String, required: true, index: { unique: true } },
  password: { type: String, required: true },
});

// What we do before we change the password
// I'll probably need to abstract this a bit more in the case we update another 
//   user preferene and don't care about the password for that action.
UserSchema.pre('findOneAndUpdate', function (next) {
  var query = this;

  // Not sure what this is for exactly because, because the 'user' is actually a query...
  //   However it was included in the example on mongodb for how to properly authenticate
  // if (!user.isModified('password')) return next();

  // Get some salt
  bcrypt.genSalt(SALT_WORK_FACTOR, function (err, salt) {
    if (err) throw err;

    // Hash the password
    bcrypt.hash(query._update.password, salt, function (err, hash) {
      if (err) throw err;

      // Save it in before we attempt to store it
      query._update.password = hash;
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