package dkgencryptkey

import (
	"testing"

	"github.com/DeAI-Artist/MintAI/core/services/keystore/keys"
)

func TestDKGEncryptKeys_ExportImport(t *testing.T) {
	keys.RunKeyExportImportTestcase(t, createKey, decryptKey)
}

func createKey() (keys.KeyType, error) {
	return New()
}

func decryptKey(keyJSON []byte, password string) (keys.KeyType, error) {
	return FromEncryptedJSON(keyJSON, password)
}
