package handler

import (
	"crypto/rand"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"net/smtp"
	"os"
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

//CreateToken func
func CreateToken(userid string) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userid
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("ACCESS_SECRET"))
	if err != nil {
		return "", err
	}
	return token, nil
}

//CheckPasswordHash func
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//HashPassword func
func HashPassword(password string) string {
	bytesPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return ""
	}
	return string(bytesPass)
}

//ValidEmail func
func ValidEmail(email string) bool {

	if !emailRegex.MatchString(email) {
		return false
	}

	return true
}

//EnviarCorreo func
func EnviarCorreo(to string, id string, codVerf string) bool {

	if !ValidEmail(to) {
		return false
	}

	from := "app.sidoc@sidocsa.com"
	pass := os.Getenv("SIDOC_EMAIL_PASS")
	fromM := fmt.Sprintf("From: <%s>\r\n", from)
	toM := fmt.Sprintf("To: <%s>\r\n", "recipient@gmail.com")
	subject := "Subject: Cambio de contraseña\r\n"
	body := fmt.Sprintf(
		`<!DOCTYPE html>
		<style>
			body {
			   font-family: "HelveticaNeue-Light", "Helvetica Neue Light", "Helvetica Neue", Helvetica, Arial, "Lucida Grande", sans-serif; 
			   font-weight: 300;
			}
		</style>
		<html>
			<body><p>Hola buen d&iacute;a,</p>
		<p>El siguiente correo es para realizar un cambio de contraseña.</strong> Por favor digite en la app este codigo de verf.</p>
		<table table style="border-collapse: collapse; background-color: #fF6FE49; border-style: solid;" border="1">
		<tbody>
		<tr>
		<td>Correo</td>
		<td>Codigo de verf</td>
		</tr>
		<tr>
		<td>%s</td>
		<td>%s</td>
		</tr>
		</tbody>
		</table>
		<p>Sí usted no ha hecho esta solicitud, por favor haga caso omiso a este correo.</p>
		<p>Correo generado automaticamente.</p></body>
		</html>`, to, codVerf)

	msg := fromM + toM + subject + "Content-Type: text/html; charset=\"utf-8\"\r\n\r\n" + body + "\r\n"

	auth := smtp.PlainAuth("", from, pass, "mail.sidocsa.com")
	// Send the email to user

	if err := SendMailTLS("mail.sidocsa.com:465", auth, from, []string{to}, []byte(msg)); err != nil {
		return false
	}
	return true
}

//EncodeToString func
func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

//SendMailTLS func
func SendMailTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	tlsconfig := &tls.Config{ServerName: host}
	if err = validateLine(from); err != nil {
		return err
	}
	for _, recp := range to {
		if err = validateLine(recp); err != nil {
			return err
		}
	}
	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return err
	}
	defer conn.Close()
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	defer c.Close()
	if err = c.Hello("localhost"); err != nil {
		return err
	}
	if err = c.Auth(auth); err != nil {
		return err
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

// validateLine checks to see if a line has CR or LF as per RFC 5321
func validateLine(line string) error {
	if strings.ContainsAny(line, "\n\r") {
		return errors.New("a line must not contain CR or LF")
	}
	return nil
}
