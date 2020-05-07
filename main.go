package main

import (
	"fmt"
	prt "github.com/alexbrainman/printer"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	_ = godotenv.Load()
	var re = regexp.MustCompile(os.Getenv("PRINTER"))
	printers, err := prt.ReadNames()
	if err != nil {
		panic(err)
	}
	printer := ""
	for _, p := range printers {
		if re.MatchString(p) {
			printer = p
		}
	}
	if printer != "" {
		log.Println("found ", printer)
		content, err := ioutil.ReadFile(os.Getenv("FILE"))
		if err != nil {
			panic(err)
		}
		printContent("raw", content)
	} else {
		log.Println("printer not found")
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func printContent(datatype string, content []byte) {
	name, err := prt.Default() // returns name of Default Printer as string
	checkErr(err)
	fmt.Println(name)

	p, err := prt.Open(name) // Opens the named printer and returns a *Printer
	checkErr(err)

	err = p.StartDocument("test", datatype)
	checkErr(err)

	err = p.StartPage() // begin a new page
	checkErr(err)

	n, err := p.Write(content) // Send some text to the printer
	checkErr(err)
	fmt.Println("Num of bytes written to printer:", n)

	err = p.EndPage() // end of page
	checkErr(err)

	err = p.EndDocument() // end of document
	checkErr(err)

	err = p.Close() // close the resource
	checkErr(err)

}
