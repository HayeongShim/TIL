package main

import (
	"fmt"

	"github.com/HayeongShim/learngo/dict"
)

//-----------------BANK & DICTIONARY PROJECTS-----------------//

func main() {
	dictionary := dict.Dictionary{"first": "First word"}

	result_1, err_1 := dictionary.Search("second")
	if err_1 != nil {
		fmt.Println(err_1)
	} else {
		fmt.Println(result_1)
	}

	err_2 := dictionary.Add("second", "Second word")
	if err_2 != nil {
		fmt.Println(err_2)
	}

	err_3 := dictionary.UpdateDef("second", "Fake! Third word")
	if err_3 != nil {
		fmt.Println(err_3)
	}

	result_2, err_4 := dictionary.Search("second")
	if err_4 != nil {
		fmt.Println(err_4)
	} else {
		fmt.Println(result_2)
	}

	err_5 := dictionary.Delete("first")
	if err_5 != nil {
		fmt.Println(err_5)
	}

	result_3, err_6 := dictionary.Search("first")
	if err_6 != nil {
		fmt.Println(err_6)
	} else {
		fmt.Println(result_3)
	}
}
