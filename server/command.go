package server

import (
	"bytes"
	"fmt"

	"github.com/tidwall/resp"
)

const (
	GET  = string("GET")
	SET  = string("SET")
	PING = string("PING")
)

type Command interface{}

type SetCommand struct {
	Key, Val []byte
}

type GetCommand struct {
	Key []byte
}

type PingCommand struct {
	Value string
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
