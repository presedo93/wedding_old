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

type createProfileBody struct {
	Name             string    `json:"name" binding:"required"`
	Phone            string    `json:"phone" binding:"required,e164"`
	Email            string    `json:"email" binding:"required,email"`
	PictureUrl       string    `json:"picture_url"`
	ID               uuid.UUID `json:"id" binding:"required,uuid4"`
	CompletedProfile bool      `json:"completed_profile"`
	AddedGuests      bool      `json:"added_guests"`
	AddedSongs       bool      `json:"added_songs"`
	AddedPictures    bool      `json:"added_pictures"`
}

func (s *Server) createProfile(c *gin.Context) {
	var body createProfileBody
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

	arg := db.CreateProfileParams{
		ID:               userID.(uuid.UUID),
		Name:             body.Name,
		Phone:            body.Phone,
		Email:            body.Email,
		PictureUrl:       pgtype.Text{String: body.PictureUrl, Valid: body.PictureUrl != ""},
		CompletedProfile: body.CompletedProfile,
		AddedGuests:      body.AddedGuests,
		AddedSongs:       body.AddedSongs,
		AddedPictures:    body.AddedPictures,
	}

	profile, err := s.store.CreateProfile(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (s *Server) getProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		err := errors.New("userID not found in context")
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	profile, err := s.store.GetProfile(c, userID.(uuid.UUID))
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, profile)
}

type getProfilesForm struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// TODO: Needs to verify that user is an admin (group).
func (s *Server) getProfiles(c *gin.Context) {
	var form getProfilesForm

	if err := c.ShouldBindQuery(&form); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetProfilesParams{
		Limit:  form.PageSize,
		Offset: (form.PageID - 1) * form.PageSize,
	}

	profiles, err := s.store.GetProfiles(c, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, profiles)
}

type updateProfileBody struct {
	Name             string `json:"name"`
	Phone            string `json:"phone" binding:"omitempty,e164"`
	Email            string `json:"email" binding:"omitempty,email"`
	PictureUrl       string `json:"picture_url" binding:"omitempty,http_url"`
	CompletedProfile bool   `json:"completed_profile"`
	AddedGuests      bool   `json:"added_guests"`
	AddedSongs       bool   `json:"added_songs"`
	AddedPictures    bool   `json:"added_pictures"`
}

func (s *Server) updateProfile(c *gin.Context) {
	var body updateProfileBody

	userID, exists := c.Get("userID")
	if !exists {
		err := errors.New("userID not found in context")
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateProfileParams{
		ID:               userID.(uuid.UUID),
		Name:             pgtype.Text{String: body.Name, Valid: body.Name != ""},
		Phone:            pgtype.Text{String: body.Phone, Valid: body.Phone != ""},
		Email:            pgtype.Text{String: body.Email, Valid: body.Email != ""},
		PictureUrl:       pgtype.Text{String: body.PictureUrl, Valid: body.PictureUrl != ""},
		CompletedProfile: pgtype.Bool{Bool: body.CompletedProfile, Valid: body.CompletedProfile},
		AddedGuests:      pgtype.Bool{Bool: body.AddedGuests, Valid: body.AddedGuests},
		AddedSongs:       pgtype.Bool{Bool: body.AddedSongs, Valid: body.AddedSongs},
		AddedPictures:    pgtype.Bool{Bool: body.AddedPictures, Valid: body.AddedPictures},
	}

	profile, err := s.store.UpdateProfile(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (s *Server) deleteProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		err := errors.New("userID not found in context")
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	err := s.store.DeleteProfile(c, userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, "Profile deleted")
}
