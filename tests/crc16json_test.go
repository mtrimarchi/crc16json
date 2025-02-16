package crc16json_test

import (
	"testing"

	"github.com/mtrimarchi/crc16json"
)

func TestCRC16Calculate(t *testing.T) {
	data := []byte("test data")
	expected := uint16(0x73d2)
	result := crc16json.CRC16Calculate(data, crc16json.CRC16Seed)
	if result != expected {
		t.Errorf("Expected %04x, got %04x", expected, result)
	}
}

func TestAddCRC(t *testing.T) {
	input := `{"data":"example","CRC_16":"0x0000"}`
	expected := `{"data":"example","CRC_16":"0x016f"}`
	result := crc16json.AddCRC(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
