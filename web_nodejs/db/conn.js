const mongoose = require('mongoose');

const mongoString = process.env.MONGODB_CONNSTRING;
mongoose.connect(mongoString);
const db = mongoose.connection;

db.on('error', (err) => {
    console.log(err);
});

db.once('connected', () => {
    console.log('MongoDB connected');
});

module.exports = db;