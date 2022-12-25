const express = require('express');

const app = express();
const port = 3000;

const db = require('./db/conn');
const auth = require('./auth/routes');

app.use(express.json())
app.use('/auth', auth);

app.listen(port, () => {
    console.log(`App running on port ${port}`);
});