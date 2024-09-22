const { Telegraf } = require('telegraf');
require('dotenv').config(); // Load variables from .env file

const BOT_TOKEN = process.env.BOT_TOKEN;
const WEB_APP_URL = process.env.WEB_APP_URL;

const bot = new Telegraf(BOT_TOKEN);

bot.start((ctx) => ctx.reply('Welcome to MeraMoney!', {
    reply_markup: {
        keyboard: [
            [{ text: 'Start', web_app: { url: WEB_APP_URL } }],
        ]
    }
}));

bot.launch();