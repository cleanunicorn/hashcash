package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"time"
)

func main() {
	prefix := "1:20:1807201527:daniel@gmail.com::McMybZIhxKXu57jd:"

	start := time.Now()
	counter := 0
	for {
		counter++

		h := sha1.New()
		io.WriteString(h, prefix+base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", counter))))
		hash := h.Sum(nil)

		hashStr := fmt.Sprintf("%x", hash)
		if hashStr[:5] == "00000" {
			fmt.Println(counter, hashStr)
			break
		}
	}
	fmt.Println("Time:", time.Since(start))
}
