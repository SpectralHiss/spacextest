package main_test

import (
	//"net/http"

	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
	//. "github.com/SpectralHiss/TabeoTest"
)

var _ = Describe("Space flight booking", func() {
	Context("When the request is missing Key parameters", func() {
		Context("when it is missing launchpadID", func() {
			BeforeEach(func() {

			})

			It("should error", func() {

			})
		})
	})
	XDescribe("creating new reservation", func() {
		Context("When creating a valid new reservation with a good desination mapping, no clash with SpaceX", func() {
			//http.PostForm("http://localhost/book") // url.Values{"firstName": "Houssem",
			// 	"lastName":      "ElFekih",
			// 	"Gender":        "Male",
			// 	"Birthday":      time.Now().Format(time.RFC3339),
			// 	"LaunchPadId":   1,
			// 	"DestinationID": 3,
			// 	"LaunchDate":    time.Now.Add(time.Day * 3)},

		})

		Context("When creating a valid new reservation with a good desination mapping, clashing with SpaceX", func() {})

		Context("When creating a valid new reservation with a bad desination mapping", func() {})

	})

})
