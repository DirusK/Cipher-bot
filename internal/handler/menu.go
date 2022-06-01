package handler

import (
	tele "gopkg.in/telebot.v3"

	"cipher-bot/pkg/cipher"
)

func (h Handler) OnStart(ctx tele.Context) error {
	return ctx.Send(
		h.layout.Text(ctx, "start"),
		h.layout.Markup(ctx, "menu"),
		tele.NoPreview,
		tele.OneTimeKeyboard,
	)
}

func (h Handler) OnMenu(ctx tele.Context) error {
	return ctx.Send(
		h.layout.Text(ctx, "menu"),
		h.layout.Markup(ctx, "menu"),
		tele.NoPreview,
		tele.OneTimeKeyboard,
	)
}

func (h Handler) OnHelp(ctx tele.Context) error {
	return ctx.Send(
		h.layout.Text(ctx, "help"),
	)
}

func (h Handler) OnGenerate(ctx tele.Context) error {
	key, err := cipher.GenerateKeyHex(32)
	if err != nil {
		return err
	}

	return ctx.Send(
		h.layout.Text(ctx, "generated-key", key),
	)
}
