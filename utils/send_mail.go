
package utils

import (
	"fmt"
	"net/smtp"
	"c2nofficialsitebackend/config"
)

func SendMail() {
	from := "contactus@c2nshop.com"
	password := config.Env.ZOHO_MAIL_PASSWORD // app password if 2FA is enabled

	to := []string{"harshsason2000@gmail.com"}
	smtpHost := "smtp.zoho.in"
	smtpPort := "587"

	message := []byte("From: C2N Shop <contactus@c2nshop.com>\r\n" +
		"To: harshsason2000@gmail.com\r\n" +
		"Subject: Hello from Go via Zoho SMTP\r\n" +
		"\r\n" +
		"This is a test email sent from Go using Zoho's SMTP server.\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully.")
}