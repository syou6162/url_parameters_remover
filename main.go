package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

type strSlice []string

func (ss *strSlice) String() string {
	return fmt.Sprintf("%s", *ss)
}

func (ss *strSlice) Set(v string) error {
	*ss = append(*ss, v)
	return nil
}

func removeUnnecessaryParams(origUrl string, parametersToBeRemoved []string) (string, error) {
	u, err := url.Parse(origUrl)
	if err != nil {
		return origUrl, err
	}

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return origUrl, err
	}

	for _, p := range parametersToBeRemoved {
		q.Del(p)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}

var parametersToBeRemoved strSlice

func main() {
	flag.Var(&parametersToBeRemoved, "p", "URL parameter to be removed")
	flag.Parse()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		url, err := removeUnnecessaryParams(s.Text(), parametersToBeRemoved)
		if err != nil {
			continue
		}
		fmt.Println(url)
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
}
