// package main

// import (
// 	"log"
// 	"mtii-backend/config"
// 	"mtii-backend/controllers"
// 	"mtii-backend/middlewares"
// 	"mtii-backend/migrations"
// 	"mtii-backend/repositories"
// 	"mtii-backend/routes"
// 	"mtii-backend/services"
// 	"os"
// 	"sync"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// func main() {
// 	var (
// 		db *gorm.DB = config.SetUpDatabaseConnection()

// 		userRepository          repositories.UserRepository          = repositories.NewUserRepository(db)
// 		platformRepository      repositories.PlatformRepository      = repositories.NewPlatformRepository(db)
// 		statusRepository        repositories.StatusRepository        = repositories.NewStatusRepository(db)
// 		paymentMethodRepository repositories.PaymentMethodRepository = repositories.NewPaymentMethodRepository(db)
// 		salePersonRepository    repositories.SalePersonRepository    = repositories.NewSalePersonRepository(db)
// 		channelRepository       repositories.ChannelRepository       = repositories.NewChannelRepository(db)
// 		bankRepository          repositories.BankRepository          = repositories.NewBankRepository(db)
// 		receiverRepository      repositories.ReceiverRepository      = repositories.NewReceiverRepository(db)
// 		incomeRepository        repositories.IncomeRepository        = repositories.NewIncomeRepository(db)
// 		detailRepository        repositories.DetailRepository        = repositories.NewDetailRepository(db)

// 		tokenService         services.TokenService         = services.NewTokenService()
// 		userService          services.UserService          = services.NewUserService(tokenService, userRepository)
// 		platformService      services.PlatformService      = services.NewPlatformService(platformRepository)
// 		statusService        services.StatusService        = services.NewStatusService(statusRepository)
// 		paymentMethodService services.PaymentMethodService = services.NewPaymentMethodService(paymentMethodRepository)
// 		salePersonService    services.SalePersonService    = services.NewSalePersonService(salePersonRepository)
// 		channelService       services.ChannelService       = services.NewChannelService(channelRepository)
// 		bankService          services.BankService          = services.NewBankService(bankRepository)
// 		receiverService      services.ReceiverService      = services.NewReceiverService(receiverRepository)
// 		incomeService        services.IncomeService        = services.NewIncomeService(incomeRepository)
// 		detailService        services.DetailService        = services.NewDetailService(detailRepository)

// 		userController          controllers.UserController          = controllers.NewUserController(tokenService, userService)
// 		platformController      controllers.PlatformController      = controllers.NewPlatformController(tokenService, platformService)
// 		statusController        controllers.StatusController        = controllers.NewStatusController(tokenService, statusService)
// 		paymentMethodController controllers.PaymentMethodController = controllers.NewPaymentMethodController(tokenService, paymentMethodService)
// 		salePersonController    controllers.SalePersonController    = controllers.NewSalePersonController(tokenService, salePersonService)
// 		channelController       controllers.ChannelController       = controllers.NewChannelController(tokenService, channelService)
// 		bankController          controllers.BankController          = controllers.NewBankController(tokenService, bankService)
// 		receiverController      controllers.ReceiverController      = controllers.NewReceiverController(tokenService, receiverService)
// 		incomeController        controllers.IncomeController        = controllers.NewIncomeController(tokenService, incomeService)
// 		detailController        controllers.DetailController        = controllers.NewDetailController(tokenService, detailService)
// 	)

// 	server := gin.Default()
// 	server.Use(middlewares.CORSMiddleware())
// 	routes.Router(
// 		server,
// 		userController,
// 		platformController,
// 		statusController,
// 		paymentMethodController,
// 		salePersonController,
// 		channelController,
// 		bankController,
// 		receiverController,
// 		incomeController,
// 		detailController,
// 		tokenService,
// 	)

// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go func() {
// 		defer wg.Done()
// 		if err := migrations.Migrate(db); err != nil {
// 			log.Fatalf("error migration: %v", err)
// 		}
// 	}()

// 	// go func() {
// 	// 	defer wg.Done()
// 	// 	if err := migrations.Seeder(db); err != nil {
// 	// 		log.Fatalf("error migration seeder: %v", err)
// 	// 	}
// 	// }()

