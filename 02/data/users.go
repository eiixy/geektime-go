package data

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type User struct {
	Id     uint
	Name   string
	Avatar string
	Sex    uint8
	Age    uint8
}

func (d *Dao) GetUsers(ids []uint) (users []User, err error) {
	var _ids = make([]string, 0)
	for _, id := range ids {
		_ids = append(_ids, strconv.Itoa(int(id)))
	}
	sql := "select id,name,age,sex from users where id in (?)"
	rows, err := d.db.Query(sql, strings.Join(_ids, ","))
	for err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("GetUsers error: \r\nsql:%s\r\nbindings:%v", sql, ids))
	}

	var user User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Sex); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("GetUsers scan error: \r\nsql:%s\r\nbindings:%v", sql, ids))
		}
		users = append(users, user)
	}
	return users, nil
}
