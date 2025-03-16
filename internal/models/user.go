package models

import (
	"database/sql"
	"time"
)

type User struct {
	AddedId            int32          `db:"added_id"`
	Id                 int64          `db:"id,key"`
	Name               string         `db:"name"`
	Username           sql.NullString `db:"username"`
	PhoneNumber        string         `db:"phone_number"`
	Avatar             sql.NullInt64  `db:"avatar"`
	Bio                sql.NullString `db:"bio"`
	Gender             sql.NullString `db:"gender"`
	Status             string         `db:"status"`
	CreatedAt          *time.Time     `db:"created_at"`
	ModifiedAt         *time.Time     `db:"modified_at"`
	EnterprisesId      int64          `db:"enterprises_id"`
	IsDeleted          int            `json:"is_deleted" db:"is_deleted"`
	CheckEnterprisesId int64          `db:"check_enterprises_id"`
}

func (u *User) TableName() string {
	return "user"
}
