package patient

import (
	addressDomain "api-buddy/domain/address"
	facilityDomain "api-buddy/domain/facility"
	areaDomain "api-buddy/domain/facility/area"
	patientDomain "api-buddy/domain/patient"
	userDomain "api-buddy/domain/user"
	serviceCodeDomain "api-buddy/domain/visit_info/service_code"
	"context"
)

type CreatePatientUseCase struct {
	patientRepository     patientDomain.PatientRepository
	facilityRepository    facilityDomain.FacilityRepository
	areaRepository        areaDomain.AreaRepository
	addressRepository     addressDomain.AddressRepository
	userRepository        userDomain.UserRepository
	serviceCodeRepository serviceCodeDomain.ServiceCodeRepository
}

func NewCreatePatientUseCase(
	patientRepository patientDomain.PatientRepository,
	facilityRepository facilityDomain.FacilityRepository,
	areaRepository areaDomain.AreaRepository,
	addressRepository addressDomain.AddressRepository,
	userRepository userDomain.UserRepository,
	serviceCodeRepository serviceCodeDomain.ServiceCodeRepository,
) *CreatePatientUseCase {
	return &CreatePatientUseCase{
		patientRepository:     patientRepository,
		facilityRepository:    facilityRepository,
		areaRepository:        areaRepository,
		addressRepository:     addressRepository,
		userRepository:        userRepository,
		serviceCodeRepository: serviceCodeRepository,
	}
}

type CreatePatientUseCaseInputDto struct {
	Name            string
	PreferredTime   *patientDomain.PreferredTime
	PreferredGender *patientDomain.PreferredGender
	ServiceCodeID   string
	AddressID       string
	AreaID          string
	AssignedStaffID string
	FacilityID      string
}

type CreatePatientUseCaseOutputDto struct {
	ID              string
	Name            string
	PreferredTime   *patientDomain.PreferredTime
	PreferredGender *patientDomain.PreferredGender
	ServiceCode     string
	Address         string
	Area            string
	AssignedStaff   string
	Facility        string
}

func (uc *CreatePatientUseCase) Run(ctx context.Context, input CreatePatientUseCaseInputDto) (*CreatePatientUseCaseOutputDto, error) {

	// IDからそれぞれのエンティティを取得
	facility, err := uc.facilityRepository.FindByID(ctx, input.FacilityID)
	if err != nil {
		return nil, err
	}

	area, err := uc.areaRepository.FindByID(ctx, input.AreaID)
	if err != nil {
		return nil, err
	}

	address, err := uc.addressRepository.FindByID(ctx, input.AddressID)
	if err != nil {
		return nil, err
	}

	serviceCode, err := uc.serviceCodeRepository.FindByID(ctx, input.ServiceCodeID)
	if err != nil {
		return nil, err
	}

	assignedStaff, err := uc.userRepository.FindByID(ctx, input.AssignedStaffID)
	if err != nil {
		return nil, err
	}

	// ドメインモデルを作成
	patient, err := patientDomain.NewPatient(
		input.Name,
		input.PreferredTime,
		input.PreferredGender,
		serviceCode,
		address,
		area,
		assignedStaff,
		facility,
	)

	if err != nil {
		return nil, err
	}

	// リポジトリに保存
	err = uc.patientRepository.Create(ctx, patient)
	if err != nil {
		return nil, err
	}

	var joinedAddress string
	if patient.Address != nil {
		joinedAddress = patient.Address.JoinAddress()
	}

	// DTOに変換
	output := &CreatePatientUseCaseOutputDto{
		ID:              patient.ID,
		Name:            patient.Name,
		PreferredTime:   patient.PreferredTime,
		PreferredGender: patient.PreferredGender,
		ServiceCode:     patient.ServiceCode.Code,
		Address:         joinedAddress,
		Area:            patient.Area.Name,
		AssignedStaff:   patient.AssignedStaff.Username,
		Facility:        patient.Facility.Name,
	}

	return output, nil

}
