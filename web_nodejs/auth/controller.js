const bcrypt = require('bcrypt');
const jwt = require('jsonwebtoken');
const fs = require('fs');

const saltRounds = parseInt(process.env.SALT_ROUNDS);

const User = require('./model');

const createUser = async (email, username, password) => {
    const salt = bcrypt.genSaltSync(saltRounds);
    const hashed = await bcrypt.hash(password, salt);

    const user = new User({
        username: username,
        password: hashed,
        email: email
    });

    const dataToSave = await user.save();

    return dataToSave;
};

const getUsers = async () => {
    const users = await User.find().select('email username roles');

    return users;
}

const login = async (username, password) => {
    const user = await User.findOne({
        username: username
    });

    return bcrypt.compare(password, user.password).then((result) => {
        if (result) {
            const data = {
                'username': user.username,
                'email': user.email,
                'roles': user.roles
            };

            const token = generateToken(data);

            return token;
        }
    }).catch(err => console.log(err.message));
}

const generateToken = (data) => {
    const privateKey = fs.readFileSync(process.env.PRIVATE_KEY);

    const token = jwt.sign({
        data
    }, privateKey, { expiresIn: parseInt(process.env.TOKEN_EXPIRES), algorithm: process.env.JWT_ALGORITHM });

    return token;
}

const verifyToken = (token) => {
    const publicKey = fs.readFileSync(process.env.PUBLIC_KEY);

    const decoded = jwt.verify(token, publicKey, { algorithms: [process.env.JWT_ALGORITHM] });

    return decoded;
}

const getRoles = async (username) => {
    const user = await User.findOne({
       username: username
    });

    return user.roles;
}

const setRoles = async (username, roles) => {
    const user = await User.findOne({
        username: username
    });

    user.roles = [];
    for (var i in roles) {
        user.roles.push(roles[i]);
    }

    const dataToSave = await user.save();

    return dataToSave;
}

module.exports = { createUser, getUsers, login, getRoles, setRoles, verifyToken };