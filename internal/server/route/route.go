package route

import (
	"api-buddy/infrastructure/mysql/repository"
	"api-buddy/presentation/auth"
	departmentPre "api-buddy/presentation/facility/department"
	positionPre "api-buddy/presentation/facility/position"
	teamPre "api-buddy/presentation/facility/team"
	"api-buddy/presentation/health_handler"
	policyPre "api-buddy/presentation/policy"
	"api-buddy/presentation/settings"
	userPre "api-buddy/presentation/user"
	departmentUse "api-buddy/usecase/facility/department"
	positionUse "api-buddy/usecase/facility/position"
	teamUse "api-buddy/usecase/facility/team"
	policyUse "api-buddy/usecase/policy"
	userUse "api-buddy/usecase/user"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	h := auth.NewHandler(
		userUse.NewCreateUserUseCase(userRepository, facilityRepository, departmentRepository, positionRepository, teamRepository),
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
