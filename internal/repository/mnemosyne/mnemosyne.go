package mnemosyne

import (
	"github.com/gin-gonic/gin"
)

type Mnemosyne struct {
}

func NewMnemosyne() *Mnemosyne {
	return &Mnemosyne{}
}

func (c *Mnemosyne) Test(ctx *gin.Context) (err error) {
	_ = ctx

	return
}
