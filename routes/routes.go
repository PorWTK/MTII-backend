package routes

import (
	"mtii-backend/controllers"
	"mtii-backend/middlewares"
	"mtii-backend/services"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func Router(
	route *gin.Engine,
	UserController controllers.UserController,
	PlatformController controllers.PlatformController,
	StatusController controllers.StatusController,
	PaymentMethodController controllers.PaymentMethodController,
	SalePersonController controllers.SalePersonController,
	ChannelController controllers.ChannelController,
	BankController controllers.BankController,
	ReceiverController controllers.ReceiverController,
	IncomeController controllers.IncomeController,
	DetailController controllers.DetailController,
	tokenService services.TokenService,
) {

	// 1) Register CORS *before* your routes:
	route.Use(cors.New(cors.Config{
		// Replace with your frontend’s actual URL
		AllowOrigins:     []string{"https://mtii-production.up.railway.app/"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	userRoutes := route.Group("/api/user")
	{
		userRoutes.POST("/login", UserController.LoginUser)
		userRoutes.POST("/logout", middlewares.Authenticate(tokenService), UserController.LogoutUser)
	}

	platformRoutes := route.Group("/api/platform")
	{
		platformRoutes.GET("/", middlewares.Authenticate(tokenService), PlatformController.GetAllPlatform)
		platformRoutes.GET("/:platform_id", middlewares.Authenticate(tokenService), PlatformController.GetPlatformById)
		platformRoutes.POST("/", middlewares.Authenticate(tokenService), PlatformController.CreatePlatform)
		platformRoutes.PATCH("/:platform_id", middlewares.Authenticate(tokenService), PlatformController.UpdatePlatform)
		platformRoutes.DELETE("/:platform_id", middlewares.Authenticate(tokenService), PlatformController.DeletePlatform)
	}

	statusRoutes := route.Group("/api/status")
	{
		statusRoutes.GET("/", middlewares.Authenticate(tokenService), StatusController.GetAllStatus)
		statusRoutes.GET("/:status_id", middlewares.Authenticate(tokenService), StatusController.GetStatusById)
		statusRoutes.POST("/", middlewares.Authenticate(tokenService), StatusController.CreateStatus)
		statusRoutes.PATCH("/:status_id", middlewares.Authenticate(tokenService), StatusController.UpdateStatus)
		statusRoutes.DELETE("/:status_id", middlewares.Authenticate(tokenService), StatusController.DeleteStatus)
	}

	paymentMethodRoutes := route.Group("/api/payment_method")
	{
		paymentMethodRoutes.GET("/", middlewares.Authenticate(tokenService), PaymentMethodController.GetAllPaymentMethod)
		paymentMethodRoutes.GET("/:payment_method_id", middlewares.Authenticate(tokenService), PaymentMethodController.GetPaymentMethodById)
		paymentMethodRoutes.POST("/", middlewares.Authenticate(tokenService), PaymentMethodController.CreatePaymentMethod)
		paymentMethodRoutes.PATCH("/:payment_method_id", middlewares.Authenticate(tokenService), PaymentMethodController.UpdatePaymentMethod)
		paymentMethodRoutes.DELETE("/:payment_method_id", middlewares.Authenticate(tokenService), PaymentMethodController.DeletePaymentMethod)
	}

	salePersonRoutes := route.Group("/api/sale_person")
	{
		salePersonRoutes.GET("/", middlewares.Authenticate(tokenService), SalePersonController.GetAllSalePerson)
		salePersonRoutes.GET("/:sale_person_id", middlewares.Authenticate(tokenService), SalePersonController.GetSalePersonById)
		salePersonRoutes.POST("/", middlewares.Authenticate(tokenService), SalePersonController.CreateSalePerson)
		salePersonRoutes.PATCH("/:sale_person_id", middlewares.Authenticate(tokenService), SalePersonController.UpdateSalePerson)
		salePersonRoutes.DELETE("/:sale_person_id", middlewares.Authenticate(tokenService), SalePersonController.DeleteSalePerson)
	}

	channelRoutes := route.Group("/api/channel")
	{
		channelRoutes.GET("/", middlewares.Authenticate(tokenService), ChannelController.GetAllChannel)
		channelRoutes.GET("/:channel_id", middlewares.Authenticate(tokenService), ChannelController.GetChannelById)
		channelRoutes.POST("/", middlewares.Authenticate(tokenService), ChannelController.CreateChannel)
		channelRoutes.PATCH("/:channel_id", middlewares.Authenticate(tokenService), ChannelController.UpdateChannel)
		channelRoutes.DELETE("/:channel_id", middlewares.Authenticate(tokenService), ChannelController.DeleteChannel)
	}

	bankRoutes := route.Group("/api/bank")
	{
		bankRoutes.GET("/", middlewares.Authenticate(tokenService), BankController.GetAllBank)
		bankRoutes.GET("/:bank_id", middlewares.Authenticate(tokenService), BankController.GetBankById)
		bankRoutes.POST("/", middlewares.Authenticate(tokenService), BankController.CreateBank)
		bankRoutes.PATCH("/:bank_id", middlewares.Authenticate(tokenService), BankController.UpdateBank)
		bankRoutes.DELETE("/:bank_id", middlewares.Authenticate(tokenService), BankController.DeleteBank)
	}

	receiverRoutes := route.Group("/api/receiver")
	{
		receiverRoutes.GET("/", middlewares.Authenticate(tokenService), ReceiverController.GetAllReceiver)
		receiverRoutes.GET("/:receiver_id", middlewares.Authenticate(tokenService), ReceiverController.GetReceiverById)
		receiverRoutes.POST("/", middlewares.Authenticate(tokenService), ReceiverController.CreateReceiver)
		receiverRoutes.PATCH("/:receiver_id", middlewares.Authenticate(tokenService), ReceiverController.UpdateReceiver)
		receiverRoutes.DELETE("/:receiver_id", middlewares.Authenticate(tokenService), ReceiverController.DeleteReceiver)
	}

	incomeRoutes := route.Group("/api/income")
	{
		incomeRoutes.GET("/", middlewares.Authenticate(tokenService), IncomeController.GetAllIncome)
		incomeRoutes.GET("/:income_invoice_id_number", middlewares.Authenticate(tokenService), IncomeController.GetIncomeByInvoiceIdNumber)
		incomeRoutes.POST("/", middlewares.Authenticate(tokenService), IncomeController.CreateIncome)
		incomeRoutes.PATCH("/:income_invoice_id_number", middlewares.Authenticate(tokenService), IncomeController.UpdateIncome)
		incomeRoutes.DELETE("/:income_invoice_id_number", middlewares.Authenticate(tokenService), IncomeController.DeleteIncome)
	}

	detailRoutes := route.Group("/api/detail")
	{
		detailRoutes.GET("/", middlewares.Authenticate(tokenService), DetailController.GetAllDetail)
		detailRoutes.GET("/:detail_id", middlewares.Authenticate(tokenService), DetailController.GetDetailById)
		detailRoutes.POST("/", middlewares.Authenticate(tokenService), DetailController.CreateDetail)
		detailRoutes.PATCH("/:detail_id", middlewares.Authenticate(tokenService), DetailController.UpdateDetail)
		detailRoutes.DELETE("/:detail_id", middlewares.Authenticate(tokenService), DetailController.DeleteDetail)
	}
}
