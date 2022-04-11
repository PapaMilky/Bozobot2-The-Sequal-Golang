package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func R34api(terms string, page string) []byte {
	url := "https://api.rule34.xxx/index.php?page=dapi&s=post&q=index&json=1&limit=2&tags=" + strings.Replace(terms, " ", "_", -1) + "&pid=" + page

	client := http.Client{
		Timeout: time.Second * 6,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "r34-fetcher")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}

func RandomUrban() []byte {

	url := "https://api.urbandictionary.com/v0/random"

	client := http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "urban-fetcher")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}

func UrbanLookup(lookup string) []byte {

	url := "https://api.urbandictionary.com/v0/define?term=" + strings.Replace(lookup, " ", "_", -1)

	client := http.Client{
		Timeout: time.Second * 3,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "urban-fetcher")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}
