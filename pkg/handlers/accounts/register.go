package accounts

import (
	"strconv"

	"github.com/eXoterr/StrimaNET_Backend/internal/store"
	"github.com/eXoterr/StrimaNET_Backend/internal/store/models"
	"github.com/eXoterr/StrimaNET_Backend/pkg/handlers/adapters"
	"github.com/eXoterr/StrimaNET_Backend/pkg/logger"
)

func Register(ctx adapters.HTTPContext) {
	store := store.GetStore()
	log := logger.GetLogger()
	user := models.User{
		Login:    "user",
		Password: "test",
		Email:    "email@mail.ru",
	}
	usr, err := store.User().Create(user)
	if err != nil {
		log.Error(logger.LoggingData{
			Data: err.Error(),
		})
		return
	}

	ctx.ResponseWriter.Write(
		[]byte(strconv.Itoa(usr.ID)),
	)
}
