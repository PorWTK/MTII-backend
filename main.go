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
	"mtii-backend/config"
	"mtii-backend/controllers"
	"mtii-backend/middlewares"
	"mtii-backend/migrations"
	"mtii-backend/repositories"
	"mtii-backend/routes"
	"mtii-backend/services"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Set up the database connection
	db := config.SetUpDatabaseConnection()

	// 2. Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	platRepo := repositories.NewPlatformRepository(db)
	statRepo := repositories.NewStatusRepository(db)
	payRepo := repositories.NewPaymentMethodRepository(db)
	saleRepo := repositories.NewSalePersonRepository(db)
	chanRepo := repositories.NewChannelRepository(db)
	bankRepo := repositories.NewBankRepository(db)
	recvRepo := repositories.NewReceiverRepository(db)
	incRepo := repositories.NewIncomeRepository(db)
	detRepo := repositories.NewDetailRepository(db)

	// 3. Initialize services
	tokenSvc := services.NewTokenService()
	userSvc := services.NewUserService(tokenSvc, userRepo)
	platSvc := services.NewPlatformService(platRepo)
	statSvc := services.NewStatusService(statRepo)
	paySvc := services.NewPaymentMethodService(payRepo)
	saleSvc := services.NewSalePersonService(saleRepo)
	chanSvc := services.NewChannelService(chanRepo)
	bankSvc := services.NewBankService(bankRepo)
	recvSvc := services.NewReceiverService(recvRepo)
	incSvc := services.NewIncomeService(incRepo)
	detSvc := services.NewDetailService(detRepo)

	// 4. Initialize controllers
	userCtrl := controllers.NewUserController(tokenSvc, userSvc)
	platCtrl := controllers.NewPlatformController(tokenSvc, platSvc)
	statCtrl := controllers.NewStatusController(tokenSvc, statSvc)
	payCtrl := controllers.NewPaymentMethodController(tokenSvc, paySvc)
	saleCtrl := controllers.NewSalePersonController(tokenSvc, saleSvc)
	chanCtrl := controllers.NewChannelController(tokenSvc, chanSvc)
	bankCtrl := controllers.NewBankController(tokenSvc, bankSvc)
	recvCtrl := controllers.NewReceiverController(tokenSvc, recvSvc)
	incCtrl := controllers.NewIncomeController(tokenSvc, incSvc)
	detCtrl := controllers.NewDetailController(tokenSvc, detSvc)

	// 5. Set up Gin server with CORS
	server := gin.Default()
	server.Use(middlewares.CORSMiddleware())

	// 6. Register routes
	routes.Router(
		server,
		userCtrl,
		platCtrl,
		statCtrl,
		payCtrl,
		saleCtrl,
		chanCtrl,
		bankCtrl,
		recvCtrl,
		incCtrl,
		detCtrl,
		tokenSvc,
	)

	// 7. Run migrations (always)
	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("Migration error: %v", err)
	}

	// 8. Optionally run seeder
	if os.Getenv("SKIP_SEEDER") != "true" {
		if err := migrations.Seeder(db); err != nil {
			log.Fatalf("Seeder error: %v", err)
		}
	}

	// 9. Start the server on the given PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	server.Run(
		":" + port,
	)
}
