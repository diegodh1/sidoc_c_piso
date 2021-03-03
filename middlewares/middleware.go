package middlewares



import (
	"github.com/gin-gonic/gin"
	"strings"
	"gorm.io/gorm"
	)

// Add more
var options = [3]string{"CREAR_USUARIOS", "EDITAR_USUARIOS", "BUSCAR_USUARIO_ERP"}

//TokenMiddleware func
func TokenMiddleware(option byte, db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        reqToken := c.Request.Header.Get("Authorization")
        splitToken := strings.Split(reqToken, "Bearer")
        if len(splitToken) != 2 {
            c.AbortWithStatusJSON(401, gin.H{"payload": nil, "message": "No se envío Token", "status": 401})
            return
        }
        reqToken = strings.TrimSpace(splitToken[1])
        id := decodeJWT(reqToken)
        if id == "" {
			c.AbortWithStatusJSON(401, gin.H{"payload": nil, "message": "Tokén no válido", "status": 401})
            return
		}
		var rols []string
		db.Raw("SELECT app_submenu_id FROM user_permissions WHERE app_user_id = ?", id).Scan(&rols)   
		for _, rol := range rols{
			if (strings.TrimSpace(rol) == options[option]){
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(401, gin.H{"payload": nil, "message": "Su Rol no puede ejecutar esta acción", "status": 401})
    }
}