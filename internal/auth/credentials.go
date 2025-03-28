package auth

import "errors"

func (a *Authenticator) verifyCredentials(login, password string) error {

	if login == "admin" && password == "secret" {
		return nil
	}
	return errors.New("wrong credentials")
}
