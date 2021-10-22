package domain

type Config interface {
	Init() error
	Print() string
	Validate() error
}
