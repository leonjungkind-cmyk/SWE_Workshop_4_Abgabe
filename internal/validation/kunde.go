package validation

// AdresseCreateRequest is the expected JSON body part for creating a new Adresse.
// ID is assigned by the database and is not part of the request.
type AdresseCreateRequest struct {
	// Strasse is required.
	Strasse string `json:"strasse" binding:"required,max=80"`

	// Hausnummer is required. It is a string because house numbers can contain letters, e.g. "12a".
	Hausnummer string `json:"hausnummer" binding:"required,max=10"`

	// PLZ is required and must contain exactly 5 numeric characters.
	PLZ string `json:"plz" binding:"required,len=5,numeric"`

	// Ort is required.
	Ort string `json:"ort" binding:"required,max=80"`
}

// BestellungCreateRequest is the expected JSON body part for creating a new Bestellung.
// ID is assigned by the database and is not part of the request.
type BestellungCreateRequest struct {
	// Produktname is required.
	Produktname string `json:"produktname" binding:"required,max=80"`

	// Menge is required and must be at least 1.
	Menge int `json:"menge" binding:"required,min=1"`
}

// KundeCreateRequest is the expected JSON body for creating a new Kunde.
// ID, Version and foreign keys are assigned by the application/database and are not part of the request.
type KundeCreateRequest struct {
	// Nachname is required and limited to the column size of "nachname".
	Nachname string `json:"nachname" binding:"required,max=40"`

	// Email must be a syntactically valid email address.
	Email string `json:"email" binding:"required,email,max=40"`

	// Username is required and limited to the column size of "username".
	Username string `json:"username" binding:"required,max=40"`

	// Adresse is required when creating a complete customer record.
	Adresse AdresseCreateRequest `json:"adresse" binding:"required"`

	// Bestellungen are optional when creating a customer.
	// If provided, every Bestellung in the array is validated.
	Bestellungen []BestellungCreateRequest `json:"bestellungen,omitempty" binding:"omitempty,dive"`
}
