package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const (
	kleineLetters = "abcdefghijklmnopqrstuvwxyz"
	HoofdLetters  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Cijfers       = "0123456789"
	Tekens        = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
	AlleKarakters = kleineLetters + HoofdLetters + Cijfers + Tekens
	//minimale wachtwoodlengte voor de leesbaarheid en configureerbaarheid.
	//opties tekens en cijfers toevoegen.
)

func main() {
	fmt.Print("Kies de lengte van het wachtwoord: ")
	var lengte int
	fmt.Scanln(&lengte)
	//input gescant

	combinatie := AlleKarakters
	var minLengte int = 8
	if lengte >= minLengte {

		combinatie = kleineLetters + Voorkeur()

		Password := GenerateRandomString(lengte, combinatie)

		fmt.Println(Password)
	} else {
		fmt.Print("Je wachtwoord moet minimaal 8 tekens zijn.")
	}
}

func Voorkeur() string {
	var (
		Hoofdletters string
		Nummers      string
		Symbolen     string
	)

	combinatie := strings.Builder{}

	fmt.Println("Wilt u hoofdletters? [y|n] ")
	fmt.Scanln(&Hoofdletters)

	fmt.Println("Wilt u nummers? [y|n] ")
	fmt.Scanln(&Nummers)

	fmt.Println("Wilt u symbolen? [y|n] ")
	fmt.Scanln(&Symbolen)

	strings.ToLower(Hoofdletters)
	strings.ToLower(Nummers)
	strings.ToLower(Symbolen)

	if Hoofdletters == "y" {
		combinatie.WriteString(string(HoofdLetters))
	}

	if Nummers == "y" {
		combinatie.WriteString(string(Cijfers))
	}

	if Symbolen == "y" {
		combinatie.WriteString(string(Tekens))
	}
	return combinatie.String()
}

func GenerateRandomString(WachtwoordLengte int, combinatie string) string {
	ret := make([]byte, WachtwoordLengte)
	for i := 0; i < WachtwoordLengte; i++ {
		random, err := rand.Int(rand.Reader, big.NewInt(int64(len(combinatie))))
		if err != nil {
			fmt.Println(err)
		}
		ret[i] = combinatie[random.Int64()]
	}

	return string(ret)
}
