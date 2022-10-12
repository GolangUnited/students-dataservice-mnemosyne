package rest

import (
	"github.com/NEKETSKY/mnemosyne/pkg/log"
	"github.com/gin-gonic/gin"
)

func (h *Handler) logger(c *gin.Context) {
	l := log.LoggerFromContext(h.ctx)
	c.Set(log.LoggerKey, l)
}
