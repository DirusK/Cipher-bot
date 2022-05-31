package handler

import (
	"strconv"

	tele "gopkg.in/telebot.v3"

	"cipher-bot/ent/request"
	"cipher-bot/ent/user"
)

func getUserID(ctx tele.Context) int {
	userID := ctx.Get(user.FieldID)
	return userID.(int)
}

func getUserIDBytes(ctx tele.Context) []byte {
	userID := ctx.Get(user.FieldID)
	return []byte(strconv.Itoa(userID.(int)))
}

func (h Handler) deleteActiveRequest(ctx tele.Context) error {
	_, err := h.client.Request.
		Delete().
		Where(
			request.StatusEQ(request.StatusActive),
			request.UserIDEQ(getUserID(ctx)),
		).
		Exec(h.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (h Handler) setExpiredActiveRequest(ctx tele.Context) error {
	return h.client.Request.
		Update().
		SetStatus(request.StatusExpired).
		Where(
			request.StatusEQ(request.StatusActive),
			request.UserIDEQ(getUserID(ctx)),
		).
		Exec(h.ctx)
}
