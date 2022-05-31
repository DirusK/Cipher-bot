package handler

import (
	tele "gopkg.in/telebot.v3"

	"cipher-bot/ent/request"
)

func (h Handler) OnCrypt(reqType request.Type) func(ctx tele.Context) error {
	return func(ctx tele.Context) error {
		_, err := h.client.Request.
			Delete().
			Where(request.StatusEQ(request.StatusActive)).
			Exec(h.ctx)
		if err != nil {
			return err
		}

		err = h.client.Request.
			Create().
			SetType(reqType).
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
}
