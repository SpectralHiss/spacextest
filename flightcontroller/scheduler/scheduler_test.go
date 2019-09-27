package scheduler_test

import (
	. "github.com/SpectralHiss/spacextest/flightcontroller/scheduler"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test suite 2", func() {
	It("Should SUCCESS", func() {})
})

// func GenerateFakeSchedule() Schedule {
// 	// assuming launchpads are serial
// 	numLaunchPads := math.Random()
// 	numDestinations := 8

// 	scheduleMap := make(map[int]map[Day]DestinationID)

// 	for i := 0; i < numLaunchPads; i++ {
// 		scheduleMap[i] = make(map[Day]DestinationID)
// 		for idx, val := range rotatingPerm7Of8() {
// 			scheduleMap[Day(idx)] = DestinationId(val)
// 		}

// 	}

// 	return scheduleMap
// }

// // GenerateScheduleReturns(map[int]map[scheduler.Day]scheduler.DestinationID{
// // 	0: map[scheduler.Day]scheduler.DestinationID{
// // 		0: 0,
// // 		1: 1,
// // 		2: 2,
// // 		3: 3,
// // 		4: 4,
// // 		5: 5,
// // 		6: 6,
// // 	},
// // })

// func rotatingPerm7Of8() []int {
// 	// swap 8 times random idxes
// 	var initialConfig [8]int
// 	for i := 0; i < 8; i++ {
// 		initialConfig[i] = i
// 	}

// 	view := initialConfig[:]
// 	k := rand.Int() % 8
// 	for ; k >= 0; k-- {
// 		i := rand.Int() % 8
// 		j := rand.Int() % 8
// 		temp := view[i]
// 		view[i] = view[j]
// 		view[j] = temp
// 	}

// 	return view[:7]

// }

// func giveValidTicketDetails(scheduleMap Schedule) ticketdetails.TicketDetails {
// 	number := rand.Int()

// 	//TODO CONTINUE HERE
// 	var defaultTicketRequest = ticketdetails.TicketDetails{
// 		FirstName:     "Houssem",
// 		LastName:      "El Fekih",
// 		Gender:        "Male",
// 		Birthday:      time.Now(),
// 		LaunchpadID:   0,
// 		DestinationID: 0,
// 		LaunchDate:    time.Now()}

// }
