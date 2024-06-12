mod commands;

use std::env;
use std::path::Path;
use std::io;

use dotenv::from_path;

use serenity::async_trait;
use serenity::model::channel::Message;
use serenity::prelude::*;
struct Handler;

#[async_trait]
impl EventHandler for Handler {
    async fn message(&self, ctx: Context, msg: Message) {
        if msg.channel_id.to_string() != "1019465229651955752" {
            return;
        }
    }
}


#[tokio::main]
async fn main() {
    println!("Path to .env: ");
    let mut env_path: String = String::new();
    io::stdin().read_line(&mut env_path).expect("Failed to open file");
    from_path(Path::new(env_path.replace("\\", "/").trim())).ok();
    // Login with a bot token from the environment
    let token = env::var("DISCORD_TOKEN").expect("Expected a token in the environment");
    // Set gateway intents, which decides what events the bot will be notified about
    let intents = GatewayIntents::GUILD_MESSAGES
        | GatewayIntents::DIRECT_MESSAGES
        | GatewayIntents::MESSAGE_CONTENT;

    // Create a new instance of the Client, logging in as a bot.
    let mut client =
       Client::builder(&token, intents).event_handler(Handler).await.expect("Err creating client");
    // Start listening for events by starting a single shard
    if let Err(why) = client.start().await {
        println!("Client error: {why:?}");
    }
}