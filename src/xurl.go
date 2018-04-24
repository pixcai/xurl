package main

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

const (
	To30Bit = 0x3fffffff
	ToIndex = 0x0000003d
	Table   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GetChecksum(url string) string {
	hash := md5.New()
	hash.Write([]byte(url))

	return hex.EncodeToString(hash.Sum(nil))
}

func Transform(longUrl string) (shortUrls [4]string, err error) {
	checksum := GetChecksum(longUrl)

	for i := 0; i < 32; i += 8 {
		value, err := strconv.ParseInt(checksum[i:i+8], 16, 64)
		if err != nil {
			return shortUrls, err
		}

		value = int64(To30Bit) & value
		shortUrl := make([]byte, 6)
		for j := 0; j < 6; j++ {
			shortUrl[j] = Table[ToIndex&value]
			value >>= 5
		}
		shortUrls[i/8] = string(shortUrl)
	}

	return shortUrls, nil
}
