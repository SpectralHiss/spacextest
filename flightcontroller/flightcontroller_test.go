package flightcontroller_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"

	. "github.com/SpectralHiss/spacextest/flightcontroller"
	"github.com/SpectralHiss/spacextest/flightcontroller/persister/persisterfakes"
	"github.com/SpectralHiss/spacextest/flightcontroller/scheduler"

	"github.com/SpectralHiss/spacextest/flightcontroller/scheduler/schedulerfakes"
	"github.com/SpectralHiss/spacextest/flightcontroller/spacexquerier/spacexquerierfakes"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller suite")
}

var _ = Describe("Space flight booking", func() {

	var defaultTicketRequest = TicketDetails{
		FirstName:     "Houssem",
		LastName:      "El Fekih",
		Gender:        "Male",
		Birthday:      time.Now(),
		LaunchpadID:   0,
		DestinationID: 0,
		LaunchDate:    time.Now()}

	Describe("creating new reservation", func() {
		Context("When creating a valid new reservation with a good destination mapping,no clash with SpaceX", func() {
			var FakePersister persisterfakes.FakeSaver
			var FakeSpaceXQuerier spacexquerierfakes.FakeSpaceXQuerier
			var FakeScheduler schedulerfakes.Scheduler

			FakeScheduler.GenerateScheduleReturns(map[int]map[scheduler.Day]scheduler.DestinationID{
				0: map[scheduler.Day]scheduler.DestinationID{
					0: 0,
					1: 1,
					2: 2,
					3: 3,
					4: 4,
					5: 5,
					6: 6,
				},
			})
		})

		FakeSpaceXQuerier.LaunchPossibleReturns(true)

		fakePersister.SaveReturns(nil)

		controller = NewFlightController(FakePersister, FakeSpaceXQuerier, FakeScheduler)

		It("succeeds in saving the reservation", func() {
			Expect(FakeSpaceXQuerier.LaunchPossibleCallCount()).To(Equal(1))
			Expect(FakeScheduler.GenerateScheduleCallCount()).To(Equal(1))
			Expect(FakePersister.SaveCallCount()).To(Equal(1))

			Expect(FakePersister.SaveArgsForCall(0)).To(Equal(defaultTicketRequest))

			Expect(controller.Reserve(defaultTicketRequest)).ToBeTrue()
		})

	})

	Context("When creating a valid new reservation with a good destination mapping, clashing with SpaceX", func() {})

	Context("When creating a valid new reservation with a bad destination mapping", func() {})

})
