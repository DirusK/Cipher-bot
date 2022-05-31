package handler

import (
	tele "gopkg.in/telebot.v3"

	"cipher-bot/ent/request"
)

func (h Handler) OnAlgorithm(ctx tele.Context) error {
	err := h.client.Request.
		Update().
		SetAlgorithm(request.Algorithm(ctx.Callback().Data)).
		Exec(h.ctx)
	if err != nil {
		return err
	}

	return ctx.Edit(
		h.layout.Text(ctx, "key-request"),
		h.layout.Markup(ctx, "key"),
	)
}
