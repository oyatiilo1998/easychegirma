const User = require("../../models/user");

const UserStorage = {
    get: async (tg_user_id) => {
        let query = {};
        if (tg_user_id){
            query.tg_user_id = tg_user_id;
        }

        let user = await User.findOne(query)

        return user
    },
    getAll: async (params) => {
        let users = await User.find(params)
        return users
    },
    getOrCreate: async (tg_user_id, username) => {
        let query = {};
        if (tg_user_id){
            query.tg_user_id = tg_user_id;
        }

        let user = await User.findOne(query)
        if (user == null) {
            user = new User();
            user.tg_user_id = tg_user_id
            user.username = username
            await user.save();
        }

        if (user.username != username) {
            user.username = username
            user.save()
        }
        
        return user
    },
    changeStep: async (tg_user_id, step) => {
        await User.updateOne({tg_user_id: tg_user_id}, {step: step});
    },
    changeField: async (tg_user_id, fieldName, fieldVariable) => {
        let query = {};
        if (tg_user_id){
            query[fieldName] = fieldVariable;
        }
        await User.updateOne({tg_user_id: tg_user_id}, query);
    },
};

module.exports = UserStorage;
