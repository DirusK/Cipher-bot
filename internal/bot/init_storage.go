package bot

import (
	"context"
	"encoding/hex"

	"github.com/dgraph-io/badger/v3"
	_ "github.com/jackc/pgx/stdlib"
	db "gitlab.com/go-insomnia/database/ent"

	"cipher-bot/ent"
	"cipher-bot/pkg/printer"
)

func (b *Bot) initClient() {
	var err error

	b.client = ent.NewClient(ent.Driver(db.NewDriver(b.config.Storage.Postgres)))
	if err = b.client.Schema.Create(context.Background()); err != nil {
		printer.Fatal("ENT", "can't create schema resources", err)
	}

	cipherKey, err := hex.DecodeString(b.config.Storage.Badger.CipherKey)
	if err != nil {
		printer.Fatal("CIPHER-KEY", "can't decode hex key", err)
	}

	b.cache, err = badger.Open(
		badger.DefaultOptions("").
			WithInMemory(true).
			WithEncryptionKey(cipherKey).
			WithIndexCacheSize(100 << 20), // 100 MB
	)
	if err != nil {
		printer.Fatal("BADGER", "can't create schema resources", err)
	}
}
