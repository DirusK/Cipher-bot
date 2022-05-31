package handler

import (
	"encoding/hex"
	"fmt"

	"github.com/dgraph-io/badger/v3"
	tele "gopkg.in/telebot.v3"

	"cipher-bot/ent"
	"cipher-bot/ent/predicate"
	"cipher-bot/ent/request"
	"cipher-bot/pkg/cipher"
)

const defaultKeyLength = 32

func (h Handler) OnText(ctx tele.Context) error {
	done, err := h.processManualKey(ctx)
	if err != nil {
		return checkDone(done, err)
	}

	done, err = h.processCipher(ctx)
	if err != nil {
		return checkDone(done, err)
	}

	return ctx.Send(h.layout.Text(ctx, "not-understand"))
}

func (h Handler) processManualKey(ctx tele.Context) (bool, error) {
	requestFilter := []predicate.Request{
		request.UserIDEQ(getUserID(ctx)),
		request.StatusEQ(request.StatusActive),
		request.KeyModeEQ(request.KeyModeManual),
		request.Or(request.ManualKeyValidation(false), request.ManualKeyValidationIsNil()),
	}

	exist, err := h.client.Request.Query().Where(requestFilter...).Exist(h.ctx)
	if err != nil {
		return false, err
	}

	if !exist {
		return false, nil
	}

	key, err := hex.DecodeString(ctx.Message().Text)
	if err != nil {
		h.logger.Debugf("user %s: hex decode key: invalid format key", ctx.Sender().Username)
		return false, ctx.Send(h.layout.Text(ctx, "invalid-hex-key"))
	}

	if len(key) != 32 {
		h.logger.Debugf("user %s: invalid length key", ctx.Sender().Username)
		return false, ctx.Send(h.layout.Text(ctx, "invalid-length-key"))
	}

	if err = h.cache.Update(func(txn *badger.Txn) error {
		err = h.client.Request.
			Update().
			SetManualKeyValidation(true).
			Where(requestFilter...).
			Exec(h.ctx)
		if err != nil {
			return err
		}

		e := badger.NewEntry(getUserIDBytes(ctx), key).WithTTL(h.keyTTL)

		return txn.SetEntry(e)
	}); err != nil {
		return false, err
	}

	return true, nil
}

func (h Handler) processCipher(ctx tele.Context) (bool, error) {
	req, err := h.client.Request.
		Query().
		Where(
			request.UserIDEQ(getUserID(ctx)),
			request.StatusEQ(request.StatusActive),
			request.TypeEQ(request.TypeEncryption),
		).
		Only(h.ctx)
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			return false, nil
		default:
			return false, err
		}
	}

	key, err := h.prepareKeyByMode(ctx, *req.KeyMode)
	if err != nil {
		return false, err
	}

	switch req.Type {
	case request.TypeEncryption:
		cipherText, err := h.encrypt(ctx, *req.Algorithm, key)
		if err != nil {
			return false, err
		}

		result := struct {
			Key        string
			CipherText string
		}{
			Key:        key,
			CipherText: cipherText,
		}

		if err = ctx.Send(h.layout.Text(ctx, "encrypt-result", result)); err != nil {
			return false, err
		}

	case request.TypeDecryption:
		plainText, err := h.decrypt(ctx, *req.Algorithm, key)
		if err != nil {
			return false, err
		}

		if err = ctx.Send(h.layout.Text(ctx, "decrypt-result", plainText)); err != nil {
			return false, err
		}

	default:
		return false, fmt.Errorf("unsupported type: %s", req.Type)
	}

	return true, h.client.Request.
		Update().
		SetStatus(request.StatusDone).
		Where(request.IDEQ(req.ID)).
		Exec(h.ctx)
}

func (h Handler) prepareKeyByMode(ctx tele.Context, mode request.KeyMode) (string, error) {
	var key string

	switch mode {
	case request.KeyModeAuto:
		return cipher.GenerateKey(defaultKeyLength)
	case request.KeyModeManual:
		err := h.cache.Update(func(txn *badger.Txn) error {
			cacheKey := getUserIDBytes(ctx)

			item, err := txn.Get(cacheKey)
			if err != nil {
				return err
			}

			err = item.Value(func(val []byte) error {
				key = hex.EncodeToString(val)
				return nil
			})
			if err != nil {
				return err
			}

			return txn.Delete(cacheKey)
		})
		if err != nil {
			return "", err
		}

		return key, nil

	default:
		return "", nil
	}
}

func (h Handler) encrypt(ctx tele.Context, algorithm request.Algorithm, key string) (string, error) {
	plainText := ctx.Message().Text

	switch algorithm {
	case request.AlgorithmAES:
		return cipher.EncryptAES(plainText, key)
	case request.AlgorithmRC4:
		return cipher.EncryptRC4(plainText, key)
	default:
		return "", fmt.Errorf("unsupported encrypt algorithm: %s", algorithm)
	}
}

func (h Handler) decrypt(ctx tele.Context, algorithm request.Algorithm, key string) (string, error) {
	cipherText := ctx.Message().Text

	switch algorithm {
	case request.AlgorithmAES:
		return cipher.DecryptAES(cipherText, key)
	case request.AlgorithmRC4:
		return cipher.DecryptRC4(cipherText, key)
	default:
		return "", fmt.Errorf("unsupported decrypt algorithm: %s", algorithm)
	}
}
