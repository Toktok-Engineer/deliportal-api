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
	divisionRepository     repository.DivisionRepository     = repository.NewDivisionRepository(db)
	departmentRepository   repository.DepartmentRepository   = repository.NewDepartmentRepository(db)
	sectionRepository      repository.SectionRepository      = repository.NewSectionRepository(db)
	positionRepository     repository.PositionRepository     = repository.NewPositionRepository(db)
	locationRepository     repository.LocationRepository     = repository.NewLocationRepository(db)
	employeeRepository     repository.EmployeeRepository     = repository.NewEmployeeRepository(db)
	userRepository         repository.UserRepository         = repository.NewUserRepository(db)
	formRepository         repository.FormRepository         = repository.NewFormRepository(db)
	formTypeRepository     repository.FormTypeRepository     = repository.NewFormTypeRepository(db)
	roleRepository         repository.RoleRepository         = repository.NewRoleRepository(db)
	userRoleRepository     repository.UserRoleRepository     = repository.NewUserRoleRepository(db)
	roleFormRepository     repository.RoleFormRepository     = repository.NewRoleFormRepository(db)
	businessUnitRepository repository.BusinessUnitRepository = repository.NewBusinessUnitRepository(db)
	authService            service.AuthService               = service.NewAuthService(userRepository)
	divisionService        service.DivisionService           = service.NewDivisionService(divisionRepository)
	departmentService      service.DepartmentService         = service.NewDepartmentService(departmentRepository)
	sectionService         service.SectionService            = service.NewSectionService(sectionRepository)
	positionService        service.PositionService           = service.NewPositionService(positionRepository)
	locationService        service.LocationService           = service.NewLocationService(locationRepository)
	employeeService        service.EmployeeService           = service.NewEmployeeService(employeeRepository)
	userService            service.UserService               = service.NewUserService(userRepository)
	formService            service.FormService               = service.NewFormService(formRepository)
	formTypeService        service.FormTypeService           = service.NewFormTypeService(formTypeRepository)
	roleService            service.RoleService               = service.NewRoleService(roleRepository)
	userRoleService        service.UserRoleService           = service.NewUserRoleService(userRoleRepository)
	roleFormService        service.RoleFormService           = service.NewRoleFormService(roleFormRepository)
	businessUnitService    service.BusinessUnitService       = service.NewBusinessUnitService(businessUnitRepository)
	jwtService             service.JWTService                = service.NewJWTService()
	authController         controller.AuthController         = controller.NewAuthController(authService, jwtService)
	divisionController     controller.DivisionController     = controller.NewDivisionController(divisionService, jwtService)
	departmentController   controller.DepartmentController   = controller.NewDepartmentController(departmentService, jwtService)
	sectionController      controller.SectionController      = controller.NewSectionController(sectionService, jwtService)
	positionController     controller.PositionController     = controller.NewPositionController(positionService, jwtService)
	locationController     controller.LocationController     = controller.NewLocationController(locationService, jwtService)
	employeeController     controller.EmployeeController     = controller.NewEmployeeController(employeeService, jwtService)
	userController         controller.UserController         = controller.NewUserController(userService, jwtService)
	formController         controller.FormController         = controller.NewFormController(formService, jwtService)
	formTypeController     controller.FormTypeController     = controller.NewFormTypeController(formTypeService, jwtService)
	roleController         controller.RoleController         = controller.NewRoleController(roleService, jwtService)
	userRoleController     controller.UserRoleController     = controller.NewUserRoleController(userRoleService, jwtService)
	roleFormController     controller.RoleFormController     = controller.NewRoleFormController(roleFormService, jwtService)
	businessUnitController controller.BusinessUnitController = controller.NewBusinessUnitController(businessUnitService, jwtService)
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

	departmentGroup := r.Group("/api/department", middleware.AuthorizeJWT(jwtService))
	{
		departmentGroup.GET("/", departmentController.FindDepartments)
		departmentGroup.GET("/:id", departmentController.FindDepartmentById)
		departmentGroup.GET("/exc/:divId/:id", departmentController.FindExcDepartment)
		departmentGroup.GET("/byDivision/:divId", departmentController.FindDepartmentByDivId)
		departmentGroup.POST("/", departmentController.InsertDepartment)
		departmentGroup.PUT("/:id", departmentController.UpdateDepartment)
		departmentGroup.DELETE("/:id", departmentController.DeleteDepartment)
	}

	sectionGroup := r.Group("/api/section", middleware.AuthorizeJWT(jwtService))
	{
		sectionGroup.GET("/", sectionController.FindSections)
		sectionGroup.GET("/:id", sectionController.FindSectionById)
		sectionGroup.GET("/exc/:depId/:id", sectionController.FindExcSection)
		sectionGroup.GET("/byDepartment/:depId", sectionController.FindSectionByDepId)
		sectionGroup.POST("/", sectionController.InsertSection)
		sectionGroup.PUT("/:id", sectionController.UpdateSection)
		sectionGroup.DELETE("/:id", sectionController.DeleteSection)
	}

	positionGroup := r.Group("/api/position", middleware.AuthorizeJWT(jwtService))
	{
		positionGroup.GET("/", positionController.FindPositions)
		positionGroup.GET("/:id", positionController.FindPositionById)
		positionGroup.GET("/exc/:id", positionController.FindExcPosition)
		positionGroup.POST("/", positionController.InsertPosition)
		positionGroup.PUT("/:id", positionController.UpdatePosition)
		positionGroup.DELETE("/:id", positionController.DeletePosition)
	}

	locationGroup := r.Group("/api/location", middleware.AuthorizeJWT(jwtService))
	{
		locationGroup.GET("/", locationController.FindLocations)
		locationGroup.GET("/:id", locationController.FindLocationById)
		locationGroup.GET("/exc/:id", locationController.FindExcLocation)
		locationGroup.POST("/", locationController.InsertLocation)
		locationGroup.PUT("/:id", locationController.UpdateLocation)
		locationGroup.DELETE("/:id", locationController.DeleteLocation)
	}

	employeeGroup := r.Group("/api/employee", middleware.AuthorizeJWT(jwtService))
	{
		employeeGroup.GET("/", employeeController.FindEmployees)
		employeeGroup.GET("/:id", employeeController.FindEmployeeById)
		employeeGroup.GET("/byNik/:nik", employeeController.FindEmployeeByNik)
		employeeGroup.GET("/exc/:id", employeeController.FindExcEmployee)
		employeeGroup.POST("/", employeeController.InsertEmployee)
		employeeGroup.PUT("/:id", employeeController.UpdateEmployee)
		employeeGroup.DELETE("/:id", employeeController.DeleteEmployee)
	}

	userGroup := r.Group("/api/user", middleware.AuthorizeJWT(jwtService))
	{
		userGroup.GET("/", userController.FindUsers)
		userGroup.GET("/:id", userController.FindUserById)
		userGroup.GET("/byNUName/:uName", userController.FindUserByUName)
		userGroup.GET("/exc/:id", userController.FindExcUser)
		userGroup.POST("/", userController.InsertUser)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.DELETE("/:id", userController.DeleteUser)
	}

	formGroup := r.Group("/api/form", middleware.AuthorizeJWT(jwtService))
	{
		formGroup.GET("/", formController.FindForms)
		formGroup.GET("/formJoinRole/:uId/:fpId", formController.FindFormJoinRole)
		formGroup.GET("/formByRole/:uId", formController.FindFormByRole)
		formGroup.GET("/formByType/:tyId", formController.FindFormByType)
		formGroup.GET("/excFormByType/:tyId/:id", formController.FindExcFormByType)
		formGroup.GET("/:id", formController.FindFormById)
		formGroup.GET("/formByFormTypeId/:ftId", formController.FindFormByFormTypeId)
		formGroup.GET("/excForm/:ftId/:id", formController.FindExcForm)
		formGroup.GET("/formHead/:ftId", formController.FindFormHead)
		formGroup.GET("/formHeadDetail/:id", formController.FindFormHeadDetail)
		formGroup.GET("/excFormHead/:id", formController.FindExcFormHead)
		formGroup.GET("/excFormOnly/:id", formController.FindExcFormOnly)
		formGroup.POST("/", formController.InsertForm)
		formGroup.PUT("/:id", formController.UpdateForm)
		formGroup.DELETE("/:id", formController.DeleteForm)
	}

	formTypeGroup := r.Group("/api/FormType", middleware.AuthorizeJWT(jwtService))
	{
		formTypeGroup.GET("/", formTypeController.FindFormTypes)
		formTypeGroup.GET("/:id", formTypeController.FindFormTypeById)
		formTypeGroup.GET("/exc/:id", formTypeController.FindExcFormType)
		formTypeGroup.POST("/", formTypeController.InsertFormType)
		formTypeGroup.PUT("/:id", formTypeController.UpdateFormType)
		formTypeGroup.DELETE("/:id", formTypeController.DeleteFormType)
	}

	roleGroup := r.Group("/api/role", middleware.AuthorizeJWT(jwtService))
	{
		roleGroup.GET("/", roleController.FindRoles)
		roleGroup.GET("/:id", roleController.FindRoleById)
		roleGroup.GET("/exc/:id", roleController.FindExcRole)
		roleGroup.POST("/", roleController.InsertRole)
		roleGroup.PUT("/:id", roleController.UpdateRole)
		roleGroup.DELETE("/:id", roleController.DeleteRole)
	}

	userRoleGroup := r.Group("/api/UserRole", middleware.AuthorizeJWT(jwtService))
	{
		userRoleGroup.GET("/", userRoleController.FindUserRoles)
		userRoleGroup.GET("/:id", userRoleController.FindUserRoleById)
		userRoleGroup.GET("/byUserId/:uid", userRoleController.FindUserRoleByUserId)
		userRoleGroup.GET("/exc/:id/:uid", userRoleController.FindExcUserRole)
		userRoleGroup.GET("/excOnly/:id", userRoleController.FindExcUserRoleOnly)
		userRoleGroup.POST("/", userRoleController.InsertUserRole)
		userRoleGroup.PUT("/:id", userRoleController.UpdateUserRole)
		userRoleGroup.DELETE("/:id", userRoleController.DeleteUserRole)
	}

	roleFormGroup := r.Group("/api/RoleForm", middleware.AuthorizeJWT(jwtService))
	{
		roleFormGroup.GET("/", roleFormController.FindRoleForms)
		roleFormGroup.GET("/:id", roleFormController.FindRoleFormById)
		roleFormGroup.GET("/byFormId/:fid/:rid", roleFormController.FindRoleFormByFormId)
		roleFormGroup.GET("/exc/:id/:rid", roleFormController.FindExcRoleForm)
		roleFormGroup.GET("/excOnly/:id", roleFormController.FindExcRoleFormOnly)
		roleFormGroup.POST("/", roleFormController.InsertRoleForm)
		roleFormGroup.PUT("/:id", roleFormController.UpdateRoleForm)
		roleFormGroup.DELETE("/:id", roleFormController.DeleteRoleForm)
	}

	businessUnitGroup := r.Group("/api/BusinessUnit", middleware.AuthorizeJWT(jwtService))
	{
		businessUnitGroup.GET("/", businessUnitController.FindBusinessUnits)
		businessUnitGroup.GET("/:id", businessUnitController.FindBusinessUnitById)
		businessUnitGroup.GET("/exc/:id", businessUnitController.FindExcBusinessUnit)
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
