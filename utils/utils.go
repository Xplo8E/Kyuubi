package utils

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

func CheckDomainFormat(dr string, d string) bool {
	a := false
	re := regexp.MustCompile(dr)
	if re.MatchString(d) {
		a = true
	}
	return a
}

func GoDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	return os.Getenv(key)
}

func Filecreate(rs string, fp string) (string, *os.File) {
	trf, err := os.Create(fp)
	log.Println(len(rs))
	if err != nil {
		log.Println("Error in file creation", err)
	}
	fmt.Fprintln(trf, rs)

	return trf.Name(), trf
}
