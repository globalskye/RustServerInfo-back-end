package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("avatar.bin")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}

func RetrieveROM(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	return bytes, err
}
