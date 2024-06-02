package db

import (
	"os"
	"path"

	"github.com/Yoseph-code/haken/internal/fs"
	safeMap "github.com/Yoseph-code/haken/internal/safe_map"
)

const (
	defaultSourceName string = "default"
	extFile           string = ".db"
	defaultDir        string = "haken"
)

type DB struct {
	sourceName string

	sm *safeMap.SafeMap[string]
}

func NewSource(sourceName ...string) (*DB, error) {
	s := defaultSourceName

	if len(sourceName) > 0 {
		s = sourceName[0]
	}

	pwd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	dirname := path.Join(pwd, defaultDir)
	filename := path.Join(pwd, defaultDir, s+extFile)

	if ok := fs.IsDirExist(dirname); !ok {
		err := fs.CreateDir(dirname)

		if err != nil {
			return nil, err
		}
	}

	if ok := fs.IsFileExist(filename); !ok {
		err := fs.CreateFile(filename)

		if err != nil {
			return nil, err
		}
	}

	return &DB{
		sourceName: filename,
		sm:         safeMap.NewSafeMap[string](),
	}, nil
}
