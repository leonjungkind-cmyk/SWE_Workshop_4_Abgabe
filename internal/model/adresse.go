package model

// Adresse maps to the existing "adresse" table in the "kunde" schema.
type Adresse struct {
	ID         uint64 `gorm:"column:id;primaryKey" json:"id"`
	Strasse    string `gorm:"column:strasse;not null" json:"strasse"`
	Hausnummer string `gorm:"column:hausnummer;not null" json:"hausnummer"`
	PLZ        string `gorm:"column:plz;not null" json:"plz"`
	Ort        string `gorm:"column:ort;not null" json:"ort"`
	KundeID    uint64 `gorm:"column:kunde_id;not null" json:"kunde_id"`

	Kunde *Kunde `gorm:"foreignKey:KundeID;references:ID" json:"kunde,omitempty"`
}

func (Adresse) TableName() string {
	return "kunde.adresse"
}
