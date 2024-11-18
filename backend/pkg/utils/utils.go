package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateOTP() (string, error) {
	const otpLength = 4
	var otp string

	for i := 0; i < otpLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		otp += num.String()
	}
	return otp, nil
}

func GenerateRandomElements(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return ""
		}
		b[i] = charset[idx.Int64()]
	}
	return string(b)
}

func GetEmailDomain(email string) string {
	index := strings.LastIndex(email, "@")
	if index == -1 {
		return ""
	}
	return email[index+1:]
}

func GetEmailProvider(emailDomain string) string {
	providers := map[string]string{
		"gmail.com": "smtp.gmail.com",
		"hotmail.com": "smtp-mail.outlook.com",
		"outlook.com": "smtp-mail.outlook.com",
	}
	return providers[emailDomain]
}

func SendEmail(email, emailProvider, body string, subject string) error {
    services := map[string]struct {
        from     string
        password string
        host     string
        port     int
    }{
        "smtp.gmail.com": {
            from:     os.Getenv("GMAIL_EMAIL"),
            password: os.Getenv("GMAIL_PASSWORD"),
            host:     "smtp.gmail.com",
            port:     587,
        },
        "smtp-mail.outlook.com": {
            from:     os.Getenv("OUTLOOK_EMAIL"),
            password: os.Getenv("OUTLOOK_PASSWORD"),
            host:     "smtp-mail.outlook.com",
            port:     587,
        },
    }

    service, exists := services[emailProvider]
    if !exists {
        return fmt.Errorf("unsupported email provider: %s", emailProvider)
    }

    m := gomail.NewMessage()
    m.SetHeader("From", service.from)
    m.SetHeader("To", email)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)

    
    d := gomail.NewDialer(service.host, service.port, service.from, service.password)

    if err := d.DialAndSend(m); err != nil {
        return fmt.Errorf("failed to send email: %v", err)
    }

    return nil
}