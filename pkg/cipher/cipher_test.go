package cipher

import (
	"testing"
)

func TestEncryptDecryptAES(t *testing.T) {
	type args struct {
		plainText string
		key       string
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
				key:       "349b03562fe84f8b8e6b5492762d0246f9a2a51c81436cba70e93fd21691bd38",
			},
			wantErr: false,
		},
		{
			name: "[fail] invalid key length",
			args: args{
				plainText: "text",
				key:       "349b03562fed",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encrypted, err := EncryptAES(tt.args.plainText, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptAES() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			decrypted, err := DecryptAES(encrypted, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptAES() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.args.plainText != decrypted {
				t.Errorf("EncryptDecryptAES() got = %v, want %v", decrypted, tt.args.plainText)
			}

			t.Log("Plain text: ", tt.args.plainText)
			t.Log("Cipher text: ", encrypted)
			t.Log("Key: ", tt.args.key)
		})
	}
}

func TestEncryptDecryptRC4(t *testing.T) {
	type args struct {
		plainText string
		key       string
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
				key:       "349b03562fe84f8b8e6b5492762d0246f9a2a51c81436cba70e93fd21691bd38",
			},
			wantErr: false,
		},
		{
			name: "[fail] invalid key length",
			args: args{
				plainText: "text",
				key:       "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encrypted, err := EncryptRC4(tt.args.plainText, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptRC4() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			decrypted, err := DecryptRC4(encrypted, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptRC4() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.args.plainText != decrypted {
				t.Errorf("EncryptDecryptRC4() got = %v, want %v", decrypted, tt.args.plainText)
			}

			t.Log("Plain text: ", tt.args.plainText)
			t.Log("Cipher text: ", encrypted)
			t.Log("Key: ", tt.args.key)
		})
	}
}

func TestGenerate32BytesKey(t *testing.T) {
	got, err := GenerateKey(32)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Generated key hex: %s", got)
}
