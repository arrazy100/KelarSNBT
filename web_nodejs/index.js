const express = require('express');

const app = express();
const port = 3000;

const cors = require('cors')

const corsOptions = {
    origin: 'http://localhost:3002',
    optionsSuccessStatus: 200,
    credentials: true
}
app.use(cors(corsOptions))

require('./db/conn');
const auth = require('./auth/routes');

app.use(express.json())
app.use('/auth', cors(corsOptions), auth);

app.listen(port, () => {
    console.log(`App running on port ${port}`);
});