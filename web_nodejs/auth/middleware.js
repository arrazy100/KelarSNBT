const { verifyToken } = require('./controller');

const authMiddleware = (req, res, next) => {
    const token = req.body.token || req.query.token || req.headers['x-access-token'];

    try {   
        const decoded = verifyToken(token);
        req.user = decoded;
    }
    catch (err) {
        return res.status(401).send({message: err.message});
    }

    return next();
}

const authAdminMiddleware = (req, res, next) => {
    const token = req.body.token || req.query.token || req.headers['x-access-token'];

    try {   
        const decoded = verifyToken(token);
        req.user = decoded;
    }
    catch (err) {
        return res.status(401).send({message: err.message});
    }

    if (!req.user.data.roles.includes('admin')) {
        return res.status(403).send({message: 'You don`t have permission'});
    }

    return next();
}

module.exports = { authMiddleware, authAdminMiddleware };