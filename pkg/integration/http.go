// +build integration

package integration

import (
	"fmt"
	"github.com/gavv/httpexpect"
	"net/http"
	"sync"
	"testing"
	"time"
)

const (
	baseUrl = "http://localhost%s"
	schema  = `{
		"docs": "array"
	}`
)

type Client struct {
	baseUrl string
	e       *httpexpect.Expect
	t       *testing.T
}

func newClient(t *testing.T, port string) *Client {
	http := httpexpect.WithConfig(httpexpect.Config{
		BaseURL: getBaseUrl(port),
		Client: &http.Client{
			Jar:     httpexpect.NewJar(),
			Timeout: time.Second * 30,
		},
		// use fatal failures
		Reporter: httpexpect.NewAssertReporter(t),
		// use verbose logging
		Printers: []httpexpect.Printer{
		},
	})
	return &Client{
		baseUrl: getBaseUrl(port),
		e:       http,
		t:       t,
	}
}

func (c *Client) testGet(url string) {
	request := c.e.GET(url).WithURL(c.baseUrl)
	response := request.Expect()
	//TODO create a logic to validate schemas
	//response.JSON().Schema(schema)
	if response.Raw().StatusCode != 200 {
		fmt.Printf("\n%s - %s\n", response.Raw().Status, url)
	}
	response.Status(http.StatusOK)
}

func (c *Client) doTests(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	if isExcluded(path) {
		return
	}
	url := addFixtures(path)
	c.testGet(url)
}

func getBaseUrl(port string) string {
	return fmt.Sprintf(baseUrl, port)
}
