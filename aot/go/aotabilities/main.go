package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "go.opencensus.io/tag"
	_ "golang.org/x/crypto/acme"

	"go.verytable.online/aot/aot/go/eren"
	"go.verytable.online/aot/aot/go/mikasa"
)

var abilities = map[string][]string{
	"Eren":   eren.Abilities(),
	"Mikasa": mikasa.Abilities(),
}

func main() {
	name := flag.String("name", "", "AoT character name")
	flag.Parse()

	s, ok := abilities[*name]
	if !ok {
		_, _ = fmt.Fprintf(os.Stderr, "name %q not found", *name)
		os.Exit(1)
	}

	fmt.Printf("%s's abilities: %v", *name, s)
}
