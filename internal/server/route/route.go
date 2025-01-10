package route

import (
	"api-buddy/infrastructure/mysql/repository"
	addressPre "api-buddy/presentation/address"
	"api-buddy/presentation/auth"
	areaPre "api-buddy/presentation/facility/area"
	departmentPre "api-buddy/presentation/facility/department"
	positionPre "api-buddy/presentation/facility/position"
	teamPre "api-buddy/presentation/facility/team"
	"api-buddy/presentation/health_handler"
	policyPre "api-buddy/presentation/policy"
	schedulePre "api-buddy/presentation/schedule"
	scheduleTypePre "api-buddy/presentation/schedule/schedule_type"
	"api-buddy/presentation/settings"
	userPre "api-buddy/presentation/user"
	serviceCodePre "api-buddy/presentation/visit_info/service_code"
	visitCategoryPre "api-buddy/presentation/visit_info/visit_category"
	addressUse "api-buddy/usecase/address"
	areaUse "api-buddy/usecase/facility/area"
	departmentUse "api-buddy/usecase/facility/department"
	positionUse "api-buddy/usecase/facility/position"
	teamUse "api-buddy/usecase/facility/team"
	policyUse "api-buddy/usecase/policy"
	scheduleUse "api-buddy/usecase/schedule"
	recurringScheduleUse "api-buddy/usecase/schedule/recurring_schedule"
	scheduleTypeUse "api-buddy/usecase/schedule/schedule_type"
	userUse "api-buddy/usecase/user"
	serviceCodeUse "api-buddy/usecase/visit_info/service_code"
	visitCategoryUse "api-buddy/usecase/visit_info/visit_category"

	visitInfoDomain "api-buddy/domain/visit_info"
	routeDomain "api-buddy/domain/visit_info/route"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	patientPre "api-buddy/presentation/patient"
	patientUse "api-buddy/usecase/patient"
)

