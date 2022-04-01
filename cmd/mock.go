// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package cmd

import (
	"github.com/dotzero/git-profile/config"
	"sync"
)

// storageMock is a mock implementation of storage.
//
// 	func TestSomethingThatUsesstorage(t *testing.T) {
//
// 		// make and configure a mocked storage
// 		mockedstorage := &storageMock{
// 			DeleteFunc: func(profile string, value string) bool {
// 				panic("mock out the Delete method")
// 			},
// 			DeleteProfileFunc: func(profile string) bool {
// 				panic("mock out the DeleteProfile method")
// 			},
// 			LenFunc: func() int {
// 				panic("mock out the Len method")
// 			},
// 			LoadFunc: func(filename string) error {
// 				panic("mock out the Load method")
// 			},
// 			LookupFunc: func(name string) ([]config.Entry, bool) {
// 				panic("mock out the Lookup method")
// 			},
// 			NamesFunc: func() []string {
// 				panic("mock out the Names method")
// 			},
// 			SaveFunc: func(filename string) error {
// 				panic("mock out the Save method")
// 			},
// 			StoreFunc: func(profile string, key string, value string)  {
// 				panic("mock out the Store method")
// 			},
// 		}
//
// 		// use mockedstorage in code that requires storage
// 		// and then make assertions.
//
// 	}
type storageMock struct {
	// DeleteFunc mocks the Delete method.
	DeleteFunc func(profile string, value string) bool

	// DeleteProfileFunc mocks the DeleteProfile method.
	DeleteProfileFunc func(profile string) bool

	// LenFunc mocks the Len method.
	LenFunc func() int

	// LoadFunc mocks the Load method.
	LoadFunc func(filename string) error

	// LookupFunc mocks the Lookup method.
	LookupFunc func(name string) ([]config.Entry, bool)

	// NamesFunc mocks the Names method.
	NamesFunc func() []string

	// SaveFunc mocks the Save method.
	SaveFunc func(filename string) error

	// StoreFunc mocks the Store method.
	StoreFunc func(profile string, key string, value string)

	// calls tracks calls to the methods.
	calls struct {
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Profile is the profile argument value.
			Profile string
			// Value is the value argument value.
			Value string
		}
		// DeleteProfile holds details about calls to the DeleteProfile method.
		DeleteProfile []struct {
			// Profile is the profile argument value.
			Profile string
		}
		// Len holds details about calls to the Len method.
		Len []struct {
		}
		// Load holds details about calls to the Load method.
		Load []struct {
			// Filename is the filename argument value.
			Filename string
		}
		// Lookup holds details about calls to the Lookup method.
		Lookup []struct {
			// Name is the name argument value.
			Name string
		}
		// Names holds details about calls to the Names method.
		Names []struct {
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// Filename is the filename argument value.
			Filename string
		}
		// Store holds details about calls to the Store method.
		Store []struct {
			// Profile is the profile argument value.
			Profile string
			// Key is the key argument value.
			Key string
			// Value is the value argument value.
			Value string
		}
	}
	lockDelete        sync.RWMutex
	lockDeleteProfile sync.RWMutex
	lockLen           sync.RWMutex
	lockLoad          sync.RWMutex
	lockLookup        sync.RWMutex
	lockNames         sync.RWMutex
	lockSave          sync.RWMutex
	lockStore         sync.RWMutex
}

