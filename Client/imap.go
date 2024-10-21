package client

import (
	"bytes"
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

type ImapClient struct {
	cImap *client.Client
}

func (f *ImapClient) LoginImap(e string, p string) {
	c, err := client.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		panic("Erro iniciar client Imap")
	}
	if err := c.Login(e, p); err != nil {
		panic("Erro Login imap" + err.Error())
	}
	f.cImap = c
}

type ReturnData struct {
	Subject string
	To      string
	From    string
	Body    string
}

func (f *ImapClient) GetEmails(e string, p string) []ReturnData {

	// List mailboxes
	done := make(chan error, 1)

	mbox, err := f.cImap.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}

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

	var data []ReturnData
	i := 0

    for msg := range messages {
        // Verificar se msg.Envelope ou Body é nulo
        if msg == nil || msg.Envelope == nil {
            continue
        }

        var body string
        if msg.Body != nil {
            bodyReader := msg.GetBody(&imap.BodySectionName{})
            if bodyReader != nil {
                buf := new(bytes.Buffer)
                buf.ReadFrom(bodyReader)
                body = buf.String()
            }
        }

        // Verificar se os campos From e To existem
      
        newData := ReturnData{
            Subject: msg.Envelope.Subject,
            From:    msg.Envelope.To,
            To:      msg.Envelope.From,
            Body:    body, // Corpo convertido para string
        }
        data = append(data, newData)
        i++
    }

	if err := <-done; err != nil {
		log.Fatal(err)
	}
	return data
}

func (f *ImapClient) ListEmailBox() []string {
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- f.cImap.List("", "*", mailboxes)
	}()

	var data []string

	i := 0
	for m := range mailboxes {
		data = append(data, m.Name)
		i++
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}
	return data
}
