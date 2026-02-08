package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/santzin/gin-tattoo/internal/models"
)

// ListCuriosities godoc
//
//	@Summary		List tattoo curiosities
//	@Description	Returns all tattoo curiosities and interesting facts
//	@Tags			curiosities
//	@Produce		json
//	@Success		200	{array}		models.Curiosity
//	@Failure		500	{object}	map[string]string
//	@Router			/api/v1/curiosities [get]
func (h *H) ListCuriosities(c *gin.Context) {
	rows, err := h.DB.Query(c.Request.Context(),
		"SELECT id, title, content, category FROM curiosities ORDER BY id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query curiosities"})
		return
	}
	defer rows.Close()

	curiosities := make([]models.Curiosity, 0)
	for rows.Next() {
		var cur models.Curiosity
		if err := rows.Scan(&cur.ID, &cur.Title, &cur.Content, &cur.Category); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to scan curiosity"})
			return
		}
		curiosities = append(curiosities, cur)
	}

	c.JSON(http.StatusOK, curiosities)
}

// GetCuriosity godoc
//
//	@Summary		Get a tattoo curiosity
//	@Description	Returns details of a specific curiosity by ID
//	@Tags			curiosities
//	@Produce		json
//	@Param			id	path		int	true	"Curiosity ID"
//	@Success		200	{object}	models.Curiosity
//	@Failure		400	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Router			/api/v1/curiosities/{id} [get]
func (h *H) GetCuriosity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var cur models.Curiosity
	err = h.DB.QueryRow(c.Request.Context(),
		"SELECT id, title, content, category FROM curiosities WHERE id = $1", id,
	).Scan(&cur.ID, &cur.Title, &cur.Content, &cur.Category)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "curiosity not found"})
		return
	}

	c.JSON(http.StatusOK, cur)
}
