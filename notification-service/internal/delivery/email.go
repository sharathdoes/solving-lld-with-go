package delivery

import (
	"crypto/tls"
	"log"

	"gopkg.in/gomail.v2"

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
	log.Printf("this is an email %s to %s " , event.Title, event.Email )
	m := gomail.NewMessage()
    m.SetHeader("From", e.From)
    m.SetHeader("To", event.Email)
    m.SetHeader("Subject", event.Title)
    m.SetBody("text/plain", event.Message)

    // Gomail handles the dialer, TLS, and RFC formatting for you
    d := gomail.NewDialer(e.SMTPHost, 587, e.Username, e.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
    if err := d.DialAndSend(m); err != nil {
        log.Printf("Failed: %v", err)
        return
    }
    log.Printf("✅ Success! Sent to %s", event.Email)

	log.Print("Email sent to :", event.Email )

}