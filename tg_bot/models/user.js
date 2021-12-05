const mongoose = require('mongoose')
const steps = require('../config/bot_steps')

const User = mongoose.model('User', {
    tg_user_id: {
        type: String,
        required: true,
    },
    username: {
        type: String,
    },
    login: {
        type: String,
    },
    is_authorized: {
        type: Boolean,
        default: false,
    },
    user_id: {
        type: String,
    },
    step: {
        type: String,
        required: true,
        default: steps.ENTER_LOGIN,
        enum: [
            steps.ENTER_LOGIN,
            steps.ENTER_PASSWORD,
            steps.MAIN_MENU,
            steps.ENTER_SALE_NAME,
            steps.ENTER_SALE_URL,
            steps.SELECT_CATEGORY,
            steps.ENTER_DISCOUNT_AMOUNT,
            steps.ENTER_PRICE_BEFORE,
            steps.ENTER_PRICE_AFTER,
            steps.UPLOAD_PHOTO,
            steps.CONFIRM_SALE,
        ],
    },
    created_at: {
        type: Date,
        default: Date.now,
    },
    draft_sale: {},
})

module.exports = User