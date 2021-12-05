const Bot = require("./bot")
const steps = require("../config/bot_steps");
const logger = require("../config/logger.js");
const utils = require("./utils");

const handlers = {
    handleStartCommand: async (ctx) => {
        try {
            let bot = new Bot(ctx)
            await bot.init(ctx.from)

            // let data = {
            //     name: "test",
            //     url: "test",
            //     category_id: "test",
            //     discount_amount: parseFloat(31),
            //     image_url: "https://store.storeimages.cdn-apple.com/4668/as-images.apple.com/is/iphone-12-family-select-2021?wid=940&hei=1112&fmt=jpeg&qlt=80&.v=1617135051000",
            //     // description: ,
            //     // image_url: ,
            //     // from_time: ,
            //     // to_time: ,
              
            //     price_after: parseFloat(40000),
            //     price_before: parseFloat(30000),
            // }
            // utils.sendSaleToChannel(ctx, data)
            if (bot.user.is_authorized) {
                bot.displayMainMenu() 
            } else {
                bot.displayEnterLoginMenu()
            }
        }
        catch (e) {
            logger.error("Error while handling start command: " + e)

            ctx.reply("Xatolik sodir bo'ldi, botni qayta ishga tushuring!\n"
                + "-----\n" 
                + "Произошла ошибка, перезапустите бота!"
                + "\n\n👉 /start")
        }
    },

    handleTextMessage: async (ctx) => {
        try {
            let bot = new Bot(ctx)
            await bot.init()
        
            text = ctx.message.text
            switch(bot.user.step) {
                case steps.ENTER_LOGIN:
                    bot.handleEnterLoginMenu(text)
                    break;
                case steps.ENTER_PASSWORD:
                    bot.handleEnterPasswordMenu(text)
                    break;
                case steps.MAIN_MENU:
                    bot.displayMainMenu()
                    break;
                case steps.ENTER_SALE_NAME:
                    bot.handleEnterSaleNameMenu(text)
                    break;
                case steps.ENTER_SALE_URL:
                    bot.handleEnterSaleUrlMenu(text)
                    break;
                case steps.SELECT_CATEGORY:
                    bot.displaySelectCategoryMenu()
                    break;
                case steps.ENTER_DISCOUNT_AMOUNT:
                    bot.handleEnterDiscountAmountMenu(text)
                    break;
                case steps.ENTER_PRICE_BEFORE:
                    bot.handleEnterPriceBeforeMenu(text)
                    break;
                case steps.ENTER_PRICE_AFTER:
                    bot.handleEnterPriceAfterMenu(text)
                    break;
                case steps.CONFIRM_SALE:
                    bot.displayConfirmSaleMenu()
                    break;
                default:
                    bot.displayEnterLoginMenu()
            }
        } catch(e) {
            logger.error("Error while handling text message: " + e)

            ctx.reply("Xatolik sodir bo'ldi, botni qayta ishga tushuring!\n"
                + "-----\n" 
                + "Произошла ошибка, перезапустите бота!"
                + "\n\n👉 /start")
        }
    },
    handleInlineQuery: async (ctx) => {
        try {
            ctx.answerCbQuery()
            
            let text = ctx.update.callback_query.data
        
            let bot = new Bot(ctx)
            await bot.init()
            
            switch(bot.user.step) {
                case steps.ENTER_PASSWORD:
                    bot.handleEnterPasswordMenu(text)
                    break;
                case steps.MAIN_MENU:
                    bot.handleMainMenu(text)
                    break;
                case steps.ENTER_SALE_NAME:
                    bot.handleEnterSaleNameMenu(text)
                    break;
                case steps.ENTER_SALE_URL:
                    bot.handleEnterSaleUrlMenu(text)
                    break;
                case steps.SELECT_CATEGORY:
                    bot.handleSelectCategoryMenu(text)
                    break;
                case steps.ENTER_DISCOUNT_AMOUNT:
                    bot.handleEnterDiscountAmountMenu(text)
                    break;
                case steps.ENTER_PRICE_BEFORE:
                    bot.handleEnterPriceBeforeMenu(text)
                    break;
                case steps.ENTER_PRICE_AFTER:
                    bot.handleEnterPriceAfterMenu(text)
                    break;
                case steps.CONFIRM_SALE:
                    bot.handleConfirmSaleMenu(text)
                    break;
                default:
                    bot.displayEnterLoginMenu()
            }
        } catch(e) {
            logger.error("Error while handling inline message: " + e)

            ctx.reply("Xatolik sodir bo'ldi, botni qayta ishga tushuring!\n"
                + "-----\n" 
                + "Произошла ошибка, перезапустите бота!"
                + "\n\n👉 /start")
        }
    },
}

module.exports = handlers;
