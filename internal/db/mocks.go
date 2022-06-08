// Code generated by go-mockgen 1.2.0; DO NOT EDIT.

package db

import (
	"context"
	"sync"
)

// MockAccessTokensStore is a mock implementation of the AccessTokensStore
// interface (from the package gogs.io/gogs/internal/db) used for unit
// testing.
type MockAccessTokensStore struct {
	// CreateFunc is an instance of a mock function object controlling the
	// behavior of the method Create.
	CreateFunc *AccessTokensStoreCreateFunc
	// DeleteByIDFunc is an instance of a mock function object controlling
	// the behavior of the method DeleteByID.
	DeleteByIDFunc *AccessTokensStoreDeleteByIDFunc
	// GetBySHA1Func is an instance of a mock function object controlling
	// the behavior of the method GetBySHA1.
	GetBySHA1Func *AccessTokensStoreGetBySHA1Func
	// ListFunc is an instance of a mock function object controlling the
	// behavior of the method List.
	ListFunc *AccessTokensStoreListFunc
	// TouchFunc is an instance of a mock function object controlling the
	// behavior of the method Touch.
	TouchFunc *AccessTokensStoreTouchFunc
}

// NewMockAccessTokensStore creates a new mock of the AccessTokensStore
// interface. All methods return zero values for all results, unless
// overwritten.
func NewMockAccessTokensStore() *MockAccessTokensStore {
	return &MockAccessTokensStore{
		CreateFunc: &AccessTokensStoreCreateFunc{
			defaultHook: func(context.Context, int64, string) (r0 *AccessToken, r1 error) {
				return
			},
		},
		DeleteByIDFunc: &AccessTokensStoreDeleteByIDFunc{
			defaultHook: func(context.Context, int64, int64) (r0 error) {
				return
			},
		},
		GetBySHA1Func: &AccessTokensStoreGetBySHA1Func{
			defaultHook: func(context.Context, string) (r0 *AccessToken, r1 error) {
				return
			},
		},
		ListFunc: &AccessTokensStoreListFunc{
			defaultHook: func(context.Context, int64) (r0 []*AccessToken, r1 error) {
				return
			},
		},
		TouchFunc: &AccessTokensStoreTouchFunc{
			defaultHook: func(context.Context, int64) (r0 error) {
				return
			},
		},
	}
}

// NewStrictMockAccessTokensStore creates a new mock of the
// AccessTokensStore interface. All methods panic on invocation, unless
// overwritten.
func NewStrictMockAccessTokensStore() *MockAccessTokensStore {
	return &MockAccessTokensStore{
		CreateFunc: &AccessTokensStoreCreateFunc{
			defaultHook: func(context.Context, int64, string) (*AccessToken, error) {
				panic("unexpected invocation of MockAccessTokensStore.Create")
			},
		},
		DeleteByIDFunc: &AccessTokensStoreDeleteByIDFunc{
			defaultHook: func(context.Context, int64, int64) error {
				panic("unexpected invocation of MockAccessTokensStore.DeleteByID")
			},
		},
		GetBySHA1Func: &AccessTokensStoreGetBySHA1Func{
			defaultHook: func(context.Context, string) (*AccessToken, error) {
				panic("unexpected invocation of MockAccessTokensStore.GetBySHA1")
			},
		},
		ListFunc: &AccessTokensStoreListFunc{
			defaultHook: func(context.Context, int64) ([]*AccessToken, error) {
				panic("unexpected invocation of MockAccessTokensStore.List")
			},
		},
		TouchFunc: &AccessTokensStoreTouchFunc{
			defaultHook: func(context.Context, int64) error {
				panic("unexpected invocation of MockAccessTokensStore.Touch")
			},
		},
	}
}

// NewMockAccessTokensStoreFrom creates a new mock of the
// MockAccessTokensStore interface. All methods delegate to the given
// implementation, unless overwritten.
func NewMockAccessTokensStoreFrom(i AccessTokensStore) *MockAccessTokensStore {
	return &MockAccessTokensStore{
		CreateFunc: &AccessTokensStoreCreateFunc{
			defaultHook: i.Create,
		},
		DeleteByIDFunc: &AccessTokensStoreDeleteByIDFunc{
			defaultHook: i.DeleteByID,
		},
		GetBySHA1Func: &AccessTokensStoreGetBySHA1Func{
			defaultHook: i.GetBySHA1,
		},
		ListFunc: &AccessTokensStoreListFunc{
			defaultHook: i.List,
		},
		TouchFunc: &AccessTokensStoreTouchFunc{
			defaultHook: i.Touch,
		},
	}
}

