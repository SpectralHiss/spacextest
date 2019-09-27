package spacexquerier_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scheduler suite")
}

var _ = Describe("SpaceX api integration", func() {
	BeforeEach(func() {
		newSpaceXServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.RequestURI == "/getall" {
				bytebuff, _ := ioutil.ReadFile("fixtures/getall.json")
				w.Header().Set("Content-Type", "text/json")
				w.Write(bytebuff)
			}
		}))
	})
})
