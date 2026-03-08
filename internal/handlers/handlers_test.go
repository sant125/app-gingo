package handlers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/santzin/gin-tattoo/internal/handlers"
	"github.com/santzin/gin-tattoo/internal/models"
)

// ── Mock DB ────────────────────────────────────────────────────────────────────

type mockDB struct {
	queryFn    func(ctx context.Context, sql string, args ...any) (handlers.Rows, error)
	queryRowFn func(ctx context.Context, sql string, args ...any) handlers.Row
	pingErr    error
}

func (m *mockDB) Query(ctx context.Context, sql string, args ...any) (handlers.Rows, error) {
	if m.queryFn != nil {
		return m.queryFn(ctx, sql, args...)
	}
	return &mockRows{}, nil
}

func (m *mockDB) QueryRow(ctx context.Context, sql string, args ...any) handlers.Row {
	if m.queryRowFn != nil {
		return m.queryRowFn(ctx, sql, args...)
	}
	return &mockRow{}
}

func (m *mockDB) Ping(ctx context.Context) error { return m.pingErr }

// mockRow implements handlers.Row.
type mockRow struct {
	scanFn func(dest ...any) error
}

func (r *mockRow) Scan(dest ...any) error {
	if r.scanFn != nil {
		return r.scanFn(dest...)
	}
	return nil
}

// mockRows implements handlers.Rows.
type mockRows struct {
	rows    []func(dest ...any) error
	current int
	err     error
}

func (m *mockRows) Next() bool {
	if m.current < len(m.rows) {
		m.current++
		return true
	}
	return false
}

func (m *mockRows) Scan(dest ...any) error { return m.rows[m.current-1](dest...) }
func (m *mockRows) Close()                 {}
func (m *mockRows) Err() error             { return m.err }

// ── Router ─────────────────────────────────────────────────────────────────────

func newRouter(h *handlers.H) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/api/v1/styles", h.ListStyles)
	r.GET("/api/v1/styles/:id", h.GetStyle)
	r.GET("/api/v1/curiosities", h.ListCuriosities)
	r.GET("/api/v1/curiosities/:id", h.GetCuriosity)
	r.GET("/health", h.HealthCheck)
	return r
}

// ── HealthCheck ────────────────────────────────────────────────────────────────

func TestHealthCheck_OK(t *testing.T) {
	h := &handlers.H{DB: &mockDB{}}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	newRouter(h).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var body map[string]string
	json.NewDecoder(w.Body).Decode(&body)
	if body["status"] != "ok" {
		t.Fatalf("expected status ok, got %q", body["status"])
	}
}

func TestHealthCheck_DBDown(t *testing.T) {
	h := &handlers.H{DB: &mockDB{pingErr: fmt.Errorf("connection refused")}}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	newRouter(h).ServeHTTP(w, req)

	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected 503, got %d", w.Code)
	}
}

// ── Styles ─────────────────────────────────────────────────────────────────────

func TestListStyles_OK(t *testing.T) {
	h := &handlers.H{DB: &mockDB{
		queryFn: func(_ context.Context, _ string, _ ...any) (handlers.Rows, error) {
			return &mockRows{rows: []func(dest ...any) error{
				func(dest ...any) error {
					*dest[0].(*int) = 1
					*dest[1].(*string) = "Blackwork"
					*dest[2].(*string) = "Desc"
					*dest[3].(*string) = "Europe"
					*dest[4].(*string) = "high"
					return nil
				},
			}}, nil
		},
	}}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/styles", nil)
	newRouter(h).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var styles []models.Style
	json.NewDecoder(w.Body).Decode(&styles)
	if len(styles) == 0 {
		t.Fatal("expected at least one style")
	}
}

func TestListStyles_QueryError(t *testing.T) {
	h := &handlers.H{DB: &mockDB{
		queryFn: func(_ context.Context, _ string, _ ...any) (handlers.Rows, error) {
			return nil, fmt.Errorf("db error")
		},
	}}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/styles", nil)
	newRouter(h).ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}

