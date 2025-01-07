package visit_info

import (
	patientDomain "api-buddy/domain/patient"
	userDomain "api-buddy/domain/user"
	routeDomain "api-buddy/domain/visit_info/route"
	serviceCodeDomain "api-buddy/domain/visit_info/service_code"
	visitCategoryDomain "api-buddy/domain/visit_info/visit_category"
	"context"

	"gorm.io/gorm"
)

type VisitInfoService struct {
	visitInfoRepository     VisitInfoRepository
	patientRepository       patientDomain.PatientRepository
	userRepository          userDomain.UserRepository
	serviceCodeRepository   serviceCodeDomain.ServiceCodeRepository
	routeService            *routeDomain.RouteService
	visitCategoryRepository visitCategoryDomain.VisitCategoryRepository
}

func NewVisitInfoService(
	visitInfoRepository VisitInfoRepository,
	patientRepository patientDomain.PatientRepository,
	userRepository userDomain.UserRepository,
	serviceCodeRepository serviceCodeDomain.ServiceCodeRepository,
	routeService *routeDomain.RouteService,
	visitCategoryRepository visitCategoryDomain.VisitCategoryRepository,
) *VisitInfoService {
	return &VisitInfoService{
		visitInfoRepository:     visitInfoRepository,
		patientRepository:       patientRepository,
		userRepository:          userRepository,
		serviceCodeRepository:   serviceCodeRepository,
		routeService:            routeService,
		visitCategoryRepository: visitCategoryRepository,
	}
}

type VisitInfoModel struct {
	PatientID        string
	AssignStaffID    string
	CompanionID      *string
	Route            *routeDomain.RouteModel
	ServiceCodeID    string
	VisitCategoryIDs []*string
}

func (s *VisitInfoService) CreateVisitInfo(ctx context.Context, tx *gorm.DB, visitInfoModel VisitInfoModel) (*VisitInfo, error) {
	patient, err := s.patientRepository.FindByID(ctx, visitInfoModel.PatientID)
	if err != nil {
		return nil, err
	}

	// 担当スタッフの取得とバリデーション
	assignStaff, err := s.userRepository.FindByID(ctx, visitInfoModel.AssignStaffID)
	if err != nil {
		return nil, err
	}

	// サービスコードの取得とバリデーション
	serviceCode, err := s.serviceCodeRepository.FindByID(ctx, visitInfoModel.ServiceCodeID)
	if err != nil {
		return nil, err
	}

	// option処理
	options := []VisitInfoOption{}
	if visitInfoModel.Route != nil {
		// ルートの作成
		route, err := s.routeService.CreateRoute(ctx, tx, visitInfoModel.Route)
		if err != nil {
			return nil, err
		}

		options = append(options, WithRoute(route))
	}

	// 訪問情報のオプションを作成
	if visitInfoModel.CompanionID != nil {
		companion, err := s.userRepository.FindByID(ctx, *visitInfoModel.CompanionID)
		if err != nil {
			return nil, err
		}
		options = append(options, WithCompanion(companion))
	}

	// 訪問カテゴリの取得
	if visitInfoModel.VisitCategoryIDs != nil {
		visitCategories := []*visitCategoryDomain.VisitCategory{}
		for _, visitCategoryID := range visitInfoModel.VisitCategoryIDs {

			visitCategory, err := s.visitCategoryRepository.FindByID(ctx, *visitCategoryID)
			if err != nil {
				return nil, err
			}
			visitCategories = append(visitCategories, visitCategory)
		}
		options = append(options, WithVisitCategories(visitCategories))
	}

	// 訪問情報の作成
	visitInfo, err := NewVisitInfo(patient, assignStaff, serviceCode, options...)
	if err != nil {
		return nil, err
	}

	// 訪問情報の保存
	if err := s.visitInfoRepository.Create(ctx, tx, visitInfo); err != nil {
		return nil, err
	}

	return visitInfo, nil
}
