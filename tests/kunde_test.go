package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"swe-workshop-api/internal/handler"
	"swe-workshop-api/internal/model"
)

// fakeKundeRepository is an in-memory repository.KundeRepository used to
// exercise the Kunde handlers without a real database connection.
type fakeKundeRepository struct {
	kunden []model.Kunde
}

func (f *fakeKundeRepository) GetAll() ([]model.Kunde, error) {
	return f.kunden, nil
}

func (f *fakeKundeRepository) GetByID(id uint64) (model.Kunde, bool, error) {
	for _, k := range f.kunden {
		if k.ID == id {
			return k, true, nil
		}
	}
	return model.Kunde{}, false, nil
}

func (f *fakeKundeRepository) Create(kunde *model.Kunde) error {
	kunde.ID = uint64(len(f.kunden) + 1)
	f.kunden = append(f.kunden, *kunde)
	return nil
}

// newKundeTestRouter builds a minimal router with the same Kunde routes as
// the real server, backed by a fakeKundeRepository instead of PostgreSQL.
func newKundeTestRouter(repo *fakeKundeRepository) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	kundeHandler := handler.NewKundeHandler(repo)
	router.GET("/api/public/kunden", kundeHandler.GetAll)
	router.POST("/api/secured/kunden", kundeHandler.Create)
	return router
}

func TestGetAllKunden(t *testing.T) {
	repo := &fakeKundeRepository{kunden: []model.Kunde{
		{ID: 1, Nachname: "Müller", Email: "mueller@example.de", Username: "mueller"},
	}}
	router := newKundeTestRouter(repo)

	req := httptest.NewRequest(http.MethodGet, "/api/public/kunden", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.JSONEq(t, `[{"id":1,"nachname":"Müller","email":"mueller@example.de","username":"mueller","version":0}]`, recorder.Body.String())
}

func TestCreateKundeValidBody(t *testing.T) {
	repo := &fakeKundeRepository{}
	router := newKundeTestRouter(repo)

	body := []byte(`{"nachname":"Schmidt","email":"schmidt@example.de","username":"schmidt"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/secured/kunden", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Len(t, repo.kunden, 1)
}

func TestCreateKundeInvalidBody(t *testing.T) {
	repo := &fakeKundeRepository{}
	router := newKundeTestRouter(repo)

	body := []byte(`{"nachname":"","email":"not-an-email","username":""}`)
	req := httptest.NewRequest(http.MethodPost, "/api/secured/kunden", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Empty(t, repo.kunden)
}
