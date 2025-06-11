package utils

import (
	"crypto/rand"
	"encoding/binary"
)

// 生成随机的 int32
func RandInt32() (int32, error) {
	var buf [4]byte // int32 是 4 字节
	_, err := rand.Read(buf[:])
	if err != nil {
		return 0, err
	}
	// 将字节转换为 int32（大端序）
	return int32(binary.BigEndian.Uint32(buf[:])), nil
}
