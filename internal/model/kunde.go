package model

// Kunde maps to the existing "kunde" table in the "kunde" schema.
// The table is created by the existing database scripts, not by GORM AutoMigrate.
type Kunde struct {
	ID       uint64 `gorm:"column:id;primaryKey" json:"id"`
	Nachname string `gorm:"column:nachname;size:40;not null" json:"nachname"`
	Email    string `gorm:"column:email;size:40;not null;unique" json:"email"`
	Username string `gorm:"column:username;size:40;not null;unique" json:"username"`
	Version  int    `gorm:"column:version;not null;default:0" json:"version"`

	Adresse      *Adresse     `gorm:"foreignKey:KundeID;references:ID" json:"adresse,omitempty"`
	Bestellungen []Bestellung `gorm:"foreignKey:KundeID;references:ID" json:"bestellungen,omitempty"`
}

func (Kunde) TableName() string {
	return "kunde.kunde"
}
