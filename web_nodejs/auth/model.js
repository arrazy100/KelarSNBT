const mongoose = require('mongoose');

const userSchema = new mongoose.Schema({
    name: {
        first: String,
        last: String
    },
    email: {
        type: String,
        required: true
    },
    username: {
        type: String,
        required: true,
        unique: true
    },
    password: {
        type: String,
        required: true
    },
    roles: [String]
});

const User = mongoose.model('User', userSchema);

module.exports = User;