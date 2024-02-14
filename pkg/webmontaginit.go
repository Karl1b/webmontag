package webmontag

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type options struct {
	breaktimeS    int
	maxSearchDeep int
	maxGoRoutines int
}

var Options options

type webserveroptions struct {
	Key      string
	Port     string
	isWithDb bool
}

var Webserveroptions webserveroptions

func init() {
	// Checks for mandatory env variables
	godotenv.Load(".env")
	key := os.Getenv("KEY")
	if key == "" {
		log.Fatal("KEY is empty")
	}
	Webserveroptions.Key = key
	port := os.Getenv("PORT")
	if key == "" {
		log.Fatal("PORT is empty")
	}
	Webserveroptions.Port = port

	Options = options{
		breaktimeS:    30, // default values
		maxSearchDeep: 5,  // default values
		maxGoRoutines: 5,  // default values
	}

	// Checks for optional env variables

	if breaktimeS, err := strconv.Atoi(os.Getenv("BREAKTIMES")); err == nil {
		Options.breaktimeS = breaktimeS
	}

	if maxSearchDeep, err := strconv.Atoi(os.Getenv("MAXSEARCHDEEP")); err == nil {
		Options.maxSearchDeep = maxSearchDeep
	}

	if maxGoRoutines, err := strconv.Atoi(os.Getenv("MAXGOROUTINES")); err == nil {
		Options.maxGoRoutines = maxGoRoutines
	}

}
