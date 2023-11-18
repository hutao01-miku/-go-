package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", nil)
}
