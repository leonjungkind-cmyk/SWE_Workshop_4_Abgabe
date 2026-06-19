package model

// Bestellung maps to the existing "bestellung" table in the "kunde" schema.
type Bestellung struct {
	ID          uint64 `gorm:"column:id;primaryKey" json:"id"`
	Produktname string `gorm:"column:produktname;not null" json:"produktname"`
	Menge       int    `gorm:"column:menge;not null" json:"menge"`
	KundeID     uint64 `gorm:"column:kunde_id;not null" json:"kunde_id"`

	Kunde *Kunde `gorm:"foreignKey:KundeID;references:ID" json:"kunde,omitempty"`
}

func (Bestellung) TableName() string {
	return "kunde.bestellung"
}
