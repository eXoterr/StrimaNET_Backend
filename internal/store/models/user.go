package models

type User struct {
	ID          int    `db:"id"`
	Login       string `db:"login"`
	Password    string `db:"password"`
	Email       string `db:"email"`
	IsBanned    bool   `db:"is_banned"`
	BannedUntil string `db:"banned_until"`
	BansCount   int    `db:"bans_count"`
}