func InitRoute(api *gin.Engine) {
	api.Use(settings.ErrorHandler())

	v1 := api.Group("/v1")
	// ヘルスチェック
	v1.GET("/health", health_handler.HealthCheck)

	{
		authRoute(v1)
		policyRoute(v1)
		positionRoute(v1)
		teamRoute(v1)
		departmentRoute(v1)
		userRoute(v1)
		addressRoute(v1)
		areaRoute(v1)
		scheduleRoute(v1)
		scheduleTypeRoute(v1)
		serviceCodeRoute(v1)
		visitCategoryRoute(v1)
		patientRoute(v1)
	}

	// Swagger
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func authRoute(r *gin.RouterGroup) {
	userRepository := repository.NewUserRepository()
	facilityRepository := repository.NewFacilityRepository()
	departmentRepository := repository.NewDepartmentRepository()
	positionRepository := repository.NewPositionRepository()
	teamRepository := repository.NewTeamRepository()
	areaRepository := repository.NewAreaRepository()
	h := auth.NewHandler(
		userUse.NewCreateUserUseCase(userRepository, facilityRepository, departmentRepository, positionRepository, teamRepository, areaRepository),
	)
	group := r.Group("/auth")
	group.POST("/signin", h.SignIn)
	group.POST("/signup", h.SignUp)
}

func departmentRoute(r *gin.RouterGroup) {
	departmentRepository := repository.NewDepartmentRepository()
	h := departmentPre.NewHandler(
		departmentUse.NewFindDepartmentUseCase(departmentRepository),
		departmentUse.NewFetchDepartmentsUseCase(departmentRepository),
	)
	group := r.Group("/departments")
	group.GET("/:department_id", h.FindById)

	group = r.Group("/facilities/:facility_id/departments")
	group.GET("", h.FetchByFacilityId)
}

func positionRoute(r *gin.RouterGroup) {
	positionRepository := repository.NewPositionRepository()
	policyRepository := repository.NewPolicyRepository()
	h := positionPre.NewHandler(
		positionUse.NewCreatePositionUseCase(positionRepository, policyRepository),
		positionUse.NewFindPositionUseCase(positionRepository),
		positionUse.NewFetchPositionsUseCase(positionRepository),
	)
	group := r.Group("/positions")
	group.GET("/:position_id", h.FindById)

	group = r.Group("/facilities/:facility_id/positions")
	group.POST("", h.CreateByFacilityId)
	group.GET("", h.FetchByFacilityId)
}

func teamRoute(r *gin.RouterGroup) {
	teamRepository := repository.NewTeamRepository()
	h := teamPre.NewHandler(
		teamUse.NewCreateTeamUseCase(teamRepository),
		teamUse.NewFindTeamUseCase(teamRepository),
		teamUse.NewFetchTeamsUseCase(teamRepository),
	)
	group := r.Group("/teams")
	group.GET("/:team_id", h.FindByID)

	group = r.Group("/facilities/:facility_id/teams")
	group.POST("", h.CreateByFacilityId)
	group.GET("", h.FetchByFacilityId)
}

func policyRoute(r *gin.RouterGroup) {
	policyRepository := repository.NewPolicyRepository()
	h := policyPre.NewHandler(
		policyUse.NewCreatePolicyUseCase(policyRepository),
		policyUse.NewFindPolicyUseCase(policyRepository),
		policyUse.NewFetchPoliciesUseCase(policyRepository),
	)
	group := r.Group("/policies")
	group.POST("", h.Create)
	group.GET("/:policy_id", h.FindById)
	group.GET("", h.Fetch)
}

func userRoute(r *gin.RouterGroup) {
	userRepository := repository.NewUserRepository()
	h := userPre.NewHandler(
		userUse.NewFindUserUseCase(userRepository),
		userUse.NewFetchUsersUseCase(userRepository),
	)
	group := r.Group("/users")
	group.GET("/:user_id", h.FindByUserId)

	group = r.Group("/facilities/:facility_id/users")
	group.GET("", h.FetchByFacilityId)
}

func addressRoute(r *gin.RouterGroup) {
	addressRepository := repository.NewAddressRepository()
	h := addressPre.NewHandler(
		addressUse.NewCreateAddressUseCase(addressRepository),
		addressUse.NewFindAddressUseCase(addressRepository),
		addressUse.NewFetchAddressUseCase(addressRepository),
	)
	group := r.Group("/addresses")
	group.POST("", h.Create)
	group.GET("", h.Fetch)
	group.GET("/:address_id", h.FindById)

}

func areaRoute(r *gin.RouterGroup) {
	addressRepository := repository.NewAddressRepository()
	areaRepository := repository.NewAreaRepository()
	h := areaPre.NewHandler(
		areaUse.NewCreateAreaUseCase(areaRepository, addressRepository),
		areaUse.NewFindAreaUseCase(areaRepository),
		areaUse.NewFetchAreaUseCase(areaRepository),
	)
	group := r.Group("/areas")
	group.POST("", h.Create)
	group.GET("/:area_id", h.FindById)

	group = r.Group("/facilities/:facility_id/areas")
	group.GET("", h.FetchByFacilityId)
}

func scheduleRoute(r *gin.RouterGroup) {
	scheduleRepository := repository.NewScheduleRepository()
	facilityRepository := repository.NewFacilityRepository()
	scheduleTypeRepository := repository.NewScheduleTypeRepository()
	userRepository := repository.NewUserRepository()
	recurringScheduleRepository := repository.NewRecurringScheduleRepository()
	recurringRuleRepository := repository.NewRecurringRuleRepository()
	addressRepository := repository.NewAddressRepository()
	routeRepository := repository.NewRouteRepository()
	patientRepository := repository.NewPatientRepository()
	visitInfoRepository := repository.NewVisitInfoRepository()
	serviceCodeRepository := repository.NewServiceCodeRepository()
	visitCategoryRepository := repository.NewVisitCategoryRepository()
	routeService := routeDomain.NewRouteService(
		routeRepository,
		addressRepository,
	)
	visitInfoService := visitInfoDomain.NewVisitInfoService(
		visitInfoRepository,
		patientRepository,
		userRepository,
		serviceCodeRepository,
		routeService,
		visitCategoryRepository,
	)
	h := schedulePre.NewHandler(
		scheduleUse.NewCreateScheduleUseCase(scheduleRepository, facilityRepository, scheduleTypeRepository, userRepository, recurringScheduleRepository, visitInfoService),
		recurringScheduleUse.NewCreateRecurringScheduleUseCase(recurringRuleRepository, facilityRepository, scheduleTypeRepository, userRepository, recurringScheduleRepository, visitInfoService),
	)
	group := r.Group("/facilities/:facility_id/schedules")
	group.POST("", h.CreateSchedule)

	group = r.Group("/facilities/:facility_id/schedules/recurring")
	group.POST("", h.CreateRecurringSchedule)
}

func scheduleTypeRoute(r *gin.RouterGroup) {
	scheduleTypeRepository := repository.NewScheduleTypeRepository()
	h := scheduleTypePre.NewHandler(
		scheduleTypeUse.NewFetchScheduleTypesUseCase(scheduleTypeRepository),
	)
	group := r.Group("/facilities/:facility_id/schedules/schedule_types")
	group.GET("", h.FetchScheduleTypes)
}

func serviceCodeRoute(r *gin.RouterGroup) {
	serviceCodeRepository := repository.NewServiceCodeRepository()
	h := serviceCodePre.NewHandler(
		serviceCodeUse.NewFetchServiceCodesUseCase(serviceCodeRepository),
	)
	group := r.Group("/visit_infos/service_codes")
	group.GET("", h.FetchServiceCodes)
}

func visitCategoryRoute(r *gin.RouterGroup) {
	visitCategoryRepository := repository.NewVisitCategoryRepository()
	h := visitCategoryPre.NewHandler(
		visitCategoryUse.NewFetchVisitCategoriesUseCase(visitCategoryRepository),
	)
	group := r.Group("/visit_infos/visit_categories")
	group.GET("", h.FetchVisitCategories)
}

func patientRoute(r *gin.RouterGroup) {
	patientRepository := repository.NewPatientRepository()
	facilityRepository := repository.NewFacilityRepository()
	areaREpository := repository.NewAreaRepository()
	addressRepository := repository.NewAddressRepository()
	userRepository := repository.NewUserRepository()
	serviceCodeRepository := repository.NewServiceCodeRepository()

	h := patientPre.NewHandler(
		patientUse.NewCreatePatientUseCase(patientRepository, facilityRepository, areaREpository, addressRepository, userRepository, serviceCodeRepository),
		patientUse.NewFetchPatientUseCase(patientRepository),
	)
	group := r.Group("/facilities/:facility_id/patients")
	group.POST("", h.CreatePatient)
	group.GET("", h.FetchByFacilityId)
}
