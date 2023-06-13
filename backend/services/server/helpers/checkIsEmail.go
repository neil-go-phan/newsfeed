package helpers

import "net/mail"

func CheckIsEmail(checkedString string) bool {
	_, err := mail.ParseAddress(checkedString)
	return err == nil 
}