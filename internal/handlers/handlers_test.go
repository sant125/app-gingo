package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/santzin/gin-tattoo/internal/data"
	"github.com/santzin/gin-tattoo/internal/models"
)

func newTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// stub handlers using in-memory data (no DB needed in unit tests)
	r.GET("/api/v1/styles", func(c *gin.Context) {
		c.JSON(http.StatusOK, data.Styles)
	})
	r.GET("/api/v1/curiosities", func(c *gin.Context) {
		c.JSON(http.StatusOK, data.Curiosities)
	})
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}

func TestHealthCheck(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	newTestRouter().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var body map[string]string
	json.NewDecoder(w.Body).Decode(&body)
	if body["status"] != "ok" {
		t.Fatalf("expected status ok, got %q", body["status"])
	}
}

func TestListStyles(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/styles", nil)
	newTestRouter().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var styles []models.Style
	json.NewDecoder(w.Body).Decode(&styles)
	if len(styles) == 0 {
		t.Fatal("expected at least one style")
	}
}

func TestListCuriosities(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/curiosities", nil)
	newTestRouter().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var curiosities []models.Curiosity
	json.NewDecoder(w.Body).Decode(&curiosities)
	if len(curiosities) == 0 {
		t.Fatal("expected at least one curiosity")
	}
}
