package handler

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"regexp"
	"os"
	"crypto/rand"
	"fmt"
	"io"
	"github.com/go-gomail/gomail"
	)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func CreateToken(userid string, rols []string) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userid
	atClaims["rols"] = strings.Join(rols,", ")
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("ACCESS_SECRET"))
	if err != nil {
		return "", err
	}
	return token, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) string {
	bytesPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return ""
	}
	return string(bytesPass)
}

func ValidEmail(email string) bool {

	if !emailRegex.MatchString(email) {
		return false
	}

	return true
}

func EnviarCorreo(to string, id string,codVerf string) bool {
	
	if(!ValidEmail(to)){
		return false
	}
	
    from := "app.sidoc@sidocsa.com"
    pass := os.Getenv("SIDOC_EMAIL_PASS")

    m := gomail.NewMessage()
    m.SetHeader("From", from)
    m.SetHeader("To", to)
    m.SetHeader("Subject", "Cambio contraseña")
    m.SetBody("text/html", fmt.Sprintf(
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
    <td>%s</td>
    <td>%s</td>
    </tr>
    </tbody>
    </table>
    <p>Sí usted no ha hecho esta solicitud, por favor haga caso omiso a este correo.</p>
    <p>Correo generado automaticamente.</p></body>
    </html>`, to, codVerf))

    // Send the email to user
    d := gomail.NewPlainDialer("mail.sidocsa.com", 993, from, pass)
    if err := d.DialAndSend(m); err != nil {
		fmt.Print("JajajaSISAS")
		fmt.Print(err)
        return false
    }
    return true
}

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

