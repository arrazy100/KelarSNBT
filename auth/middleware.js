const { verifyToken } = require('./controller');
const jwt = require('jsonwebtoken');

const authMiddleware = (req, res, next) => {
    const token = req.body.token || req.query.token || req.headers['x-access-token'];

    try {   
        const decoded = jwt.verify(token, process.env.SECRET_KEY);
        req.user = decoded;
    }
    catch (err) {
        return res.status(401).send({message: 'Invalid token'});
    }

    return next();
}

module.exports = { authMiddleware };