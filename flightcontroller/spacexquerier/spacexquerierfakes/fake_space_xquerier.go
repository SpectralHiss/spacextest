// Code generated by counterfeiter. DO NOT EDIT.
package spacexquerierfakes

import (
	"sync"
	"time"

	"github.com/SpectralHiss/spacextest/flightcontroller/querytypes"
	"github.com/SpectralHiss/spacextest/flightcontroller/spacexquerier"
)

type FakeSpaceXQuerier struct {
	GetLaunchIdsStub        func() []querytypes.LaunchPadID
	getLaunchIdsMutex       sync.RWMutex
	getLaunchIdsArgsForCall []struct {
	}
	getLaunchIdsReturns struct {
		result1 []querytypes.LaunchPadID
	}
	getLaunchIdsReturnsOnCall map[int]struct {
		result1 []querytypes.LaunchPadID
	}
	LaunchPossibleStub        func(querytypes.LaunchPadID, time.Time) bool
	launchPossibleMutex       sync.RWMutex
	launchPossibleArgsForCall []struct {
		arg1 querytypes.LaunchPadID
		arg2 time.Time
	}
	launchPossibleReturns struct {
		result1 bool
	}
	launchPossibleReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpaceXQuerier) GetLaunchIds() []querytypes.LaunchPadID {
	fake.getLaunchIdsMutex.Lock()
	ret, specificReturn := fake.getLaunchIdsReturnsOnCall[len(fake.getLaunchIdsArgsForCall)]
	fake.getLaunchIdsArgsForCall = append(fake.getLaunchIdsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetLaunchIds", []interface{}{})
	fake.getLaunchIdsMutex.Unlock()
	if fake.GetLaunchIdsStub != nil {
		return fake.GetLaunchIdsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getLaunchIdsReturns
	return fakeReturns.result1
}

func (fake *FakeSpaceXQuerier) GetLaunchIdsCallCount() int {
	fake.getLaunchIdsMutex.RLock()
	defer fake.getLaunchIdsMutex.RUnlock()
	return len(fake.getLaunchIdsArgsForCall)
}

func (fake *FakeSpaceXQuerier) GetLaunchIdsCalls(stub func() []querytypes.LaunchPadID) {
	fake.getLaunchIdsMutex.Lock()
	defer fake.getLaunchIdsMutex.Unlock()
	fake.GetLaunchIdsStub = stub
}

func (fake *FakeSpaceXQuerier) GetLaunchIdsReturns(result1 []querytypes.LaunchPadID) {
	fake.getLaunchIdsMutex.Lock()
	defer fake.getLaunchIdsMutex.Unlock()
	fake.GetLaunchIdsStub = nil
	fake.getLaunchIdsReturns = struct {
		result1 []querytypes.LaunchPadID
	}{result1}
}

func (fake *FakeSpaceXQuerier) GetLaunchIdsReturnsOnCall(i int, result1 []querytypes.LaunchPadID) {
	fake.getLaunchIdsMutex.Lock()
	defer fake.getLaunchIdsMutex.Unlock()
	fake.GetLaunchIdsStub = nil
	if fake.getLaunchIdsReturnsOnCall == nil {
		fake.getLaunchIdsReturnsOnCall = make(map[int]struct {
			result1 []querytypes.LaunchPadID
		})
	}
	fake.getLaunchIdsReturnsOnCall[i] = struct {
		result1 []querytypes.LaunchPadID
	}{result1}
}

func (fake *FakeSpaceXQuerier) LaunchPossible(arg1 querytypes.LaunchPadID, arg2 time.Time) bool {
	fake.launchPossibleMutex.Lock()
	ret, specificReturn := fake.launchPossibleReturnsOnCall[len(fake.launchPossibleArgsForCall)]
	fake.launchPossibleArgsForCall = append(fake.launchPossibleArgsForCall, struct {
		arg1 querytypes.LaunchPadID
		arg2 time.Time
	}{arg1, arg2})
	fake.recordInvocation("LaunchPossible", []interface{}{arg1, arg2})
	fake.launchPossibleMutex.Unlock()
	if fake.LaunchPossibleStub != nil {
		return fake.LaunchPossibleStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.launchPossibleReturns
	return fakeReturns.result1
}

func (fake *FakeSpaceXQuerier) LaunchPossibleCallCount() int {
	fake.launchPossibleMutex.RLock()
	defer fake.launchPossibleMutex.RUnlock()
	return len(fake.launchPossibleArgsForCall)
}

func (fake *FakeSpaceXQuerier) LaunchPossibleCalls(stub func(querytypes.LaunchPadID, time.Time) bool) {
	fake.launchPossibleMutex.Lock()
	defer fake.launchPossibleMutex.Unlock()
	fake.LaunchPossibleStub = stub
}

func (fake *FakeSpaceXQuerier) LaunchPossibleArgsForCall(i int) (querytypes.LaunchPadID, time.Time) {
	fake.launchPossibleMutex.RLock()
	defer fake.launchPossibleMutex.RUnlock()
	argsForCall := fake.launchPossibleArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSpaceXQuerier) LaunchPossibleReturns(result1 bool) {
	fake.launchPossibleMutex.Lock()
	defer fake.launchPossibleMutex.Unlock()
	fake.LaunchPossibleStub = nil
	fake.launchPossibleReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeSpaceXQuerier) LaunchPossibleReturnsOnCall(i int, result1 bool) {
	fake.launchPossibleMutex.Lock()
	defer fake.launchPossibleMutex.Unlock()
	fake.LaunchPossibleStub = nil
	if fake.launchPossibleReturnsOnCall == nil {
		fake.launchPossibleReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.launchPossibleReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeSpaceXQuerier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getLaunchIdsMutex.RLock()
	defer fake.getLaunchIdsMutex.RUnlock()
	fake.launchPossibleMutex.RLock()
	defer fake.launchPossibleMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSpaceXQuerier) recordInvocation(key string, args []interface{}) {
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

var _ spacexquerier.SpaceXQuerier = new(FakeSpaceXQuerier)
