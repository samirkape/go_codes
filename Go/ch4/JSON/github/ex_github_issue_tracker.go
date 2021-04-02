// Sort github issues as per the age

package main

import (
	"fmt"
	"log"
	"samgithub"
	"strconv"
	"strings"
	"time"
)

func main() {
	issue := []string{"repo:golang/go is:open json decoder"}
	// result, err := github.SearchIssues(os.Args[1:])
	result, err := samgithub.SearchIssues(issue)
	if err != nil {
		log.Fatal(err)
	}
	sortbyage(result)
}

func sortbyage(result *samgithub.IssuesSearchResult) {

	year, month, _ := nowDate()

	fmt.Printf("%d issues sorted by age:\n", result.TotalCount)
	fmt.Println("More than a year ago")
	for _, item := range result.Items {
		createdAt := fmt.Sprintf("%.10s", item.CreatedAt)
		gyear, _, _ := gitDate(createdAt)
		if year > gyear {
			fmt.Printf("#%-5d %9.9s %.55s %.10s\n",
				item.Number, item.User.Login, item.Title, createdAt)
		}
	}

	fmt.Println("Less than a month ago")
	for _, item := range result.Items {
		createdAt := fmt.Sprintf("%.10s", item.CreatedAt)
		gyear, gmonth, _ := gitDate(createdAt)
		if year == gyear && (month-gmonth <= 1) {
			fmt.Printf("#%-5d %9.9s %.55s %.10s\n",
				item.Number, item.User.Login, item.Title, createdAt)
		}
	}

	fmt.Println("Less than a year ago")
	for _, item := range result.Items {
		createdAt := fmt.Sprintf("%.10s", item.CreatedAt)
		gyear, _, _ := gitDate(createdAt)
		if year-gyear <= 1 {
			fmt.Printf("#%-5d %9.9s %.55s %.10s\n",
				item.Number, item.User.Login, item.Title, createdAt)
		}
	}
}

func gitDate(fullDate string) (int, int, int) {
	dt := strings.Split(fullDate, "-")
	y, _ := strconv.Atoi(dt[0])
	m, _ := strconv.Atoi(dt[1])
	d, _ := strconv.Atoi(dt[2])
	return y, m, d
}

func nowDate() (int, int, int) {
	year, tmonth, day := time.Now().Local().Date()
	month := int(tmonth)
	return year, month, day
}
