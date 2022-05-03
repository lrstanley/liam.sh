// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"time"

	"github.com/go-faster/jx"
)

type CreateUserReq struct {
	CreateTime time.Time "json:\"create_time\""
	UpdateTime time.Time "json:\"update_time\""
	UserID     int       "json:\"user_id\""
	Login      string    "json:\"login\""
	Name       OptString "json:\"name\""
	AvatarURL  OptString "json:\"avatar_url\""
	Email      OptString "json:\"email\""
	Location   OptString "json:\"location\""
	Bio        OptString "json:\"bio\""
}

// DeleteUserNoContent is response for DeleteUser operation.
type DeleteUserNoContent struct{}

func (*DeleteUserNoContent) deleteUserRes() {}

type ListUserOKApplicationJSON []UserList

func (ListUserOKApplicationJSON) listUserRes() {}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

type R400 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R400) createUserRes() {}
func (*R400) deleteUserRes() {}
func (*R400) listUserRes()   {}
func (*R400) readUserRes()   {}
func (*R400) updateUserRes() {}

type R404 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R404) deleteUserRes() {}
func (*R404) listUserRes()   {}
func (*R404) readUserRes()   {}
func (*R404) updateUserRes() {}

type R409 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R409) createUserRes() {}
func (*R409) deleteUserRes() {}
func (*R409) listUserRes()   {}
func (*R409) readUserRes()   {}
func (*R409) updateUserRes() {}

type R500 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R500) createUserRes() {}
func (*R500) deleteUserRes() {}
func (*R500) listUserRes()   {}
func (*R500) readUserRes()   {}
func (*R500) updateUserRes() {}

type UpdateUserReq struct {
	UpdateTime OptDateTime "json:\"update_time\""
	Login      OptString   "json:\"login\""
	Name       OptString   "json:\"name\""
	AvatarURL  OptString   "json:\"avatar_url\""
	Email      OptString   "json:\"email\""
	Location   OptString   "json:\"location\""
	Bio        OptString   "json:\"bio\""
}

// Ref: #/components/schemas/UserCreate
type UserCreate struct {
	ID         int       "json:\"id\""
	CreateTime time.Time "json:\"create_time\""
	UpdateTime time.Time "json:\"update_time\""
	UserID     int       "json:\"user_id\""
	Login      string    "json:\"login\""
	Name       OptString "json:\"name\""
	AvatarURL  OptString "json:\"avatar_url\""
	Email      OptString "json:\"email\""
	Location   OptString "json:\"location\""
	Bio        OptString "json:\"bio\""
}

func (*UserCreate) createUserRes() {}

// Ref: #/components/schemas/UserList
type UserList struct {
	ID         int       "json:\"id\""
	CreateTime time.Time "json:\"create_time\""
	UpdateTime time.Time "json:\"update_time\""
	UserID     int       "json:\"user_id\""
	Login      string    "json:\"login\""
	Name       OptString "json:\"name\""
	AvatarURL  OptString "json:\"avatar_url\""
	Email      OptString "json:\"email\""
	Location   OptString "json:\"location\""
	Bio        OptString "json:\"bio\""
}

// Ref: #/components/schemas/UserRead
type UserRead struct {
	ID         int       "json:\"id\""
	CreateTime time.Time "json:\"create_time\""
	UpdateTime time.Time "json:\"update_time\""
	UserID     int       "json:\"user_id\""
	Login      string    "json:\"login\""
	Name       OptString "json:\"name\""
	AvatarURL  OptString "json:\"avatar_url\""
	Email      OptString "json:\"email\""
	Location   OptString "json:\"location\""
	Bio        OptString "json:\"bio\""
}

func (*UserRead) readUserRes() {}

// Ref: #/components/schemas/UserUpdate
type UserUpdate struct {
	ID         int       "json:\"id\""
	CreateTime time.Time "json:\"create_time\""
	UpdateTime time.Time "json:\"update_time\""
	UserID     int       "json:\"user_id\""
	Login      string    "json:\"login\""
	Name       OptString "json:\"name\""
	AvatarURL  OptString "json:\"avatar_url\""
	Email      OptString "json:\"email\""
	Location   OptString "json:\"location\""
	Bio        OptString "json:\"bio\""
}

func (*UserUpdate) updateUserRes() {}