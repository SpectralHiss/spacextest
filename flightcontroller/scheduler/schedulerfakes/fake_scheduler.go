// Code generated by counterfeiter. DO NOT EDIT.
package schedulerfakes

import (
	"sync"

	"github.com/SpectralHiss/spacextest/flightcontroller/scheduler"
)

type FakeScheduler struct {
	GenerateScheduleStub        func() map[int]map[scheduler.Day]scheduler.DestinationID
	generateScheduleMutex       sync.RWMutex
	generateScheduleArgsForCall []struct {
	}
	generateScheduleReturns struct {
		result1 map[int]map[scheduler.Day]scheduler.DestinationID
	}
	generateScheduleReturnsOnCall map[int]struct {
		result1 map[int]map[scheduler.Day]scheduler.DestinationID
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeScheduler) GenerateSchedule() map[int]map[scheduler.Day]scheduler.DestinationID {
	fake.generateScheduleMutex.Lock()
	ret, specificReturn := fake.generateScheduleReturnsOnCall[len(fake.generateScheduleArgsForCall)]
	fake.generateScheduleArgsForCall = append(fake.generateScheduleArgsForCall, struct {
	}{})
	fake.recordInvocation("GenerateSchedule", []interface{}{})
	fake.generateScheduleMutex.Unlock()
	if fake.GenerateScheduleStub != nil {
		return fake.GenerateScheduleStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.generateScheduleReturns
	return fakeReturns.result1
}

func (fake *FakeScheduler) GenerateScheduleCallCount() int {
	fake.generateScheduleMutex.RLock()
	defer fake.generateScheduleMutex.RUnlock()
	return len(fake.generateScheduleArgsForCall)
}

func (fake *FakeScheduler) GenerateScheduleCalls(stub func() map[int]map[scheduler.Day]scheduler.DestinationID) {
	fake.generateScheduleMutex.Lock()
	defer fake.generateScheduleMutex.Unlock()
	fake.GenerateScheduleStub = stub
}

func (fake *FakeScheduler) GenerateScheduleReturns(result1 map[int]map[scheduler.Day]scheduler.DestinationID) {
	fake.generateScheduleMutex.Lock()
	defer fake.generateScheduleMutex.Unlock()
	fake.GenerateScheduleStub = nil
	fake.generateScheduleReturns = struct {
		result1 map[int]map[scheduler.Day]scheduler.DestinationID
	}{result1}
}

func (fake *FakeScheduler) GenerateScheduleReturnsOnCall(i int, result1 map[int]map[scheduler.Day]scheduler.DestinationID) {
	fake.generateScheduleMutex.Lock()
	defer fake.generateScheduleMutex.Unlock()
	fake.GenerateScheduleStub = nil
	if fake.generateScheduleReturnsOnCall == nil {
		fake.generateScheduleReturnsOnCall = make(map[int]struct {
			result1 map[int]map[scheduler.Day]scheduler.DestinationID
		})
	}
	fake.generateScheduleReturnsOnCall[i] = struct {
		result1 map[int]map[scheduler.Day]scheduler.DestinationID
	}{result1}
}

func (fake *FakeScheduler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateScheduleMutex.RLock()
	defer fake.generateScheduleMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeScheduler) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ scheduler.Scheduler = new(FakeScheduler)
