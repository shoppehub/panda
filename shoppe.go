package shoppe

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stnc/pongo2gin"
)

func New() *gin.Engine {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)

	// 使用默认中间件（logger和recovery）
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.HTMLRender = pongo2gin.TemplatePath("templates")

	// api.Create(r)

	return r
}

func Start(r *gin.Engine) {
	// 启动服务，并默认监听4000端口
	r.Run(":" + os.Getenv(PORT))
}
