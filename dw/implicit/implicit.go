package implicit

import "fmt"

type Phone []int

// Phone 타입에 String 메소드를 binding 할 수 있음. (Kotlin 의 extension 과 언어 기능적으로 동일해 보임)
func (p Phone) String() string {
	text := ""
	for _, digit := range p {
		text += string(rune('0' + digit))
	}
	return text
}

// Stringer : interface 선언, 그러나 Phone 타입에 명시적으로 interface 를 구현한다는 표시를 안해줘도 된다.
type Stringer interface {
	String() string
}

// Print : Stringer interface 를 인자로 받는 Print 함수.
func Print(s Stringer) {
	fmt.Print(s.String())
}
