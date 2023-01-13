package main

import (
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()

	m.SetHeader("From", "applicatieopdracht@gmail.com")
	m.SetHeader("To", "zemaetibark@hotmail.com")
	m.SetHeader("Subject", "Even testen voor applicatie ;)")
	m.SetBody("text/plain", "de body testen")

	d := gomail.NewDialer("smtp.gmail.com", 587, "applicatieopdracht@gmail.com", "wgftnptcxmownaus")

	d.DialAndSend(m)
	fmt.Println("mail gestuurd!")

	//stuur mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
