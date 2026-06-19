package validation

// KundeCreateRequest is the expected JSON body for creating a new Kunde.
// ID and Version are assigned by the database and are not part of the
// request.
type KundeCreateRequest struct {
	// Nachname is required and limited to the column size of "nachname".
	Nachname string `json:"nachname" binding:"required,max=40"`
	// Email must be a syntactically valid email address.
	Email string `json:"email" binding:"required,email,max=40"`
	// Username is required and limited to the column size of "username".
	Username string `json:"username" binding:"required,max=40"`
}
