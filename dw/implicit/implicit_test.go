package implicit

// Phone 타입은 Stringer 인터페이스를 명시적으로 구현하지 않는다. 그러나 Print 의 인자로 사용 가능하다.
// Go 는 implicit interface (structural interface) 를 사용한다.
// Duck typing 이라고도 불린다.
// Python 과 같은 interpreter language 에서 지원하는 유연성을 compile language 에서 지원한다!
func ExamplePrint() {
	phone := &Phone{1, 2, 3, 4, 5, 6, 7}
	Print(phone)
	// Output: 1234567
}
