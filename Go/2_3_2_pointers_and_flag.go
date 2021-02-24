//This program will use flag package to parse CLA 

package main

import(
	"fmt"
	"flag"
	"strings"
)

var n = flag.Bool( "n", false, "Omit trailing newline" )
var sep = flag.String( "s", " ", "Separator" )

func main() {
	
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}

}