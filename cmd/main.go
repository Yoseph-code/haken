package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Permissions struct {
	Get    bool
	Post   bool
	Delete bool
	Put    bool
}

type UserRole struct {
	Username string
	Password string
	Roles    map[string]Permissions
}

// Função para calcular o hash SHA-256 de uma string
func sha256Hash(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

// Função para converter o mapa de roles para uma string JSON
func rolesToJSON(roles map[string]Permissions) (string, error) {
	data, err := json.Marshal(roles)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Função para converter uma string JSON para um mapa de roles
func jsonToRoles(data string) (map[string]Permissions, error) {
	var roles map[string]Permissions
	err := json.Unmarshal([]byte(data), &roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// Função para adicionar um usuário com roles ao buffer
func addUserToBuffer(username, password string, roles map[string]Permissions, buffer *strings.Builder) {
	hashedPassword := sha256Hash(password)
	rolesJSON, _ := rolesToJSON(roles)
	buffer.WriteString(username + ":" + hashedPassword + ":" + rolesJSON + "\n")
}

// Função para criptografar dados e salvar no arquivo
func encryptAndSave(key, iv []byte, data string) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("could not create cipher: %v", err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	copy(ciphertext[:aes.BlockSize], iv)

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(data))

	return ioutil.WriteFile("users.enc", ciphertext, 0644)
}

// Função para descriptografar dados do arquivo
func decryptAndLoad(key []byte, iv []byte) (string, error) {
	ciphertext, err := ioutil.ReadFile("users.enc")
	if err != nil {
		return "", fmt.Errorf("could not read file: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("could not create cipher: %v", err)
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	copy(iv, ciphertext[:aes.BlockSize])
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

// Função para verificar se um usuário existe com a senha e roles fornecidos
func checkUser(username, password string, roles map[string]Permissions, buffer string) bool {
	hashedPassword := sha256Hash(password)
	lines := strings.Split(buffer, "\n")

	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) == 3 {
			storedUsername := parts[0]
			storedPassword := parts[1]
			storedRolesJSON := parts[2]

			if storedUsername == username && storedPassword == hashedPassword {
				storedRoles, err := jsonToRoles(storedRolesJSON)
				if err != nil {
					continue
				}
				if equalRoles(storedRoles, roles) {
					return true
				}
			}
		}
	}

	return false
}

// Função para comparar dois mapas de roles
func equalRoles(a, b map[string]Permissions) bool {
	if len(a) != len(b) {
		return false
	}
	for key, valA := range a {
		valB, exists := b[key]
		if !exists || valA != valB {
			return false
		}
	}
	return true
}

func main() {
	key := []byte("mysecretkey12345678901234567890") // Chave de 256 bits (32 bytes)
	iv := make([]byte, aes.BlockSize)

	fmt.Println("Descriptografando e carregando dados...")
	buffer, err := decryptAndLoad(key, iv)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Erro ao carregar dados: %v\n", err)
			return
		}
		// Se o arquivo não existe, inicializa o buffer vazio
		buffer = ""
	}
	fmt.Println("Dados carregados:")
	fmt.Println(buffer)

	var username, password string
	var rolesInput string

	fmt.Print("Enter username: ")
	fmt.Scan(&username)
	fmt.Print("Enter password: ")
	fmt.Scan(&password)

	fmt.Println("Enter roles (format: role1:get=1,post=1,delete=0,put=1;role2:get=1,post=0,delete=1,put=0):")
	fmt.Scan(&rolesInput)

	// Parse rolesInput
	roles := make(map[string]Permissions)

	for _, role := range strings.Split(rolesInput, ";") {
		if role == "" {
			continue
		}

		parts := strings.Split(role, ":")

		if len(parts) != 2 {
			fmt.Println("Invalid role format")
			return
		}

		roleName := parts[0]
		permissionsStr := parts[1]

		var perm Permissions
		for _, permPart := range strings.Split(permissionsStr, ",") {
			permData := strings.Split(permPart, "=")
			if len(permData) != 2 {
				fmt.Println("Invalid permission format")
				return
			}
			switch permData[0] {
			case "get":
				perm.Get = permData[1] == "1"
			case "post":
				perm.Post = permData[1] == "1"
			case "delete":
				perm.Delete = permData[1] == "1"
			case "put":
				perm.Put = permData[1] == "1"
			}
		}
		roles[roleName] = perm
	}

	var buf strings.Builder
	buf.WriteString(buffer)
	addUserToBuffer(username, password, roles, &buf)

	fmt.Println("Criptografando e salvando dados...")
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Fprintf(os.Stderr, "Could not generate IV: %v\n", err)
		return
	}
	if err := encryptAndSave(key, iv, buf.String()); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao salvar dados: %v\n", err)
		return
	}

	fmt.Println("Verificando usuário...")
	if checkUser(username, password, roles, buf.String()) {
		fmt.Printf("User %s verified successfully.\n", username)
	} else {
		fmt.Printf("User %s not found or incorrect password or roles.\n", username)
	}
}

// package main

// import (
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"crypto/rand"
// 	"crypto/sha256"
// 	"encoding/hex"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"os"
// 	"strings"
// )

// // Função para calcular o hash SHA-256 de uma string
// func sha256Hash(str string) string {
// 	hash := sha256.Sum256([]byte(str))
// 	return hex.EncodeToString(hash[:])
// }

// // Função para adicionar um usuário com roles ao buffer
// func addUserToBuffer(username, password, roles string, buffer *strings.Builder) {
// 	hashedPassword := sha256Hash(password)
// 	buffer.WriteString(username + ":" + hashedPassword + ":" + roles + "\n")
// }

