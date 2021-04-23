package advance

import (
	"errors"
	"fmt"
	"strconv"
)

func convertStringToInt(data string) (int, error) {
	number, err := strconv.Atoi(data)

	if err != nil {
		return 0, errors.New("not a number")
	} else {
		return number, nil
	}
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("error occured :", r)
	} else {
		fmt.Println("app running perfect")
	}
}

func execute(val string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error occured :", r)
		} else {
			fmt.Println("app running perfect")
		}
	}()

	number, err := strconv.Atoi(val)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(number)
	}
}

func errPanicRecover() {
	// error

	var data = []string{"afis", "10", "20"}
	for _, val := range data {
		execute(val)
	}
	// IIFE recover
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("error occured :", r)
	// 	} else {
	// 		fmt.Println("app running perfect")
	// 	}
	// }()

	// var input string
	// fmt.Print("input some number : ")
	// fmt.Scanln(&input)

	// result, e := convertStringToInt(input)

	// if e != nil {
	// 	panic(e.Error())
	// 	// fmt.Println("eksekusi code selesai")
	// } else {
	// 	fmt.Println(result)
	// 	// fmt.Println("eksekusi code selesai")
	// }
}
