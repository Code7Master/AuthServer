package initializers

import "AuthServer/models"

func SyncDatabase() {
	DB.AutoMigrate(new(models.User))
}
