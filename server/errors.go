// errors are in a dedicated file, because they are used in more than one place
package server

import "errors"

var ErrInvalidForm = errors.New("the form has invalid syntax")
var ErrSavingToDatabase = errors.New("error saving to the database")
var ErrWrongValue = errors.New("error converting or getting value")

// should not occur during runtime
var ErrCatalogueNotFound = errors.New("cake catalogue not found on the server")
