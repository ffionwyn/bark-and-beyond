package pages

import (
	"log"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

func HandleFormSubmission(c *gin.Context) {
    name := c.PostForm("name")
    email := c.PostForm("email")
    subject := c.PostForm("subject")
    message := c.PostForm("message")

    smtpHost := os.Getenv("SMTP_HOST")
    smtpPort := os.Getenv("SMTP_PORT")
    smtpUsername := os.Getenv("SMTP_USERNAME")
    smtpPassword := os.Getenv("SMTP_PASSWORD")

    body := "Name: " + name + "\n"
    body += "Email: " + email + "\n\n"
    body += "Subject: " + subject + "\n\n"
    body += "Message:\n" + message

    auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUsername, []string{"barkbeyonddogs@gmail.com"}, []byte(body))
    if err != nil {
        log.Printf("Error sending email: %v", err)
        c.JSON(500, gin.H{"message": "Internal server error"})
        return
    }

    c.JSON(200, gin.H{"message": "Email sent successfully"})
}