const { Markup } = require('telegraf')

const keyboards = {
    removeKeyboard: Markup.removeKeyboard(),
    backKeyboard: Markup.inlineKeyboard([
                Markup.button.callback("⬅️ Orqaga", 'back'),
    ]),
    mainMenuKeyboard:  Markup.inlineKeyboard([
        [
            Markup.button.callback("🎁 Mening chegirmalarim", 'my_sales'),
            Markup.button.callback("🔖 Chegirma yaratish", 'create_sale'),
        ]
    ]),
    selectCategoryMenuKeyboard: (categories) => {
        return new Promise((resolve, reject) => {
            let arr = [];
            categories.forEach(category => {
                arr.push([
                    Markup.button.callback(category.name, 'category_id' + category.id)
                ])
            });

            arr.push([
                Markup.button.callback("⬅️ Orqaga", 'back')
            ])
            let keyboard = Markup.inlineKeyboard(arr)
            resolve(keyboard)
        })
    },
    confirmSaleMenuKeyboard: Markup.inlineKeyboard([
        [
            Markup.button.callback("❌ Bekor qilish", 'cancel_sale'),
            Markup.button.callback("✅ Tasdiqlash", 'confirm_sale'),
        ],
        [
            Markup.button.callback("⬅️ Orqaga", 'back'),
        ]
    ]),
}

module.exports = keyboards;