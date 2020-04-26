package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

func (c *Config) decrypt() {
	if c.CryptoKey == "" {
		return
	}
	c.CryptoKey = strings.Repeat(c.CryptoKey, (16/len(c.CryptoKey) + 1))[:16]

	if c.URL != "" {
		raw, err := base64.StdEncoding.DecodeString(c.URL)
		if err != nil {
			c.Log.Fatalf("fatal: cannot decode URL bytes")
		}
		dec, err := Decrypt(raw, []byte(c.CryptoKey))
		if err != nil {
			c.Log.Errorf("error: cannot decrypt URL with key")
			c.Log.Fatalf("fatal: your cryptokey [--key] is likely wrong.")
		}
		c.URL = string(dec)
	}

	if c.Username != "" {
		raw, err := base64.StdEncoding.DecodeString(c.Username)
		if err != nil {
			c.Log.Fatalf("fatal: cannot decode username bytes")
		}
		dec, err := Decrypt(raw, []byte(c.CryptoKey))
		if err != nil {
			c.Log.Errorf("error: cannot decrypt username bytes")
			c.Log.Fatalf("fatal: your cryptokey [--key] is likely wrong.")
		}

		c.Username = string(dec)
	}

	if c.Password != "" {
		raw, err := base64.StdEncoding.DecodeString(c.Password)
		if err != nil {
			c.Log.Fatalf("fatal: cannot decode password bytes")
		}
		dec, err := Decrypt(raw, []byte(c.CryptoKey))
		if err != nil {
			c.Log.Errorf("error: cannot decrypt password bytes")
			c.Log.Fatalf("fatal: your cryptokey [--key] is likely wrong.")
		}
		c.Password = string(dec)
	}

	return
}

func (c *Config) encrypt() {
	if c.CryptoKey == "" {
		c.Log.Fatalf("fatal: cannot encrypt config without --key.")
		return
	}

	c.CryptoKey = strings.Repeat(c.CryptoKey, (16/len(c.CryptoKey) + 1))[:16]

	r, err := Encrypt([]byte(c.Username), []byte(c.CryptoKey))
	if err != nil {
		c.Log.Fatalf("fatal: cannot encrypt username")
	}
	c.Username = string(base64.StdEncoding.EncodeToString([]byte(r)))

	r, err = Encrypt([]byte(c.Password), []byte(c.CryptoKey))
	if err != nil {
		c.Log.Fatalf("fatal: cannot encrypt password")
	}
	c.Password = string(base64.StdEncoding.EncodeToString([]byte(r)))

	r, err = Encrypt([]byte(c.URL), []byte(c.CryptoKey))
	if err != nil {
		c.Log.Fatalf("fatal: cannot encrypt URL")
	}
	c.URL = string(base64.StdEncoding.EncodeToString([]byte(r)))
}

// Encrypt takes the plaintext unencrypted and the key (byte arrays) and performs default AES encryption
func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// Decrypt takes the cyphered text and the key (byte arrays) and performs default AES decryption
func Decrypt(crypt []byte, key []byte) (out []byte, err error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return
	}

	nonceSize := gcm.NonceSize()
	if len(crypt) < nonceSize {
		err = fmt.Errorf("crypt text too short")
		return
	}

	nonce, crypt := crypt[:nonceSize], crypt[nonceSize:]
	out, err = gcm.Open(nil, nonce, crypt, nil)

	return
}
