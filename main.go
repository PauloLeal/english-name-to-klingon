package main

import (
	"fmt"
	"os"
	"flag"
	"github.com/PauloLeal/english-name-to-klingon/stapiminiclient"
	"github.com/PauloLeal/english-name-to-klingon/translator"
	"strings"
)

func main() {
	flag.Usage = func() { //since I'm not using flags, only command line arguments, I'm implementing usage method
		fmt.Printf("Usage: %s <name_in_english>\n", os.Args[0])
	}

	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	name := strings.Join(os.Args[1:], " ")

	klingon, err := translator.EnglishToKlingon(name)

	if err != nil {
		fmt.Println(err)
		os.Exit(1) //not using panic just because too verbose output for this error
	}

	fmt.Println(klingon)
	fmt.Println(stapiminiclient.GetCharacterSpecie(name))
}
