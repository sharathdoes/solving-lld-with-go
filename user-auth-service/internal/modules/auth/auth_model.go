package auth

import "time"

type User struct {
  ID        string
  username 	string
  Email     string
  Password  string
  CreatedAt time.Time
}
