const request = require('request');
const logger = require("../config/logger.js");
const cfg = require("../config");

const utils = {
    authorizeWithBackend(data) {
        return new Promise((resolve, reject) => {
            request({
                // url: cfg.BACKEND_URL + "/v1/user-exists",
                url: cfg.BACKEND_URL + "/v1/user-exists?login="+data.login + "&password="+data.password,
                method: "POST",
                json: true,
            }, function (error, response, body) {
                if (error) {
                    reject(error)
                }
                resolve(body)
            });
        })
    },
    createSaleInBackend(data) {
        return new Promise((resolve, reject) => {
            request({
                url: cfg.BACKEND_URL + "/v1/product",
                method: "POST",
                json: true,
                body: data,
            }, function (error, response, body) {
                if (error) {
                    reject(error)
                }
                resolve(body)
            });
        })
    },
    getCategoriesFromBackend() {
        return new Promise((resolve, reject) => {
            request({
                url: cfg.BACKEND_URL + "/v1/category",
                method: "GET",
                json: true,
            }, function (error, response, body) {
                if (error) {
                    reject(error)
                }
                resolve(body)
            });
        })
    },
    getSalesFromBackend(userID) {
        return new Promise((resolve, reject) => {
            request({
                url: cfg.BACKEND_URL + "/v1/product?owner_id"+userID,
                method: "GET",
                json: true,
            }, function (error, response, body) {
                if (error) {
                    reject(error)
                }
                resolve(body)
            });
        })
    },
    async sendSales(ctx, userID){
        let sales = await this.getSalesFromBackend(userID)
        await sales.products.forEach(async (sale) => {
            let text = '<b>' + sale.name + '</b>' + '\n\n'
                + '<b>' + 'Chegirma' + '</b>' + ': ' + sale.discount_amount + '%\n'
                + '<b>' + 'Asl narxi' + '</b>' + ': ' + this.formatPrice(sale.price_before) + ' so\'m\n'
                + '<b>' + 'Chegirmadagi narxi' + '</b>' + ': ' + this.formatPrice(sale.price_after) + ' so\'m\n'
                + '<b>' + 'Boshanish sanasi' + '</b>' + ': ' + this.formatDate(sale.from_time) + '\n'
                + '<b>' + 'Amal qilish sanasi' + '</b>' + ': ' + this.formatDate(sale.to_time) + '\n'
                + '<b>' + 'Link' + '</b>' + ': ' + sale.url + '\n'
            // await ctx.replyWithHTML(text)

            try {
                if (sale.image_url != "" && sale.image_url != undefined) {
                    await ctx.telegram.sendPhoto(ctx.from.id, 
                        sale.image_url,
                        { caption: text, parse_mode: "html"})
                } else {
                    await ctx.telegram.sendMessage(ctx.from.id, text, {parse_mode: "html"})
                }
            } catch (error){
                console.log(error)
            }
            await utils.sleep(100)
        });
        await utils.sleep(500)
    },
    sleep(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    },
    formatDate(date) {
        if (date == "" || date == undefined) {
            return ""
        }
        var d = new Date(date),
            month = '' + (d.getMonth() + 1),
            day = '' + d.getDate(),
            year = d.getFullYear();
    
        if (month.length < 2) 
            month = '0' + month;
        if (day.length < 2) 
            day = '0' + day;
    
        return [day, month, year].join('.');
    },
    formatPrice(price) {
        let result = ""
        let priceStr = price.toString()
        while (priceStr.length > 0) {
            let subStr = ""

            if (priceStr.length > 3) {
                subStr = priceStr.substr(-3)
                priceStr = priceStr.substr(0, priceStr.length - 3)
            } else {
                subStr = priceStr
                priceStr = ""
            }

            if (result == "") {
                result = subStr
            } else {
                result = subStr + " " + result
            }
        }
    
        return result;
    },
    async sendSaleToChannel(ctx, sale){
       
        let text = '<b>' + sale.name + '</b>' + '\n\n'
            + '<b>' + 'Chegirma' + '</b>' + ': ' + sale.discount_amount + '%\n'
            + '<b>' + 'Asl narxi' + '</b>' + ': ' + this.formatPrice(sale.price_before) + ' so\'m\n'
            + '<b>' + 'Chegirmadagi narxi' + '</b>' + ': ' + this.formatPrice(sale.price_after) + ' so\'m\n'
            + '<b>' + 'Boshanish sanasi' + '</b>' + ': ' + this.formatDate(sale.from_time) + '\n'
            + '<b>' + 'Amal qilish sanasi' + '</b>' + ': ' + this.formatDate(sale.to_time) + '\n'
            + '<b>' + 'Link' + '</b>' + ': ' + sale.url + '\n'
       
        try {
            if (sale.image_url != "" && sale.image_url != undefined) {
                await ctx.telegram.sendPhoto(cfg.CHANNEL_ID, 
                    sale.image_url,
                    { caption: text, parse_mode: "html"})
            } else {
                await ctx.telegram.sendMessage(cfg.CHANNEL_ID, text, {parse_mode: "html"})
            }
        } catch (error){
            console.log(error)
        }
    },

}

module.exports = utils;