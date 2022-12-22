const { sequelize, Sequelize } = require('../models');

const bcrypt = require('bcrypt');
const { where } = require('sequelize');
const saltRounds = 10;

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
    const users = await User.findAll();

    return users;
}

const login = async (username, password) => {
    const user = await User.findOne({
        where: {
            username: username
        }
    });

    const sukses = bcrypt.compareSync(password, user.password);

    return sukses;
}

module.exports = { createUser, getUsers, login };