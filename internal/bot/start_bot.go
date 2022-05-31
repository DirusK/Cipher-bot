package bot

import (
	"context"

	telemw "gopkg.in/telebot.v3/middleware"

	"cipher-bot/internal/middleware"
)

func startBot(ctx context.Context, bot *Bot) {
	bot.telegram.Use(middleware.Logger(bot.logger))
	bot.telegram.Use(bot.layout.Middleware("ru"))
	bot.telegram.Use(middleware.Register(bot.client, bot.logger))
	bot.telegram.Use(telemw.AutoRespond())

	bot.registerCommands()

	// graceful shutdown telegram.
	go func() {
		<-ctx.Done()
		bot.telegram.Stop()
	}()

	bot.telegram.Start()
}
