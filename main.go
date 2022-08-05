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
	db                                     *gorm.DB                                          = config.ConnectDataBase()
	divisionRepository                     repository.DivisionRepository                     = repository.NewDivisionRepository(db)
	departmentRepository                   repository.DepartmentRepository                   = repository.NewDepartmentRepository(db)
	sectionRepository                      repository.SectionRepository                      = repository.NewSectionRepository(db)
	positionRepository                     repository.PositionRepository                     = repository.NewPositionRepository(db)
	locationRepository                     repository.LocationRepository                     = repository.NewLocationRepository(db)
	employeeRepository                     repository.EmployeeRepository                     = repository.NewEmployeeRepository(db)
	userRepository                         repository.UserRepository                         = repository.NewUserRepository(db)
	formRepository                         repository.FormRepository                         = repository.NewFormRepository(db)
	formTypeRepository                     repository.FormTypeRepository                     = repository.NewFormTypeRepository(db)
	roleRepository                         repository.RoleRepository                         = repository.NewRoleRepository(db)
	userRoleRepository                     repository.UserRoleRepository                     = repository.NewUserRoleRepository(db)
	roleFormRepository                     repository.RoleFormRepository                     = repository.NewRoleFormRepository(db)
	businessUnitRepository                 repository.BusinessUnitRepository                 = repository.NewBusinessUnitRepository(db)
	companyManagementTypeRepository        repository.CompanyManagementTypeRepository        = repository.NewCompanyManagementTypeRepository(db)
	licenseTypeRepository                  repository.LicenseTypeRepository                  = repository.NewLicenseTypeRepository(db)
	emailQueueTypeRepository               repository.EmailQueueTypeRepository               = repository.NewEmailQueueTypeRepository(db)
	emailQueueRepository                   repository.EmailQueueRepository                   = repository.NewEmailQueueRepository(db)
	emailQueueReferenceRepository          repository.EmailQueueReferenceRepository          = repository.NewEmailQueueReferenceRepository(db)
	companyRepository                      repository.CompanyRepository                      = repository.NewCompanyRepository(db)
	companyShareholderRepository           repository.CompanyShareholderRepository           = repository.NewCompanyShareholderRepository(db)
	companyManagementRepository            repository.CompanyManagementRepository            = repository.NewCompanyManagementRepository(db)
	companyLicenseRepository               repository.CompanyLicenseRepository               = repository.NewCompanyLicenseRepository(db)
	companyLicenseRenewalTracingRepository repository.CompanyLicenseRenewalTracingRepository = repository.NewCompanyLicenseRenewalTracingRepository(db)
	userCompanyRestrictionRepository       repository.UserCompanyRestrictionRepository       = repository.NewUserCompanyRestrictionRepository(db)
	authService                            service.AuthService                               = service.NewAuthService(userRepository)
	divisionService                        service.DivisionService                           = service.NewDivisionService(divisionRepository)
	departmentService                      service.DepartmentService                         = service.NewDepartmentService(departmentRepository)
	sectionService                         service.SectionService                            = service.NewSectionService(sectionRepository)
	positionService                        service.PositionService                           = service.NewPositionService(positionRepository)
	locationService                        service.LocationService                           = service.NewLocationService(locationRepository)
	employeeService                        service.EmployeeService                           = service.NewEmployeeService(employeeRepository)
	userService                            service.UserService                               = service.NewUserService(userRepository)
	formService                            service.FormService                               = service.NewFormService(formRepository)
	formTypeService                        service.FormTypeService                           = service.NewFormTypeService(formTypeRepository)
	roleService                            service.RoleService                               = service.NewRoleService(roleRepository)
	userRoleService                        service.UserRoleService                           = service.NewUserRoleService(userRoleRepository)
	roleFormService                        service.RoleFormService                           = service.NewRoleFormService(roleFormRepository)
	businessUnitService                    service.BusinessUnitService                       = service.NewBusinessUnitService(businessUnitRepository)
	companyManagementTypeService           service.CompanyManagementTypeService              = service.NewCompanyManagementTypeService(companyManagementTypeRepository)
	licenseTypeService                     service.LicenseTypeService                        = service.NewLicenseTypeService(licenseTypeRepository)
	companyLicenseRenewalTracingService    service.CompanyLicenseRenewalTracingService       = service.NewCompanyLicenseRenewalTracingService(companyLicenseRenewalTracingRepository)
	emailQueueTypeService                  service.EmailQueueTypeService                     = service.NewEmailQueueTypeService(emailQueueTypeRepository)
	emailQueueService                      service.EmailQueueService                         = service.NewEmailQueueService(emailQueueRepository)
	emailQueueReferenceService             service.EmailQueueReferenceService                = service.NewEmailQueueReferenceService(emailQueueReferenceRepository)
	companyService                         service.CompanyService                            = service.NewCompanyService(companyRepository)
	companyShareholderService              service.CompanyShareholderService                 = service.NewCompanyShareholderService(companyShareholderRepository)
	companyManagementService               service.CompanyManagementService                  = service.NewCompanyManagementService(companyManagementRepository)
	companyLicenseService                  service.CompanyLicenseService                     = service.NewCompanyLicenseService(companyLicenseRepository)
	userCompanyRestrictionService          service.UserCompanyRestrictionService             = service.NewUserCompanyRestrictionService(userCompanyRestrictionRepository)
	jwtService                             service.JWTService                                = service.NewJWTService()
	authController                         controller.AuthController                         = controller.NewAuthController(authService, jwtService)
	divisionController                     controller.DivisionController                     = controller.NewDivisionController(divisionService, jwtService)
	departmentController                   controller.DepartmentController                   = controller.NewDepartmentController(departmentService, jwtService)
	sectionController                      controller.SectionController                      = controller.NewSectionController(sectionService, jwtService)
	positionController                     controller.PositionController                     = controller.NewPositionController(positionService, jwtService)
	locationController                     controller.LocationController                     = controller.NewLocationController(locationService, jwtService)
	employeeController                     controller.EmployeeController                     = controller.NewEmployeeController(employeeService, jwtService)
	userController                         controller.UserController                         = controller.NewUserController(userService, jwtService)
	formController                         controller.FormController                         = controller.NewFormController(formService, jwtService)
	formTypeController                     controller.FormTypeController                     = controller.NewFormTypeController(formTypeService, jwtService)
	roleController                         controller.RoleController                         = controller.NewRoleController(roleService, jwtService)
	userRoleController                     controller.UserRoleController                     = controller.NewUserRoleController(userRoleService, jwtService)
	roleFormController                     controller.RoleFormController                     = controller.NewRoleFormController(roleFormService, jwtService)
	businessUnitController                 controller.BusinessUnitController                 = controller.NewBusinessUnitController(businessUnitService, jwtService)
	companyManagementTypeController        controller.CompanyManagementTypeController        = controller.NewCompanyManagementTypeController(companyManagementTypeService, jwtService)
	licenseTypeController                  controller.LicenseTypeController                  = controller.NewLicenseTypeController(licenseTypeService, jwtService)
	emailQueueTypeController               controller.EmailQueueTypeController               = controller.NewEmailQueueTypeController(emailQueueTypeService, jwtService)
	emailQueueController                   controller.EmailQueueController                   = controller.NewEmailQueueController(emailQueueService, jwtService)
	emailQueueReferenceController          controller.EmailQueueReferenceController          = controller.NewEmailQueueReferenceController(emailQueueReferenceService, jwtService)
	companyController                      controller.CompanyController                      = controller.NewCompanyController(companyService, jwtService)
	companyShareholderController           controller.CompanyShareholderController           = controller.NewCompanyShareholderController(companyShareholderService, jwtService)
	companyManagementController            controller.CompanyManagementController            = controller.NewCompanyManagementController(companyManagementService, jwtService)
	companyLicenseController               controller.CompanyLicenseController               = controller.NewCompanyLicenseController(companyLicenseService, jwtService)
	companyLicenseRenewalTracingController controller.CompanyLicenseRenewalTracingController = controller.NewCompanyLicenseRenewalTracingController(companyLicenseRenewalTracingService, jwtService)
	userCompanyRestrictionController       controller.UserCompanyRestrictionController       = controller.NewUserCompanyRestrictionController(userCompanyRestrictionService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()

	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/checkEx", authController.CheckExisting)
		authGroup.POST("/sendMail", authController.SendMail)
		authGroup.POST("/password/:username/:email", authController.UpdateDataPassword)
		authGroup.POST("/upReq/:username/:email", authController.UpdateDataRequest)
	}

	divisionGroup := r.Group("/api/division", middleware.AuthorizeJWT(jwtService))
	{
		divisionGroup.GET("/", divisionController.CountDivisionAll)
		divisionGroup.GET("/all", divisionController.FindDivisions)
		divisionGroup.GET("/list/:limit/:offset/:order/:dir", divisionController.FindDivisionsOffset)
		divisionGroup.GET("/search/:limit/:offset/:order/:dir/:search", divisionController.SearchDivision)
		divisionGroup.GET("/countSearch/:search", divisionController.CountSearchDivision)
		divisionGroup.GET("/:id", divisionController.FindDivisionById)
		divisionGroup.GET("/exc/:id", divisionController.FindExcDivision)
		divisionGroup.POST("/", divisionController.InsertDivision)
		divisionGroup.PUT("/:id", divisionController.UpdateDivision)
		divisionGroup.DELETE("/:id", divisionController.DeleteDivision)
	}

	departmentGroup := r.Group("/api/department", middleware.AuthorizeJWT(jwtService))
	{
		departmentGroup.GET("/", departmentController.CountDepartmentAll)
		departmentGroup.GET("/all", departmentController.FindDepartments)
		departmentGroup.GET("/list/:limit/:offset/:order/:dir", departmentController.FindDepartmentsOffset)
		departmentGroup.GET("/search/:limit/:offset/:order/:dir/:search", departmentController.SearchDepartment)
		departmentGroup.GET("/countSearch/:search", departmentController.CountSearchDepartment)
		departmentGroup.GET("/:id", departmentController.FindDepartmentById)
		departmentGroup.GET("/exc/:divId/:id", departmentController.FindExcDepartment)
		departmentGroup.GET("/byDivision/:divId", departmentController.FindDepartmentByDivId)
		departmentGroup.POST("/", departmentController.InsertDepartment)
		departmentGroup.PUT("/:id", departmentController.UpdateDepartment)
		departmentGroup.DELETE("/:id", departmentController.DeleteDepartment)
	}

	sectionGroup := r.Group("/api/section", middleware.AuthorizeJWT(jwtService))
	{
		sectionGroup.GET("/", sectionController.CountSectionAll)
		sectionGroup.GET("/all", sectionController.FindSections)
		sectionGroup.GET("/list/:limit/:offset/:order/:dir", sectionController.FindSectionsOffset)
		sectionGroup.GET("/search/:limit/:offset/:order/:dir/:search", sectionController.SearchSection)
		sectionGroup.GET("/countSearch/:search", sectionController.CountSearchSection)
		sectionGroup.GET("/:id", sectionController.FindSectionById)
		sectionGroup.GET("/exc/:depId/:id", sectionController.FindExcSection)
		sectionGroup.GET("/byDepartment/:depId", sectionController.FindSectionByDepId)
		sectionGroup.POST("/", sectionController.InsertSection)
		sectionGroup.PUT("/:id", sectionController.UpdateSection)
		sectionGroup.DELETE("/:id", sectionController.DeleteSection)
	}

	positionGroup := r.Group("/api/position", middleware.AuthorizeJWT(jwtService))
	{
		positionGroup.GET("/", positionController.CountPositionAll)
		positionGroup.GET("/all", positionController.FindPositions)
		positionGroup.GET("/list/:limit/:offset/:order/:dir", positionController.FindPositionsOffset)
		positionGroup.GET("/search/:limit/:offset/:order/:dir/:search", positionController.SearchPosition)
		positionGroup.GET("/countSearch/:search", positionController.CountSearchPosition)
		positionGroup.GET("/:id", positionController.FindPositionById)
		positionGroup.GET("/exc/:id", positionController.FindExcPosition)
		positionGroup.POST("/", positionController.InsertPosition)
		positionGroup.PUT("/:id", positionController.UpdatePosition)
		positionGroup.DELETE("/:id", positionController.DeletePosition)
	}

	locationGroup := r.Group("/api/location", middleware.AuthorizeJWT(jwtService))
	{
		locationGroup.GET("/", locationController.CountLocationAll)
		locationGroup.GET("/all", locationController.FindLocations)
		locationGroup.GET("/list/:limit/:offset/:order/:dir", locationController.FindLocationsOffset)
		locationGroup.GET("/search/:limit/:offset/:order/:dir/:search", locationController.SearchLocation)
		locationGroup.GET("/countSearch/:search", locationController.CountSearchLocation)
		locationGroup.GET("/:id", locationController.FindLocationById)
		locationGroup.GET("/exc/:id", locationController.FindExcLocation)
		locationGroup.POST("/", locationController.InsertLocation)
		locationGroup.PUT("/:id", locationController.UpdateLocation)
		locationGroup.DELETE("/:id", locationController.DeleteLocation)
	}

	employeeGroup := r.Group("/api/employee", middleware.AuthorizeJWT(jwtService))
	{
		employeeGroup.GET("/", employeeController.CountEmployeeAll)
		employeeGroup.GET("/all", employeeController.FindEmployees)
		employeeGroup.GET("/list/:limit/:offset/:order/:dir", employeeController.FindEmployeesOffset)
		employeeGroup.GET("/search/:limit/:offset/:order/:dir/:search", employeeController.SearchEmployee)
		employeeGroup.GET("/countSearch/:search", employeeController.CountSearchEmployee)
		employeeGroup.GET("/:id", employeeController.FindEmployeeById)
		employeeGroup.GET("/byNik/:nik", employeeController.FindEmployeeByNik)
		employeeGroup.GET("/exc/:id", employeeController.FindExcEmployee)
		employeeGroup.POST("/", employeeController.InsertEmployee)
		employeeGroup.PUT("/:id", employeeController.UpdateEmployee)
		employeeGroup.DELETE("/:id", employeeController.DeleteEmployee)
	}

	userGroup := r.Group("/api/user", middleware.AuthorizeJWT(jwtService))
	{
		userGroup.GET("/", userController.CountUserAll)
		userGroup.GET("/full", userController.FindUsersAll)
		userGroup.GET("/all", userController.FindUsers)
		userGroup.GET("/list/:limit/:offset/:order/:dir", userController.FindUsersOffset)
		userGroup.GET("/search/:limit/:offset/:order/:dir/:search", userController.SearchUser)
		userGroup.GET("/countSearch/:search", userController.CountSearchUser)
		userGroup.GET("/:id", userController.FindUserById)
		userGroup.GET("/full/:id", userController.FindUsersAll)
		userGroup.GET("/byNUName/:uName", userController.FindUserByUName)
		userGroup.GET("/exc/:id", userController.FindExcUser)
		userGroup.POST("/", userController.InsertUser)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.DELETE("/:id", userController.DeleteUser)
	}

	formGroup := r.Group("/api/form", middleware.AuthorizeJWT(jwtService))
	{
		formGroup.GET("/", formController.CountFormAll)
		formGroup.GET("/all", formController.FindForms)
		formGroup.GET("/list/:limit/:offset/:order/:dir", formController.FindFormsOffset)
		formGroup.GET("/search/:limit/:offset/:order/:dir/:search", formController.SearchForm)
		formGroup.GET("/countSearch/:search", formController.CountSearchForm)
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
		formTypeGroup.GET("/", formTypeController.CountFormTypeAll)
		formTypeGroup.GET("/all", formTypeController.FindFormTypes)
		formTypeGroup.GET("/list/:limit/:offset/:order/:dir", formTypeController.FindFormTypesOffset)
		formTypeGroup.GET("/search/:limit/:offset/:order/:dir/:search", formTypeController.SearchFormType)
		formTypeGroup.GET("/countSearch/:search", formTypeController.CountSearchFormType)
		formTypeGroup.GET("/:id", formTypeController.FindFormTypeById)
		formTypeGroup.GET("/exc/:id", formTypeController.FindExcFormType)
		formTypeGroup.POST("/", formTypeController.InsertFormType)
		formTypeGroup.PUT("/:id", formTypeController.UpdateFormType)
		formTypeGroup.DELETE("/:id", formTypeController.DeleteFormType)
	}

	roleGroup := r.Group("/api/role", middleware.AuthorizeJWT(jwtService))
	{
		roleGroup.GET("/", roleController.CountRoleAll)
		roleGroup.GET("/all", roleController.FindRoles)
		roleGroup.GET("/list/:limit/:offset/:order/:dir", roleController.FindRolesOffset)
		roleGroup.GET("/search/:limit/:offset/:order/:dir/:search", roleController.SearchRole)
		roleGroup.GET("/countSearch/:search", roleController.CountSearchRole)
		roleGroup.GET("/:id", roleController.FindRoleById)
		roleGroup.GET("/exc/:id", roleController.FindExcRole)
		roleGroup.POST("/", roleController.InsertRole)
		roleGroup.PUT("/:id", roleController.UpdateRole)
		roleGroup.DELETE("/:id", roleController.DeleteRole)
	}

	userRoleGroup := r.Group("/api/UserRole", middleware.AuthorizeJWT(jwtService))
	{
		userRoleGroup.GET("/", userRoleController.CountUserRoleAll)
		userRoleGroup.GET("/all", userRoleController.FindUserRoles)
		userRoleGroup.GET("/list/:limit/:offset/:order/:dir", userRoleController.FindUserRolesOffset)
		userRoleGroup.GET("/search/:limit/:offset/:order/:dir/:search", userRoleController.SearchUserRole)
		userRoleGroup.GET("/countSearch/:search", userRoleController.CountSearchUserRole)
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
		roleFormGroup.GET("/", roleFormController.CountRoleFormAll)
		roleFormGroup.GET("/all", roleFormController.FindRoleForms)
		roleFormGroup.GET("/list/:limit/:offset/:order/:dir", roleFormController.FindRoleFormsOffset)
		roleFormGroup.GET("/search/:limit/:offset/:order/:dir/:search", roleFormController.SearchRoleForm)
		roleFormGroup.GET("/countSearch/:search", roleFormController.CountSearchRoleForm)
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
		businessUnitGroup.GET("/", businessUnitController.CountBusinessUnitAll)
		businessUnitGroup.GET("/all", businessUnitController.FindBusinessUnits)
		businessUnitGroup.GET("/list/:limit/:offset/:order/:dir", businessUnitController.FindBusinessUnitsOffset)
		businessUnitGroup.GET("/search/:limit/:offset/:order/:dir/:search", businessUnitController.SearchBusinessUnit)
		businessUnitGroup.GET("/countSearch/:search", businessUnitController.CountSearchBusinessUnit)
		businessUnitGroup.GET("/:id", businessUnitController.FindBusinessUnitById)
		businessUnitGroup.GET("/exc/:id", businessUnitController.FindExcBusinessUnit)
		businessUnitGroup.POST("/", businessUnitController.InsertBusinessUnit)
		businessUnitGroup.PUT("/:id", businessUnitController.UpdateBusinessUnit)
		businessUnitGroup.DELETE("/:id", businessUnitController.DeleteBusinessUnit)
	}

	companyManagementTypeGroup := r.Group("/api/CompanyManagementType", middleware.AuthorizeJWT(jwtService))
	{
		companyManagementTypeGroup.GET("/", companyManagementTypeController.CountCompanyManagementTypeAll)
		companyManagementTypeGroup.GET("/all", companyManagementTypeController.FindCompanyManagementTypes)
		companyManagementTypeGroup.GET("/list/:limit/:offset/:order/:dir", companyManagementTypeController.FindCompanyManagementTypesOffset)
		companyManagementTypeGroup.GET("/search/:limit/:offset/:order/:dir/:search", companyManagementTypeController.SearchCompanyManagementType)
		companyManagementTypeGroup.GET("/countSearch/:search", companyManagementTypeController.CountSearchCompanyManagementType)
		companyManagementTypeGroup.GET("/:id", companyManagementTypeController.FindCompanyManagementTypeById)
		companyManagementTypeGroup.GET("/exc/:id", companyManagementTypeController.FindExcCompanyManagementType)
		companyManagementTypeGroup.POST("/", companyManagementTypeController.InsertCompanyManagementType)
		companyManagementTypeGroup.PUT("/:id", companyManagementTypeController.UpdateCompanyManagementType)
		companyManagementTypeGroup.DELETE("/:id", companyManagementTypeController.DeleteCompanyManagementType)
	}

	licenseTypeGroup := r.Group("/api/LicenseType", middleware.AuthorizeJWT(jwtService))
	{
		licenseTypeGroup.GET("/", licenseTypeController.CountLicenseTypeAll)
		licenseTypeGroup.GET("/all", licenseTypeController.FindLicenseTypes)
		licenseTypeGroup.GET("/list/:limit/:offset/:order/:dir", licenseTypeController.FindLicenseTypesOffset)
		licenseTypeGroup.GET("/search/:limit/:offset/:order/:dir/:search", licenseTypeController.SearchLicenseType)
		licenseTypeGroup.GET("/countSearch/:search", licenseTypeController.CountSearchLicenseType)
		licenseTypeGroup.GET("/:id", licenseTypeController.FindLicenseTypeById)
		licenseTypeGroup.GET("/exc/:id", licenseTypeController.FindExcLicenseType)
		licenseTypeGroup.POST("/", licenseTypeController.InsertLicenseType)
		licenseTypeGroup.PUT("/:id", licenseTypeController.UpdateLicenseType)
		licenseTypeGroup.DELETE("/:id", licenseTypeController.DeleteLicenseType)
	}

	emailQueueTypeGroup := r.Group("/api/EmailQueueType", middleware.AuthorizeJWT(jwtService))
	{
		emailQueueTypeGroup.GET("/", emailQueueTypeController.CountEmailQueueTypeAll)
		emailQueueTypeGroup.GET("/all", emailQueueTypeController.FindEmailQueueTypes)
		emailQueueTypeGroup.GET("/list/:limit/:offset/:order/:dir", emailQueueTypeController.FindEmailQueueTypesOffset)
		emailQueueTypeGroup.GET("/search/:limit/:offset/:order/:dir/:search", emailQueueTypeController.SearchEmailQueueType)
		emailQueueTypeGroup.GET("/countSearch/:search", emailQueueTypeController.CountSearchEmailQueueType)
		emailQueueTypeGroup.GET("/:id", emailQueueTypeController.FindEmailQueueTypeById)
		emailQueueTypeGroup.GET("/exc/:id", emailQueueTypeController.FindExcEmailQueueType)
		emailQueueTypeGroup.POST("/", emailQueueTypeController.InsertEmailQueueType)
		emailQueueTypeGroup.PUT("/:id", emailQueueTypeController.UpdateEmailQueueType)
		emailQueueTypeGroup.DELETE("/:id", emailQueueTypeController.DeleteEmailQueueType)
	}

	emailQueueGroup := r.Group("/api/EmailQueue", middleware.AuthorizeJWT(jwtService))
	{
		emailQueueGroup.GET("/", emailQueueController.CountEmailQueueAll)
		emailQueueGroup.GET("/all", emailQueueController.FindEmailQueues)
		emailQueueGroup.GET("/list/:limit/:offset/:order/:dir", emailQueueController.FindEmailQueuesOffset)
		emailQueueGroup.GET("/search/:limit/:offset/:order/:dir/:search", emailQueueController.SearchEmailQueue)
		emailQueueGroup.GET("/countSearch/:search", emailQueueController.CountSearchEmailQueue)
		emailQueueGroup.GET("/:id", emailQueueController.FindEmailQueueById)
		emailQueueGroup.GET("/byStatus/:status", emailQueueController.FindEmailQueueByStatus)
		emailQueueGroup.GET("/exc/:id", emailQueueController.FindExcEmailQueue)
		emailQueueGroup.POST("/", emailQueueController.InsertEmailQueue)
		emailQueueGroup.PUT("/:id", emailQueueController.UpdateEmailQueue)
		emailQueueGroup.DELETE("/:id", emailQueueController.DeleteEmailQueue)
	}

	emailQueueReferenceGroup := r.Group("/api/EmailQueueReference", middleware.AuthorizeJWT(jwtService))
	{
		emailQueueReferenceGroup.GET("/", emailQueueReferenceController.CountEmailQueueReferenceAll)
		emailQueueReferenceGroup.GET("/all", emailQueueReferenceController.FindEmailQueueReferences)
		emailQueueReferenceGroup.GET("/list/:limit/:offset/:order/:dir", emailQueueReferenceController.FindEmailQueueReferencesOffset)
		emailQueueReferenceGroup.GET("/search/:limit/:offset/:order/:dir/:search", emailQueueReferenceController.SearchEmailQueueReference)
		emailQueueReferenceGroup.GET("/countSearch/:search", emailQueueReferenceController.CountSearchEmailQueueReference)
		emailQueueReferenceGroup.GET("/:id", emailQueueReferenceController.FindEmailQueueReferenceById)
		emailQueueReferenceGroup.GET("/exc/:id", emailQueueReferenceController.FindExcEmailQueueReference)
		emailQueueReferenceGroup.POST("/", emailQueueReferenceController.InsertEmailQueueReference)
		emailQueueReferenceGroup.PUT("/:id", emailQueueReferenceController.UpdateEmailQueueReference)
		emailQueueReferenceGroup.DELETE("/:id", emailQueueReferenceController.DeleteEmailQueueReference)
	}

	companyGroup := r.Group("/api/company", middleware.AuthorizeJWT(jwtService))
	{
		companyGroup.GET("/", companyController.CountCompanyAll)
		companyGroup.GET("/all", companyController.FindCompanys)
		companyGroup.GET("/list/:limit/:offset/:order/:dir/:companyID", companyController.FindCompanysOffset)
		companyGroup.GET("/search/:limit/:offset/:order/:dir/:search/:companyID", companyController.SearchCompany)
		companyGroup.GET("/countSearch/:search/:companyID", companyController.CountSearchCompany)
		companyGroup.GET("/app/count/:companyId", companyController.CountCompanyApprove)
		companyGroup.GET("/app/:limit/:offset/:order/:dir/:companyId", companyController.FindCompanyApprove)
		companyGroup.GET("/app/search/:limit/:offset/:order/:dir/:search/:companyId", companyController.SearchCompanyApprove)
		companyGroup.GET("/app/countSearch/:search/:companyId", companyController.CountSearchCompanyApprove)
		companyGroup.GET("/:id", companyController.FindCompanyById)
		companyGroup.GET("/exc/:id", companyController.FindExcCompany)
		companyGroup.POST("/", companyController.InsertCompany)
		companyGroup.PUT("/:id", companyController.UpdateCompany)
		companyGroup.PUT("/approved/:id", companyController.UpdateCompanyApprove)
		companyGroup.PUT("/deactived/:id", companyController.UpdateCompanyDeactive)
		companyGroup.DELETE("/:id", companyController.DeleteCompany)
	}

	companyShareholderGroup := r.Group("/api/companyShareholder", middleware.AuthorizeJWT(jwtService))
	{
		companyShareholderGroup.GET("/count/:companyId", companyShareholderController.CountCompanyShareholderAll)
		companyShareholderGroup.GET("/all/:companyId", companyShareholderController.FindCompanyShareholders)
		companyShareholderGroup.GET("/list/:limit/:offset/:order/:dir/:companyId", companyShareholderController.FindCompanyShareholdersOffset)
		companyShareholderGroup.GET("/search/:limit/:offset/:order/:dir/:search/:companyId", companyShareholderController.SearchCompanyShareholder)
		companyShareholderGroup.GET("/countSearch/:search/:companyId", companyShareholderController.CountSearchCompanyShareholder)
		companyShareholderGroup.GET("/byCompany/:id", companyShareholderController.FindCompanyShareholderByCompanyId)
		companyShareholderGroup.GET("/:id", companyShareholderController.FindCompanyShareholderById)
		companyShareholderGroup.GET("/exc/:id", companyShareholderController.FindExcCompanyShareholder)
		companyShareholderGroup.POST("/", companyShareholderController.InsertCompanyShareholder)
		companyShareholderGroup.PUT("/:id", companyShareholderController.UpdateCompanyShareholder)
		companyShareholderGroup.DELETE("/:id", companyShareholderController.DeleteCompanyShareholder)
	}

	companyManagementGroup := r.Group("/api/companyManagement", middleware.AuthorizeJWT(jwtService))
	{
		companyManagementGroup.GET("/count/:companyId", companyManagementController.CountCompanyManagementAll)
		companyManagementGroup.GET("/all/:companyId", companyManagementController.FindCompanyManagements)
		companyManagementGroup.GET("/list/:limit/:offset/:order/:dir/:companyId", companyManagementController.FindCompanyManagementsOffset)
		companyManagementGroup.GET("/search/:limit/:offset/:order/:dir/:search/:companyId", companyManagementController.SearchCompanyManagement)
		companyManagementGroup.GET("/countSearch/:search/:companyId", companyManagementController.CountSearchCompanyManagement)
		companyManagementGroup.GET("/:id", companyManagementController.FindCompanyManagementById)
		companyManagementGroup.GET("/exc/:id", companyManagementController.FindExcCompanyManagement)
		companyManagementGroup.GET("/byCompany/:id", companyManagementController.FindCompanyManagementByCompanyId)
		companyManagementGroup.POST("/", companyManagementController.InsertCompanyManagement)
		companyManagementGroup.PUT("/:id", companyManagementController.UpdateCompanyManagement)
		companyManagementGroup.DELETE("/:id", companyManagementController.DeleteCompanyManagement)
	}

	companyLicenseGroup := r.Group("/api/companyLicense", middleware.AuthorizeJWT(jwtService))
	{
		companyLicenseGroup.GET("/byCompany/:id", companyLicenseController.FindCompanyLicenseByCompanyId)
		companyLicenseGroup.GET("/count/:companyId", companyLicenseController.CountCompanyLicenseAll)
		companyLicenseGroup.GET("/all", companyLicenseController.FindCompanyLicenses)
		companyLicenseGroup.GET("/list/:limit/:offset/:order/:dir/:companyId", companyLicenseController.FindCompanyLicensesOffset)
		companyLicenseGroup.GET("/search/:limit/:offset/:order/:dir/:search/:companyId", companyLicenseController.SearchCompanyLicense)
		companyLicenseGroup.GET("/countSearch/:search/:companyId", companyLicenseController.CountSearchCompanyLicense)
		companyLicenseGroup.GET("/app/count/:companyId", companyLicenseController.CountCompanyLicenseApp)
		companyLicenseGroup.GET("/app/:limit/:offset/:order/:dir/:companyId", companyLicenseController.FindCompanyLicensesApp)
		companyLicenseGroup.GET("/app/search/:limit/:offset/:order/:dir/:search/:companyId", companyLicenseController.SearchCompanyLicenseApp)
		companyLicenseGroup.GET("/app/countSearch/:search/:companyId", companyLicenseController.CountSearchCompanyLicenseApp)
		companyLicenseGroup.GET("/exp/count/:companyId", companyLicenseController.CountExpCompanyLicense)
		companyLicenseGroup.GET("/exp/:limit/:offset/:order/:dir/:companyId", companyLicenseController.FindExpCompanyLicenses)
		companyLicenseGroup.GET("/exp/search/:limit/:offset/:order/:dir/:search/:companyId", companyLicenseController.SearchExpCompanyLicense)
		companyLicenseGroup.GET("/exp/countSearch/:search/:companyId", companyLicenseController.CountSearchExpCompanyLicense)
		companyLicenseGroup.GET("/:id", companyLicenseController.FindCompanyLicenseById)
		companyLicenseGroup.GET("/exc/:id", companyLicenseController.FindExcCompanyLicense)
		companyLicenseGroup.GET("/countFull/:companyId", companyLicenseController.CountCompanyLicenseFull)
		companyLicenseGroup.GET("/listFull/:limit/:offset/:order/:dir/:companyId", companyLicenseController.FindCompanyLicensesOffsetFull)
		companyLicenseGroup.GET("/searchFull/:limit/:offset/:order/:dir/:search/:companyId", companyLicenseController.SearchCompanyLicenseFull)
		companyLicenseGroup.GET("/countSearchFull/:search/:companyId", companyLicenseController.CountSearchCompanyLicenseFull)
		companyLicenseGroup.POST("/", companyLicenseController.InsertCompanyLicense)
		companyLicenseGroup.PUT("/:id", companyLicenseController.UpdateCompanyLicense)
		companyLicenseGroup.PUT("/status/:id", companyLicenseController.UpdateCompanyLicenseStatus)
		companyLicenseGroup.PUT("/statusDeactive/:id", companyLicenseController.UpdateCompanyLicenseDeactive)
		companyLicenseGroup.PUT("/renewal/:id", companyLicenseController.UpdateCompanyLicenseApprovedRenewalStatus)
		companyLicenseGroup.PUT("/remark/:id", companyLicenseController.UpdateCompanyRemark)
		companyLicenseGroup.DELETE("/:id", companyLicenseController.DeleteCompanyLicense)
	}

	companyLicenseRenewalTracingGroup := r.Group("/api/CompanyLicenseRenewalTracing", middleware.AuthorizeJWT(jwtService))
	{
		companyLicenseRenewalTracingGroup.GET("/", companyLicenseRenewalTracingController.FindCompanyLicenseRenewalTracings)
		companyLicenseRenewalTracingGroup.GET("/:id", companyLicenseRenewalTracingController.FindCompanyLicenseRenewalTracingById)
		companyLicenseRenewalTracingGroup.GET("/exc/:id", companyLicenseRenewalTracingController.FindExcCompanyLicenseRenewalTracing)
		companyLicenseRenewalTracingGroup.POST("/", companyLicenseRenewalTracingController.InsertCompanyLicenseRenewalTracing)
		companyLicenseRenewalTracingGroup.PUT("/:id", companyLicenseRenewalTracingController.UpdateCompanyLicenseRenewalTracing)
		companyLicenseRenewalTracingGroup.DELETE("/:id", companyLicenseRenewalTracingController.DeleteCompanyLicenseRenewalTracing)
	}

	usercompanyrestrictionGroup := r.Group("/api/userCompanyRestriction", middleware.AuthorizeJWT(jwtService))
	{
		usercompanyrestrictionGroup.GET("/", userCompanyRestrictionController.CountUserCompanyRestrictionAll)
		usercompanyrestrictionGroup.GET("/all", userCompanyRestrictionController.FindUserCompanyRestrictions)
		usercompanyrestrictionGroup.GET("/list/:limit/:offset/:order/:dir", userCompanyRestrictionController.FindUserCompanyRestrictionsOffset)
		usercompanyrestrictionGroup.GET("/search/:limit/:offset/:order/:dir/:search", userCompanyRestrictionController.SearchUserCompanyRestriction)
		usercompanyrestrictionGroup.GET("/countSearch/:search", userCompanyRestrictionController.CountSearchUserCompanyRestriction)
		usercompanyrestrictionGroup.GET("/:id", userCompanyRestrictionController.FindUserCompanyRestrictionById)
		usercompanyrestrictionGroup.GET("/user/:uid", userCompanyRestrictionController.FindUserCompanyRestrictionByUserId)
		usercompanyrestrictionGroup.GET("/exc/:id", userCompanyRestrictionController.FindExcUserCompanyRestriction)
		usercompanyrestrictionGroup.POST("/", userCompanyRestrictionController.InsertUserCompanyRestriction)
		usercompanyrestrictionGroup.PUT("/:id", userCompanyRestrictionController.UpdateUserCompanyRestriction)
		usercompanyrestrictionGroup.DELETE("/:id", userCompanyRestrictionController.DeleteUserCompanyRestriction)
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
