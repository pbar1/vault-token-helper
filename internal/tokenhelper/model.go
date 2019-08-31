package tokenhelper

type TokenHelper interface {
	Get() ([]byte, error)
	Store([]byte) error
	Erase() error
}
