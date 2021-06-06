package data

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string
	Avatar string
	Sex    uint8
	Age    uint8
}

func (d *Dao) GetUsers(ids []uint) (*[]User, error) {
	users := &[]User{}
	d.db.Where(ids).Find(&users)
	return users, errors.Wrap(d.db.Error, fmt.Sprintf("GetUsers error: %v", ids))
}
