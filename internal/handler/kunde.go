package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"swe-workshop-api/internal/model"
	"swe-workshop-api/internal/repository"
	"swe-workshop-api/internal/validation"
)

// KundeHandler holds the Gin handlers for the Kunde REST resource.
type KundeHandler struct {
	repo repository.KundeRepository
}

// NewKundeHandler creates a KundeHandler backed by the given repository.
func NewKundeHandler(repo repository.KundeRepository) *KundeHandler {
	return &KundeHandler{repo: repo}
}

// GetAll handles GET /api/public/kunden and returns every Kunde as a JSON array.
func (h *KundeHandler) GetAll(c *gin.Context) {
	kunden, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load kunden"})
		return
	}

	c.JSON(http.StatusOK, kunden)
}

// GetByID handles GET /api/public/kunden/:id and returns a single Kunde.
func (h *KundeHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return
	}

	kunde, found, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load kunde"})
		return
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "kunde not found"})
		return
	}

	c.JSON(http.StatusOK, kunde)
}

// Create handles POST /api/secured/kunden.
// It validates the request body via Gin binding and then creates a new Kunde
// including Adresse and optional Bestellungen.
func (h *KundeHandler) Create(c *gin.Context) {
	var req validation.KundeCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bestellungen := make([]model.Bestellung, 0, len(req.Bestellungen))
	for _, bestellungReq := range req.Bestellungen {
		bestellungen = append(bestellungen, model.Bestellung{
			Produktname: bestellungReq.Produktname,
			Menge:       bestellungReq.Menge,
		})
	}

	kunde := model.Kunde{
		Nachname: req.Nachname,
		Email:    req.Email,
		Username: req.Username,
		Adresse: &model.Adresse{
			Strasse:    req.Adresse.Strasse,
			Hausnummer: req.Adresse.Hausnummer,
			PLZ:        req.Adresse.PLZ,
			Ort:        req.Adresse.Ort,
		},
		Bestellungen: bestellungen,
	}

	if err := h.repo.Create(&kunde); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create kunde"})
		return
	}

	c.JSON(http.StatusCreated, kunde)
}
