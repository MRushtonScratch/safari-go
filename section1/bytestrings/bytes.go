package bytestrings

import (
	"fmt"
	"bytes"
	"bufio"
)

func workWithBuffer() error {
	rawString := "It's easy to encode unicode into a byte array ♥"

	b := Buffer(rawString)

	fmt.Println(b.String())

	s, err := toString(b)
	if err != nil {
		return err
	}

	fmt.Println(s)

	reader := bytes.NewReader([]byte(rawString))
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil
}