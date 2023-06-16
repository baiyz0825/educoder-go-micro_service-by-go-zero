package model

import (
	"gorm.io/gen"
)

// only used to User
type UserMethod interface {
	// where(id=@id)
	FindByID(id int64) (gen.T, error)

	// select * from @@table where name=@userName
	FindByUserName(userName string) ([]gen.T, error)

	// update users
	//	{{set}}
	//		{{if passwd != ""}}
	//			password=@passwd
	//		{{end}}
	//	{{end}}
	// where id=@id
	UpdateUserName(passwd string, id int) error
}
