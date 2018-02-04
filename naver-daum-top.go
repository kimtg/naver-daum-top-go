package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func slurp(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	res.Body.Close()
	return string(body), nil
}

// return all specific groups
func reGroups(re *regexp.Regexp, text string, group int) []string {
	result := []string{}
	found := re.FindAllStringSubmatch(text, -1)
	for _, v := range found {
		result = append(result, v[group])
	}
	return result
}

func printReGroupsSlurp(title, url, re string, group int) {
	s, err := slurp(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(title, strings.Join(reGroups(regexp.MustCompile(re), s, group), ", "))
	}
}

func main() {
	const interval = 5
	fmt.Println("Refreshes every", interval, "minutes.")
	for {
		fmt.Println(time.Now())
		printReGroupsSlurp("Naver:", "https://www.naver.com", "<span class=\"ah_k\">(.+?)</span>\n</a>\n</li>", 1)
		printReGroupsSlurp("Daum:", "http://www.daum.net", "class=\"link_issue\" tabindex.*?>(.+?)</a>", 1)
		time.Sleep(interval * time.Minute)
	}
}
