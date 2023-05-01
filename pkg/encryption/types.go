package encryption

type encConfig struct {
	Cost int64
}

type EncItf interface {
	HashSalt(data []byte) (string, error)

	ValidateHashData(data, hashedData []byte) (bool, error)
}
