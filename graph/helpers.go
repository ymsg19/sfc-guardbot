package graph

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

const (
	TOKEN_LENGTH       = 8 // should be a multiple of 4.
	TOKEN_BYTES_LENGTH = TOKEN_LENGTH * 3 / 4
)

type Token [TOKEN_BYTES_LENGTH]byte

// Returns string converted from Token.
func (b *Token) ToString() string {
	return base64.URLEncoding.EncodeToString(b[:])
}

// Returns slice converted from Token.
func (b *Token) ToSlice() []byte {
	return b[:]
}

func generateToken() (*Token, error) {
	count := 0
	for {
		token, err := func() (*Token, error) {
			token := Token{}
			_, err := rand.Read(token.ToSlice())
			if err != nil {
				return nil, err
			}
			return &token, nil
		}()
		if err == nil {
			return token, nil
		}
		if count > 10 {
			return nil, err
		}
		count++
	}
}

func sendTokenByMail(email, token string) error {
	const DOMAIN = "sfc.guardbot.online"

	apikey := os.Getenv("MAILGUN_PRIVATE_API_KEY")
	if apikey == "" {
		return fmt.Errorf("MAILGUN_PRIVATE_API_KEY is not set")
	}

	mg := mailgun.NewMailgun(DOMAIN, apikey)

	sender := "Keio SFC '22 Discord Guild <noreply@" + DOMAIN + ">"
	subject := "[Keio SFC '22 Discord Guild] メールアドレスを認証してください。 - Please verify your email address."
	body := "Keio SFC '22 Discord サーバーの利用を開始するために、以下のコードを入力してメールアドレスを認証してください: \n" + token +
		"\n\n" +
		"To start using Keio SFC '22 Discord Guild, Please enter the following code to verify your email: \n" + token +
		"\n\n" +
		"このメールは自動送信されています。返信することはできません。\n" +
		"This email is automatically sent. Please do not reply."
	recipient := email

	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		return err
	}

	fmt.Printf("mail was sent. ID: %s Resp: %s\n", id, resp)
	return nil
}
