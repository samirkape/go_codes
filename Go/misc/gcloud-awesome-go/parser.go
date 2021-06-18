package mybot

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Open file specified in  filename and return its handle
func FileHandle(filename string) *os.File {
	awsm, err := os.Open(filename)
	if err != nil {
		fmt.Println("cannot read file")
		os.Exit(-1)
	}
	return awsm
}

// ParsePkgs works on *bufio.Scanner which is pointer to location in a file.
// It  fills up Meta struct based on certain pattern in which the input markdown file is written.
func ParsePkgs(pkg_meta *Meta, title string, scanner *bufio.Scanner) {
	scanner.Scan()
	scanner.Scan()
	t := scanner.Text()
	if strings.Contains(t, "#") {
		var rpkg_meta Package
		ParsePkgs(&rpkg_meta.Details, t, scanner)
	} else {
		pkg_meta.SubTitle = t
	}
	scanner.Scan()
	n := 0
	var sub_pkgs = make([]string, 0)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			if n == 0 {
				continue
			}
			pkg_meta.Title = title
			pkg_meta.Count = n
			pkg_meta.Line.FullLink = sub_pkgs
			return
		}
		sub_pkgs = append(sub_pkgs, t)
		n++
	}
}

// GetSlice is a driver function that gets filehandler as an input,
// reads file line-by-line and based on some pattern it calls ParsePkgs
// for further parsing. It returns []Package which is a complete meta-data
// associated with the package.
func GetSlice(file *os.File) []Package {
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Count(text, "#") == 2 && strings.HasPrefix(text, "##") {
			var pkg_meta Package
			ParsePkgs(&pkg_meta.Details, text, scanner)
			if pkg_meta.Details.Count != 0 {
				pkgs = append(pkgs, pkg_meta)
			}
			total++
		}
	}
	return pkgs
}

// TrimString is a post-processing function that divides an input strings into,
// name -- package name
// url  -- package url
// info  -- a short info about the package
func TrimString(raw string) (string, string, string) {
	sre := regexp.MustCompile(`\[(.*?)\]`)
	rre := regexp.MustCompile(`\((.*?)\)`)
	_name := sre.FindAllString(raw, -1)
	_url := rre.FindAllString(raw, -1)
	if _name == nil || _url == nil {
		return "", "", ""
	}
	name := strings.Trim(_name[0], "[")
	name = strings.Trim(name, "]")
	url := strings.Trim(_url[0], "(")
	url = strings.Trim(url, ")")
	info := strings.Split(raw, "- ")
	if len(info) <= 1 {
		return name, url, ""
	}
	fmt.Println(info)
	return name, url, info[1]
}

// Split is a driver function for splitting the Line from []Package
// it calls TrimString for splitting and handles a creation and appending of
// a result into a LinkDetails struct.
func Split(pkgs []Package) {
	for i, e := range pkgs {
		var TmpLinks []SplitLink
		for _, q := range e.Details.Line.FullLink {
			name, url, info := TrimString(q)
			LD := SplitLink{Name: name, URL: url, Info: info}
			TmpLinks = append(TmpLinks, LD)
		}
		pkgs[i].Details.Line.LinkDetails = append(e.Details.Line.LinkDetails, TmpLinks...)
	}
}
