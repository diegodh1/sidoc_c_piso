package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
	"time"

	"gorm.io/gorm"
)

//DecryptPass
func decryptPass(encryptedString string) (decryptedString string) {

	key := []byte("integrappsssssss")
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}

//EncryptPass func
func encryptPass(stringToEncrypt string) (encryptedString string) {
	//Since the key is in string, we need to convert decode it to bytes
	key := []byte("integrappsssssss")
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

//CreateUser func
func CreateUser(user *AppUser, db *gorm.DB) Response {
	switch {
	case user.AppUserID == "":
		return Response{Payload: nil, Message: "El ID de usuario es obligatorio", Status: 400}
	case user.AppUserName == "":
		return Response{Payload: nil, Message: "La contraseña no puede estar vacía", Status: 400}
	case user.AppUserEmail == "":
		return Response{Payload: nil, Message: "El nombre es obligatorio", Status: 400}
	case user.AppUserPassword == "" || !strings.Contains(user.AppUserEmail, "@"):
		return Response{Payload: nil, Message: "El correo es obligatorio", Status: 4000}
	case user.AppUserEmail == "" || !strings.Contains(user.AppUserEmail, "@"):
		return Response{Payload: nil, Message: "El correo es obligatorio", Status: 4000}
	case user.AppUserErpID == -1:
		return Response{Payload: nil, Message: "Debe seleccionar un usuario de ERP", Status: 4000}
	default:
		user.AppUserName = strings.ToUpper(user.AppUserName)
		user.AppUserLastName = strings.ToUpper(user.AppUserLastName)
		user.AppUserCdate = time.Now()
		user.AppUserPassword = encryptPass(user.AppUserPassword)
		if err := db.Create(&user).Error; err != nil {
			if strings.Contains(err.Error(), "PRIMARY KEY") {
				return Response{Payload: nil, Message: "El registro ya existe en el sistema", Status: 400}
			}
			return Response{Payload: nil, Message: "No se pudo crear el registro", Status: 500}
		}
		return Response{Payload: nil, Message: "Registro Realizado!", Status: 201}
	}
}

//GetUsersERP FUNC
func GetUsersERP(db *gorm.DB) Response {
	usuarios := []UsuariosErp{}
	db.Find(&usuarios)
	return Response{Payload: usuarios, Message: "OK", Status: 200}
}
