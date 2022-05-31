package middleware

import (
	"regexp"
	"time"

	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

var RegexpText = regexp.MustCompile(`[\p{L}|\s]+`)

const KeySensitive = "sensitive"

func Logger(l *zap.SugaredLogger) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			start := time.Now()

			if err := next(c); err != nil {
				return err
			}

			params := []interface{}{
				"duration", time.Since(start).Seconds(),
				"username", c.Sender().Username,
			}

			if c.Message() != nil && c.Get(KeySensitive) == nil {
				params = append(params, "message", RegexpText.FindString(c.Message().Text))
			}

			if c.Query() != nil {
				params = append(params, "inline query", RegexpText.FindString(c.Query().Text))
			}

			l.Infow("completed request", params...)

			return nil
		}
	}
}
