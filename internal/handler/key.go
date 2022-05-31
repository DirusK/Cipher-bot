package handler

import (
	tele "gopkg.in/telebot.v3"

	"cipher-bot/ent/request"
)

func (h Handler) OnManual(ctx tele.Context) error {
	err := h.client.Request.
		Update().
		SetKeyMode(request.KeyModeManual).
		Where(
			request.UserIDEQ(getUserID(ctx)),
			request.StatusEQ(request.StatusActive),
		).
		Exec(h.ctx)
	if err != nil {
		return err
	}

	return ctx.Edit(
		h.layout.Text(ctx, "cipher-key-request"),
	)
}

func (h Handler) OnAuto(ctx tele.Context) error {
	err := h.client.Request.
		Update().
		SetKeyMode(request.KeyModeAuto).
		Where(
			request.UserIDEQ(getUserID(ctx)),
			request.StatusEQ(request.StatusActive),
		).
		Exec(h.ctx)
	if err != nil {
		return err
	}

	return ctx.Edit(
		h.layout.Text(ctx, "cipher-text-request"),
	)
}
