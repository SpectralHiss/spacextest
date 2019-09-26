package main_test

import (
	"net/http"
	"net/url"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/SpectralHiss/TabeoTest"
)

var _ = Describe("Space flight booking", func() {

	BeforeEach(func(){
		server = NewSpaceServer(database.NewModel())
	})

	Describe("creating new reservation",func(){
		Context("When creating a valid new reservation with a good desination mapping, no clash with SpaceX", func() {
			http.PostForm("http://localhost/book",
				url.Values{"firstName": "Houssem",
					"lastName": "ElFekih",
					"Gender":   "Male",
					"Birthday": time.Now().Format(time.RFC3339)},
					"LaunchPadId": 1,
					"DestinationID": 3,
					"LaunchDate": time.Now.Add(time.Day * 3)
				)
		})


		Context("When creating a valid new reservation with a good desination mapping, clashing with SpaceX", func() {}

		Context("When creating a valid new reservation with a bad desination mapping", func() {}


	})

})
