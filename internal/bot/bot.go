package bot

import (
	"context"

	"github.com/dgraph-io/badger/v3"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"

	"cipher-bot/ent"
	"cipher-bot/internal/config"
	"cipher-bot/internal/handler"
)

type (
	// Bot main application structure.
	Bot struct {
		telegram *tele.Bot
		client   *ent.Client
		cache    *badger.DB
		config   *config.Config
		logger   *zap.SugaredLogger
		layout   *layout.Layout
		handler  handler.Handler
	}
)

func New(configPath string) *Bot {
	bot := new(Bot)
	bot.config = config.New(configPath)

	bot.initLogger()
	bot.initTelebot()
	bot.initClient()

	bot.handler = handler.New(
		bot.client,
		bot.cache,
		bot.config.Storage.Badger.KeyTTL,
		bot.telegram,
		bot.layout,
		bot.logger,
	)

	return bot
}

func (b *Bot) Run(ctx context.Context) {
	b.runWorkers(ctx)
}
