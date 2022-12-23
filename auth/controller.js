const { sequelize, Sequelize } = require('../models');
const bcrypt = require('bcrypt');
const jwt = require('jsonwebtoken');

const saltRounds = process.env.SALT_ROUNDS;

const User = require('../models/user')(sequelize, Sequelize.DataTypes);

const createUser = async (email, username, password) => {
    const salt = bcrypt.genSaltSync(saltRounds);
    const hashed = await bcrypt.hash(password, salt);

    const user = await User.create({
        username: username,
        password: hashed,
        email: email
    });

    return user;
};

const getUsers = async () => {
    const users = await User.findAll({
        attributes: ['firstName', 'lastName', 'email', 'username']
    });

    return users;
}

const login = async (username, password) => {
    const user = await User.findOne({
        where: {
            username: username
        },
    });

    return bcrypt.compare(password, user.password).then((result) => {
        if (result) {
            const data = {
                'username': user.username,
                'email': user.email
            };

            const token = generateToken(data);

            return token;
        }
    }).catch(err => console.log(err.message));
}

const generateToken = (data) => {
    const token = jwt.sign({
        data: data
    }, process.env.SECRET_KEY, { expiresIn: '1h' });

    return token;
}

module.exports = { createUser, getUsers, login};