package advance

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randomString(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())

	b := make([]rune, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))] // letters[rand.Intn(15)]
	}

	return string(b)
	// 10 huruf , sisanya angka

	// encryption
}

func random() {
	data := randomString(15)

	// fmt.Println(time.Now().UTC().UnixNano())
	// fmt.Println(time.Now().Local().UnixNano())

	fmt.Println(data)
	// belajar random
	// rand.Seed(time.Now().UTC().UnixNano())

	// fmt.Println("random 1 ", rand.Int())
	// fmt.Println("random 2 ", rand.Int())
	// fmt.Println("random 3 ", rand.Int())
}
