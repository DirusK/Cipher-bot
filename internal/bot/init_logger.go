package bot

import (
	"go.uber.org/zap"

	"cipher-bot/pkg/printer"
)

func (b *Bot) initLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		printer.Fatal("LOGGER", "can't initialize zap logger", err)
	}

	defer func() {
		_ = logger.Sync()
	}()

	b.logger = logger.Sugar()
}
