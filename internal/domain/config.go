package domain

//go:generate mockgen -destination=../mocks/mock_config.go -package=mocks -source=./config.go
type Config interface {
	Init() error
	Print() string
	Validate() error
}
