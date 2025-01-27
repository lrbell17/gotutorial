package main

import (
	"flag"
	"fmt"
	"math/rand"
)

var (
	passLen  = 16
	numPw    = 1
	lowers   = []rune("abcdefghijklmnopqrstuvwxyz")
	uppers   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers  = []rune("0123456789")
	specials = []rune("-=~!@#$%^&*()_+[]\\{}|;':\",./<>?")
)

func main() {

	if passLen < 8 || numPw < 1 {
		flag.Usage()
		return
	}

	for i := range numPw {
		passChan := randomPassProducer(passLen)
		password := randomPassConsumer(passChan)
		fmt.Printf("Password %v: %v\n", i+1, password)
	}
}

func randomPassProducer(l int) (pass chan rune) {
	pass = make(chan rune, l)
	go func() {
		for i := 0; i < l; i++ {
			randLowerIdx := rand.Intn(len(lowers))
			randUpperIdx := rand.Intn(len(uppers))
			randNumberIdx := rand.Intn(len(numbers))
			randSpecialIdx := rand.Intn(len(specials))

			select {
			case pass <- lowers[randLowerIdx]:
			case pass <- lowers[randLowerIdx]:
			case pass <- lowers[randLowerIdx]:
			case pass <- lowers[randLowerIdx]:

			case pass <- uppers[randUpperIdx]:
			case pass <- uppers[randUpperIdx]:
			case pass <- uppers[randUpperIdx]:

			case pass <- numbers[randNumberIdx]:
			case pass <- numbers[randNumberIdx]:

			case pass <- specials[randSpecialIdx]:
			}
		}
		close(pass)
	}()
	return
}

func randomPassConsumer(in chan rune) (password string) {
	for v := range in {
		password += string(v)
	}
	return
}

func init() {
	flag.IntVar(&passLen, "l", passLen, "Password length (min 8)")
	flag.IntVar(&numPw, "n", numPw, "Number of passwords (min 1)")
	flag.Parse()
}
