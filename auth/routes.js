const express = require('express');
const { getUsers, createUser, login } = require('./methods');

const router = express.Router();

router.get('/lists', (req, res) => {
    getUsers().then(users => {
        res.status(200).json(users);
    }).catch(err => {
        res.status(500).json({message: err.message});
    });
});

router.post('/create', (req, res) => {
    const { email, username, password } = req.body;

    createUser(email, username, password).then(user => {
        res.status(200).json(user);
    }).catch(err => {
        res.status(500).json({message: err.message});
    });
});

router.post('/login', async (req, res) => {
    const { username, password } = req.body;

    const sukses = await login(username, password);

    if (sukses) res.status(200).json({message: 'Login berhasil!'});
    else res.status(500).json({message: 'Username atau password salah'});
})

module.exports = router;