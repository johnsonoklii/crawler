package collect

import (
	"bufio"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

type BaseFetch struct {
}

func (b BaseFetch) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)

	e := DeterminEncoding(bodyReader)

	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}
