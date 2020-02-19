package main

import (
	"flag"
	"log"
	"os"
	"path"

	"github.com/t94j0/czds"
)

var (
	// flags
	username = flag.String("username", "", "username to authenticate with")
	password = flag.String("password", "", "password to authenticate with")
	outFile  = flag.String("file", "report.csv", "filename to save report to, '-' for stdout")
	verbose  = flag.Bool("verbose", false, "enable verbose logging")

	client *czds.Client
)

func v(format string, v ...interface{}) {
	if *verbose {
		log.Printf(format, v...)
	}
}

func checkFlags() {
	flag.Parse()
	flagError := false
	if len(*username) == 0 {
		log.Printf("must pass username")
		flagError = true
	}
	if len(*password) == 0 {
		log.Printf("must pass password")
		flagError = true
	}
	if flagError {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	checkFlags()

	client = czds.NewClient(*username, *password)

	// validate credentials
	v("Authenticating to %s", client.AuthURL)
	err := client.Authenticate()
	if err != nil {
		log.Fatal(err)
	}

	out := os.Stdout
	if *outFile != "-" {
		v("Saving to %s", *outFile)
		dir := path.Dir(*outFile)
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Create(*outFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		out = file
	} else {
		v("Printing to StdOut")
	}

	// CSV report to out
	err = client.DownloadAllRequests(out)
	if err != nil {
		log.Fatal(err)
	}
}
