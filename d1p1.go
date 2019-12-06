package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func inputFileCreator(dayCode string, inputString string) bool {
	fp, err := os.Create(fmt.Sprintf("./input/d%s", dayCode))
	if err != nil {
		return false
	}

	defer fp.Close()

	_, err = fp.WriteString(inputString)
	if err != nil {
		return false
	}

	return true
}

func inputGetter(dayCode string) string {
	godotenv.Load()
	myCookie := os.Getenv("COOKIE")

	inputURL := fmt.Sprintf("https://adventofcode.com/2019/day/%s/input", dayCode)
	req, _ := http.NewRequest("GET", inputURL, nil)

	req.Header.Add("authority", "adventofcode.com")
	req.Header.Add("cache-control", "max-age=0,no-cache")
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("referer", inputURL)
	req.Header.Add("cookie", fmt.Sprintf("session=%s", myCookie))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func solutionDriver(dayCode string) {
	fp, _ := os.Open(fmt.Sprintf("./input/d%s", dayCode))
	defer fp.Close()

	totalSum := 0

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		inputPayload, _ := strconv.Atoi(scanner.Text())
		output := inputPayload/3 - 2
		totalSum += output
	}

	fmt.Println(totalSum)
}

func main() {
	dayCode := os.Args[1]

	inputFileStatus := inputFileCreator(dayCode, inputGetter(dayCode))
	if inputFileStatus != true {
		return
	}

	solutionDriver(dayCode)
}
