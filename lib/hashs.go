package lib

import (
	"hash/crc32"
)

// HashUint функция для хеша
func HashUint(data interface{}) (uint32, error) {
	hashData := crc32.NewIEEE()
	jsonRead, err := new(JSONFile).JSONEncode(data)
	if err != nil {
		return 0, err
	}
	hashData.Write([]byte(jsonRead))
	return hashData.Sum32(), nil
}
