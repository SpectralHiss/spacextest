package FlightController_test

import . "github.com/SpectralHiss/TabeoTest/FightController"
import "github.com/SpectralHiss/TabeoTest/FightController/persister"


var _ = Describe("Space flight booking", func() {


	Describe("creating new reservation",func(){
		Context("When creating a valid new reservation with a good desination mapping, no clash with SpaceX", func() {
			var fakePersister fakePersister
			var FakeSpaceXQuerier FakeSpaceXQuerier

			fakePersister.SaveReturns(nil)

			FakeSpaceXQuerier.LaunchPossibleReturns(true)

			controller = NewFlightController(fakePersister, FakeSpaceXQuerier, FlightController.DefaultDesinationMapping)
		}
	


		Context("When creating a valid new reservation with a good desination mapping, clashing with SpaceX", func() {}

		Context("When creating a valid new reservation with a bad desination mapping", func() {}


	})

})

