package analyzers

import (
	"fmt"
	"github.com/ericr/solanalyzer/sources"
	"io/ioutil"
	"net/http"
)

// Analyzer is an interface for all analyzers.
type Analyzer interface {
	Name() string
	ID() string
	Execute(*sources.Source) ([]*Issue, error)
}

func getUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte(""), fmt.Errorf("Got non-200 HTTP response from %s", url)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return bodyBytes, nil
}