// AccessTokensStoreCreateFunc describes the behavior when the Create method
// of the parent MockAccessTokensStore instance is invoked.
type AccessTokensStoreCreateFunc struct {
	defaultHook func(context.Context, int64, string) (*AccessToken, error)
	hooks       []func(context.Context, int64, string) (*AccessToken, error)
	history     []AccessTokensStoreCreateFuncCall
	mutex       sync.Mutex
}

// Create delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockAccessTokensStore) Create(v0 context.Context, v1 int64, v2 string) (*AccessToken, error) {
	r0, r1 := m.CreateFunc.nextHook()(v0, v1, v2)
	m.CreateFunc.appendCall(AccessTokensStoreCreateFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the Create method of the
// parent MockAccessTokensStore instance is invoked and the hook queue is
// empty.
func (f *AccessTokensStoreCreateFunc) SetDefaultHook(hook func(context.Context, int64, string) (*AccessToken, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Create method of the parent MockAccessTokensStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *AccessTokensStoreCreateFunc) PushHook(hook func(context.Context, int64, string) (*AccessToken, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *AccessTokensStoreCreateFunc) SetDefaultReturn(r0 *AccessToken, r1 error) {
	f.SetDefaultHook(func(context.Context, int64, string) (*AccessToken, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *AccessTokensStoreCreateFunc) PushReturn(r0 *AccessToken, r1 error) {
	f.PushHook(func(context.Context, int64, string) (*AccessToken, error) {
		return r0, r1
	})
}

func (f *AccessTokensStoreCreateFunc) nextHook() func(context.Context, int64, string) (*AccessToken, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *AccessTokensStoreCreateFunc) appendCall(r0 AccessTokensStoreCreateFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of AccessTokensStoreCreateFuncCall objects
// describing the invocations of this function.
func (f *AccessTokensStoreCreateFunc) History() []AccessTokensStoreCreateFuncCall {
	f.mutex.Lock()
	history := make([]AccessTokensStoreCreateFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// AccessTokensStoreCreateFuncCall is an object that describes an invocation
// of method Create on an instance of MockAccessTokensStore.
type AccessTokensStoreCreateFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int64
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *AccessToken
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c AccessTokensStoreCreateFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c AccessTokensStoreCreateFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// AccessTokensStoreDeleteByIDFunc describes the behavior when the
// DeleteByID method of the parent MockAccessTokensStore instance is
// invoked.
type AccessTokensStoreDeleteByIDFunc struct {
	defaultHook func(context.Context, int64, int64) error
	hooks       []func(context.Context, int64, int64) error
	history     []AccessTokensStoreDeleteByIDFuncCall
	mutex       sync.Mutex
}

// DeleteByID delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockAccessTokensStore) DeleteByID(v0 context.Context, v1 int64, v2 int64) error {
	r0 := m.DeleteByIDFunc.nextHook()(v0, v1, v2)
	m.DeleteByIDFunc.appendCall(AccessTokensStoreDeleteByIDFuncCall{v0, v1, v2, r0})
	return r0
}

// SetDefaultHook sets function that is called when the DeleteByID method of
// the parent MockAccessTokensStore instance is invoked and the hook queue
// is empty.
func (f *AccessTokensStoreDeleteByIDFunc) SetDefaultHook(hook func(context.Context, int64, int64) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// DeleteByID method of the parent MockAccessTokensStore instance invokes
// the hook at the front of the queue and discards it. After the queue is
// empty, the default hook function is invoked for any future action.
func (f *AccessTokensStoreDeleteByIDFunc) PushHook(hook func(context.Context, int64, int64) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *AccessTokensStoreDeleteByIDFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, int64, int64) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *AccessTokensStoreDeleteByIDFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, int64, int64) error {
		return r0
	})
}

func (f *AccessTokensStoreDeleteByIDFunc) nextHook() func(context.Context, int64, int64) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *AccessTokensStoreDeleteByIDFunc) appendCall(r0 AccessTokensStoreDeleteByIDFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of AccessTokensStoreDeleteByIDFuncCall objects
// describing the invocations of this function.
func (f *AccessTokensStoreDeleteByIDFunc) History() []AccessTokensStoreDeleteByIDFuncCall {
	f.mutex.Lock()
	history := make([]AccessTokensStoreDeleteByIDFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// AccessTokensStoreDeleteByIDFuncCall is an object that describes an
// invocation of method DeleteByID on an instance of MockAccessTokensStore.
type AccessTokensStoreDeleteByIDFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int64
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 int64
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c AccessTokensStoreDeleteByIDFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c AccessTokensStoreDeleteByIDFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// AccessTokensStoreGetBySHA1Func describes the behavior when the GetBySHA1
// method of the parent MockAccessTokensStore instance is invoked.
type AccessTokensStoreGetBySHA1Func struct {
	defaultHook func(context.Context, string) (*AccessToken, error)
	hooks       []func(context.Context, string) (*AccessToken, error)
	history     []AccessTokensStoreGetBySHA1FuncCall
	mutex       sync.Mutex
}

// GetBySHA1 delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockAccessTokensStore) GetBySHA1(v0 context.Context, v1 string) (*AccessToken, error) {
	r0, r1 := m.GetBySHA1Func.nextHook()(v0, v1)
	m.GetBySHA1Func.appendCall(AccessTokensStoreGetBySHA1FuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetBySHA1 method of
// the parent MockAccessTokensStore instance is invoked and the hook queue
// is empty.
func (f *AccessTokensStoreGetBySHA1Func) SetDefaultHook(hook func(context.Context, string) (*AccessToken, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetBySHA1 method of the parent MockAccessTokensStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *AccessTokensStoreGetBySHA1Func) PushHook(hook func(context.Context, string) (*AccessToken, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *AccessTokensStoreGetBySHA1Func) SetDefaultReturn(r0 *AccessToken, r1 error) {
	f.SetDefaultHook(func(context.Context, string) (*AccessToken, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *AccessTokensStoreGetBySHA1Func) PushReturn(r0 *AccessToken, r1 error) {
	f.PushHook(func(context.Context, string) (*AccessToken, error) {
		return r0, r1
	})
}

func (f *AccessTokensStoreGetBySHA1Func) nextHook() func(context.Context, string) (*AccessToken, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *AccessTokensStoreGetBySHA1Func) appendCall(r0 AccessTokensStoreGetBySHA1FuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of AccessTokensStoreGetBySHA1FuncCall objects
// describing the invocations of this function.
func (f *AccessTokensStoreGetBySHA1Func) History() []AccessTokensStoreGetBySHA1FuncCall {
	f.mutex.Lock()
	history := make([]AccessTokensStoreGetBySHA1FuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// AccessTokensStoreGetBySHA1FuncCall is an object that describes an
// invocation of method GetBySHA1 on an instance of MockAccessTokensStore.
type AccessTokensStoreGetBySHA1FuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *AccessToken
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c AccessTokensStoreGetBySHA1FuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c AccessTokensStoreGetBySHA1FuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// AccessTokensStoreListFunc describes the behavior when the List method of
// the parent MockAccessTokensStore instance is invoked.
type AccessTokensStoreListFunc struct {
	defaultHook func(context.Context, int64) ([]*AccessToken, error)
	hooks       []func(context.Context, int64) ([]*AccessToken, error)
	history     []AccessTokensStoreListFuncCall
	mutex       sync.Mutex
}

// List delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockAccessTokensStore) List(v0 context.Context, v1 int64) ([]*AccessToken, error) {
	r0, r1 := m.ListFunc.nextHook()(v0, v1)
	m.ListFunc.appendCall(AccessTokensStoreListFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the List method of the
// parent MockAccessTokensStore instance is invoked and the hook queue is
// empty.
func (f *AccessTokensStoreListFunc) SetDefaultHook(hook func(context.Context, int64) ([]*AccessToken, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// List method of the parent MockAccessTokensStore instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *AccessTokensStoreListFunc) PushHook(hook func(context.Context, int64) ([]*AccessToken, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *AccessTokensStoreListFunc) SetDefaultReturn(r0 []*AccessToken, r1 error) {
	f.SetDefaultHook(func(context.Context, int64) ([]*AccessToken, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *AccessTokensStoreListFunc) PushReturn(r0 []*AccessToken, r1 error) {
	f.PushHook(func(context.Context, int64) ([]*AccessToken, error) {
		return r0, r1
	})
}

func (f *AccessTokensStoreListFunc) nextHook() func(context.Context, int64) ([]*AccessToken, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *AccessTokensStoreListFunc) appendCall(r0 AccessTokensStoreListFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of AccessTokensStoreListFuncCall objects
// describing the invocations of this function.
func (f *AccessTokensStoreListFunc) History() []AccessTokensStoreListFuncCall {
	f.mutex.Lock()
	history := make([]AccessTokensStoreListFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// AccessTokensStoreListFuncCall is an object that describes an invocation
// of method List on an instance of MockAccessTokensStore.
type AccessTokensStoreListFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int64
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []*AccessToken
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c AccessTokensStoreListFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c AccessTokensStoreListFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// AccessTokensStoreTouchFunc describes the behavior when the Touch method
// of the parent MockAccessTokensStore instance is invoked.
type AccessTokensStoreTouchFunc struct {
	defaultHook func(context.Context, int64) error
	hooks       []func(context.Context, int64) error
	history     []AccessTokensStoreTouchFuncCall
	mutex       sync.Mutex
}

// Touch delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockAccessTokensStore) Touch(v0 context.Context, v1 int64) error {
	r0 := m.TouchFunc.nextHook()(v0, v1)
	m.TouchFunc.appendCall(AccessTokensStoreTouchFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Touch method of the
// parent MockAccessTokensStore instance is invoked and the hook queue is
// empty.
func (f *AccessTokensStoreTouchFunc) SetDefaultHook(hook func(context.Context, int64) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Touch method of the parent MockAccessTokensStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *AccessTokensStoreTouchFunc) PushHook(hook func(context.Context, int64) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *AccessTokensStoreTouchFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, int64) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *AccessTokensStoreTouchFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, int64) error {
		return r0
	})
}

func (f *AccessTokensStoreTouchFunc) nextHook() func(context.Context, int64) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *AccessTokensStoreTouchFunc) appendCall(r0 AccessTokensStoreTouchFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of AccessTokensStoreTouchFuncCall objects
// describing the invocations of this function.
func (f *AccessTokensStoreTouchFunc) History() []AccessTokensStoreTouchFuncCall {
	f.mutex.Lock()
	history := make([]AccessTokensStoreTouchFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// AccessTokensStoreTouchFuncCall is an object that describes an invocation
// of method Touch on an instance of MockAccessTokensStore.
type AccessTokensStoreTouchFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int64
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c AccessTokensStoreTouchFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c AccessTokensStoreTouchFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// MockPermsStore is a mock implementation of the PermsStore interface (from
// the package gogs.io/gogs/internal/db) used for unit testing.
type MockPermsStore struct {
	// AccessModeFunc is an instance of a mock function object controlling
	// the behavior of the method AccessMode.
	AccessModeFunc *PermsStoreAccessModeFunc
	// AuthorizeFunc is an instance of a mock function object controlling
	// the behavior of the method Authorize.
	AuthorizeFunc *PermsStoreAuthorizeFunc
	// SetRepoPermsFunc is an instance of a mock function object controlling
	// the behavior of the method SetRepoPerms.
	SetRepoPermsFunc *PermsStoreSetRepoPermsFunc
}

// NewMockPermsStore creates a new mock of the PermsStore interface. All
// methods return zero values for all results, unless overwritten.
func NewMockPermsStore() *MockPermsStore {
	return &MockPermsStore{
		AccessModeFunc: &PermsStoreAccessModeFunc{
			defaultHook: func(context.Context, int64, int64, AccessModeOptions) (r0 AccessMode) {
				return
			},
		},
		AuthorizeFunc: &PermsStoreAuthorizeFunc{
			defaultHook: func(context.Context, int64, int64, AccessMode, AccessModeOptions) (r0 bool) {
				return
			},
		},
		SetRepoPermsFunc: &PermsStoreSetRepoPermsFunc{
			defaultHook: func(context.Context, int64, map[int64]AccessMode) (r0 error) {
				return
			},
		},
	}
}

// NewStrictMockPermsStore creates a new mock of the PermsStore interface.
// All methods panic on invocation, unless overwritten.
func NewStrictMockPermsStore() *MockPermsStore {
	return &MockPermsStore{
		AccessModeFunc: &PermsStoreAccessModeFunc{
			defaultHook: func(context.Context, int64, int64, AccessModeOptions) AccessMode {
				panic("unexpected invocation of MockPermsStore.AccessMode")
			},
		},
		AuthorizeFunc: &PermsStoreAuthorizeFunc{
			defaultHook: func(context.Context, int64, int64, AccessMode, AccessModeOptions) bool {
				panic("unexpected invocation of MockPermsStore.Authorize")
			},
		},
		SetRepoPermsFunc: &PermsStoreSetRepoPermsFunc{
			defaultHook: func(context.Context, int64, map[int64]AccessMode) error {
				panic("unexpected invocation of MockPermsStore.SetRepoPerms")
			},
		},
	}
}

// NewMockPermsStoreFrom creates a new mock of the MockPermsStore interface.
// All methods delegate to the given implementation, unless overwritten.
func NewMockPermsStoreFrom(i PermsStore) *MockPermsStore {
	return &MockPermsStore{
		AccessModeFunc: &PermsStoreAccessModeFunc{
			defaultHook: i.AccessMode,
		},
		AuthorizeFunc: &PermsStoreAuthorizeFunc{
			defaultHook: i.Authorize,
		},
		SetRepoPermsFunc: &PermsStoreSetRepoPermsFunc{
			defaultHook: i.SetRepoPerms,
		},
	}
}

// PermsStoreAccessModeFunc describes the behavior when the AccessMode
// method of the parent MockPermsStore instance is invoked.
type PermsStoreAccessModeFunc struct {
	defaultHook func(context.Context, int64, int64, AccessModeOptions) AccessMode
	hooks       []func(context.Context, int64, int64, AccessModeOptions) AccessMode
	history     []PermsStoreAccessModeFuncCall
	mutex       sync.Mutex
}

// AccessMode delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockPermsStore) AccessMode(v0 context.Context, v1 int64, v2 int64, v3 AccessModeOptions) AccessMode {
	r0 := m.AccessModeFunc.nextHook()(v0, v1, v2, v3)
	m.AccessModeFunc.appendCall(PermsStoreAccessModeFuncCall{v0, v1, v2, v3, r0})
	return r0
}

// SetDefaultHook sets function that is called when the AccessMode method of
// the parent MockPermsStore instance is invoked and the hook queue is
// empty.
func (f *PermsStoreAccessModeFunc) SetDefaultHook(hook func(context.Context, int64, int64, AccessModeOptions) AccessMode) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// AccessMode method of the parent MockPermsStore instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *PermsStoreAccessModeFunc) PushHook(hook func(context.Context, int64, int64, AccessModeOptions) AccessMode) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *PermsStoreAccessModeFunc) SetDefaultReturn(r0 AccessMode) {
	f.SetDefaultHook(func(context.Context, int64, int64, AccessModeOptions) AccessMode {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *PermsStoreAccessModeFunc) PushReturn(r0 AccessMode) {
	f.PushHook(func(context.Context, int64, int64, AccessModeOptions) AccessMode {
		return r0
	})
}

func (f *PermsStoreAccessModeFunc) nextHook() func(context.Context, int64, int64, AccessModeOptions) AccessMode {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *PermsStoreAccessModeFunc) appendCall(r0 PermsStoreAccessModeFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of PermsStoreAccessModeFuncCall objects
// describing the invocations of this function.
func (f *PermsStoreAccessModeFunc) History() []PermsStoreAccessModeFuncCall {
	f.mutex.Lock()
	history := make([]PermsStoreAccessModeFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// PermsStoreAccessModeFuncCall is an object that describes an invocation of
// method AccessMode on an instance of MockPermsStore.
type PermsStoreAccessModeFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int64
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 int64
	// Arg3 is the value of the 4th argument passed to this method
	// invocation.
	Arg3 AccessModeOptions
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 AccessMode
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c PermsStoreAccessModeFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2, c.Arg3}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c PermsStoreAccessModeFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// PermsStoreAuthorizeFunc describes the behavior when the Authorize method
// of the parent MockPermsStore instance is invoked.
type PermsStoreAuthorizeFunc struct {
	defaultHook func(context.Context, int64, int64, AccessMode, AccessModeOptions) bool
	hooks       []func(context.Context, int64, int64, AccessMode, AccessModeOptions) bool
	history     []PermsStoreAuthorizeFuncCall
	mutex       sync.Mutex
}

// Authorize delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockPermsStore) Authorize(v0 context.Context, v1 int64, v2 int64, v3 AccessMode, v4 AccessModeOptions) bool {
	r0 := m.AuthorizeFunc.nextHook()(v0, v1, v2, v3, v4)
	m.AuthorizeFunc.appendCall(PermsStoreAuthorizeFuncCall{v0, v1, v2, v3, v4, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Authorize method of
// the parent MockPermsStore instance is invoked and the hook queue is
// empty.
func (f *PermsStoreAuthorizeFunc) SetDefaultHook(hook func(context.Context, int64, int64, AccessMode, AccessModeOptions) bool) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Authorize method of the parent MockPermsStore instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *PermsStoreAuthorizeFunc) PushHook(hook func(context.Context, int64, int64, AccessMode, AccessModeOptions) bool) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *PermsStoreAuthorizeFunc) SetDefaultReturn(r0 bool) {
	f.SetDefaultHook(func(context.Context, int64, int64, AccessMode, AccessModeOptions) bool {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *PermsStoreAuthorizeFunc) PushReturn(r0 bool) {
	f.PushHook(func(context.Context, int64, int64, AccessMode, AccessModeOptions) bool {
		return r0
	})
}

func (f *PermsStoreAuthorizeFunc) nextHook() func(context.Context, int64, int64, AccessMode, AccessModeOptions) bool {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *PermsStoreAuthorizeFunc) appendCall(r0 PermsStoreAuthorizeFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of PermsStoreAuthorizeFuncCall objects
// describing the invocations of this function.
func (f *PermsStoreAuthorizeFunc) History() []PermsStoreAuthorizeFuncCall {
	f.mutex.Lock()
	history := make([]PermsStoreAuthorizeFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// PermsStoreAuthorizeFuncCall is an object that describes an invocation of
// method Authorize on an instance of MockPermsStore.
type PermsStoreAuthorizeFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int64
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 int64
	// Arg3 is the value of the 4th argument passed to this method
	// invocation.
	Arg3 AccessMode
	// Arg4 is the value of the 5th argument passed to this method
	// invocation.
	Arg4 AccessModeOptions
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 bool
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c PermsStoreAuthorizeFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2, c.Arg3, c.Arg4}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c PermsStoreAuthorizeFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// PermsStoreSetRepoPermsFunc describes the behavior when the SetRepoPerms
// method of the parent MockPermsStore instance is invoked.
type PermsStoreSetRepoPermsFunc struct {
	defaultHook func(context.Context, int64, map[int64]AccessMode) error
	hooks       []func(context.Context, int64, map[int64]AccessMode) error
	history     []PermsStoreSetRepoPermsFuncCall
	mutex       sync.Mutex
}

// SetRepoPerms delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockPermsStore) SetRepoPerms(v0 context.Context, v1 int64, v2 map[int64]AccessMode) error {
	r0 := m.SetRepoPermsFunc.nextHook()(v0, v1, v2)
	m.SetRepoPermsFunc.appendCall(PermsStoreSetRepoPermsFuncCall{v0, v1, v2, r0})
	return r0
}

// SetDefaultHook sets function that is called when the SetRepoPerms method
// of the parent MockPermsStore instance is invoked and the hook queue is
// empty.
func (f *PermsStoreSetRepoPermsFunc) SetDefaultHook(hook func(context.Context, int64, map[int64]AccessMode) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// SetRepoPerms method of the parent MockPermsStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *PermsStoreSetRepoPermsFunc) PushHook(hook func(context.Context, int64, map[int64]AccessMode) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *PermsStoreSetRepoPermsFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, int64, map[int64]AccessMode) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *PermsStoreSetRepoPermsFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, int64, map[int64]AccessMode) error {
		return r0
	})
}

func (f *PermsStoreSetRepoPermsFunc) nextHook() func(context.Context, int64, map[int64]AccessMode) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *PermsStoreSetRepoPermsFunc) appendCall(r0 PermsStoreSetRepoPermsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of PermsStoreSetRepoPermsFuncCall objects
// describing the invocations of this function.
func (f *PermsStoreSetRepoPermsFunc) History() []PermsStoreSetRepoPermsFuncCall {
	f.mutex.Lock()
	history := make([]PermsStoreSetRepoPermsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// PermsStoreSetRepoPermsFuncCall is an object that describes an invocation
// of method SetRepoPerms on an instance of MockPermsStore.
type PermsStoreSetRepoPermsFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int64
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 map[int64]AccessMode
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c PermsStoreSetRepoPermsFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c PermsStoreSetRepoPermsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
