use serenity::all::{Context, EventHandler, Interaction, Ready};
use serenity::async_trait;
use serenity::builder::{CreateCommand, CreateCommandOption as CreateOption};
use serenity::model::application::{Command, CommandOptionType};

struct CommandHandler;

#[async_trait]
impl EventHandler for CommandHandler {
    async fn interaction_create(&self, ctx: Context, interaction: Interaction) {
        if let Interaction::Command(command) = interaction {
            let name = command.data.name.as_str();
            let content = match name {
                "ping" => Some(&self::ping::run(&command.data.options))
            };
        }
    }

    async fn ready(&self, ctx: Context, ready: Ready) {

    }
}


