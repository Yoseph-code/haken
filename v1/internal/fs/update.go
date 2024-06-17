package fs

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

type KeyValue struct {
	Key   [32]byte // Tamanho fixo para a chave
	Value [64]byte // Tamanho fixo para o valor
}

func main() {
	// Abrir o arquivo binário para leitura e escrita
	file, err := os.OpenFile("seuarquivo.bin", os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	// Defina a chave e o novo valor
	searchKey := "chave"
	newValue := "novo_valor"

	// Converter a chave e o valor em bytes, garantindo que tenham o tamanho correto
	var keyBytes [32]byte
	copy(keyBytes[:], searchKey)
	var valueBytes [64]byte
	copy(valueBytes[:], newValue)

	// Tamanho total de cada entrada (tamanho da chave + tamanho do valor)
	entrySize := binary.Size(KeyValue{})

	// Tamanho do arquivo
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("Erro ao obter informações do arquivo: %v", err)
	}
	fileSize := fileInfo.Size()

	// Verificar todas as entradas no arquivo para encontrar a chave e atualizar o valor
	for i := int64(0); i < fileSize; i += int64(entrySize) {
		// Ler uma entrada do arquivo
		var entry KeyValue
		_, err := file.ReadAt(entry.Key[:], i)
		if err != nil {
			log.Fatalf("Erro ao ler a entrada do arquivo: %v", err)
		}

		// Se a chave corresponder à chave de pesquisa, atualize o valor
		if string(entry.Key[:]) == searchKey {
			_, err := file.WriteAt(valueBytes[:], i+int64(binary.Size(entry.Key)))
			if err != nil {
				log.Fatalf("Erro ao escrever o novo valor no arquivo: %v", err)
			}
			fmt.Printf("Valor atualizado para a chave '%s'.\n", searchKey)
			return
		}
	}

	fmt.Printf("Chave '%s' não encontrada no arquivo.\n", searchKey)
}
