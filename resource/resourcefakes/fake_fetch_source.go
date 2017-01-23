// This file was generated by counterfeiter
package resourcefakes

import (
	"os"
	"sync"

	"github.com/concourse/atc/resource"
)

type FakeFetchSource struct {
	IsInitializedStub        func() (bool, error)
	isInitializedMutex       sync.RWMutex
	isInitializedArgsForCall []struct{}
	isInitializedReturns     struct {
		result1 bool
		result2 error
	}
	LockNameStub        func() (string, error)
	lockNameMutex       sync.RWMutex
	lockNameArgsForCall []struct{}
	lockNameReturns     struct {
		result1 string
		result2 error
	}
	VersionedSourceStub        func() resource.VersionedSource
	versionedSourceMutex       sync.RWMutex
	versionedSourceArgsForCall []struct{}
	versionedSourceReturns     struct {
		result1 resource.VersionedSource
	}
	InitializeStub        func(signals <-chan os.Signal, ready chan<- struct{}) error
	initializeMutex       sync.RWMutex
	initializeArgsForCall []struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}
	initializeReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFetchSource) IsInitialized() (bool, error) {
	fake.isInitializedMutex.Lock()
	fake.isInitializedArgsForCall = append(fake.isInitializedArgsForCall, struct{}{})
	fake.recordInvocation("IsInitialized", []interface{}{})
	fake.isInitializedMutex.Unlock()
	if fake.IsInitializedStub != nil {
		return fake.IsInitializedStub()
	} else {
		return fake.isInitializedReturns.result1, fake.isInitializedReturns.result2
	}
}

func (fake *FakeFetchSource) IsInitializedCallCount() int {
	fake.isInitializedMutex.RLock()
	defer fake.isInitializedMutex.RUnlock()
	return len(fake.isInitializedArgsForCall)
}

func (fake *FakeFetchSource) IsInitializedReturns(result1 bool, result2 error) {
	fake.IsInitializedStub = nil
	fake.isInitializedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeFetchSource) LockName() (string, error) {
	fake.lockNameMutex.Lock()
	fake.lockNameArgsForCall = append(fake.lockNameArgsForCall, struct{}{})
	fake.recordInvocation("LockName", []interface{}{})
	fake.lockNameMutex.Unlock()
	if fake.LockNameStub != nil {
		return fake.LockNameStub()
	} else {
		return fake.lockNameReturns.result1, fake.lockNameReturns.result2
	}
}

func (fake *FakeFetchSource) LockNameCallCount() int {
	fake.lockNameMutex.RLock()
	defer fake.lockNameMutex.RUnlock()
	return len(fake.lockNameArgsForCall)
}

func (fake *FakeFetchSource) LockNameReturns(result1 string, result2 error) {
	fake.LockNameStub = nil
	fake.lockNameReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFetchSource) VersionedSource() resource.VersionedSource {
	fake.versionedSourceMutex.Lock()
	fake.versionedSourceArgsForCall = append(fake.versionedSourceArgsForCall, struct{}{})
	fake.recordInvocation("VersionedSource", []interface{}{})
	fake.versionedSourceMutex.Unlock()
	if fake.VersionedSourceStub != nil {
		return fake.VersionedSourceStub()
	} else {
		return fake.versionedSourceReturns.result1
	}
}

func (fake *FakeFetchSource) VersionedSourceCallCount() int {
	fake.versionedSourceMutex.RLock()
	defer fake.versionedSourceMutex.RUnlock()
	return len(fake.versionedSourceArgsForCall)
}

func (fake *FakeFetchSource) VersionedSourceReturns(result1 resource.VersionedSource) {
	fake.VersionedSourceStub = nil
	fake.versionedSourceReturns = struct {
		result1 resource.VersionedSource
	}{result1}
}

func (fake *FakeFetchSource) Initialize(signals <-chan os.Signal, ready chan<- struct{}) error {
	fake.initializeMutex.Lock()
	fake.initializeArgsForCall = append(fake.initializeArgsForCall, struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}{signals, ready})
	fake.recordInvocation("Initialize", []interface{}{signals, ready})
	fake.initializeMutex.Unlock()
	if fake.InitializeStub != nil {
		return fake.InitializeStub(signals, ready)
	} else {
		return fake.initializeReturns.result1
	}
}

func (fake *FakeFetchSource) InitializeCallCount() int {
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	return len(fake.initializeArgsForCall)
}

func (fake *FakeFetchSource) InitializeArgsForCall(i int) (<-chan os.Signal, chan<- struct{}) {
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	return fake.initializeArgsForCall[i].signals, fake.initializeArgsForCall[i].ready
}

func (fake *FakeFetchSource) InitializeReturns(result1 error) {
	fake.InitializeStub = nil
	fake.initializeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFetchSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isInitializedMutex.RLock()
	defer fake.isInitializedMutex.RUnlock()
	fake.lockNameMutex.RLock()
	defer fake.lockNameMutex.RUnlock()
	fake.versionedSourceMutex.RLock()
	defer fake.versionedSourceMutex.RUnlock()
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFetchSource) recordInvocation(key string, args []interface{}) {
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

var _ resource.FetchSource = new(FakeFetchSource)
