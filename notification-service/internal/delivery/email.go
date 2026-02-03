package delivery

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"notification-service/internal/domain"
)
type EmailSender struct {
	SMTPHost string
	SMTPPort string
	Username string
	Password string
	From     string
}

func (e EmailSender) Send(event domain.NotificationEvent){
	log.Printf("this is an email %s " , event.Title )
	addr := e.SMTPHost + ":" + e.SMTPPort
    auth := smtp.PlainAuth("", e.Username, e.Password, e.SMTPHost)

    // 2. Build the message with a buffer to ensure clean byte handling
    var msg bytes.Buffer
    msg.WriteString(fmt.Sprintf("From: %s\r\n", e.From))
    msg.WriteString(fmt.Sprintf("To: %s\r\n", event.Email))
    msg.WriteString(fmt.Sprintf("Subject: %s\r\n", event.Title))
    msg.WriteString("MIME-Version: 1.0\r\n")
    msg.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
    msg.WriteString("\r\n") // Critical empty line
    msg.WriteString(event.Message)

    // 3. Send
    err := smtp.SendMail(addr, auth, e.From, []string{event.Email}, msg.Bytes())
	   // 3. Attempt the send
	if err!=nil {
		log.Print("Email send Failed", err)
		return
	}

	log.Print("Email sent to :", event.Email )

}