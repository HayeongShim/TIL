```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}


// func main : Go 프로그램의 시작점이 되는 부분.
// 컴파일러는 main package와 그 안에 있는 main function을 먼저 찾아 실행시킴.

// vscode에서 자동완성이 안되어서 settings.json에 아래 내용을 추가했다.
// "go.autocompleteUnimportedPackages": true

// Println, SayHello(예시) 등은 대문자로 시작하는 function이기 때문에
// export된 것이다.

```