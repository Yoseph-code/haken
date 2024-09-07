package users

const (
	Read       int = 1 << iota // 1
	Write                      // 2
	Overwrite                  // 4
	Remove                     // 8
	SuperAdmin                 // 16
)
