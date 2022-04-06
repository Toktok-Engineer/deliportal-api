package main

import (
	"deliportal-api/config"
	"deliportal-api/controller"
	"deliportal-api/middleware"
	"deliportal-api/repository"
	"deliportal-api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                     *gorm.DB                          = config.ConnectDataBase()
	userRepository         repository.UserRepository         = repository.NewUserRepository(db)
	businessUnitRepository repository.BusinessUnitRepository = repository.NewBusinessUnitRepository(db)
	divisionRepository     repository.DivisionRepository     = repository.NewDivisionRepository(db)
	authService            service.AuthService               = service.NewAuthService(userRepository)
	businessUnitService    service.BusinessUnitService       = service.NewBusinessUnitService(businessUnitRepository)
	divisionService        service.DivisionService           = service.NewDivisionService(divisionRepository)
	jwtService             service.JWTService                = service.NewJWTService()
	authController         controller.AuthController         = controller.NewAuthController(authService, jwtService)
	businessUnitController controller.BusinessUnitController = controller.NewBusinessUnitController(businessUnitService, jwtService)
	divisionController     controller.DivisionController     = controller.NewDivisionController(divisionService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()

	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/register", authController.Register)
	}

	divisionGroup := r.Group("/api/division", middleware.AuthorizeJWT(jwtService))
	{
		divisionGroup.GET("/", divisionController.FindDivisions)
		divisionGroup.GET("/:id", divisionController.FindDivisionById)
		divisionGroup.GET("/exc/:id", divisionController.FindExcDivision)
		divisionGroup.POST("/", divisionController.InsertDivision)
		divisionGroup.PUT("/:id", divisionController.UpdateDivision)
		divisionGroup.DELETE("/:id", divisionController.DeleteDivision)
	}

	businessUnitGroup := r.Group("/api/businessunit", middleware.AuthorizeJWT(jwtService))
	{
		businessUnitGroup.GET("/", businessUnitController.FindBusinessUnits)
		businessUnitGroup.GET("/:id", businessUnitController.FindBusinessUnitById)
		businessUnitGroup.POST("/", businessUnitController.InsertBusinessUnit)
		businessUnitGroup.PUT("/:id", businessUnitController.UpdateBusinessUnit)
		businessUnitGroup.DELETE("/:id", businessUnitController.DeleteBusinessUnit)
	}

	tokenGroup := r.Group("/api/token", middleware.AuthorizeJWTRefreshToken(jwtService))
	{
		tokenGroup.POST("/renew", authController.RenewToken)
	}

	// Handle error response when a route is not defined
	r.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})

	r.Run()
}
