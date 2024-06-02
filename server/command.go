package server

import (
	"bytes"
	"fmt"

	"github.com/tidwall/resp"
)

const (
	OK     string = "OK"
	EMPTY  string = ""
	ERR    string = "ERR"
	CREATE string = "CREATE"
	READ   string = "READ"
	UPDATE string = "UPDATE"
	REMOVE string = "REMOVE"
	PING   string = "PING"
)

type Command interface{}

type CreateCommand struct {
	Key, Val string
}

type ReadCommand struct {
	Key string
}

type UpdateCommand struct {
	Key, Val string
}

type RemoveCommand struct {
	Key string
}

type PingCommand struct {
	Val string
}

func RespWriteMap(m map[string]string) []byte {
	buf := &bytes.Buffer{}
	buf.WriteString("%" + fmt.Sprintf("%d\r\n", len(m)))
	rw := resp.NewWriter(buf)
	for k, v := range m {
		rw.WriteString(k)
		rw.WriteString(":" + v)
	}
	return buf.Bytes()
}
