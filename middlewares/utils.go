package middlewares

import jwt "github.com/dgrijalva/jwt-go"
import "strings"

func decodeJWT(tokenReq string) (string, []string) {
	token, err := jwt.Parse(tokenReq, func(token *jwt.Token) (interface{}, error) {
	   return []byte("ACCESS_SECRET"), nil
	})
	if err != nil {
	   return "", nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
  	if ok && token.Valid {
    	val_id_u, ok := claims["user_id"].(string)
		if !ok {
			return "", nil
		}
		val_rols, err := claims["rols"].(string)
		if !err {
			return "", nil
		}
		resuRols := strings.Split(val_rols, ",")
		return val_id_u, resuRols
	}	

	return "", nil
}