// 	if os.Getenv("SKIP_SEEDER") != "true" {
// 		go func() {
// 			defer wg.Done()
// 			if err := migrations.Seeder(db); err != nil {
// 				log.Fatalf("error migration seeder: %v", err)
// 			}
// 		}()
// 	} else {
// 		wg.Done() // skip and mark as done
// 	}

// 	wg.Wait()

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8888"
// 	}
// 	server.Run(":" + port)
// }

package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"mtii-backend/config"
	"mtii-backend/controllers"
	"mtii-backend/middlewares"
	"mtii-backend/migrations"
	"mtii-backend/repositories"
	"mtii-backend/routes"
	"mtii-backend/services"
)

func main() {
	// ─────────────────────────────────────────────────────────────
	// 1. Database connection
	// ─────────────────────────────────────────────────────────────
	db := config.SetUpDatabaseConnection()

	// ─────────────────────────────────────────────────────────────
	// 2. Repositories, services, controllers (unchanged)
	// ─────────────────────────────────────────────────────────────
	userRepo := repositories.NewUserRepository(db)
	platformRepo := repositories.NewPlatformRepository(db)
	statusRepo := repositories.NewStatusRepository(db)
	paymentRepo := repositories.NewPaymentMethodRepository(db)
	salePersonRepo := repositories.NewSalePersonRepository(db)
	channelRepo := repositories.NewChannelRepository(db)
	bankRepo := repositories.NewBankRepository(db)
	receiverRepo := repositories.NewReceiverRepository(db)
	incomeRepo := repositories.NewIncomeRepository(db)
	detailRepo := repositories.NewDetailRepository(db)

	tokenSvc := services.NewTokenService()
	userSvc := services.NewUserService(tokenSvc, userRepo)
	platformSvc := services.NewPlatformService(platformRepo)
	statusSvc := services.NewStatusService(statusRepo)
	paymentSvc := services.NewPaymentMethodService(paymentRepo)
	salePersonSvc := services.NewSalePersonService(salePersonRepo)
	channelSvc := services.NewChannelService(channelRepo)
	bankSvc := services.NewBankService(bankRepo)
	receiverSvc := services.NewReceiverService(receiverRepo)
	incomeSvc := services.NewIncomeService(incomeRepo)
	detailSvc := services.NewDetailService(detailRepo)

	userCtrl := controllers.NewUserController(tokenSvc, userSvc)
	platformCtrl := controllers.NewPlatformController(tokenSvc, platformSvc)
	statusCtrl := controllers.NewStatusController(tokenSvc, statusSvc)
	paymentCtrl := controllers.NewPaymentMethodController(tokenSvc, paymentSvc)
	salePersonCtrl := controllers.NewSalePersonController(tokenSvc, salePersonSvc)
	channelCtrl := controllers.NewChannelController(tokenSvc, channelSvc)
	bankCtrl := controllers.NewBankController(tokenSvc, bankSvc)
	receiverCtrl := controllers.NewReceiverController(tokenSvc, receiverSvc)
	incomeCtrl := controllers.NewIncomeController(tokenSvc, incomeSvc)
	detailCtrl := controllers.NewDetailController(tokenSvc, detailSvc)

	// ─────────────────────────────────────────────────────────────
	// 3. Run migrations (always) and seeder (optionally)
	// ─────────────────────────────────────────────────────────────
	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("migration error: %v", err)
	}

	if os.Getenv("SKIP_SEEDER") != "true" {
		if err := migrations.Seeder(db); err != nil {
			log.Fatalf("seeder error: %v", err)
		}
	}

	// ─────────────────────────────────────────────────────────────
	// 4. Gin engine with CORS, then routes
	// ─────────────────────────────────────────────────────────────
	server := gin.Default()
	server.Use(middlewares.CORSMiddleware())

	routes.Router(
		server,
		userCtrl,
		platformCtrl,
		statusCtrl,
		paymentCtrl,
		salePersonCtrl,
		channelCtrl,
		bankCtrl,
		receiverCtrl,
		incomeCtrl,
		detailCtrl,
		tokenSvc,
	)

	// ─────────────────────────────────────────────────────────────
	// 5. Start HTTP server
	// ─────────────────────────────────────────────────────────────
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	log.Printf("Listening on :%s", port)
	if err := server.Run(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
