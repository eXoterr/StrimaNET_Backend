package store

import "github.com/eXoterr/StrimaNET_Backend/internal/store/postgres/user"

var instance Database

func GetStore() Database {
	return instance
}

func New() Database {
	if instance != nil {
		return instance
	}

	instance = &PgDatabase{
		DB: nil,
	}

	return instance
}

func (db *PgDatabase) User() user.UserStore {
	return user.New(db.DB)
}
