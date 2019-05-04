package main

import (
	"encoding/json"
	"fmt"

	"github.com/dahengzhang/mysql/control"
)

func showAll() {
	result, err := control.SelectAll()
	if err != nil {
		panic(err.Error())
	}
	r, err := json.Marshal(result)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(r))
}

func searchOne() {
	result, err := control.SelectDB(1)
	if err != nil {
		panic(err.Error())
	}

	r, err := json.Marshal(result)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(r))
}

func updateData() {
	err := control.UpdateDB(2, "流浪地球法案", "刹车时代，逃逸时代，流浪时代，新太阳时代")
	if err != nil {
		panic(err.Error())
	}
}

func deleteData() {
	err := control.DeleteDB(1)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	showAll()
	updateData()
	showAll()
}
