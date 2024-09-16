package bids

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateBidRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h *Implementation) UpdateBid(c *gin.Context) {
	bidId := c.Param("bidId")
	username := c.Param("username")
	var req UpdateBidRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	updatedBid, err := h.bidService.EditBid(bidId, username, req.Name, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedBid)
}
