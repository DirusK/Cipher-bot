package cipher

import (
	"encoding/hex"
	"testing"
)

func TestEncryptDecryptAES(t *testing.T) {
	type args struct {
		plainText string
		keyLength int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] common flow",
			args: args{
				plainText: "text",
				keyLength: 32,
			},
			wantErr: false,
		},
		{
			name: "[fail] invalid key length",
			args: args{
				plainText: "text",
				keyLength: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GenerateKeyBytes(tt.args.keyLength)
			if err != nil {
				t.Errorf("GenerateKeyBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			encrypted, err := EncryptAES(tt.args.plainText, key)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptAES() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			decrypted, err := DecryptAES(encrypted, key)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptAES() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.args.plainText != decrypted {
				t.Errorf("EncryptDecryptAES() got = %v, want %v", decrypted, tt.args.plainText)
			}

			t.Log("Plain text: ", tt.args.plainText)
			t.Log("Cipher text: ", encrypted)
			t.Log("Key: ", hex.EncodeToString(key))
		})
	}
}

func TestEncryptDecryptRC4(t *testing.T) {
	type args struct {
		plainText string
		keyLength int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] common flow",
			args: args{
				plainText: "text",
				keyLength: 32,
			},
			wantErr: false,
		},
		{
			name: "[fail] invalid key length",
			args: args{
				plainText: "text",
				keyLength: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GenerateKeyBytes(tt.args.keyLength)
			if err != nil {
				t.Errorf("GenerateKeyBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			encrypted, err := EncryptRC4(tt.args.plainText, key)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptRC4() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			decrypted, err := DecryptRC4(encrypted, key)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptRC4() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.args.plainText != decrypted {
				t.Errorf("EncryptDecryptRC4() got = %v, want %v", decrypted, tt.args.plainText)
			}

			t.Log("Plain text: ", tt.args.plainText)
			t.Log("Cipher text: ", encrypted)
			t.Log("Key: ", hex.EncodeToString(key))
		})
	}
}

func TestGenerate32BytesKeyHex(t *testing.T) {
	got, err := GenerateKeyHex(32)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Generated key hex: %s", got)
}
