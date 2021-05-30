package random

import "math/rand"

var Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var UpperAndNumbers = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var Uppers = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

type UUID [16]byte

func String(n int) string {

	s := make([]rune, n)
	for i := range s {
		s[i] = Letters[rand.Intn(len(Letters))]
	}
	return string(s)
}
func StringUpperAndNumber(n int) string {

	s := make([]rune, n)
	for i := range s {
		s[i] = UpperAndNumbers[rand.Intn(len(UpperAndNumbers))]
	}
	return string(s)
}

func StringUpper(n int) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = Uppers[rand.Intn(len(Uppers))]
	}
	return string(s)
}
