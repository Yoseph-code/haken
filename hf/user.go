package hf

import (
	"bytes"
	"errors"

	"github.com/Yoseph-code/haken/internal/users"
)

var (
	breakLine = []byte("\n")
)

var (
	ErrUserExists = errors.New("user already exists")
)

type UserHakenFile struct {
	*HakenFile
}

func NewUserHakenFile() *UserHakenFile {
	return &UserHakenFile{
		HakenFile: NewHakenFile(),
	}
}

func (uhf *UserHakenFile) DB() error {
	uhf.SetFileName(encFile)

	return uhf.CreateOrRead()
}

func (uhf *UserHakenFile) LoadUsers() (map[string]users.User, error) {
	content, err := uhf.ReadAll()

	if err != nil {
		return nil, err
	}

	data := make(map[string]users.User)

	for _, u := range bytes.SplitN(content, breakLine, -1) {
		splits := bytes.SplitN(u, []byte(":"), 3)

		user := users.NewUser("", "", 0)

		for _, s := range splits {
			switch {
			case user.Name() == "":
				user = users.NewUser(string(s), "", 0)
			case user.Secret() == "":
				user = users.NewUser(user.Name(), string(s), 0)
			case user.Flag() == 0:
				user = users.NewUser(user.Name(), user.Secret(), int(s[0]))
			}
		}

		if user.Name() == "" {
			continue
		}

		data[user.Name()] = *user
	}

	return data, nil
}

func (uhf *UserHakenFile) SaveUser(user *users.User) error {
	users, err := uhf.LoadUsers()

	if err != nil {
		return err
	}

	if _, ok := users[user.Name()]; ok {
		return ErrUserExists
	}

	var b bytes.Buffer

	b.Write(user.Bytes())
	b.Write(breakLine)

	return uhf.AppendToFile(b.Bytes())
}
