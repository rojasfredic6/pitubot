package data

import "fmt"

type AuthData struct {
	Token      string
	Auth       string
	MessageUrl string
}

var AuthDataInfo AuthData = AuthData{Token: "EAARUv7hbyVkBAI7Y122FnwokkJE0aZAZAF33Id72wGBT3L7mAafJKNUddmsJVvW5L0NyEZAYJ2n258fGxCJ5Hr7zhcykjNA6joroawWgp8EhbBsHZA4RDyYW8dtJ2lkMKBcoQlqgWVPINiQGnaZCMIgs2ZAJCsAByYZBrJv4JPy59il1T5kWbd0RPWmsFhBIowZD", Auth: "P1tuP4nB0t", MessageUrl: "https://graph.facebook.com/v16.0/me/messages"}

func (a *AuthData) SendURL() string {
	return fmt.Sprintf("%v?access_token=%v", a.MessageUrl, a.Token)
}
