package main

import (
	"database/sql"
	"fmt"
	"geektime-go/02/code"
	"geektime-go/02/data"

	"github.com/pkg/errors"
)

//

func main() {
	// 01 error
	fmt.Println("===== 01 =====")
	code.Check(-5)
	code.Check(0)
	code.Check(5)

	// 02 error type & opaque errors
	fmt.Println("===== 02 =====")
	code.Test02()

	// 03 warp error
	fmt.Println("===== 03 =====")
	code.Test03()

	// 作业
	fmt.Println("===== 04 =====")
	dao := data.New()
	users, err := dao.GetUsers([]uint{1, 2, 3})

	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
		return
	}
	fmt.Printf("users: %v", users)
}
