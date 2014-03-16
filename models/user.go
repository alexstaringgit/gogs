// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dchest/scrypt"
	"github.com/gogits/gogs/modules/base"
	git "github.com/libgit2/git2go"
)

var UserPasswdSalt string

func init() {
	UserPasswdSalt = base.Cfg.MustValue("security", "USER_PASSWD_SALT")
}

// User types.
const (
	UT_INDIVIDUAL = iota + 1
	UT_ORGANIZATION
)

// Login types.
const (
	LT_PLAIN = iota + 1
	LT_LDAP
)

// A User represents the object of individual and member of organization.
type User struct {
	Id            int64
	LowerName     string `xorm:"unique not null"`
	Name          string `xorm:"unique not null"`
	Email         string `xorm:"unique not null"`
	Passwd        string `xorm:"not null"`
	LoginType     int
	Type          int
	NumFollowers  int
	NumFollowings int
	NumStars      int
	NumRepos      int
	Avatar        string `xorm:"varchar(2048) not null"`
	AvatarEmail   string `xorm:"not null"`
	Location      string
	Website       string
	Created       time.Time `xorm:"created"`
	Updated       time.Time `xorm:"updated"`
}

func (user *User) HomeLink() string {
	return "/user/" + user.LowerName
}

func (user *User) AvatarLink() string {
	return "http://1.gravatar.com/avatar/" + user.Avatar
}

// A Follow represents
type Follow struct {
	Id       int64
	UserId   int64     `xorm:"unique(s)"`
	FollowId int64     `xorm:"unique(s)"`
	Created  time.Time `xorm:"created"`
}

var (
	ErrUserOwnRepos     = errors.New("User still have ownership of repositories")
	ErrUserAlreadyExist = errors.New("User already exist")
	ErrUserNotExist     = errors.New("User does not exist")
	ErrEmailAlreadyUsed = errors.New("E-mail already used")
)

// IsUserExist checks if given user name exist,
// the user name should be noncased unique.
func IsUserExist(name string) (bool, error) {
	return orm.Get(&User{LowerName: strings.ToLower(name)})
}

func IsEmailUsed(email string) (bool, error) {
	return orm.Get(&User{Email: email})
}

func (user *User) NewGitSig() *git.Signature {
	return &git.Signature{
		Name:  user.Name,
		Email: user.Email,
		When:  time.Now(),
	}
}

// RegisterUser creates record of a new user.
func RegisterUser(user *User) (err error) {
	isExist, err := IsUserExist(user.Name)
	if err != nil {
		return err
	} else if isExist {
		return ErrUserAlreadyExist
	}

	isExist, err = IsEmailUsed(user.Email)
	if err != nil {
		return err
	} else if isExist {
		return ErrEmailAlreadyUsed
	}

	user.LowerName = strings.ToLower(user.Name)
	user.Avatar = base.EncodeMd5(user.Email)
	user.AvatarEmail = user.Email
	if err = user.EncodePasswd(); err != nil {
		return err
	}
	if _, err = orm.Insert(user); err != nil {
		return err
	}

	if err = os.MkdirAll(UserPath(user.Name), os.ModePerm); err != nil {

		if _, err := orm.Id(user.Id).Delete(&User{}); err != nil {
			return errors.New(fmt.Sprintf(
				"both create userpath %s and delete table record faild", user.Name))
		}
		return err
	}
	return nil
}

// UpdateUser updates user's information.
func UpdateUser(user *User) (err error) {
	_, err = orm.Id(user.Id).Update(user)
	return err
}

// DeleteUser completely deletes everything of the user.
func DeleteUser(user *User) error {
	// Check ownership of repository.
	count, err := GetRepositoryCount(user)
	if err != nil {
		return errors.New("modesl.GetRepositories: " + err.Error())
	} else if count > 0 {
		return ErrUserOwnRepos
	}

	// TODO: check issues, other repos' commits

	// Delete all feeds.
	if _, err = orm.Delete(&Action{UserId: user.Id}); err != nil {
		return err
	}

	// Delete all SSH keys.
	keys := make([]PublicKey, 0, 10)
	if err = orm.Find(&keys, &PublicKey{OwnerId: user.Id}); err != nil {
		return err
	}
	for _, key := range keys {
		if err = DeletePublicKey(&key); err != nil {
			return err
		}
	}

	// Delete user directory.
	if err = os.RemoveAll(UserPath(user.Name)); err != nil {
		return err
	}

	_, err = orm.Delete(user)
	// TODO: delete and update follower information.
	return err
}

// EncodePasswd encodes password to safe format.
func (user *User) EncodePasswd() error {
	newPasswd, err := scrypt.Key([]byte(user.Passwd), []byte(UserPasswdSalt), 16384, 8, 1, 64)
	user.Passwd = fmt.Sprintf("%x", newPasswd)
	return err
}

func UserPath(userName string) string {
	return filepath.Join(RepoRootPath, userName)
}

func GetUserByKeyId(keyId int64) (*User, error) {
	user := new(User)
	has, err := orm.Sql("select a.* from user as a, public_key as b where a.id = b.owner_id and b.id=?", keyId).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		err = errors.New("not exist key owner")
		return nil, err
	}
	return user, nil
}

func GetUserById(id int64) (*User, error) {
	user := new(User)
	has, err := orm.Id(id).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrUserNotExist
	}
	return user, nil
}

func GetUserByName(name string) (*User, error) {
	if len(name) == 0 {
		return nil, ErrUserNotExist
	}
	user := &User{
		LowerName: strings.ToLower(name),
	}
	has, err := orm.Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrUserNotExist
	}
	return user, nil
}

// LoginUserPlain validates user by raw user name and password.
func LoginUserPlain(name, passwd string) (*User, error) {
	user := User{LowerName: strings.ToLower(name), Passwd: passwd}
	if err := user.EncodePasswd(); err != nil {
		return nil, err
	}

	has, err := orm.Get(&user)
	if !has {
		err = ErrUserNotExist
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FollowUser marks someone be another's follower.
func FollowUser(userId int64, followId int64) error {
	session := orm.NewSession()
	defer session.Close()
	session.Begin()
	_, err := session.Insert(&Follow{UserId: userId, FollowId: followId})
	if err != nil {
		session.Rollback()
		return err
	}
	_, err = session.Exec("update user set num_followers = num_followers + 1 where id = ?", followId)
	if err != nil {
		session.Rollback()
		return err
	}
	_, err = session.Exec("update user set num_followings = num_followings + 1 where id = ?", userId)
	if err != nil {
		session.Rollback()
		return err
	}
	return session.Commit()
}

// UnFollowUser unmarks someone be another's follower.
func UnFollowUser(userId int64, unFollowId int64) error {
	session := orm.NewSession()
	defer session.Close()
	session.Begin()
	_, err := session.Delete(&Follow{UserId: userId, FollowId: unFollowId})
	if err != nil {
		session.Rollback()
		return err
	}
	_, err = session.Exec("update user set num_followers = num_followers - 1 where id = ?", unFollowId)
	if err != nil {
		session.Rollback()
		return err
	}
	_, err = session.Exec("update user set num_followings = num_followings - 1 where id = ?", userId)
	if err != nil {
		session.Rollback()
		return err
	}
	return session.Commit()
}
