package repository

import (
	"errors"

	"gorm.io/gorm"

	"swe-workshop-api/internal/model"
)

// KundeRepository defines the persistence operations for model.Kunde. It is
// an interface so handlers can be tested against a fake implementation
// without a real database connection.
type KundeRepository interface {
	// GetAll loads every Kunde record.
	GetAll() ([]model.Kunde, error)
	// GetByID loads a single Kunde by its ID. The second return value is
	// false if no matching record exists.
	GetByID(id uint64) (model.Kunde, bool, error)
	// Create persists a new Kunde. The database assigns its ID.
	Create(kunde *model.Kunde) error
}

// gormKundeRepository is the GORM-based KundeRepository implementation used
// in production, backed by the existing PostgreSQL "kunde" table.
type gormKundeRepository struct {
	db *gorm.DB
}

// NewKundeRepository creates a KundeRepository backed by the given GORM
// connection.
func NewKundeRepository(db *gorm.DB) KundeRepository {
	return &gormKundeRepository{db: db}
}

// GetAll loads every Kunde record.
func (r *gormKundeRepository) GetAll() ([]model.Kunde, error) {
	var kunden []model.Kunde
	if err := r.db.Find(&kunden).Error; err != nil {
		return nil, err
	}
	return kunden, nil
}

// GetByID loads a single Kunde by its ID. The second return value is false
// if no matching record exists.
func (r *gormKundeRepository) GetByID(id uint64) (model.Kunde, bool, error) {
	var kunde model.Kunde
	err := r.db.First(&kunde, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Kunde{}, false, nil
	}
	if err != nil {
		return model.Kunde{}, false, err
	}
	return kunde, true, nil
}

// Create persists a new Kunde. The database assigns its ID.
func (r *gormKundeRepository) Create(kunde *model.Kunde) error {
	return r.db.Create(kunde).Error
}
