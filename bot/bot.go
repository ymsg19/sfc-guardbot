package bot

import (
	"errors"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/ymsg19/sfc-guardbot/bot/handlers"
	"github.com/ymsg19/sfc-guardbot/ent"
)

type DiscordBot struct {
	Client *discordgo.Session
	DB     *ent.Client
}

// Returns a created DiscordBot struct.
func NewDiscordBot(db *ent.Client) (*DiscordBot, error) {
	if db == nil {
		return nil, errors.New("passed database client is nil")
	}

	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		return nil, err
	}

	return &DiscordBot{
		Client: discord,
		DB:     db,
	}, nil
}

// Start the bot. Returns a error when it raised.
func (bot *DiscordBot) Start() error {
	handler := handlers.NewHandlers(bot.DB)
	handler.ResigsterAll(bot.Client)

	bot.Client.Identify.Intents = discordgo.IntentsAll

	log.Printf("connecting to discord")
	if err := bot.Client.Open(); err != nil {
		return err
	}
	log.Printf("successfully connected to discord")

	if err := bot.Client.UpdateGameStatus(1, "UNDER_GUARD"); err != nil {
		return err
	}

	return nil
}

// Stop the bot. Returns a error when it raised.
func (bot *DiscordBot) Stop() error {
	return bot.Client.Close()
}
