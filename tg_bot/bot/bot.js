const userStorage = require('../storage/mongo/user')
const keyboards = require('./keyboards')
const steps = require('../config/bot_steps')
const utils = require('./utils')
const cfg = require('../config')

class Bot {
    constructor(ctx) {
        this.ctx = ctx
        this.tg_user_id = ctx.from.id
    }
    async init() {
        this.user = await userStorage.getOrCreate(this.tg_user_id, this.ctx.from.username)
        // console.log("user-->", this.user)
    }
    
    async displayEnterLoginMenu() {
        await userStorage.changeStep(this.tg_user_id, steps.ENTER_LOGIN)
        let text
        
        text = "Assalom alaykum!\n"
            + "-----\n" 
            + "Botdan foydalanish uchun loginni kiriting!\n"

        this.ctx.reply(text)
    } 

    async handleEnterLoginMenu(text) {
        console.log("handle login menu: ", text)
        await userStorage.changeField(this.tg_user_id, 'login', text)
        this.displayEnterPasswordMenu()
    }

    async displayEnterPasswordMenu() {
        await userStorage.changeStep(this.tg_user_id, steps.ENTER_PASSWORD)
        let text
        
        text = "Maxfiy kodni kiriting!\n"

        this.ctx.reply(text, keyboards.backKeyboard)
    } 

    async handleEnterPasswordMenu(text) {
        console.log("handle enter password menu: ", text)
        switch (text) {
            case 'back':
                await this.ctx.deleteMessage()
                this.displayEnterLoginMenu()
                break
            default:        
                
                let data = {
                    login: this.user.login,
                    password: text,
                }
                console.log(data)
                let resp = await utils.authorizeWithBackend(data)
                console.log(resp)
                
                if (resp.exist) {
                    await userStorage.changeField(this.tg_user_id, 'is_authorized', true)
                    await userStorage.changeField(this.tg_user_id, 'user_id', resp.id)
                    
                    this.displayMainMenu()
                } else {
                    await this.ctx.reply("Login yoki maxfiy kod noto'g'ri.")

                    this.displayEnterPasswordMenu()
                }
        }
    }

    async displayMainMenu() {
        await userStorage.changeStep(this.tg_user_id, steps.MAIN_MENU)
        
        let text = "Asosiy menu!"

        this.ctx.reply(text, keyboards.mainMenuKeyboard)
    } 

    async handleMainMenu(text) {
        console.log("handle main menu: ", text)
        switch (text) {
            case 'my_sales':
                await this.ctx.deleteMessage()
                await utils.sendSales(this.ctx, this.user_id)
                this.displayMainMenu()
                break
            case 'create_sale':
                await this.ctx.deleteMessage()
                this.displayEnterSaleNameMenu()
                break
        }
    }


    async displayEnterSaleNameMenu() {
        await userStorage.changeStep(this.tg_user_id, steps.ENTER_SALE_NAME)
        let text
        
        text = "Chegirma nomini kiriting"

        this.ctx.reply(text, keyboards.backKeyboard)
    } 

    async handleEnterSaleNameMenu(text) {
        console.log("handle enter sale name menu: ", text)
        switch (text) {
            case 'back':
                await this.ctx.deleteMessage()
                this.displayMainMenu()
                break
            default:        
                let draftSale = {
                    name: text
                }
                await userStorage.changeField(this.tg_user_id, 'draft_sale', draftSale)
                this.displayEnterSaleUrlMenu()
        }
    }

    async displayEnterSaleUrlMenu() {
        await userStorage.changeStep(this.tg_user_id, steps.ENTER_SALE_URL)
        
        let text = "Chegirma uchun link kiriting.\n"

        this.ctx.reply(text, keyboards.backKeyboard)
    } 

    async handleEnterSaleUrlMenu(text) {
        console.log("handle enter sale url menu: ", text)
        switch (text) {
            case 'back':
                await this.ctx.deleteMessage()
                this.displayEnterSaleNameMenu()
                break
            default:    
                let draftSale = this.user.draft_sale    
                draftSale.url = text
                await userStorage.changeField(this.tg_user_id, 'draft_sale', draftSale)
                this.displaySelectCategoryMenu()
        }
    }


    async displaySelectCategoryMenu() {
        await userStorage.changeStep(this.tg_user_id, steps.SELECT_CATEGORY)
        
        let text = "Chegirma kategoriyasini tanlang"
        let resp = await utils.getCategoriesFromBackend()
       
        this.ctx.reply(text, 
            await keyboards.selectCategoryMenuKeyboard(resp.categories))
    } 

