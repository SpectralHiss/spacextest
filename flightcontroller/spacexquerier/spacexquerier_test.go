package spacexquerier_test

import "testing"

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scheduler suite")
}
