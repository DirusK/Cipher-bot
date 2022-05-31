package bot

import (
	tele "gopkg.in/telebot.v3"

	"cipher-bot/ent/request"
)

func (b *Bot) registerCommands() {
	// Main bot commands
	b.telegram.Handle("/start", b.handler.OnStart)
	b.telegram.Handle("/help", b.handler.OnHelp)
	b.telegram.Handle(tele.OnText, b.handler.OnText)

	// Crypto menu
	b.telegram.Handle(b.layout.Callback("encrypt"), b.handler.OnCrypt(request.TypeEncryption))
	b.telegram.Handle(b.layout.Callback("decrypt"), b.handler.OnCrypt(request.TypeDecryption))

	// Algorithm menu
	b.telegram.Handle(b.layout.Callback("algorithm"), b.handler.OnAlgorithm)

	// Key process menu
	b.telegram.Handle(b.layout.Callback("auto"), b.handler.OnAuto)
	b.telegram.Handle(b.layout.Callback("manual"), b.handler.OnManual)
}
