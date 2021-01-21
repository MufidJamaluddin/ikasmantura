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
var emailChannelClosed bool

func Open() {
	emailChannel = make(chan *viewmodels.EmailMessage)
	emailChannelClosed = false
	go sendMessageInBackground()
}

func Exit() {
	emailChannelClosed = true
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
		if !emailChannelClosed {
			if err := recover(); err != nil {
				log.Println("send email panic occurred:", err)
			}
		}
	}()

	var (
		sendCloser gomail.SendCloser

		err  error = nil
		open bool
		ok   bool

		dialer  *gomail.Dialer
		htmlBuf bytes.Buffer

		emailMessage *gomail.Message
		vMessage     *viewmodels.EmailMessage
	)

	if dialer, err = makeDialer(); err != nil {
		log.Println(err.Error())
		panic("Dialer is not set!")
		return
	}

	open = false
	emailMessage = gomail.NewMessage()

	for {
		select {
		case vMessage, ok = <-emailChannel:
			if !ok {
				return
			}
			if vMessage == nil {
				continue
			}
			if !open {
				if sendCloser, err = dialer.Dial(); err != nil {
					log.Println("error in connecting the email server")
					log.Println(err.Error())
				}
				open = true
			}

			htmlBuf.Reset()
			emailMessage.Reset()

			if err = utils.HtmlTemplates.ExecuteTemplate(
				&htmlBuf, "email_template.html", vMessage); err != nil {

				log.Println("error in use email HTML template")
				log.Println(err.Error())

				htmlBuf.WriteString(string(vMessage.Message))
			}

			emailMessage.SetHeader("From", os.Getenv("EMAIL_SENDER_NAME"))
			emailMessage.SetHeader("To", vMessage.To...)
			emailMessage.SetHeader("Subject", vMessage.Header)
			emailMessage.SetBody("text/html", htmlBuf.String())

			if err = gomail.Send(sendCloser, emailMessage); err != nil {
				log.Println("error in sending email")
				log.Println(err.Error())
			}

			htmlBuf.Reset()
			emailMessage.Reset()

		case <-time.After(35 * time.Second):
			if open {
				if err = sendCloser.Close(); err != nil {
					log.Println("error in closing the email server")
					log.Println(err.Error())
				}
				open = false
			}
		}
	}
}
