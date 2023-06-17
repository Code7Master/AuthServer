package initializers

import "DevelopHub/AuthServer/models"

func SyncDatabase() {
	DB.AutoMigrate(new(models.User))
}
