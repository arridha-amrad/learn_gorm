package main

import (
	"learn_gorm/models"
	"learn_gorm/repositories"
	"learn_gorm/services"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=learn_gorm_user password=learn_gorm_pwd dbname=learn_gorm_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}

	// Auto Migrate
	err = db.AutoMigrate(&models.User{}, &models.Account{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Migration completed")

	userRepo := repositories.NewUserRepo(db)
	accountRepo := repositories.NewAccountRepo()
	baseRepo := repositories.NewBaseRepo(db)

	userAccountService := services.NewUserAccountService(baseRepo, userRepo, accountRepo)

	// CREATE USER WITH ACCOUNT
	// if err := userAccountService.CreateUserWithAccount(services.CreateUserWithAccountParams{
	// 	Name:    gofakeit.Name(),
	// 	Email:   gofakeit.Email(),
	// 	Balance: math.Floor(gofakeit.Price(1000, 100000)),
	// }); err != nil {
	// 	fmt.Println(err.Error())
	// }

	// TRANSFER
	// if err := userAccountService.Transfer(services.TransferParams{
	// 	FromId: 1,
	// 	ToId:   4,
	// 	Amount: 100,
	// }); err != nil {
	// 	log.Println(err.Error())
	// }

	// FIND USER WITH HIS ACCOUNT
	if err := userAccountService.GetUser(1); err != nil {
		log.Println(err.Error())
	}
}
