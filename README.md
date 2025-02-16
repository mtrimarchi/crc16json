# crc16json

A Go library for calculating CRC-16 checksums and appending them to JSON strings.

## Installation

```sh
go get github.com/mtrimarchi/crc16json
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/mtrimarchi/crc16json"
)

func main() {
	jsonData := `{"data":"example","CRC_16":""}`
	updatedJson := crc16json.AddCRC(jsonData)
	fmt.Println(updatedJson)
}
```

## Functions

### `CRC16Calculate(data []byte, seed uint16) uint16`
Computes the CRC-16 checksum for the given byte slice using a specified seed.

### `CRC(data string) string`
Calculates the CRC-16 checksum for a JSON string, considering only the portion before the `"CRC_16"` key.

### `AddCRC(jsonString string) string`
Appends the computed CRC-16 checksum to the provided JSON string.

## Running Tests

```sh
go test ./tests
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
