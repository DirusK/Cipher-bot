package handler

import tele "gopkg.in/telebot.v3"

func (h Handler) OnHelp(ctx tele.Context) error {
	return ctx.Send(
		h.layout.Text(ctx, "help"),
	)
}
