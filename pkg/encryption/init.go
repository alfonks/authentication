package encryption

import (
	"deall-alfon/pkg/config"
	"sync"
)

var (
	dataEncryptionItfOnce sync.Once
	dataEncryptionItf     EncItf
)

func GetEncryption() EncItf {
	dataEncryptionItfOnce.Do(func() {
		dataEncryptionItf = NewEncryption(
			config.GetConfig(),
		)
	})
	return dataEncryptionItf
}
