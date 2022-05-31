package handler

import (
	tele "gopkg.in/telebot.v3"
)

func (h Handler) OnStart(ctx tele.Context) error {
	return ctx.Send(
		h.layout.Text(ctx, "start"),
		h.layout.Markup(ctx, "menu"),
		tele.NoPreview,
		tele.OneTimeKeyboard,
	)
}
