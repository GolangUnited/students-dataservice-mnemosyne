package handler

import (
	"github.com/NEKETSKY/mnemosyne/models/mnemosyne"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Mnemosyne
// @Tags Mnemosyne
// @Description Get answer by mnemosyne command
// @ID mnemosyne-request
// @Accept  json
// @Produce  json
// @Param request body mnemosyne.Request true "Body request"
// @Success 200 {object} mnemosyne.Response
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} mnemosyne.Response
// @Router /mnemosyne [post]
func (h *Handler) mnemosyneRequest(c *gin.Context) {
	var req mnemosyne.Request

	if err := c.BindJSON(&req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.Mnemosyne.Test(c, req)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
