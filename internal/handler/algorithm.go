package handler

import (
	tele "gopkg.in/telebot.v3"

	"cipher-bot/ent/request"
)

func (h Handler) OnAlgorithm(ctx tele.Context) error {
	req, err := h.client.Request.
		Query().
		Where(
			request.UserIDEQ(getUserID(ctx)),
			request.StatusEQ(request.StatusActive),
		).
		Only(h.ctx)
	if err != nil {
		return err
	}

	err = h.client.Request.
		Update().
		SetAlgorithm(request.Algorithm(ctx.Callback().Data)).
		Where(request.IDEQ(req.ID)).
		Exec(h.ctx)
	if err != nil {
		return err
	}

	switch req.Type {
	case request.TypeDecryption:
		return ctx.Edit(
			h.layout.Text(ctx, "cipher-key-request"),
		)
	case request.TypeEncryption:
		return ctx.Edit(
			h.layout.Text(ctx, "key-request"),
			h.layout.Markup(ctx, "key"),
		)
	}

	return nil
}
