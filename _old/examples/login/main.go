package main

import contracts "github.com/mazuninky/blood-contracts-go"

const (
	emailRegex = `(?i)([A-Za-z0-9!#$%&'*+\/=?^_{|.}~-]+@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?)`
	phoneRegex = `(?:(?:\+?\d{1,3}[-.\s*]?)?(?:\(?\d{3}\)?[-.\s*]?)?\d{3}[-.\s*]?\d{4,6})|(?:(?:(?:\(\+?\d{2}\))|(?:\+?\d{2}))\s*\d{2}\s*\d{3}\s*\d{4})`
)

func main() {
	emailType := contracts.MustNewRegexType(emailRegex)
	phoneType := contracts.MustNewRegexType(phoneRegex)
	loginType := phoneType.Or(emailType)

	loginPack := loginType.Pack("test@gmail.com")
	if _, err := loginPack.Unpack(); err != nil {
		panic(err)
	}
}
