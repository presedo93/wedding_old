package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	db "github.com/presedo93/wedding/back/db/sqlc"
)

func (s *Server) createGuest(c *gin.Context) {
	var req db.CreateGuestParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	guest, err := s.store.CreateGuest(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, guest)
}

type GetGuestRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getGuest(c *gin.Context) {
	var req GetGuestRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	guest, err := s.store.GetGuest(c, req.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, guest)
}

func (s *Server) getUserGuests(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		err := errors.New("userID not found in context")
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	guests, err := s.store.GetUserGuests(c, userID.(string))
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, guests)
}

type GetAllGuestsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (s *Server) getAllGuests(c *gin.Context) {
	var req GetAllGuestsRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAllGuestsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	guests, err := s.store.GetAllGuests(c, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, guests)
}

func (s *Server) updateGuest(c *gin.Context) {
	var req db.UpdateGuestParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var guestID int
	if err := c.ShouldBindUri(&guestID); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	req.ID = int64(guestID)
	guest, err := s.store.UpdateGuest(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, guest)
}

func (s *Server) deleteGuest(c *gin.Context) {
	var guestID int
	if err := c.ShouldBindUri(&guestID); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := s.store.DeleteGuest(c, int64(guestID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, "Guest deleted")
}

func (s *Server) deleteUserGuests(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		err := errors.New("userID not found in context")
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	err := s.store.DeleteUserGuest(c, userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, "All user guests deleted")
}
