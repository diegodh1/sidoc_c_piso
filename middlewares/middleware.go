package middlewares



import (
	"github.com/gin-gonic/gin"
	"strings"
	)

// Add more
var options = [3]string{"CREAR_USUARIOS", "EDITAR_USUARIOS", "ASIGNAR_PERFILES"}

//TokenMiddleware func
func TokenMiddleware(option byte) gin.HandlerFunc {
    return func(c *gin.Context) {
        reqToken := c.Request.Header.Get("Authorization")
        splitToken := strings.Split(reqToken, "Bearer")
        if len(splitToken) != 2 {
            c.AbortWithStatusJSON(401, gin.H{"payload": nil, "message": "Tokén no válido B", "status": 401})
            return
        }
        reqToken = strings.TrimSpace(splitToken[1])
        id, rols := decodeJWT(reqToken)
        if id == "" {
			c.AbortWithStatusJSON(401, gin.H{"payload": nil, "message": "Tokén no válido D", "status": 401})
            return
		}   
		for _, rol := range rols{
			if (rol == options[option]){
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(401, gin.H{"payload": nil, "message": "Su Rol no puede ejecutar esta acción", "status": 401})
    }
}