package flightcontroller_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/SpectralHiss/spacextest/flightcontroller"
	"github.com/SpectralHiss/spacextest/flightcontroller/persister/persisterfakes"
	"github.com/SpectralHiss/spacextest/flightcontroller/ticketdetails"

	"github.com/SpectralHiss/spacextest/flightcontroller/scheduler/schedulerfakes"
	"github.com/SpectralHiss/spacextest/flightcontroller/spacexquerier/spacexquerierfakes"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller suite")
}

var _ = Describe("Space flight booking", func() {

	var defaultTicketRequest = ticketdetails.TicketDetails{
		FirstName:     "Houssem",
		LastName:      "El Fekih",
		Gender:        "Male",
		Birthday:      time.Now(),
		LaunchpadID:   0,
		DestinationID: 0,
		LaunchDate:    time.Now()}

	Describe("creating new reservation", func() {
		Context("When creating a valid new reservation with a good destination mapping,no clash with SpaceX", func() {
			var FakePersister = &persisterfakes.FakeSaver{}
			var FakeSpaceXQuerier = &spacexquerierfakes.FakeSpaceXQuerier{}
			var FakeScheduler = &schedulerfakes.FakeFlightScheduler{}

			FakeScheduler.CheckScheduleReturns(true)

			FakeSpaceXQuerier.LaunchPossibleReturns(true)

			FakePersister.SaveReturns(nil)

			controller := NewFlightController(FakePersister, FakeSpaceXQuerier, FakeScheduler)

			It("succeeds in saving the reservation", func() {
				Expect(controller.Reserve(defaultTicketRequest)).To(BeNil())

				Expect(FakeSpaceXQuerier.LaunchPossibleCallCount()).To(Equal(1))
				Expect(FakeScheduler.CheckScheduleCallCount()).To(Equal(1))
				Expect(FakePersister.SaveCallCount()).To(Equal(1))

				Expect(FakePersister.SaveArgsForCall(0)).To(Equal(defaultTicketRequest))

			})
		})

	})

	Context("When creating a valid new reservation with a good destination mapping, clashing with SpaceX", func() {})

	Context("When creating a valid new reservation with a bad destination mapping", func() {})

})