// // Função para criptografar dados e salvar no arquivo
// func encryptAndSave(key, iv []byte, data string) error {
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return fmt.Errorf("could not create cipher: %v", err)
// 	}

// 	ciphertext := make([]byte, aes.BlockSize+len(data))
// 	copy(ciphertext[:aes.BlockSize], iv)

// 	stream := cipher.NewCFBEncrypter(block, iv)
// 	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(data))

// 	return ioutil.WriteFile("users.enc", ciphertext, 0644)
// }

// // Função para descriptografar dados do arquivo
// func decryptAndLoad(key []byte, iv []byte) (string, error) {
// 	ciphertext, err := ioutil.ReadFile("users.enc")
// 	if err != nil {
// 		return "", fmt.Errorf("could not read file: %v", err)
// 	}

// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return "", fmt.Errorf("could not create cipher: %v", err)
// 	}

// 	if len(ciphertext) < aes.BlockSize {
// 		return "", fmt.Errorf("ciphertext too short")
// 	}

// 	copy(iv, ciphertext[:aes.BlockSize])
// 	ciphertext = ciphertext[aes.BlockSize:]

// 	stream := cipher.NewCFBDecrypter(block, iv)
// 	stream.XORKeyStream(ciphertext, ciphertext)

// 	return string(ciphertext), nil
// }

// // Função para verificar se um usuário existe com a senha e roles fornecidos
// func checkUser(username, password, roles, buffer string) bool {
// 	hashedPassword := sha256Hash(password)
// 	lines := strings.Split(buffer, "\n")

// 	for _, line := range lines {
// 		parts := strings.Split(line, ":")
// 		if len(parts) == 3 && parts[0] == username && parts[1] == hashedPassword && parts[2] == roles {
// 			return true
// 		}
// 	}

// 	return false
// }

// func main() {
// 	key := []byte("mysecretkey12345678901234567890") // Chave de 256 bits (32 bytes)
// 	iv := make([]byte, aes.BlockSize)

// 	fmt.Println("Descriptografando e carregando dados...")
// 	buffer, err := decryptAndLoad(key, iv)
// 	if err != nil {
// 		if !os.IsNotExist(err) {
// 			fmt.Fprintf(os.Stderr, "Erro ao carregar dados: %v\n", err)
// 			return
// 		}
// 		// Se o arquivo não existe, inicializa o buffer vazio
// 		buffer = ""
// 	}
// 	fmt.Println("Dados carregados:")
// 	fmt.Println(buffer)

// 	var username, password, roles string
// 	fmt.Print("Enter username: ")
// 	fmt.Scan(&username)
// 	fmt.Print("Enter password: ")
// 	fmt.Scan(&password)
// 	fmt.Print("Enter roles (comma-separated): ")
// 	fmt.Scan(&roles)

// 	var buf strings.Builder
// 	buf.WriteString(buffer)
// 	addUserToBuffer(username, password, roles, &buf)

// 	fmt.Println("Criptografando e salvando dados...")
// 	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
// 		fmt.Fprintf(os.Stderr, "Could not generate IV: %v\n", err)
// 		return
// 	}
// 	if err := encryptAndSave(key, iv, buf.String()); err != nil {
// 		fmt.Fprintf(os.Stderr, "Erro ao salvar dados: %v\n", err)
// 		return
// 	}

// 	fmt.Println("Verificando usuário...")
// 	if checkUser(username, password, roles, buf.String()) {
// 		fmt.Printf("User %s verified successfully.\n", username)
// 	} else {
// 		fmt.Printf("User %s not found or incorrect password or roles.\n", username)
// 	}
// }

// import (
// 	"flag"
// 	"log"

// 	"github.com/Yoseph-code/haken/cli"
// 	"github.com/Yoseph-code/haken/config"
// 	"github.com/Yoseph-code/haken/db"
// 	"github.com/Yoseph-code/haken/server"
// )

// func init() {
// 	config.DefineFlags()
// }

// func main() {
// 	flag.Parse()

// 	isServer := flag.Lookup(config.Server).Value.(flag.Getter).Get().(bool)

// 	if isServer {
// 		port := flag.Lookup(config.Port).Value.(flag.Getter).Get().(uint)

// 		s := server.New(server.Config{
// 			ListenAddr: port,
// 		})

// 		f, err := db.NewDBFile(
// 			config.MainPath,
// 			config.MainFile,
// 		)

// 		if err != nil {
// 			log.Panic(err)
// 		}

// 		if ok := f.IsDBExists(); ok {
// 			s.SetDB(f)
// 		} else {
// 			file, err := f.CreateDB()

// 			if err != nil {
// 				log.Panic(err)
// 			}

// 			defer file.Close()

// 			s.SetDB(f)
// 		}

// 		if err := s.Run(); err != nil {
// 			log.Panic(err)
// 		}
// 	} else {
// 		user := flag.Lookup(config.User).Value.(flag.Getter).Get().(string)
// 		secret := flag.Lookup(config.Secret).Value.(flag.Getter).Get().(string)
// 		port := flag.Lookup(config.Port).Value.(flag.Getter).Get().(uint)

// 		c := cli.NewCli(user, secret, config.DefaultAddr, port)

// 		if err := c.Connect(); err != nil {
// 			log.Panic(err)
// 		}

// 		if err := c.Run(); err != nil {
// 			log.Panic(err)
// 		}
// 	}
// }
