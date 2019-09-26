package flightcontroller_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"

	. "github.com/SpectralHiss/spacextest/flightcontroller"
	"github.com/SpectralHiss/spacextest/flightcontroller/persister/persisterfakes"
	"github.com/SpectralHiss/spacextest/flightcontroller/scheduler"
	"github.com/SpectralHiss/spacextest/flightcontroller/ticketdetails"

	"github.com/SpectralHiss/spacextest/flightcontroller/scheduler/schedulerfakes"
	"github.com/SpectralHiss/spacextest/flightcontroller/spacexquerier/spacexquerierfakes"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller suite")
}


func rotatingPerm7Of8() []int {
	// swap 8 times random idxes 
	var initialConfig [8]int
	for i := 0 ; i < 8 ; i++ {
		initialConfig[i] = i
	}

	view := initialConfig[:]
	k:= math.Random() % 8
	for ;  k >= 0 ; k-- {
		i:= math.Random() % 8
		j:= math.Random() % 8
		temp = view[i]
		view[i] = view[j]
		view[j] = temp
	}

	return view[:7]
	
}

/*	map[int]map[scheduler.Day]scheduler.DestinationID{
	0: map[scheduler.Day]scheduler.DestinationID{
		0: 0,
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
		6: 6,
	},
}*/


func GenerateFakeSchedule() scheduler.Schedule {
	// assuming launchpads are serial
	numLaunchPads := math.Random()
	numDestinations := 8

	scheduleMap := make(map[int]map[scheduler.Day]Scheduler)
	
	for i := 0 ; i < numLaunchPads; i++ {
		scheduleMap[i] := make(map[scheduler.Day]Scheduler)
		for idx, val := range rotatingPerm7Of8() {
			scheduleMap[scheduler.Day(idx)] = scheduler.DestinationId(val)
		}

	}
 
	return scheduleMap
}

func giveValidTicketDetails(scheduleMap scheduler.Schedule) ticketdetails.TicketDetails {
	number := math.Random()
	
	//TODO CONTINUE HERE
	var defaultTicketRequest = ticketdetails.TicketDetails{
		FirstName:     "Houssem",
		LastName:      "El Fekih",
		Gender:        "Male",
		Birthday:      time.Now(),
		LaunchpadID:   0,
		DestinationID: 0,
		LaunchDate:    time.Now()}

}

var _ = Describe("Space flight booking", func() {

	Describe("creating new reservation", func() {
		Context("When creating a valid new reservation with a good destination mapping,no clash with SpaceX", func() {
			var FakePersister = &persisterfakes.FakeSaver{}
			var FakeSpaceXQuerier = &spacexquerierfakes.FakeSpaceXQuerier{}
			var FakeScheduler = &schedulerfakes.FakeScheduler{}

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

			FakeSpaceXQuerier.LaunchPossibleReturns(true)

			FakePersister.SaveReturns(nil)

			controller := NewFlightController(FakePersister, FakeSpaceXQuerier, FakeScheduler)

			It("succeeds in saving the reservation", func() {
				Expect(controller.Reserve(defaultTicketRequest)).To(Not(BeNil()))

				Expect(FakeSpaceXQuerier.LaunchPossibleCallCount()).To(Equal(1))
				//Expect(FakeScheduler.GenerateScheduleCallCount()).To(Equal(1))
				//Expect(FakePersister.SaveCallCount()).To(Equal(1))

				//Expect(FakePersister.SaveArgsForCall(0)).To(Equal(defaultTicketRequest))

			})
		})

	})

	Context("When creating a valid new reservation with a good destination mapping, clashing with SpaceX", func() {})

	Context("When creating a valid new reservation with a bad destination mapping", func() {})

})
