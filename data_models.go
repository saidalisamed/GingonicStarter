package main

import (
	"database/sql"
	"time"
)

// ExampleTable example database table model
type ExampleTable struct {
	ID          int64          `db:"id" json:"id"`
	Description string         `db:"description" json:"description"`
	Date        time.Time      `db:"date" json:"date"`
	RoleID      sql.NullInt64  `db:"roleid" json:"roleid"`
	Name        sql.NullString `db:"name" json:"name"`
}

// UserLogin sample user login database table model
type UserLogin struct {
	ID            int64          `db:"id" json:"id"`
	Email         string         `db:"email" json:"email"`
	Pass          string         `db:"pass" json:"pass"`
	FullName      string         `db:"fullname" json:"fullname"`
	RoleID        sql.NullInt64  `db:"roleid" json:"roleid"`
	RoleName      sql.NullString `db:"rolename" json:"rolename"`
	Permission    sql.NullString `db:"permission" json:"permission"`
	Staff         int64          `db:"staff" json:"staff"`
	Phone         string         `db:"phone" json:"phone"`
	Address       string         `db:"address" json:"address"`
	Notifications int64          `db:"notifications" json:"notifications"`
	Failure       int64          `db:"failure" json:"failure"`
}
