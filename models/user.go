package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	// init() gets called before main()
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// ErrNotFound is returned when a resource cannot be found
	// in the database
	ErrNotFound = errors.New("models: resource not found")
	// ErrInvalidID is returned when no resource with a give
	// id exists
	ErrInvalidID = errors.New("models: ID provided was invalid")
)

func NewUserService(psqlinfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", psqlinfo)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return &UserService{
		db: db,
	}, nil
}

type UserService struct {
	db *gorm.DB
}

// ByID will look up by the id provided
// There are three possible outcome cases
// Case 1: When the user with the id is found
// 1 - user, nil
// Case 2: When there is no user found with that id
// 2 - nil, ErrNotFound
// Case 3: When there is an error connecting with the db. This
// will probably result in a 500 error
// 3 - nil, otherError
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

// ByEmail will look up by the email provided
// There are three possible outcome cases
// Case 1: When the user with the email is found
// 1 - user, nil
// Case 2: When there is no user found with that email
// 2 - nil, ErrNotFound
// Case 3: When there is an error connecting with the db. This
// will probably result in a 500 error
// 3 - nil, otherError
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

// first will query using the provided gorm.db and it will
// get the first item returned and place it into dst. If nothing
// is found in the query it will return an ErrNotFound
//
// Not the dst HAS to be a pointer
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}

// Close closes the user service db connection
func (us *UserService) Close() error {
	return us.db.Close()
}

// DestructiveReset drops the user table and
// rebuilds it
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}

// Update will update the provided user with all the
// data given in the user object
func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

// Delete will delete the user with the give id
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(user).Error
}

// Create will create the provided user
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

// User model contains the id, created at, deleted at
// name and email
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}
