package bot

import (
	"os"

	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"

	"cipher-bot/internal/middleware"
	"cipher-bot/pkg/printer"
)

const (
	// layoutDefaultPath - default path for layout config.
	layoutDefaultPath = "configs/bot.yaml"

	// envBotToken - environment variable for bot token.
	envBotToken = "CIPHER-BOT"

	// errorMessage - internal error message for user.
	errorMessage = `Извините, у меня произошла ошибка ❌
Повторите команду ещё раз или напишите в поддержку @insomniaJp.
`
)

func (b *Bot) initTelebot() {
	var err error
	b.layout, err = layout.New(layoutDefaultPath)
	if err != nil {
		printer.Fatal("TELEBOT", "cannot init layout", err)
	}

	settings := b.layout.Settings()
	if settings.Token == "" {
		settings.Token = os.Getenv(envBotToken)
		if settings.Token == "" {
			printer.Print("TELEBOT", printer.Red("API-token for telegram is not provided"))
			os.Exit(1)
		}
	}

	b.telegram, err = tele.NewBot(settings)
	if err != nil {
		printer.Fatal("TELEBOT", "cannot init telegram", err)
	}

	b.telegram.OnError = func(err error, ctx tele.Context) {
		b.logger.Errorf(
			"telegram username: %s, command: %s, description: %v",
			ctx.Sender().Username,
			middleware.RegexpText.FindString(ctx.Message().Text),
			err,
		)

		if err = ctx.Send(errorMessage); err != nil {
			b.logger.Errorf("send error text for user: %s", err)
		}
	}
}
