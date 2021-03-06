// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/pivotal-cf/cf-redis-broker/redis/client"
)

type FakeClient struct {
	DisconnectStub        func() error
	disconnectMutex       sync.RWMutex
	disconnectArgsForCall []struct{}
	disconnectReturns struct {
		result1 error
	}
	CreateSnapshotStub        func(timeout time.Duration) error
	createSnapshotMutex       sync.RWMutex
	createSnapshotArgsForCall []struct {
		timeout time.Duration
	}
	createSnapshotReturns struct {
		result1 error
	}
	WaitUntilRedisNotLoadingStub        func(timeoutMilliseconds int) error
	waitUntilRedisNotLoadingMutex       sync.RWMutex
	waitUntilRedisNotLoadingArgsForCall []struct {
		timeoutMilliseconds int
	}
	waitUntilRedisNotLoadingReturns struct {
		result1 error
	}
	EnableAOFStub        func() error
	enableAOFMutex       sync.RWMutex
	enableAOFArgsForCall []struct{}
	enableAOFReturns struct {
		result1 error
	}
	LastRDBSaveTimeStub        func() (int64, error)
	lastRDBSaveTimeMutex       sync.RWMutex
	lastRDBSaveTimeArgsForCall []struct{}
	lastRDBSaveTimeReturns struct {
		result1 int64
		result2 error
	}
	InfoFieldStub        func(fieldName string) (string, error)
	infoFieldMutex       sync.RWMutex
	infoFieldArgsForCall []struct {
		fieldName string
	}
	infoFieldReturns struct {
		result1 string
		result2 error
	}
	InfoStub        func() (map[string]string, error)
	infoMutex       sync.RWMutex
	infoArgsForCall []struct{}
	infoReturns struct {
		result1 map[string]string
		result2 error
	}
	GetConfigStub        func(key string) (string, error)
	getConfigMutex       sync.RWMutex
	getConfigArgsForCall []struct {
		key string
	}
	getConfigReturns struct {
		result1 string
		result2 error
	}
	RDBPathStub        func() (string, error)
	rDBPathMutex       sync.RWMutex
	rDBPathArgsForCall []struct{}
	rDBPathReturns struct {
		result1 string
		result2 error
	}
	AddressStub        func() string
	addressMutex       sync.RWMutex
	addressArgsForCall []struct{}
	addressReturns struct {
		result1 string
	}
}

func (fake *FakeClient) Disconnect() error {
	fake.disconnectMutex.Lock()
	fake.disconnectArgsForCall = append(fake.disconnectArgsForCall, struct{}{})
	fake.disconnectMutex.Unlock()
	if fake.DisconnectStub != nil {
		return fake.DisconnectStub()
	} else {
		return fake.disconnectReturns.result1
	}
}

func (fake *FakeClient) DisconnectCallCount() int {
	fake.disconnectMutex.RLock()
	defer fake.disconnectMutex.RUnlock()
	return len(fake.disconnectArgsForCall)
}

