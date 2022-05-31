package handler

import (
	tele "gopkg.in/telebot.v3"

	"cipher-bot/ent/request"
)

func (h Handler) OnEncrypt(ctx tele.Context) error {
	var err error
	if err = h.deleteActiveRequest(ctx); err != nil {
		return err
	}

	err = h.client.Request.
		Create().
		SetType(request.TypeEncryption).
		SetStatus(request.StatusActive).
		SetUserID(getUserID(ctx)).
		Exec(h.ctx)
	if err != nil {
		return err
	}

	return ctx.Send(
		h.layout.Text(ctx, "cipher-algorithms"),
		h.layout.Markup(ctx, "algorithms"))
}

func (h Handler) OnDecrypt(ctx tele.Context) error {
	var err error
	if err = h.deleteActiveRequest(ctx); err != nil {
		return err
	}

	err = h.client.Request.
		Create().
		SetType(request.TypeDecryption).
		SetStatus(request.StatusActive).
		SetKeyMode(request.KeyModeManual).
		SetUserID(getUserID(ctx)).
		Exec(h.ctx)
	if err != nil {
		return err
	}

	return ctx.Send(
		h.layout.Text(ctx, "cipher-algorithms"),
		h.layout.Markup(ctx, "algorithms"))
}
