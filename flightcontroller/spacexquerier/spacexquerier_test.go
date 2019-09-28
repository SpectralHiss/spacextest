package spacexquerier_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/SpectralHiss/spacextest/flightcontroller/spacexquerier"

	. "github.com/SpectralHiss/spacextest/flightcontroller/querytypes"
)

func TestSpcxApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scheduler suite")
}

var _ = Describe("SpaceX api integration", func() {

	var tServer *httptest.Server
	var spaceXQ spacexquerier.SpaceXQueryImpl

	BeforeSuite(func() {

		tServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.RequestURI == "/v3/landpads" {
				bytebuff, _ := ioutil.ReadFile("./json_responses/getall.json")
				w.Header().Set("Content-Type", "text/json")
				w.Write(bytebuff)
			}
			if r.RequestURI == "/v3/launches/upcoming" {
				bytebuff, _ := ioutil.ReadFile("./json_responses/upcoming.json")
				w.Header().Set("Content-Type", "text/json")
				w.Write(bytebuff)
			}
		}))
		spaceXQ.ApiUrl = tServer.URL
	})

	AfterSuite(func() {
		tServer.Close()
	})
	Describe("GetLaunchIds",func(){
		Context("When GetLaunchIds is called", func() {

			It("Should return the right launchIds", func() {
				
				expectedLaunchIds := []LaunchPadID{
					"vafb_slc_4e",
					"kwajalein_atoll",
					"ccafs_slc_40",
					"stls",
					"ksc_lc_39a",
					"vafb_slc_3w",
				}
	
				Expect(spaceXQ.GetLaunchIds()).To(Equal(expectedLaunchIds))
			})
		})
	
		//TODO make error ret not doen for interest of time
		Context("If spaceX unavailable, it returns an empty list", func() {
		})
	
	})


	Describe("Check if conflicts with spacexlaunch",func(){
		
		Context("when called with a conflicting date for a launchpadID",func(){
			It("Should return false",func(){
				aDateOfLaunch , err := time.Parse(time.RFC3339, "2019-02-22T01:45:00.000Z")
				Expect(err).To(BeNil())
				ret := spaceXQ.LaunchPossible(LaunchPadID("ccafs_slc_40"),aDateOfLaunch)
				Expect(ret).To(BeFalse())
			})

		})

		Context("when called with a non-conflicting date for a launchpadID",func(){
			It("Should return false",func(){
				aDateOfLaunch , err := time.Parse(time.RFC3339, "2049-02-22T01:45:00.000Z")
				Expect(err).To(BeNil())
				ret := spaceXQ.LaunchPossible(LaunchPadID("ksc_lc_39a"),aDateOfLaunch)
				Expect(ret).To(BeTrue())
			})

		})
	})




		  
	})

