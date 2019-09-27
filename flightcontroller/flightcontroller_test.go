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
		var FakeS *persisterfakes.FakeSaver
		var FakeSpace *spacexquerierfakes.FakeSpaceXQuerier
		var FakeSched *schedulerfakes.FakeFlightScheduler

		BeforeEach(func() {

			FakeS = &persisterfakes.FakeSaver{}

			FakeSpace = &spacexquerierfakes.FakeSpaceXQuerier{}
			FakeSched = &schedulerfakes.FakeFlightScheduler{}
		})

		Context("When creating a valid new reservation with a good destination mapping,no clash with SpaceX", func() {

			BeforeEach(func() {
				FakeSched.CheckScheduleReturns(true)
				FakeSpace.LaunchPossibleReturns(true)
				FakeS.SaveReturns(nil)
			})

			It("succeeds in saving the reservation", func() {

				controller := NewFlightController(FakeS, FakeSpace, FakeSched)
				Expect(controller.Reserve(defaultTicketRequest)).To(BeNil())

				Expect(FakeSpace.LaunchPossibleCallCount()).To(Equal(1))
				Expect(FakeSched.CheckScheduleCallCount()).To(Equal(1))
				Expect(FakeS.SaveCallCount()).To(Equal(1))

				Expect(FakeS.SaveArgsForCall(0)).To(Equal(defaultTicketRequest))

			})
		})

		Context("When creating a valid new reservation with a good destination mapping, clashing with SpaceX", func() {
			BeforeEach(func() {
				FakeSched.CheckScheduleReturns(true)
				FakeSpace.LaunchPossibleReturns(false)
				FakeS.SaveReturns(nil)
			})

			It("does not save the reservation, returns an error", func() {
				controller := NewFlightController(FakeS, FakeSpace, FakeSched)
				err := controller.Reserve(defaultTicketRequest)

				Expect(err).To(Not(BeNil()))
				Expect(err).To(MatchError("Bad request conflicting with spaceX launch"))
				Expect(FakeS.SaveCallCount()).To(Equal(0))
			})

		})

		Context("When creating a valid new reservation non conforming to schedule", func() {
			BeforeEach(func() {
				FakeSched.CheckScheduleReturns(false)
				FakeSpace.LaunchPossibleReturns(true)
				FakeS.SaveReturns(nil)
			})

			It("does not save the reservation, returns an error", func() {
				controller := NewFlightController(FakeS, FakeSpace, FakeSched)
				err := controller.Reserve(defaultTicketRequest)

				Expect(err).To(Not(BeNil()))
				Expect(err).To(MatchError("Bad request, not conforming to schedule"))
				Expect(FakeS.SaveCallCount()).To(Equal(0))
			})

		})

	})

})
