```go
package main

import (
	"fmt"
	"strings"
)

// func main : Go 프로그램의 시작점이 되는 부분.
// 컴파일러는 main package와 그 안에 있는 main function을 먼저 찾아 실행시킴.

// vscode에서 자동완성이 안되어서 settings.json에 아래 내용을 추가했다.
// "go.autocompleteUnimportedPackages": true

// Println, SayHello(예시) 등은 대문자로 시작하는 function이기 때문에
// 다른 패키지에서 export가 가능하다.

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpper_naked(name string) (length int, upperName string) { // 어떤 variables를 return할지 정의해둔 상태
	length = len(name)
	upperName = strings.ToUpper(name)
	return
}

func lenAndUpper_defer(name string) (length int, upperName string) {
	defer fmt.Println("I'm done") // function이 값을 return하고 난 후 실행된다.

	length = len(name)
	upperName = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) (sum int) {
	sum = 0
	for _, number := range numbers { // range는 index와 값을 반환한다.
		sum += number
	}
	return
}

func main() {

	// ----------------------------------------------------------------
	// Go는 type language이다.
	// 축약형은 오로지 func 안에서만 가능하고 변수에만 적용 가능하다.
	// ----------------------------------------------------------------
	const name1 string = "harry" // constant
	fmt.Println(name1)

	var name2 string = "dobby" // variable
	name2 = "house-elf"
	fmt.Println(name2)

	name3 := "hedwig" // Go is going to guess the type
	fmt.Println(name3)

	fmt.Println("multiply(2,2)", multiply(2, 2))

	// ----------------------------------------------------------------
	// function의 return value가 여러개일 수 있다.
	// return value를 무시하고 싶으면 _를 사용한다.
	// ...을 사용해서 여러개의 인자를 받을 수 있다.
	// ----------------------------------------------------------------
	totalLength, upperName := lenAndUpper("Harry Potter")
	fmt.Println("totalLength:", totalLength, ", upperName:", upperName)

	totalLength2, _ := lenAndUpper("Ron Weasley")
	fmt.Println("totalLength:", totalLength2)

	repeatMe("harry", "potter", "and", "the", "chamber", "of", "secret")

	// ----------------------------------------------------------------
	// naked return : return할 variable을 굳이 명시하지 않아도 된다.
	// ----------------------------------------------------------------
	totalLength3, upperName3 := lenAndUpper_naked("Hayeong Shim")
	fmt.Println("totalLength:", totalLength3, ", upperName:", upperName3)

	// ----------------------------------------------------------------
	// defer : function이 실행되고 난 후에 코드가 실행된다는 사실이 어썸!
	// ----------------------------------------------------------------
	totalLength4, upperName4 := lenAndUpper_defer("Teddy Bear")
	fmt.Println("totalLength:", totalLength4, ", upperName:", upperName4)

	// ----------------------------------------------------------------
	// for statement
	// ----------------------------------------------------------------
	sum := superAdd(1, 2, 3, 4, 5)
	fmt.Println("superAdd:", sum)
}
```