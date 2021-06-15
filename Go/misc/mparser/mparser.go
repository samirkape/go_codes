package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const FILE = "./awesome.md"

type Package struct {
	Details Meta
}

var pkgs = make([]Package, 0)

type Meta struct {
	Name     string
	Type     []string
	SubTitle string
	Count    int
	SplitLink
}

type SplitLink struct {
	Name string
	URL  string
	Info string
}

func FileHandle(filename string) *os.File {
	awsm, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("cannot read file")
		os.Exit(-1)
	}
	return awsm
}

func ParsePkgs(pkg_meta *Meta, scanner *bufio.Scanner) {
	scanner.Scan()
	scanner.Scan()
	t := scanner.Text()
	if strings.Contains(t, "#") {
		var rpkg_meta Package
		rpkg_meta.Details.Name = t
		ParsePkgs(&rpkg_meta.Details, scanner)
	} else {
		pkg_meta.SubTitle = t
	}
	scanner.Scan()
	n := 0
	var sub_pkgs = make([]string, 0)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			return
		}
		sub_pkgs = append(sub_pkgs, t)
		n++
	}
	pkg_meta.Count = n
	pkg_meta.Type = sub_pkgs
}

func GetSlice(file *os.File) []Package {
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Count(text, "#") == 2 && strings.HasPrefix(text, "##") {
			var pkg_meta Package
			pkg_meta.Details.Name = text
			ParsePkgs(&pkg_meta.Details, scanner)
			pkgs = append(pkgs, pkg_meta)
			total++
		}
	}
	return pkgs
}

func RgexTrim(raw string) (string, string, string) {
	sre := regexp.MustCompile(`\[(.*?)\]`)
	rre := regexp.MustCompile(`\((.*?)\)`)
	_name := sre.FindAllString(raw, -1)
	_url := rre.FindAllString(raw, -1)
	name := strings.Trim(_name[0], "[")
	name = strings.Trim(name, "]")
	url := strings.Trim(_url[0], "(")
	url = strings.Trim(url, ")")
	info := strings.Split(raw, "- ")
	fmt.Println(info)
	return name, url, info[1]
}

func Split(pkgs []Package) {
	for i, _ := range pkgs {
		for _, q := range pkgs[i].Details.Type {
			name, url, info := RgexTrim(q)
			pkgs[i].Details.SplitLink.Name = name
			pkgs[i].Details.SplitLink.URL = url
			pkgs[i].Details.SplitLink.Info = info
		}
	}
}

func main() {
	file := FileHandle(FILE)
	final := GetSlice(file)
	Split(final)
	defer file.Close()
}