func (fake *FakeClient) DisconnectReturns(result1 error) {
	fake.DisconnectStub = nil
	fake.disconnectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CreateSnapshot(timeout time.Duration) error {
	fake.createSnapshotMutex.Lock()
	fake.createSnapshotArgsForCall = append(fake.createSnapshotArgsForCall, struct {
		timeout time.Duration
	}{timeout})
	fake.createSnapshotMutex.Unlock()
	if fake.CreateSnapshotStub != nil {
		return fake.CreateSnapshotStub(timeout)
	} else {
		return fake.createSnapshotReturns.result1
	}
}

func (fake *FakeClient) CreateSnapshotCallCount() int {
	fake.createSnapshotMutex.RLock()
	defer fake.createSnapshotMutex.RUnlock()
	return len(fake.createSnapshotArgsForCall)
}

func (fake *FakeClient) CreateSnapshotArgsForCall(i int) time.Duration {
	fake.createSnapshotMutex.RLock()
	defer fake.createSnapshotMutex.RUnlock()
	return fake.createSnapshotArgsForCall[i].timeout
}

func (fake *FakeClient) CreateSnapshotReturns(result1 error) {
	fake.CreateSnapshotStub = nil
	fake.createSnapshotReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) WaitUntilRedisNotLoading(timeoutMilliseconds int) error {
	fake.waitUntilRedisNotLoadingMutex.Lock()
	fake.waitUntilRedisNotLoadingArgsForCall = append(fake.waitUntilRedisNotLoadingArgsForCall, struct {
		timeoutMilliseconds int
	}{timeoutMilliseconds})
	fake.waitUntilRedisNotLoadingMutex.Unlock()
	if fake.WaitUntilRedisNotLoadingStub != nil {
		return fake.WaitUntilRedisNotLoadingStub(timeoutMilliseconds)
	} else {
		return fake.waitUntilRedisNotLoadingReturns.result1
	}
}

func (fake *FakeClient) WaitUntilRedisNotLoadingCallCount() int {
	fake.waitUntilRedisNotLoadingMutex.RLock()
	defer fake.waitUntilRedisNotLoadingMutex.RUnlock()
	return len(fake.waitUntilRedisNotLoadingArgsForCall)
}

func (fake *FakeClient) WaitUntilRedisNotLoadingArgsForCall(i int) int {
	fake.waitUntilRedisNotLoadingMutex.RLock()
	defer fake.waitUntilRedisNotLoadingMutex.RUnlock()
	return fake.waitUntilRedisNotLoadingArgsForCall[i].timeoutMilliseconds
}

func (fake *FakeClient) WaitUntilRedisNotLoadingReturns(result1 error) {
	fake.WaitUntilRedisNotLoadingStub = nil
	fake.waitUntilRedisNotLoadingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) EnableAOF() error {
	fake.enableAOFMutex.Lock()
	fake.enableAOFArgsForCall = append(fake.enableAOFArgsForCall, struct{}{})
	fake.enableAOFMutex.Unlock()
	if fake.EnableAOFStub != nil {
		return fake.EnableAOFStub()
	} else {
		return fake.enableAOFReturns.result1
	}
}

func (fake *FakeClient) EnableAOFCallCount() int {
	fake.enableAOFMutex.RLock()
	defer fake.enableAOFMutex.RUnlock()
	return len(fake.enableAOFArgsForCall)
}

func (fake *FakeClient) EnableAOFReturns(result1 error) {
	fake.EnableAOFStub = nil
	fake.enableAOFReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) LastRDBSaveTime() (int64, error) {
	fake.lastRDBSaveTimeMutex.Lock()
	fake.lastRDBSaveTimeArgsForCall = append(fake.lastRDBSaveTimeArgsForCall, struct{}{})
	fake.lastRDBSaveTimeMutex.Unlock()
	if fake.LastRDBSaveTimeStub != nil {
		return fake.LastRDBSaveTimeStub()
	} else {
		return fake.lastRDBSaveTimeReturns.result1, fake.lastRDBSaveTimeReturns.result2
	}
}

func (fake *FakeClient) LastRDBSaveTimeCallCount() int {
	fake.lastRDBSaveTimeMutex.RLock()
	defer fake.lastRDBSaveTimeMutex.RUnlock()
	return len(fake.lastRDBSaveTimeArgsForCall)
}

func (fake *FakeClient) LastRDBSaveTimeReturns(result1 int64, result2 error) {
	fake.LastRDBSaveTimeStub = nil
	fake.lastRDBSaveTimeReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) InfoField(fieldName string) (string, error) {
	fake.infoFieldMutex.Lock()
	fake.infoFieldArgsForCall = append(fake.infoFieldArgsForCall, struct {
		fieldName string
	}{fieldName})
	fake.infoFieldMutex.Unlock()
	if fake.InfoFieldStub != nil {
		return fake.InfoFieldStub(fieldName)
	} else {
		return fake.infoFieldReturns.result1, fake.infoFieldReturns.result2
	}
}

func (fake *FakeClient) InfoFieldCallCount() int {
	fake.infoFieldMutex.RLock()
	defer fake.infoFieldMutex.RUnlock()
	return len(fake.infoFieldArgsForCall)
}

func (fake *FakeClient) InfoFieldArgsForCall(i int) string {
	fake.infoFieldMutex.RLock()
	defer fake.infoFieldMutex.RUnlock()
	return fake.infoFieldArgsForCall[i].fieldName
}

func (fake *FakeClient) InfoFieldReturns(result1 string, result2 error) {
	fake.InfoFieldStub = nil
	fake.infoFieldReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Info() (map[string]string, error) {
	fake.infoMutex.Lock()
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct{}{})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		return fake.InfoStub()
	} else {
		return fake.infoReturns.result1, fake.infoReturns.result2
	}
}

func (fake *FakeClient) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *FakeClient) InfoReturns(result1 map[string]string, result2 error) {
	fake.InfoStub = nil
	fake.infoReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetConfig(key string) (string, error) {
	fake.getConfigMutex.Lock()
	fake.getConfigArgsForCall = append(fake.getConfigArgsForCall, struct {
		key string
	}{key})
	fake.getConfigMutex.Unlock()
	if fake.GetConfigStub != nil {
		return fake.GetConfigStub(key)
	} else {
		return fake.getConfigReturns.result1, fake.getConfigReturns.result2
	}
}

func (fake *FakeClient) GetConfigCallCount() int {
	fake.getConfigMutex.RLock()
	defer fake.getConfigMutex.RUnlock()
	return len(fake.getConfigArgsForCall)
}

func (fake *FakeClient) GetConfigArgsForCall(i int) string {
	fake.getConfigMutex.RLock()
	defer fake.getConfigMutex.RUnlock()
	return fake.getConfigArgsForCall[i].key
}

func (fake *FakeClient) GetConfigReturns(result1 string, result2 error) {
	fake.GetConfigStub = nil
	fake.getConfigReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) RDBPath() (string, error) {
	fake.rDBPathMutex.Lock()
	fake.rDBPathArgsForCall = append(fake.rDBPathArgsForCall, struct{}{})
	fake.rDBPathMutex.Unlock()
	if fake.RDBPathStub != nil {
		return fake.RDBPathStub()
	} else {
		return fake.rDBPathReturns.result1, fake.rDBPathReturns.result2
	}
}

func (fake *FakeClient) RDBPathCallCount() int {
	fake.rDBPathMutex.RLock()
	defer fake.rDBPathMutex.RUnlock()
	return len(fake.rDBPathArgsForCall)
}

func (fake *FakeClient) RDBPathReturns(result1 string, result2 error) {
	fake.RDBPathStub = nil
	fake.rDBPathReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Address() string {
	fake.addressMutex.Lock()
	fake.addressArgsForCall = append(fake.addressArgsForCall, struct{}{})
	fake.addressMutex.Unlock()
	if fake.AddressStub != nil {
		return fake.AddressStub()
	} else {
		return fake.addressReturns.result1
	}
}

func (fake *FakeClient) AddressCallCount() int {
	fake.addressMutex.RLock()
	defer fake.addressMutex.RUnlock()
	return len(fake.addressArgsForCall)
}

func (fake *FakeClient) AddressReturns(result1 string) {
	fake.AddressStub = nil
	fake.addressReturns = struct {
		result1 string
	}{result1}
}

var _ client.Client = new(FakeClient)
