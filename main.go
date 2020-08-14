package main

import (
	"flag"
	"fmt"
	"log"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	var flags struct {
		Sk     string
		Type   string
		Label  string
		Issuer string
	}
	flag.StringVar(&flags.Sk, "s", "", "Auth secret key")
	flag.StringVar(&flags.Type, "t", "totp", "Format type: [totp|hotp]")
	flag.StringVar(&flags.Label, "l", "", "Label of account")
	flag.StringVar(&flags.Issuer, "i", "", "Issuer")

	flag.Parse()

	if flags.Sk == "" || flags.Label == "" {
		flag.Usage()
		return
	}

	if flags.Type != "totp" && flags.Type != "hotp" {
		log.Fatal("invalid value for key format type. it should be either totp or hotp")
		return
	}

	uri := fmt.Sprintf("otpauth://%s/%s:%s?secret=%s&issuer=%s", flags.Type, flags.Issuer, flags.Label, flags.Sk, flags.Issuer)
	q, err := qrcode.New(uri, qrcode.Medium)

	if err != nil {
		log.Fatal("error occurred while generating qrcode: ", err)
		return
	}

	art := q.ToSmallString(true)
	fmt.Println(art)
}