// Delete calls DeleteFunc.
func (mock *storageMock) Delete(profile string, value string) bool {
	if mock.DeleteFunc == nil {
		panic("storageMock.DeleteFunc: method is nil but storage.Delete was just called")
	}
	callInfo := struct {
		Profile string
		Value   string
	}{
		Profile: profile,
		Value:   value,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(profile, value)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedstorage.DeleteCalls())
func (mock *storageMock) DeleteCalls() []struct {
	Profile string
	Value   string
} {
	var calls []struct {
		Profile string
		Value   string
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// DeleteProfile calls DeleteProfileFunc.
func (mock *storageMock) DeleteProfile(profile string) bool {
	if mock.DeleteProfileFunc == nil {
		panic("storageMock.DeleteProfileFunc: method is nil but storage.DeleteProfile was just called")
	}
	callInfo := struct {
		Profile string
	}{
		Profile: profile,
	}
	mock.lockDeleteProfile.Lock()
	mock.calls.DeleteProfile = append(mock.calls.DeleteProfile, callInfo)
	mock.lockDeleteProfile.Unlock()
	return mock.DeleteProfileFunc(profile)
}

// DeleteProfileCalls gets all the calls that were made to DeleteProfile.
// Check the length with:
//     len(mockedstorage.DeleteProfileCalls())
func (mock *storageMock) DeleteProfileCalls() []struct {
	Profile string
} {
	var calls []struct {
		Profile string
	}
	mock.lockDeleteProfile.RLock()
	calls = mock.calls.DeleteProfile
	mock.lockDeleteProfile.RUnlock()
	return calls
}

// Len calls LenFunc.
func (mock *storageMock) Len() int {
	if mock.LenFunc == nil {
		panic("storageMock.LenFunc: method is nil but storage.Len was just called")
	}
	callInfo := struct {
	}{}
	mock.lockLen.Lock()
	mock.calls.Len = append(mock.calls.Len, callInfo)
	mock.lockLen.Unlock()
	return mock.LenFunc()
}

// LenCalls gets all the calls that were made to Len.
// Check the length with:
//     len(mockedstorage.LenCalls())
func (mock *storageMock) LenCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockLen.RLock()
	calls = mock.calls.Len
	mock.lockLen.RUnlock()
	return calls
}

// Load calls LoadFunc.
func (mock *storageMock) Load(filename string) error {
	if mock.LoadFunc == nil {
		panic("storageMock.LoadFunc: method is nil but storage.Load was just called")
	}
	callInfo := struct {
		Filename string
	}{
		Filename: filename,
	}
	mock.lockLoad.Lock()
	mock.calls.Load = append(mock.calls.Load, callInfo)
	mock.lockLoad.Unlock()
	return mock.LoadFunc(filename)
}

// LoadCalls gets all the calls that were made to Load.
// Check the length with:
//     len(mockedstorage.LoadCalls())
func (mock *storageMock) LoadCalls() []struct {
	Filename string
} {
	var calls []struct {
		Filename string
	}
	mock.lockLoad.RLock()
	calls = mock.calls.Load
	mock.lockLoad.RUnlock()
	return calls
}

// Lookup calls LookupFunc.
func (mock *storageMock) Lookup(name string) ([]config.Entry, bool) {
	if mock.LookupFunc == nil {
		panic("storageMock.LookupFunc: method is nil but storage.Lookup was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	mock.lockLookup.Lock()
	mock.calls.Lookup = append(mock.calls.Lookup, callInfo)
	mock.lockLookup.Unlock()
	return mock.LookupFunc(name)
}

// LookupCalls gets all the calls that were made to Lookup.
// Check the length with:
//     len(mockedstorage.LookupCalls())
func (mock *storageMock) LookupCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	mock.lockLookup.RLock()
	calls = mock.calls.Lookup
	mock.lockLookup.RUnlock()
	return calls
}

// Names calls NamesFunc.
func (mock *storageMock) Names() []string {
	if mock.NamesFunc == nil {
		panic("storageMock.NamesFunc: method is nil but storage.Names was just called")
	}
	callInfo := struct {
	}{}
	mock.lockNames.Lock()
	mock.calls.Names = append(mock.calls.Names, callInfo)
	mock.lockNames.Unlock()
	return mock.NamesFunc()
}

// NamesCalls gets all the calls that were made to Names.
// Check the length with:
//     len(mockedstorage.NamesCalls())
func (mock *storageMock) NamesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockNames.RLock()
	calls = mock.calls.Names
	mock.lockNames.RUnlock()
	return calls
}

// Save calls SaveFunc.
func (mock *storageMock) Save(filename string) error {
	if mock.SaveFunc == nil {
		panic("storageMock.SaveFunc: method is nil but storage.Save was just called")
	}
	callInfo := struct {
		Filename string
	}{
		Filename: filename,
	}
	mock.lockSave.Lock()
	mock.calls.Save = append(mock.calls.Save, callInfo)
	mock.lockSave.Unlock()
	return mock.SaveFunc(filename)
}

// SaveCalls gets all the calls that were made to Save.
// Check the length with:
//     len(mockedstorage.SaveCalls())
func (mock *storageMock) SaveCalls() []struct {
	Filename string
} {
	var calls []struct {
		Filename string
	}
	mock.lockSave.RLock()
	calls = mock.calls.Save
	mock.lockSave.RUnlock()
	return calls
}

// Store calls StoreFunc.
func (mock *storageMock) Store(profile string, key string, value string) {
	if mock.StoreFunc == nil {
		panic("storageMock.StoreFunc: method is nil but storage.Store was just called")
	}
	callInfo := struct {
		Profile string
		Key     string
		Value   string
	}{
		Profile: profile,
		Key:     key,
		Value:   value,
	}
	mock.lockStore.Lock()
	mock.calls.Store = append(mock.calls.Store, callInfo)
	mock.lockStore.Unlock()
	mock.StoreFunc(profile, key, value)
}

// StoreCalls gets all the calls that were made to Store.
// Check the length with:
//     len(mockedstorage.StoreCalls())
func (mock *storageMock) StoreCalls() []struct {
	Profile string
	Key     string
	Value   string
} {
	var calls []struct {
		Profile string
		Key     string
		Value   string
	}
	mock.lockStore.RLock()
	calls = mock.calls.Store
	mock.lockStore.RUnlock()
	return calls
}

// vcsMock is a mock implementation of vcs.
//
// 	func TestSomethingThatUsesvcs(t *testing.T) {
//
// 		// make and configure a mocked vcs
// 		mockedvcs := &vcsMock{
// 			GetFunc: func(key string) (string, error) {
// 				panic("mock out the Get method")
// 			},
// 			IsRepositoryFunc: func() bool {
// 				panic("mock out the IsRepository method")
// 			},
// 			SetFunc: func(key string, value string) error {
// 				panic("mock out the Set method")
// 			},
// 		}
//
// 		// use mockedvcs in code that requires vcs
// 		// and then make assertions.
//
// 	}
type vcsMock struct {
	// GetFunc mocks the Get method.
	GetFunc func(key string) (string, error)

	// IsRepositoryFunc mocks the IsRepository method.
	IsRepositoryFunc func() bool

	// SetFunc mocks the Set method.
	SetFunc func(key string, value string) error

	// calls tracks calls to the methods.
	calls struct {
		// Get holds details about calls to the Get method.
		Get []struct {
			// Key is the key argument value.
			Key string
		}
		// IsRepository holds details about calls to the IsRepository method.
		IsRepository []struct {
		}
		// Set holds details about calls to the Set method.
		Set []struct {
			// Key is the key argument value.
			Key string
			// Value is the value argument value.
			Value string
		}
	}
	lockGet          sync.RWMutex
	lockIsRepository sync.RWMutex
	lockSet          sync.RWMutex
}

// Get calls GetFunc.
func (mock *vcsMock) Get(key string) (string, error) {
	if mock.GetFunc == nil {
		panic("vcsMock.GetFunc: method is nil but vcs.Get was just called")
	}
	callInfo := struct {
		Key string
	}{
		Key: key,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(key)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedvcs.GetCalls())
func (mock *vcsMock) GetCalls() []struct {
	Key string
} {
	var calls []struct {
		Key string
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// IsRepository calls IsRepositoryFunc.
func (mock *vcsMock) IsRepository() bool {
	if mock.IsRepositoryFunc == nil {
		panic("vcsMock.IsRepositoryFunc: method is nil but vcs.IsRepository was just called")
	}
	callInfo := struct {
	}{}
	mock.lockIsRepository.Lock()
	mock.calls.IsRepository = append(mock.calls.IsRepository, callInfo)
	mock.lockIsRepository.Unlock()
	return mock.IsRepositoryFunc()
}

// IsRepositoryCalls gets all the calls that were made to IsRepository.
// Check the length with:
//     len(mockedvcs.IsRepositoryCalls())
func (mock *vcsMock) IsRepositoryCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockIsRepository.RLock()
	calls = mock.calls.IsRepository
	mock.lockIsRepository.RUnlock()
	return calls
}

// Set calls SetFunc.
func (mock *vcsMock) Set(key string, value string) error {
	if mock.SetFunc == nil {
		panic("vcsMock.SetFunc: method is nil but vcs.Set was just called")
	}
	callInfo := struct {
		Key   string
		Value string
	}{
		Key:   key,
		Value: value,
	}
	mock.lockSet.Lock()
	mock.calls.Set = append(mock.calls.Set, callInfo)
	mock.lockSet.Unlock()
	return mock.SetFunc(key, value)
}

// SetCalls gets all the calls that were made to Set.
// Check the length with:
//     len(mockedvcs.SetCalls())
func (mock *vcsMock) SetCalls() []struct {
	Key   string
	Value string
} {
	var calls []struct {
		Key   string
		Value string
	}
	mock.lockSet.RLock()
	calls = mock.calls.Set
	mock.lockSet.RUnlock()
	return calls
}
