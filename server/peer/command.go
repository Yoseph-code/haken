package peer

import (
	"bytes"
	"errors"
)

const (
	CREATE = "CREATE"
	READ   = "READ"
	UPDATE = "UPDATE"
	REMOVE = "REMOVE"
	PING   = "PING"
	OK     = "OK"
)

type Command interface{}

type CreateCommand struct {
	Key string
	Val []byte
}

type ReadCommand struct {
	Key string
}

type UpdateCommand struct {
	Key string
	Val []byte
}

type RemoveCommand struct {
	Key string
}

type PingCommand struct {
	Val string
}

func NewCommand(fields [][]byte) (Command, error) {
	if len(fields) == 0 {
		return nil, errors.New("invalid command with 0 len")
	}

	if len(fields[0]) == 0 {
		return nil, errors.New("invalid command key with 0 len")
	}

	key := fields[0]

	switch {
	case bytes.Equal(key, []byte(PING)):
		return PingCommand{
			Val: "PONG",
		}, nil
	case bytes.Equal(key, []byte(CREATE)):
		if len(fields[1:]) < 2 {
			return nil, errors.New("invalid command")
		}

		return CreateCommand{
			Key: string(fields[1]),
			Val: bytes.Join(fields[2:], []byte{' '}),
		}, nil
	case bytes.Equal(key, []byte(READ)):
		if len(fields[1:]) < 1 {
			return nil, errors.New("invalid command")
		}

		return ReadCommand{
			Key: string(fields[1]),
		}, nil
	case bytes.Equal(key, []byte(UPDATE)):
		if len(fields[1:]) < 2 {
			return nil, errors.New("invalid command")
		}

		return UpdateCommand{
			Key: string(fields[1]),
			Val: bytes.Join(fields[2:], []byte{' '}),
		}, nil
	case bytes.Equal(key, []byte(REMOVE)):
		if len(fields[1:]) < 1 {
			return nil, errors.New("invalid command")
		}

		return RemoveCommand{
			Key: string(fields[1]),
		}, nil
	}

	// switch fields[0] {
	// case CREATE:
	// 	if len(fields[1:]) < 2 {
	// 		return nil
	// 	}
	// 	return &CreateCommand{
	// 		Key: fields[1],
	// 		Val: fields[2],
	// 	}
	// case READ:
	// 	if len(fields[1:]) < 1 {
	// 		return nil
	// 	}
	// 	return &ReadCommand{
	// 		Key: fields[1],
	// 	}
	// case UPDATE:
	// 	if len(fields[1:]) < 2 {
	// 		return nil
	// 	}
	// 	return &UpdateCommand{
	// 		Key: fields[1],
	// 		Val: fields[2],
	// 	}
	// case REMOVE:
	// 	if len(fields[1:]) < 1 {
	// 		return nil
	// 	}
	// 	return &RemoveCommand{
	// 		Key: fields[1],
	// 	}
	// case PING:
	// 	if len(fields[1]) < 1 {
	// 		return nil
	// 	}
	// 	return &PingCommand{
	// 		Val: fields[1],
	// 	}
	// }

	return nil, errors.New("invalid command")
}
