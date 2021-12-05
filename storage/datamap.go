package storage

import (
	"errors"
	"sync"
)

type User struct {
	PID   string `json:"pid"`
	Email string `json:"email"`
}

type UserDataMap struct {
	mu      sync.RWMutex
	dataMap map[string]*User
}

var (
	userDataMap UserDataMap
	once        sync.Once
)

func GetInstance() *UserDataMap {
	once.Do(func() {
		userDataMap.dataMap = make(map[string]*User)
	})
	return &userDataMap
}

func (u *UserDataMap) Set(user *User) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	if _, exist := u.dataMap[user.Email]; exist {
		return errors.New("already exist")
	}
	u.dataMap[user.Email] = user
	return nil
}

func (u *UserDataMap) GetByEmail(email string) User {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return User{
		PID:   u.dataMap[email].PID,
		Email: u.dataMap[email].Email,
	}
}

func (u *UserDataMap) DeleteByEmail(email string) {

}
