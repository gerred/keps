// Code generated by counterfeiter. DO NOT EDIT.
package sectionsfakes

import (
	sync "sync"

	sections "github.com/calebamiles/keps/pkg/keps/sections"
)

type FakeCollection struct {
	ContentDirStub        func() string
	contentDirMutex       sync.RWMutex
	contentDirArgsForCall []struct {
	}
	contentDirReturns struct {
		result1 string
	}
	contentDirReturnsOnCall map[int]struct {
		result1 string
	}
	EraseStub        func() error
	eraseMutex       sync.RWMutex
	eraseArgsForCall []struct {
	}
	eraseReturns struct {
		result1 error
	}
	eraseReturnsOnCall map[int]struct {
		result1 error
	}
	PersistStub        func() error
	persistMutex       sync.RWMutex
	persistArgsForCall []struct {
	}
	persistReturns struct {
		result1 error
	}
	persistReturnsOnCall map[int]struct {
		result1 error
	}
	SectionsStub        func() []string
	sectionsMutex       sync.RWMutex
	sectionsArgsForCall []struct {
	}
	sectionsReturns struct {
		result1 []string
	}
	sectionsReturnsOnCall map[int]struct {
		result1 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCollection) ContentDir() string {
	fake.contentDirMutex.Lock()
	ret, specificReturn := fake.contentDirReturnsOnCall[len(fake.contentDirArgsForCall)]
	fake.contentDirArgsForCall = append(fake.contentDirArgsForCall, struct {
	}{})
	fake.recordInvocation("ContentDir", []interface{}{})
	fake.contentDirMutex.Unlock()
	if fake.ContentDirStub != nil {
		return fake.ContentDirStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.contentDirReturns
	return fakeReturns.result1
}

func (fake *FakeCollection) ContentDirCallCount() int {
	fake.contentDirMutex.RLock()
	defer fake.contentDirMutex.RUnlock()
	return len(fake.contentDirArgsForCall)
}

func (fake *FakeCollection) ContentDirReturns(result1 string) {
	fake.ContentDirStub = nil
	fake.contentDirReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCollection) ContentDirReturnsOnCall(i int, result1 string) {
	fake.ContentDirStub = nil
	if fake.contentDirReturnsOnCall == nil {
		fake.contentDirReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.contentDirReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCollection) Erase() error {
	fake.eraseMutex.Lock()
	ret, specificReturn := fake.eraseReturnsOnCall[len(fake.eraseArgsForCall)]
	fake.eraseArgsForCall = append(fake.eraseArgsForCall, struct {
	}{})
	fake.recordInvocation("Erase", []interface{}{})
	fake.eraseMutex.Unlock()
	if fake.EraseStub != nil {
		return fake.EraseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.eraseReturns
	return fakeReturns.result1
}

func (fake *FakeCollection) EraseCallCount() int {
	fake.eraseMutex.RLock()
	defer fake.eraseMutex.RUnlock()
	return len(fake.eraseArgsForCall)
}

func (fake *FakeCollection) EraseReturns(result1 error) {
	fake.EraseStub = nil
	fake.eraseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCollection) EraseReturnsOnCall(i int, result1 error) {
	fake.EraseStub = nil
	if fake.eraseReturnsOnCall == nil {
		fake.eraseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.eraseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCollection) Persist() error {
	fake.persistMutex.Lock()
	ret, specificReturn := fake.persistReturnsOnCall[len(fake.persistArgsForCall)]
	fake.persistArgsForCall = append(fake.persistArgsForCall, struct {
	}{})
	fake.recordInvocation("Persist", []interface{}{})
	fake.persistMutex.Unlock()
	if fake.PersistStub != nil {
		return fake.PersistStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.persistReturns
	return fakeReturns.result1
}

func (fake *FakeCollection) PersistCallCount() int {
	fake.persistMutex.RLock()
	defer fake.persistMutex.RUnlock()
	return len(fake.persistArgsForCall)
}

func (fake *FakeCollection) PersistReturns(result1 error) {
	fake.PersistStub = nil
	fake.persistReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCollection) PersistReturnsOnCall(i int, result1 error) {
	fake.PersistStub = nil
	if fake.persistReturnsOnCall == nil {
		fake.persistReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.persistReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCollection) Sections() []string {
	fake.sectionsMutex.Lock()
	ret, specificReturn := fake.sectionsReturnsOnCall[len(fake.sectionsArgsForCall)]
	fake.sectionsArgsForCall = append(fake.sectionsArgsForCall, struct {
	}{})
	fake.recordInvocation("Sections", []interface{}{})
	fake.sectionsMutex.Unlock()
	if fake.SectionsStub != nil {
		return fake.SectionsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sectionsReturns
	return fakeReturns.result1
}

func (fake *FakeCollection) SectionsCallCount() int {
	fake.sectionsMutex.RLock()
	defer fake.sectionsMutex.RUnlock()
	return len(fake.sectionsArgsForCall)
}

func (fake *FakeCollection) SectionsReturns(result1 []string) {
	fake.SectionsStub = nil
	fake.sectionsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeCollection) SectionsReturnsOnCall(i int, result1 []string) {
	fake.SectionsStub = nil
	if fake.sectionsReturnsOnCall == nil {
		fake.sectionsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.sectionsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeCollection) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.contentDirMutex.RLock()
	defer fake.contentDirMutex.RUnlock()
	fake.eraseMutex.RLock()
	defer fake.eraseMutex.RUnlock()
	fake.persistMutex.RLock()
	defer fake.persistMutex.RUnlock()
	fake.sectionsMutex.RLock()
	defer fake.sectionsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCollection) recordInvocation(key string, args []interface{}) {
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

var _ sections.Collection = new(FakeCollection)
