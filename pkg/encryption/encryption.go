package encryption

import (
	"deall-alfon/pkg/config"
	"deall-alfon/pkg/util/converter"
	"golang.org/x/crypto/bcrypt"
)

type enc struct {
	cfg encConfig
}

func NewEncryption(cfg config.ConfigStr) EncItf {
	return &enc{
		cfg: encConfig{
			Cost: cfg.Encrpytion.Cost,
		},
	}
}

func (e *enc) HashSalt(data []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(data, converter.ToInt(e.cfg.Cost))
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (e *enc) ValidateHashData(data, hashedData []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hashedData, data)
	if err != nil {
		return false, err
	}

	return true, nil
}
