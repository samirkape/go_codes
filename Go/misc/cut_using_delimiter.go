// input -> file / string
// config -> delimiter1, delimiter2
// output -> file / console
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

const URL = "http://essential-go.programming-books.io/"

func getFile(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "href") {
			txtlines = append(txtlines, scanner.Text())
		}
	}

	file.Close()

	//for _, eachline := range txtlines {
	//	fmt.Println(eachline)
	//}
	return txtlines
}

func SplitAfter(s string) string {
	n := strings.SplitAfterN(s, "\"", 3)
	fmt.Println(n)
	return ""
}

func SliceUniqMap(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = URL + v
		j++
	}
	return s[:j]
}

func filter(str []string) (ret []string) {
	tokens := []string{"xlink", "toc-link", "use", "id", "arrow"}
	DontInsert := false
	var cut []string
	for _, e := range str {
		for _, c := range tokens {
			if strings.Contains(e, c) {
				DontInsert = true
				break
			}
		}
		if !DontInsert {
			cut = append(cut, e)
		}
		DontInsert = false
	}
	for _, e := range cut { // remove artifacts
		tmpret := strings.SplitAfterN(e, "\"", 3)
		if len(tmpret) == 3 {
			ret = append(ret, tmpret[1])
		}
	}
	ret = SliceUniqMap(ret) // remove dups
	return ret[2:66]
}

func RemoveQoute(str []string) []string {
	var newRet []string
	for _, s := range str {
		newRet = append(newRet, strings.Trim(s, "\""))
	}
	return newRet
}

func SplitString(str []string) []string {
	var tmpNewStr []string
	var Split []string
	for _, s := range str {
		tmpNewStr = strings.SplitAfter(s, "href=")

		Split = append(tmpNewStr, Split...) // to pass slice of strings to append(), use ... operator as it is a variadic function
	}
	return Split
}

func GetURLs(str []string) []string {
	//str[0] = `<use xlink:href="#arrow-not-expanded"></use></svg> <a class="toc-link" title="Conditional compilation with build tags" href="conditional-compilation-with-build-tags-d1980344374d45c082c914c2aafa50cf">Conditional compilation with build tags</a></div>`
	tmpNewStr := SplitString(str)
	newStr := filter(tmpNewStr)
	newStr = RemoveQoute(newStr)
	return newStr
}

func PrintWeb(links []string) {
	var build strings.Builder

	for _, e := range links {
		build.WriteString(e)
		build.WriteString("\n")
	}
	router := gin.New()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, build.String())
	})
	router.Run(":8080")
}

func main() {
	lines := getFile("/Users/sameer/Desktop/links.html")
	lines = GetURLs(lines)
	PrintWeb(lines)
}
