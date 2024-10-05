const { Telegraf } = require('telegraf');

const BOT_TOKEN = '7773859219:AAH7Duvqaj5rM4_kGI58mel8ZLtXfLiMG74';

const bot = new Telegraf(BOT_TOKEN);


bot.start((ctx) => ctx.reply('Welcome to MeraMoney!', {
    reply_markup: {
        inline_keyboard: [
            [ {text:'Open app', web_app: {url: 'https://finalcs50-meramoney.vercel.app/'}} ],
        ]}
}));

bot.launch();