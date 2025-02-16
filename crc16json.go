package crc16json

import (
	"fmt"
	"strings"
)

const (
	CRC16Seed = 0xFFFF
	CRC16Poly = 0x1021
)

// CRC16Calculate computes the CRC-16 checksum using the given seed.
func CRC16Calculate(data []byte, seed uint16) uint16 {
	crc := seed
	for _, b := range data {
		for i := 0x80; i > 0; i >>= 1 {
			flag := (crc & 0x8000) != 0
			crc <<= 1
			if b&byte(i) != 0 {
				crc++
			}
			if flag {
				crc ^= CRC16Poly
			}
		}
	}
	return crc
}

// CRC computes the CRC-16 checksum of a JSON string.
func CRC(data string) string {
	crcIndex := strings.LastIndex(data, "\"CRC_16\"") + len("\"CRC_16\"") + (len([]byte(data)) - len(data))
	crc := CRC16Calculate([]byte(data)[:crcIndex], CRC16Seed)
	return fmt.Sprintf("0x%04x", crc)
}

// AddCRC appends the computed CRC-16 checksum to the JSON string.
func AddCRC(jsonString string) string {
	crcValue := CRC(jsonString)
	crcIndex := strings.LastIndex(jsonString, "\"CRC_16\"") + len("\"CRC_16\":")
	return jsonString[:crcIndex] + "\"" + crcValue + "\"}"
}
