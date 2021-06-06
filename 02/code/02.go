package code

import (
	"fmt"
)

// error types 附带更多的错误信息
type MyError struct {
	Msg  string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d: %s", e.File, e.Line, e.Msg)
}

func Test02() error {
	err := test()
	// 类型断言
	switch err := err.(type) {
	case nil:
	case *MyError:
		fmt.Println("error occurred on line:", err.Line)
	default:
	}

	// Opaque errors
	if IsTemporary(err) {
		fmt.Println("err is temporary")
	}
	return test()
}

func test() error {
	return &MyError{"Something happened", "server.go", 48}
}

// Opaque errors 不透明的错误
type temporary interface {
	Temporary() bool
}

func (err MyError) Temporary() bool {
	return err.Msg == "Something happened"
}

func IsTemporary(err error) bool {
	// 断言异常类型
	te, ok := err.(temporary)
	return ok && te.Temporary()
}
