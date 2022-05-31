package middleware

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"

	"cipher-bot/ent"
	"cipher-bot/ent/user"
)

func Register(client *ent.Client, logger *zap.SugaredLogger) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			result, err := client.User.
				Query().
				Where(user.TelegramID(ctx.Sender().ID)).
				Only(context.Background())
			if err != nil {
				switch err.(type) {
				case *ent.NotFoundError:
					result, err = createUser(client, ctx)
					if err != nil {
						return err
					}

					logger.Infof("registered new user: id - %d, username - %s", result.ID, result.Username)

				default:
					return err
				}
			}

			ctx.Set(user.FieldID, result.ID)

			return next(ctx)
		}
	}
}

func createUser(client *ent.Client, ctx tele.Context) (*ent.User, error) {
	sender := ctx.Sender()

	result, err := client.User.
		Create().
		SetUsername(sender.Username).
		SetTelegramID(sender.ID).
		SetFirstName(sender.FirstName).
		SetLastName(sender.LastName).
		SetLanguage(sender.LanguageCode).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return result, nil
}
