package main_test

import (
	//"net/http"
	"github.com/SpectralHiss/spacextest/flightcontroller/spacexquerier"


	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
	//. "github.com/SpectralHiss/TabeoTest"
)

var _ = Describe("Space flight booking", func() {

	Context("When a valid spaceflight booking request is incoming",func(){
		// Cannot write this test blackbox 
		// because the scheduler needs to somehow publish its random rotating schedule
		

	})

	// TODO: not done in interest of time
	Context("When the request is missing Key parameters", func() {
		Context("when it is malformed etc..", func() {
			
			BeforeEach(func() {

			})

			It("should error", func() {

			})
		})
	})
})
