package handler

import (
	"strconv"

	tele "gopkg.in/telebot.v3"

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

func checkDone(done bool, err error) error {
	switch {
	case done:
		return nil
	default:
		return err
	}
}
