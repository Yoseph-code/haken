package hf

type DBHakenFile struct {
	*HakenFile
}

func NewDB() *DBHakenFile {
	return &DBHakenFile{
		HakenFile: NewHakenFile(),
	}
}

func (dhf *DBHakenFile) DB() error {
	dhf.SetFileName(encFile)

	return dhf.CreateOrRead()
}
