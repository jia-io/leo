package script

import (
	"leo/internal/dal"
	"leo/internal/models/user"
)

// InitDatabase
func SyncDatabase() {
	err := dal.DB.AutoMigrate(&user.User{})
	if err != nil {
		return
	}
}
