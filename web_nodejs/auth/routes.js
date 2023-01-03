const express = require('express');
const { getUsers, createUser, login, setRoles } = require('./controller');
const { authMiddleware, authAdminMiddleware } = require('./middleware');

const router = express.Router();

router.get('/lists', authAdminMiddleware, (req, res) => {
    getUsers().then(users => {
        res.status(200).json(users);
    }).catch(err => {
        res.status(500).json({message: err.message});
    });
});

router.post('/create', authAdminMiddleware, (req, res) => {
    const { email, username, password } = req.body;

    createUser(email, username, password).then(user => {
        res.status(200).json(user);
    }).catch(err => {
        res.status(500).json({message: err.message});
    });
});

router.post('/login', async (req, res) => {
    const { username, password } = req.body;

    login(username, password).then(token => {
        if (token) {
            res.cookie("accessToken", token, {
                secure: true,
                httpOnly: true,
                maxAge: 900000
            })
            res.status(200).json({accessToken: token});
        }
        else res.status(400).json({message: 'username atau password salah'});
    }).catch(err => {
        res.status(500).json({message: err.message});
    });
});

router.post('/setRoles', authAdminMiddleware, async(req, res) => {
    const { username, roles } = req.body;

    setRoles(username, roles).then(user => {
        res.status(200).json(user);
    }).catch(err => {
        res.status(500).json({message: err.message});
    });
});

module.exports = router;