package email

import (
	"backend/utils"
	"backend/viewmodels"
	"bytes"
	"github.com/go-errors/errors"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var emailChannel chan *viewmodels.EmailMessage

func Open() {
	emailChannel = make(chan *viewmodels.EmailMessage)
	go sendMessageInBackground()
}

func Exit() {
	close(emailChannel)
}

func SendMessage(message *viewmodels.EmailMessage) {
	emailChannel <- message
}

func makeDialer() (*gomail.Dialer, error) {
	var (
		dialer *gomail.Dialer

		port     int
		host     string
		usermail string
		password string

		err error
	)

	if port, err = strconv.Atoi(os.Getenv("EMAIL_PORT")); err != nil {
		return nil, err
	}

	host = strings.Trim(os.Getenv("EMAIL_HOST"), " ")
	usermail = strings.Trim(os.Getenv("EMAIL_AUTH_USER"), " ")
	password = strings.Trim(os.Getenv("EMAIL_AUTH_PASSWORD"), " ")

	if host == "" || port == 0 || usermail == "" || password == "" {
		err = errors.New("Host, Port, User Mail, and Password Must Be Set!")
	} else {
		dialer = gomail.NewDialer(host, port, usermail, password)
	}

	return dialer, err
}

func sendMessageInBackground() {

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	var (
		sendCloser gomail.SendCloser

		err  error
		open bool
		ok   bool

		dialer  *gomail.Dialer
		htmlBuf bytes.Buffer

		emailMessage gomail.Message
		vMessage     *viewmodels.EmailMessage
	)

	if dialer, err = makeDialer(); err != nil {
		panic("Dialer is not set!")
		return
	}

	open = false

	for {
		select {
		case vMessage, ok = <-emailChannel:
			if !ok {
				return
			}
			if vMessage == nil {
				return
			}
			if !open {
				if sendCloser, err = dialer.Dial(); err != nil {
					log.Println(err)
				}
				open = true
			}

			htmlBuf.Reset()

			err = utils.HtmlTemplates.ExecuteTemplate(
				&htmlBuf, "event_ticket.html", vMessage)

			emailMessage.SetHeader("From", os.Getenv("EMAIL_SENDER_NAME"))
			emailMessage.SetHeader("To", vMessage.To...)
			emailMessage.SetHeader("Subject", vMessage.Header)
			emailMessage.SetBody("text/html", htmlBuf.String())

			if err := gomail.Send(sendCloser, &emailMessage); err != nil {
				log.Println(err)
			}

			htmlBuf.Reset()
			emailMessage.Reset()

		case <-time.After(30 * time.Second):
			if open {
				if err := sendCloser.Close(); err != nil {
					log.Println(err)
				}
				open = false
			}
		}
	}
}
