const config = {
    ENVIRONMENT: getConf("NODE_ENV", "development"),
    MONGO_HOST: getConf("MONGO_HOST", "127.0.0.1"),
    MONGO_PORT: getConf("MONGO_PORT", "27017"),
    MONGO_USER: getConf("MONGO_USER", "root"),
    MONGO_PASSWORD: getConf("MONGO_PASSWORD", "root"),
    MONGO_DATABASE: getConf("MONGO_DATABASE", "chegirma_bot_db"),
    BOT_TOKEN: "5075727984:AAHE5ittFxyyzXAXSMsY0Q9r88p7h4SqSjc", 
    BACKEND_URL: "https://api.chegirma.uz",
    CHANNEL_ID: '-1001678971548',
};
  
function getConf(name, def = "") {
    if (process.env[name]) {
      return process.env[name];
    }
    return def;
  }
  
module.exports = config;