    async handleSelectCategoryMenu(text) {
        console.log("handle select category menu: ", text)
        switch (text) {
            case 'back':
                await this.ctx.deleteMessage()
                this.displayEnterSaleUrlMenu()
                break
            default:  
                if (text.slice(0, 11) == 'category_id') {
                    await this.ctx.deleteMessage()  
                    
                    let draftSale = this.user.draft_sale    
                    draftSale.category_id = text.slice(11)
                    await userStorage.changeField(this.tg_user_id, 'draft_sale', draftSale)
                    
                    this.displayEnterDiscountAmountMenu()
                } else {
                    this.displaySelectCategoryMenu()
                }    
        }
    }

    async displayEnterDiscountAmountMenu() {
        await userStorage.changeStep(this.tg_user_id, steps.ENTER_DISCOUNT_AMOUNT)
        
        let text = "Chegirma miqdori yoki foizini kiriting!"

        this.ctx.reply(text, keyboards.backKeyboard)
    } 

    async handleEnterDiscountAmountMenu(text) {
        console.log("handle enter discount amount menu: ", text)
        switch (text) {
            case 'back':
                await this.ctx.deleteMessage()
                this.displaySelectCategoryMenu()
                break
            default:                    
                let draftSale = this.user.draft_sale    
                draftSale.discount_amount = text
                await userStorage.changeField(this.tg_user_id, 'draft_sale', draftSale)
                
                this.displayEnterPriceBeforeMenu()
        }
    }

    async displayEnterPriceBeforeMenu() {
        await userStorage.changeStep(this.tg_user_id, steps.ENTER_PRICE_BEFORE)
        
        let text = "Mahsulotning asl narxini kiriting!"

        this.ctx.reply(text, keyboards.backKeyboard)
    } 

    async handleEnterPriceBeforeMenu(text) {
        console.log("handle enter price before menu: ", text)
        switch (text) {
            case 'back':
                await this.ctx.deleteMessage()
                this.displayEnterDiscountAmountMenu()
                break
            default:        
                
                let draftSale = this.user.draft_sale    
                draftSale.price_before = text
                await userStorage.changeField(this.tg_user_id, 'draft_sale', draftSale)
                
                this.displayEnterPriceAfterMenu()
        }
    }

    async displayEnterPriceAfterMenu() {
        await userStorage.changeStep(this.tg_user_id, steps.ENTER_PRICE_AFTER)
        
        let text = "Mahsuloting chegirmadagi narxini kiriting!"

        this.ctx.reply(text, keyboards.backKeyboard)
    } 

    async handleEnterPriceAfterMenu(text) {
        console.log("handle enter price after menu: ", text)
        switch (text) {
            case 'back':
                await this.ctx.deleteMessage()
                this.displayEnterPriceBeforeMenu()
                break
            default:        
                let draftSale = this.user.draft_sale    
                draftSale.price_after = text
                await userStorage.changeField(this.tg_user_id, 'draft_sale', draftSale)
            
                this.displayConfirmSaleMenu()
        }
    }

    async displayConfirmSaleMenu() {
        await userStorage.changeStep(this.tg_user_id, steps.CONFIRM_SALE)
        let text
        
        text = "Chegirmani tasdiqlang!"

        this.ctx.reply(text, keyboards.confirmSaleMenuKeyboard)
    } 

    async handleConfirmSaleMenu(text) {
        console.log("handle confirm sale menu: ", text)
        switch (text) {
            case 'back':
                await this.ctx.deleteMessage()
                this.displayEnterPriceAfterMenu()
                break
            case 'confirm_sale':
                let draftSale = this.user.draft_sale
                let data = {
                    name: draftSale.name,
                    url: draftSale.url,
                    category_id: draftSale.category_id,
                    discount_amount: parseFloat(draftSale.discount_amount),
                    // description: ,
                    // image_url: ,
                    // from_time: ,
                    // to_time: ,
                    owner_id: this.user.user_id,
                    price_after: parseFloat(draftSale.price_after),
                    price_before: parseFloat(draftSale.price_before),
                }
                let res = await utils.createSaleInBackend(data)
                console.log(res)
                if(res.id != "" && res.id != undefined) {
                    await this.ctx.deleteMessage()
                    
                    let text = "Sizning chegirmangiz muvaffaqiyatli yaratildi"
                    await this.ctx.reply(text)

                    utils.sendSaleToChannel(this.ctx, data)
                    this.displayMainMenu()
                    
                } else {
                    await this.ctx.deleteMessage()
                    await this.ctx.reply("Xatolik sodir bo'ldi, qayta urinib ko'ring.")
                    this.displayConfirmSaleMenu()
                }
                break
        }
    }
}

module.exports = Bot;
