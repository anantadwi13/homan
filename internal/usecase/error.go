package usecase

type Error interface {
	Level() ErrorLevel
	error
}

type ErrorLevel int

var (
	ErrorLevelUser   = ErrorLevel(0)
	ErrorLevelSystem = ErrorLevel(1)
)

type useCaseError struct {
	level   ErrorLevel
	message string
}

func NewErrorUser(message string) Error {
	return &useCaseError{
		level:   ErrorLevelUser,
		message: message,
	}
}

func NewErrorSystem(message string) Error {
	return &useCaseError{
		level:   ErrorLevelSystem,
		message: message,
	}
}

func WrapErrorSystem(err error) Error {
	return &useCaseError{
		level:   ErrorLevelSystem,
		message: err.Error(),
	}
}

func (u *useCaseError) Level() ErrorLevel {
	return u.level
}

func (u *useCaseError) Error() string {
	return u.message
}
