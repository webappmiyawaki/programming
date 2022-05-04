package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//err設定
	file, err := os.Open("./lesson.go")

	if err != nil {
		log.Fatalln("Error!")
	}
	defer file.Close()

	data := make([]byte, 100)

	//err設定
	//なぜ設定済みなのに「:=」で再設定できるのか？＞＞
	//式の左辺のどれかが未設定状態からだと上書き可能
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))

	if err = os.Chdir("test"); err != nil {
		log.Fatalln("Error")
	}

}
