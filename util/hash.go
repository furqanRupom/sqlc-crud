package util

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type argonParams struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLen     uint32
	KeyLen      uint32
}

var Params *argonParams = &argonParams{
	Memory:      64 * 1024,
	Parallelism: 2,
	Iterations:  3,
	SaltLen:     16,
	KeyLen:      64,
}
var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

func HashPassword(password string, p *argonParams) (string, error) {
	salt, err := generateRandomBytes(p.SaltLen)
	if err != nil {
		return "", err
	}
	if len(password) == 0 {
		return "", nil
	}
	bytePass := []byte(password)
	hash := argon2.IDKey(bytePass, salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLen)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt) // -> RawStdEncoding -> for no padding
	b64Hash := base64.RawStdEncoding.EncodeToString(hash) // -> RawStdEncoding -> for no padding

	encodeHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64Salt, b64Hash)
	return encodeHash, nil
}
func ComparePassword(password string, encodeHash string) (match bool, error error) {
	p, salt, hash, err := decodeHash(encodeHash)
	if err != nil {
		return false, err
	}
	otherHash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLen)
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}

	return false, nil
}

func decodeHash(encodeHash string) (p *argonParams, salt, hash []byte, err error) {
	values := strings.Split(encodeHash, "$")
	if len(values) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}
	var version int
	_, err = fmt.Sscanf(values[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}
	params := &argonParams{}

	_, err = fmt.Sscanf(values[3], "m=%d,t=%d,p=%d", &params.Memory, &params.Iterations, &params.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}
	salt, err = base64.RawStdEncoding.Strict().DecodeString(values[4])
	if err != nil {
		return nil, nil, nil, err
	}
	params.SaltLen = uint32(len(salt))
	hash, err = base64.RawStdEncoding.Strict().DecodeString(values[5])
	if err != nil {
		return nil, nil, nil, err
	}
	params.KeyLen = uint32(len(hash))

	return params, salt, hash, nil
}
func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}