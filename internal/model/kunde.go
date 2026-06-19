package model

// Kunde maps to the existing "kunde" table in the "kunde" schema of the
// PostgreSQL database (see deployments/postgres/init/kunde/sql). The table
// is created by create-table.sql, not by GORM AutoMigrate.
type Kunde struct {
	// ID is the database-generated primary key.
	ID uint64 `gorm:"column:id;primaryKey" json:"id"`
	// Nachname is the customer's last name.
	Nachname string `gorm:"column:nachname;size:40;not null" json:"nachname"`
	// Email is the customer's email address. Unique in the database.
	Email string `gorm:"column:email;size:40;not null;unique" json:"email"`
	// Username is the customer's login name. Unique in the database.
	Username string `gorm:"column:username;size:40;not null;unique" json:"username"`
	// Version is used for optimistic locking and defaults to 0.
	Version int `gorm:"column:version;not null;default:0" json:"version"`
}

// TableName returns the schema-qualified table name, since "kunde" lives in
// the "kunde" schema rather than the default "public" schema.
func (Kunde) TableName() string {
	return "kunde.kunde"
}
