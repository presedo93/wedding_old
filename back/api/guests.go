package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/presedo93/wedding/back/db/sqlc"
)

type createGuestBody struct {
	Name           string   `json:"name" binding:"required"`
	Phone          string   `json:"phone" binding:"required,e164"`
	Allergies      []string `json:"allergies" binding:"required"`
	IsVegetarian   bool     `json:"is_vegetarian"`
	NeedsTransport bool     `json:"needs_transport"`
}

func (s *Server) createGuest(c *gin.Context) {
	var body createGuestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		err := errors.New("userID not found in context")
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.CreateGuestParams{
		ProfileID:      userID.(uuid.UUID),
		Name:           body.Name,
		Phone:          body.Phone,
		Allergies:      body.Allergies,
		IsVegetarian:   body.IsVegetarian,
		NeedsTransport: body.NeedsTransport,
	}

	guest, err := s.store.CreateGuest(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, guest)
}

type getGuestUri struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getGuest(c *gin.Context) {
	var uri getGuestUri

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	guest, err := s.store.GetGuest(c, uri.ID)
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

	guests, err := s.store.GetUserGuests(c, userID.(uuid.UUID))
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

type getAllGuestsForm struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (s *Server) getAllGuests(c *gin.Context) {
	var form getAllGuestsForm

	if err := c.ShouldBindQuery(&form); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetGuestsParams{
		Limit:  form.PageSize,
		Offset: (form.PageID - 1) * form.PageSize,
	}

	guests, err := s.store.GetGuests(c, arg)
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

type updateGuestUri struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateGuestBody struct {
	Name           string   `json:"name"`
	Phone          string   `json:"phone" binding:"omitempty,e164"`
	Allergies      []string `json:"allergies"`
	IsVegetarian   bool     `json:"is_vegetarian"`
	NeedsTransport bool     `json:"needs_transport"`
}

func (s *Server) updateGuest(c *gin.Context) {
	var body updateGuestBody
	var uri updateGuestUri

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateGuestParams{
		ID:             uri.ID,
		Name:           pgtype.Text{String: body.Name, Valid: body.Name != ""},
		Phone:          pgtype.Text{String: body.Phone, Valid: body.Phone != ""},
		Allergies:      body.Allergies,
		IsVegetarian:   pgtype.Bool{Bool: body.IsVegetarian, Valid: body.IsVegetarian},
		NeedsTransport: pgtype.Bool{Bool: body.NeedsTransport, Valid: body.NeedsTransport},
	}

	guest, err := s.store.UpdateGuest(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, guest)
}

type deleteGuestsUri struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) deleteGuest(c *gin.Context) {
	var uri deleteGuestsUri

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := s.store.DeleteGuest(c, uri.ID)
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

	err := s.store.DeleteUserGuest(c, userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, "All user guests deleted")
}
