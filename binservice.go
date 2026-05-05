package pastebin

import (
	"math/rand"
	"pastebin/models"
	"time"
)

type BinService interface {
	GenerateId() string
	CreateBin(contents string) bool
	GetBin(id string) models.Bin
}

func GenerateId() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	id := make([]rune, 10)

	for i := range 10 {
		id[i] = letters[rand.Intn(len(letters))]
	}

	return string(id)
}

func CreateBin(contents string) models.Bin {
	return models.Bin{Id: GenerateId(), Contents: contents, CreatedAt: time.Now()}
}

func GetBin() models.Bin {
	//placeholder
	return models.Bin{Id: "", Contents: "", CreatedAt: time.Now()}
}
