package client

import (
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

type ImapClient struct {
	cImap *client.Client
}

func (f ImapClient) LoginImap(e string, p string) {
	c, err := client.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		panic("Erro iniciar client Imap")
	}
	if err := c.Login(e, p); err != nil {
		panic("Erro Login imap" + err.Error())
	}
	f.cImap = c
}

func (f ImapClient) GetEmails(e string, p string) {

	// List mailboxes
	done := make(chan error, 1)

	mbox, err := f.cImap.Select("[Gmail]/E-mails enviados", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Flags for INBOX:", mbox.Flags)

	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 3 {
		
		from = mbox.Messages - 3
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	done = make(chan error, 1)
	go func() {
		done <- f.cImap.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	log.Println("Last 4 messages:")
	for msg := range messages {
		log.Println("* " + msg.Envelope.Subject)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	log.Println("Done!")
}

func (f ImapClient) ListEmailBox() {
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- f.cImap.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}
}
