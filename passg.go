// This program generates secure passwords with different options
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"crypto/rand"
	"math/big"
)

const (
	maxLength            = 512
	minLength            = 8
	maxNumberOfPasswords = 16384
	minNumberOfPasswords = 1
	letters              = "abcdefghijklmnopqrstuvwxyz"
	digits               = "0123456789"
	special              = "=!@#$%"
)

// Command-line flags
var (
	fLength = flag.Int("len", 8, "Password length")
	fCount  = flag.Int("count", 1, "Number of passwords to generate")
	fAll    = flag.Bool("all", false, "Use all English letters (upper/lower), digits and special characters")
	fUpper  = flag.Bool("up", false, "Use upper case letters")
	fLower  = flag.Bool("low", false, "Use lower case letters")
	fDigits = flag.Bool("digits", false, "Use digits")
	fSpec   = flag.Bool("spec", false, "Use special characters")
)

func main() {
	flag.Usage = usage
	flag.Parse()
	ok := checkForBounds(*fLength, *fCount)
	if !ok {
		fmt.Fprintf(os.Stderr, "Violation of bounds:\n")
		fmt.Fprintf(os.Stderr, "%d ≤ password length ≤ %d\n", minLength, maxLength)
		fmt.Fprintf(os.Stderr, "%d ≤ password count ≤ %d\n", minNumberOfPasswords, maxNumberOfPasswords)
		os.Exit(1)
	}
	symbols := giveSymbolSet(*fAll, *fLower, *fUpper, *fDigits, *fSpec)
	if len(symbols) == 0 {
		usage()
		os.Exit(1)
	}
	pass := generatePasswords(symbols, *fLength, *fCount)
	for _, p := range pass {
		fmt.Println(p)
	}
}

// Checks if given length and number of passwords are between Min and Max
func checkForBounds(lngth, cnt int) bool {
	if lngth < minLength || lngth > maxLength ||
		cnt < 1 || cnt > maxNumberOfPasswords {
		return false
	}
	return true
}

// Returns symbol set based on user's choice of flags
func giveSymbolSet(al, lo, up, di, sp bool) string {
	if (al) {
		return letters + strings.ToUpper(letters) + digits + special
	}
	var symbols string
	if lo {
		 symbols += letters
	}
	if up {
		symbols += strings.ToUpper(letters)
	}
	if di {
		symbols += digits
	}
	if sp {
		symbols += special
	}
	return symbols
}

// Generates n=count number of passwords with specified length
func generatePasswords(symbs string, length, count int) []string {
	passwords := []string{}
	for i := 0; i < count; i++ {
		pass := generatePassword(symbs, length)
		passwords = append(passwords, pass)
	}
	return passwords
}

// Generates one password with specified length
func generatePassword(symbs string, length int) string {
	var password string;
	for i := 0; i < length; i++ {
		bIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(symbs))))
		if err != nil {
			panic(err)
		}
		index := bIndex.Int64()
		password += string(symbs[index])
	}
	return password
}

// Custom usage message
func usage() {
	fmt.Fprintf(os.Stderr, 
`
Usage: %s [-len 8] [-count 1] [-all]
Generate secure passwords with given length and symbols
 
Options:
  -len      Password length
  -count    Number of passwords to generate
  -all      Use all English letters (upper/lower), digits and special characters
  -up       Use upper case letters
  -low      Use lower case letters
  -digits   Use digits
  -spec     Use special characters
	
  -help     Print this help message

`, os.Args[0])
}