func TestListStyles_ScanError(t *testing.T) {
	h := &handlers.H{DB: &mockDB{
		queryFn: func(_ context.Context, _ string, _ ...any) (handlers.Rows, error) {
			return &mockRows{rows: []func(dest ...any) error{
				func(dest ...any) error { return fmt.Errorf("scan error") },
			}}, nil
		},
	}}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/styles", nil)
	newRouter(h).ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}

func TestGetStyle_OK(t *testing.T) {
	h := &handlers.H{DB: &mockDB{
		queryRowFn: func(_ context.Context, _ string, _ ...any) handlers.Row {
			return &mockRow{scanFn: func(dest ...any) error {
				*dest[0].(*int) = 1
				*dest[1].(*string) = "Blackwork"
				*dest[2].(*string) = "Desc"
				*dest[3].(*string) = "Europe"
				*dest[4].(*string) = "high"
				return nil
			}}
		},
	}}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/styles/1", nil)
	newRouter(h).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestGetStyle_NotFound(t *testing.T) {
	h := &handlers.H{DB: &mockDB{
		queryRowFn: func(_ context.Context, _ string, _ ...any) handlers.Row {
			return &mockRow{scanFn: func(dest ...any) error { return fmt.Errorf("not found") }}
		},
	}}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/styles/99", nil)
	newRouter(h).ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestGetStyle_InvalidID(t *testing.T) {
	h := &handlers.H{DB: &mockDB{}}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/styles/abc", nil)
	newRouter(h).ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

// ── Curiosities ────────────────────────────────────────────────────────────────

func TestListCuriosities_OK(t *testing.T) {
	h := &handlers.H{DB: &mockDB{
		queryFn: func(_ context.Context, _ string, _ ...any) (handlers.Rows, error) {
			return &mockRows{rows: []func(dest ...any) error{
				func(dest ...any) error {
					*dest[0].(*int) = 1
					*dest[1].(*string) = "Title"
					*dest[2].(*string) = "Content"
					*dest[3].(*string) = "history"
					return nil
				},
			}}, nil
		},
	}}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/curiosities", nil)
	newRouter(h).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var curiosities []models.Curiosity
	json.NewDecoder(w.Body).Decode(&curiosities)
	if len(curiosities) == 0 {
		t.Fatal("expected at least one curiosity")
	}
}

func TestListCuriosities_QueryError(t *testing.T) {
	h := &handlers.H{DB: &mockDB{
		queryFn: func(_ context.Context, _ string, _ ...any) (handlers.Rows, error) {
			return nil, fmt.Errorf("db error")
		},
	}}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/curiosities", nil)
	newRouter(h).ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}

func TestListCuriosities_ScanError(t *testing.T) {
	h := &handlers.H{DB: &mockDB{
		queryFn: func(_ context.Context, _ string, _ ...any) (handlers.Rows, error) {
			return &mockRows{rows: []func(dest ...any) error{
				func(dest ...any) error { return fmt.Errorf("scan error") },
			}}, nil
		},
	}}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/curiosities", nil)
	newRouter(h).ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}

func TestGetCuriosity_OK(t *testing.T) {
	h := &handlers.H{DB: &mockDB{
		queryRowFn: func(_ context.Context, _ string, _ ...any) handlers.Row {
			return &mockRow{scanFn: func(dest ...any) error {
				*dest[0].(*int) = 1
				*dest[1].(*string) = "Title"
				*dest[2].(*string) = "Content"
				*dest[3].(*string) = "history"
				return nil
			}}
		},
	}}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/curiosities/1", nil)
	newRouter(h).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestGetCuriosity_NotFound(t *testing.T) {
	h := &handlers.H{DB: &mockDB{
		queryRowFn: func(_ context.Context, _ string, _ ...any) handlers.Row {
			return &mockRow{scanFn: func(dest ...any) error { return fmt.Errorf("not found") }}
		},
	}}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/curiosities/99", nil)
	newRouter(h).ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestGetCuriosity_InvalidID(t *testing.T) {
	h := &handlers.H{DB: &mockDB{}}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/curiosities/abc", nil)
	newRouter(h).ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}
