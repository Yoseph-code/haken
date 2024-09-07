package users

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyName = errors.New("empty name")
)

type User struct {
	name   string
	secret string
	flag   int
}

func NewUser(name, secret string, flag int) *User {
	return &User{
		name:   name,
		secret: secret,
		flag:   flag,
	}
}

func (u *User) Valid() error {
	if u.name == "" {
		return ErrEmptyName
	}

	return nil
}

func (u *User) String() string {
	return fmt.Sprintf("%s:%s:%d", u.name, u.secret, u.flag)
}

func (u *User) Bytes() []byte {
	return []byte(u.String())
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Secret() string {
	return u.secret
}

func (u *User) Flag() int {
	return u.flag
}

func (u *User) ChangeFlag(flag int) {
	u.flag = flag
}

func (u *User) HasFlag(flag int) bool {
	return u.flag&flag != 0
}

func (u *User) SetFlag(flag int) {
	u.flag |= flag
}

func (u *User) ClearFlag(flag int) {
	u.flag &= ^flag
}

// func salt(size int) (string, error) {
// 	salt := make([]byte, size)

// 	_, err := rand.Read(salt)

// 	if err != nil {
// 		return "", err
// 	}

// 	return hex.EncodeToString(salt), nil
// }

// package main

// import (
// 	"crypto/hmac"
// 	"crypto/rand"
// 	"crypto/sha256"
// 	"encoding/gob"
// 	"encoding/hex"
// 	"fmt"
// 	"log"
// 	"os"

// 	"golang.org/x/crypto/pbkdf2"
// )

// // User representa a estrutura do usuário com nome de usuário e hash de senha
// type User struct {
// 	Username string
// 	Salt     string
// 	Hash     string
// }

// // generateSalt cria um novo salt aleatório
// func generateSalt(size int) (string, error) {
// 	salt := make([]byte, size)
// 	_, err := rand.Read(salt)
// 	if err != nil {
// 		return "", err
// 	}
// 	return hex.EncodeToString(salt), nil
// }

// // hashPassword cria um hash SCRAM-SHA-256 da senha usando um salt
// func hashPassword(password, salt string) string {
// 	saltBytes, _ := hex.DecodeString(salt)
// 	hash := pbkdf2.Key([]byte(password), saltBytes, 4096, sha256.Size, sha256.New)
// 	return hex.EncodeToString(hash)
// }

// // createUser cria um novo usuário com uma senha hash
// func createUser(username, password string) (*User, error) {
// 	salt, err := generateSalt(16)
// 	if err != nil {
// 		return nil, err
// 	}

// 	hash := hashPassword(password, salt)
// 	return &User{Username: username, Salt: salt, Hash: hash}, nil
// }

// // verifyPassword verifica a senha fornecida com o hash armazenado
// func verifyPassword(password string, user *User) bool {
// 	hash := hashPassword(password, user.Salt)
// 	return hmac.Equal([]byte(hash), []byte(user.Hash))
// }

// // saveUsers salva a lista de usuários em um arquivo binário
// func saveUsers(filename string, users []User) error {
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	encoder := gob.NewEncoder(file)
// 	return encoder.Encode(users)
// }

// // loadUsers carrega a lista de usuários de um arquivo binário
// func loadUsers(filename string) ([]User, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	var users []User
// 	decoder := gob.NewDecoder(file)
// 	err = decoder.Decode(&users)
// 	return users, err
// }

// func main() {
// 	// Criação de um novo usuário
// 	user, err := createUser("root", "root123")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Carregar usuários existentes (se houver)
// 	users, err := loadUsers("users.bin")
// 	if err != nil && !os.IsNotExist(err) {
// 		log.Fatal(err)
// 	}

// 	// Adicionar o novo usuário à lista de usuários
// 	users = append(users, *user)

// 	// Salvar os usuários no arquivo binário
// 	err = saveUsers("users.bin", users)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Usuário salvo com sucesso.")

// 	// Verificar a senha do usuário
// 	loadedUsers, err := loadUsers("users.bin")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, u := range loadedUsers {
// 		if u.Username == "root" {
// 			match := verifyPassword("root123", &u)
// 			fmt.Println("Senha correta:", match)
// 		}
// 	}
// }
