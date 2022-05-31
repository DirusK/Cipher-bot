package handler

import (
	"context"
	"time"

	"github.com/dgraph-io/badger/v3"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"

	"cipher-bot/ent"
)

type (
	Handler struct {
		ctx    context.Context
		client *ent.Client
		cache  *badger.DB
		keyTTL time.Duration
		bot    *tele.Bot
		logger *zap.SugaredLogger
		layout *layout.Layout
	}
)

func New(
	client *ent.Client,
	cache *badger.DB,
	keyTTL time.Duration,
	bot *tele.Bot,
	layout *layout.Layout,
	logger *zap.SugaredLogger,
) Handler {
	return Handler{
		ctx:    context.Background(),
		bot:    bot,
		client: client,
		cache:  cache,
		keyTTL: keyTTL,
		logger: logger,
		layout: layout,
	}
}
