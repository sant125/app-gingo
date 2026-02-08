package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/santzin/gin-tattoo/internal/models"
)

// ListStyles godoc
//
//	@Summary		List tattoo styles
//	@Description	Returns all available tattoo styles
//	@Tags			styles
//	@Produce		json
//	@Success		200	{array}		models.Style
//	@Failure		500	{object}	map[string]string
//	@Router			/api/v1/styles [get]
func (h *H) ListStyles(c *gin.Context) {
	rows, err := h.DB.Query(c.Request.Context(),
		"SELECT id, name, description, origin, popularity FROM styles ORDER BY id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query styles"})
		return
	}
	defer rows.Close()

	styles := make([]models.Style, 0)
	for rows.Next() {
		var s models.Style
		if err := rows.Scan(&s.ID, &s.Name, &s.Description, &s.Origin, &s.Popularity); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to scan style"})
			return
		}
		styles = append(styles, s)
	}

	c.JSON(http.StatusOK, styles)
}

// GetStyle godoc
//
//	@Summary		Get a tattoo style
//	@Description	Returns details of a specific tattoo style by ID
//	@Tags			styles
//	@Produce		json
//	@Param			id	path		int	true	"Style ID"
//	@Success		200	{object}	models.Style
//	@Failure		400	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Router			/api/v1/styles/{id} [get]
func (h *H) GetStyle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var s models.Style
	err = h.DB.QueryRow(c.Request.Context(),
		"SELECT id, name, description, origin, popularity FROM styles WHERE id = $1", id,
	).Scan(&s.ID, &s.Name, &s.Description, &s.Origin, &s.Popularity)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "style not found"})
		return
	}

	c.JSON(http.StatusOK, s)
}
