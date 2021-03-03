package middlewares

import jwt "github.com/dgrijalva/jwt-go"

func decodeJWT(tokenReq string) string {
	token, err := jwt.Parse(tokenReq, func(token *jwt.Token) (interface{}, error) {
	   return []byte("ACCESS_SECRET"), nil
	})
	if err != nil {
	   return ""
	}
	claims, ok := token.Claims.(jwt.MapClaims)
  	if ok && token.Valid {
    	val_id_u, ok := claims["user_id"].(string)
		if !ok {
			return ""
		}
		return val_id_u
	}	

	return ""
}
