package mailer

import "github.com/wneessen/go-mail"

type Mailer struct {
	mailerClient *mail.Client
}

func New(host string) *Mailer {
	client, _ := mail.NewClient(host)
	return &Mailer{
		mailerClient: client,
	}
}
