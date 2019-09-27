package scheduler_test

import (
	"math/rand"
	"testing"

	. "github.com/SpectralHiss/spacextest/flightcontroller/scheduler"
	"github.com/SpectralHiss/spacextest/flightcontroller/spacexquerier/spacexquerierfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scheduler suite")
}

func List(end int) []int {
	return rand.Perm(end)
}

func extractGoodSchedule(sched Schedule) (LaunchPadID, DestinationID, Day) {
	println(len(sched))
	someLaunchPad := LaunchPadID(rand.Intn(len(sched)))
	someDay := Day(rand.Intn(7))
	return someLaunchPad, sched[someLaunchPad][someDay], someDay
}

var _ = Describe("Scheduler suite", func() {
	var scheduler *Scheduler
	var fakeSpaceX *spacexquerierfakes.FakeSpaceXQuerier

	Context("When the scheduler is checked w valid LaunchPadID, DestinationID mapping for that day", func() {
		BeforeEach(func() {
			fakeSpaceX = &spacexquerierfakes.FakeSpaceXQuerier{}

			fakeSpaceX.GetLaunchIdsReturns(List(50))
			scheduler = NewScheduler(fakeSpaceX)
		})

		It("Should return false", func() {
			validLID, validDest, validDay := extractGoodSchedule(scheduler.SMap)
			isValid := scheduler.CheckSchedule(validLID, validDay, validDest)
			Expect(isValid).To(BeTrue())
			Expect(fakeSpaceX.GetLaunchIdsCallCount()).To(Equal(1))
		})

	})

	Context("When the scheduler is check with invalidLaunchPadID", func() {
		BeforeEach(func() {
			fakeSpaceX = &spacexquerierfakes.FakeSpaceXQuerier{}

			fakeSpaceX.GetLaunchIdsReturns(List(50))
			scheduler = NewScheduler(fakeSpaceX)
		})

		It("Should return false", func() {
			_, validDest, validDay := extractGoodSchedule(scheduler.SMap)
			invalidLID := LaunchPadID(51) // we have 0 to 50 in our example
			isValid := scheduler.CheckSchedule(invalidLID, validDay, validDest)
			Expect(isValid).To(Not(BeTrue()))
			Expect(fakeSpaceX.GetLaunchIdsCallCount()).To(Equal(1))

		})

	})

	Context("When the scheduler is check with a wrong destination for that day", func() {
		BeforeEach(func() {
			fakeSpaceX = &spacexquerierfakes.FakeSpaceXQuerier{}

			fakeSpaceX.GetLaunchIdsReturns(List(50))
			scheduler = NewScheduler(fakeSpaceX)
		})

		It("Should return false", func() {
			validLID, validDest, validDay := extractGoodSchedule(scheduler.SMap)

			isValid := scheduler.CheckSchedule(validLID, validDay, (validDest + 1%8))
			Expect(isValid).To(Not(BeTrue()))
			Expect(fakeSpaceX.GetLaunchIdsCallCount()).To(Equal(1))

		})

	})
	//TODO
	Context("When the scheduler is check with an invalid day", func() {})

})
