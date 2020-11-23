// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// 时间相关信息接口
type TimeModel interface {
	IsTimeModel()
}

// 经营范围的车辆或驾驶员数量
type BusinessScopeInfo struct {
	// 经营范围
	BusinessScope *string `json:"business_scope"`
	// 车辆数量
	Number *int `json:"number"`
}

// 数据页数
type DataPage struct {
	// 单个数据页中一次性返回的数据量
	Num *int `json:"num"`
	// 数据页码
	Page *int `json:"page"`
}

// 日期区间
type DateRange struct {
	// 开始时间
	Start *time.Time `json:"start"`
	// 结束时间
	End *time.Time `json:"end"`
}

// 所在县区管辖的车辆或驾驶员统计信息
type DistrictCount struct {
	// 所在县区名称
	DistrictName string `json:"district_name"`
	// 所在县区的经营范围的车辆或驾驶员的数量信息
	BusinessScopeInfoList []*BusinessScopeInfo `json:"business_scope_info_list"`
}

// 驾驶员信息数据排序
type DriverDataSort struct {
	// 排序方向
	Sort *SortDirection `json:"sort"`
	// 排序字段
	SortBy *SortableDriverField `json:"sortBy"`
}

// 驾驶员身份验证信息
type DriverIdentity struct {
	// 由golang程序生成的xid，暴露到外部使用
	IdentityID string `json:"identity_id"`
	// 身份证号码
	IDCardNum *string `json:"id_card_num"`
	// 身份证出生日期
	IDCardBirthday *time.Time `json:"id_card_birthday"`
	// 身份证签发机关
	IDCardSignGovernment *string `json:"id_card_sign_government"`
	// 身份证民族
	IDCardNation *string `json:"id_card_nation"`
	// 身份证有效起始日期
	IDCardStartDate *time.Time `json:"id_card_start_date"`
	// 身份证有效截止日期
	IDCardEndDate *time.Time `json:"id_card_end_date"`
	// 身份证正面照，云存储地址
	IDCardFrontPic *string `json:"id_card_front_pic"`
	// 身份证背面照，云存储地址
	IDCardBackPic *string `json:"id_card_back_pic"`
	// 身份证住址
	IDCardAddress *string `json:"id_card_address"`
	// 驾驶员手持身份证照片,云储存系统返回的路径
	DriverHoldingIDPhoto *string `json:"driver_holding_id_photo"`
	// 驾驶员的正面照,云储存系统返回的路径
	DriverPhoto *string `json:"driver_photo"`
	// 驾驶员签名,云储存系统返回的路径
	DriverSignature *string `json:"driver_signature"`
	// 从业资格证号码
	OccupationalNumber *string `json:"occupational_number"`
	// 从业资格证有效期至
	OccupationalExpireDate *time.Time `json:"occupational_expire_date"`
	// 从业资格证发证机构
	OccupationalIssuingAuthority *string `json:"occupational_issuing_authority"`
	// 劳动合同,云储存系统返回的完整劳动合同的图片路径
	LaborContract []*string `json:"labor_contract"`
	// 驾驶员驾驶证,云储存系统返回的路径
	DriverLicensePic *string `json:"driver_license_pic"`
	// 驾驶证发证机关
	DriverLicenseIssuingAuthority *string `json:"driver_license_issuing_authority"`
	// 年审日期（六合一）
	AnnualReviewDate *time.Time `json:"annual_review_date"`
	// 换证日期（六合一）
	RenewalDate *time.Time `json:"renewal_date"`
	// 累计积分（六合一）
	AccumulativedPoints *float64 `json:"accumulatived_points"`
	// 清分日期（六合一）
	SortingDate *time.Time `json:"sorting_date"`
	// 准驾车型（六合一）
	QuasiDrivingModels *int `json:"quasi_driving_models"`
	// 驾驶证发证所在地的省份ID
	DriverLicenseProvinceID *string `json:"driver_license_province_id"`
	// 驾驶证发证所在地的城市ID
	DriverLicenseCityID *string `json:"driver_license_city_id"`
	// 驾驶证发证所在地的区域ID
	DriverLicenseDistrictID *string `json:"driver_license_district_id"`
	// 驾驶证状态
	DriverLicenseStatus *int `json:"driver_license_status"`
	// 驾驶证初次领证日期
	DriverLicenseIssueDate *time.Time `json:"driver_license_issue_date"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 创建时间
	CreateAt time.Time `json:"create_at"`
	// 创建人
	CreateBy string `json:"create_by"`
	// 修改时间
	UpdateAt time.Time `json:"update_at"`
	// 修改人
	UpdateBy string `json:"update_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
}

func (DriverIdentity) IsTimeModel() {}

// 驾驶员信息
type DriverInfo struct {
	// 驾驶员外部编码,由golang程序生成的xid，暴露到外部使用
	DriverID string `json:"driver_id"`
	// 所在企业id
	EnterpriseID *string `json:"enterprise_id"`
	// 所在部门id
	DepartmentID *string `json:"department_id"`
	// 驾驶员身份验证信息
	DriverIdentity *DriverIdentity `json:"driver_identity"`
	// 驾驶员姓名
	DriverName *string `json:"driver_name"`
	// 手机号码
	Telephone *string `json:"telephone"`
	// 性别
	Sex int `json:"sex"`
	// 档案编号(后6位)
	FilesNumber *string `json:"files_number"`
	// 联系地址
	ContactAddress *string `json:"contact_address"`
	// 邮寄地址
	MailingAddress *string `json:"mailing_address"`
	// 是否提交
	IsSubmit *bool `json:"is_submit"`
	// 提交内容
	SubmitContent *string `json:"submit_content"`
	// 提交时间
	SubmitAt *time.Time `json:"submit_at"`
	// 提交人
	SubmitBy *string `json:"submit_by"`
	// 是否手动录入
	IsManualInput *bool `json:"is_manual_input"`
	// 是否录入
	IsInput *bool `json:"is_input"`
	// 录入时间
	InputAt *time.Time `json:"input_at"`
	// 录入人
	InputBy *string `json:"input_by"`
	// 是否校验数据
	IsCheckData *bool `json:"is_check_data"`
	// 检验时间
	CheckAt *time.Time `json:"check_at"`
	// 校验人
	CheckBy *string `json:"check_by"`
	// 驾驶员信息同步内网反馈信息
	RemarkIn *string `json:"remark_in"`
	// 内网更新时间
	UpdateTimeIn *time.Time `json:"update_time_in"`
	// 是否通过短信验证
	IsCheckSms *bool `json:"is_check_sms"`
	// 是否黑名单
	IsBlack *bool `json:"is_black"`
	// 创建时间
	CreateAt time.Time `json:"create_at"`
	// 创建人
	CreateBy string `json:"create_by"`
	// 修改时间
	UpdateAt time.Time `json:"update_at"`
	// 修改人
	UpdateBy string `json:"update_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 备注
	Remarks *string `json:"remarks"`
}

func (DriverInfo) IsTimeModel() {}

// 驾驶员信息的输入过滤
type DriverInfoFilter struct {
	// 驾驶员姓名
	DriverName *string `json:"driver_name"`
	// 驾驶员身份验证信息ID
	DriverIdentityID *string `json:"driver_identity_id"`
	// 手机号码
	Telephone *string `json:"telephone"`
	// 企业名称
	EnterpriseName *string `json:"enterprise_name"`
	// 所在县
	District *string `json:"district"`
}

// 驾驶员信息的输入
type NewDriverInfo struct {
	// 驾驶员姓名
	DriverName string `json:"driver_name"`
	// 性别
	Sex int `json:"sex"`
	// 驾驶员身份验证信息ID
	DriverIdentityID string `json:"driver_identity_id"`
	// 身份证住址
	IDCardAddress string `json:"id_card_address"`
	// 所在企业
	EnterpriseName string `json:"enterprise_name"`
	// 所在部门
	DepartmentName string `json:"department_name"`
	// 手机号码
	Telephone *string `json:"telephone"`
	// 短信验证码
	SmsVerificationCode string `json:"SMS_verification_code"`
	// 驾驶证发证所在地
	DriverLicenseAddress string `json:"driver_license_address"`
	// 档案编号(后6位)
	FileNumber string `json:"file_number"`
	// 驾驶员身份证
	IDCardPic []*graphql.Upload `json:"id_card_pic"`
	// 驾驶员驾驶证
	DriverLicense []*graphql.Upload `json:"driver_license"`
	// 劳动合同
	LaborContract []*graphql.Upload `json:"labor_contract"`
}

// 车辆驾驶员绑定信息的输入
type NewVehicleDriverBinding struct {
	// 车牌号
	LicensePlateNumber *string `json:"license_plate_number"`
	// 驾驶员姓名
	DriverName *string `json:"driver_name"`
}

// 车辆信息的输入
type NewVehicleInfo struct {
	// 车牌号
	LicensePlateNumber *string `json:"license_plate_number"`
	// 车牌颜色
	LicensePlateColor *int `json:"license_plate_color"`
	// 号牌种类
	LicensePlateType *int `json:"license_plate_type"`
	// 企业名称
	EnterpriseName *string `json:"enterprise_name"`
	// 车辆类型
	VehicleType *int `json:"vehicle_type"`
	// 经营范围
	BusinessScope *int `json:"business_scope"`
	// 车架号(后6位)
	VehicleIdentificationNumber *string `json:"vehicle_identification_number"`
	// 机动车状态
	VehicleState *int `json:"vehicle_state"`
	// 机动车所有人（六合一）
	Owner *string `json:"owner"`
	// 检验日期（六合一）
	InspectionDate *time.Time `json:"inspection_date"`
	// 报废日期(六合一)
	RetirementDate *time.Time `json:"retirement_date"`
	// 使用性质（六合一）
	UseNature *string `json:"use_nature"`
}

// 车辆信息数据排序
type VehicleDataSort struct {
	// 排序方向
	Sort *SortDirection `json:"sort"`
	// 排序字段
	SortBy *SortableVehicleField `json:"sortBy"`
}

// 车辆驾驶员绑定表
type VehicleDriverBinding struct {
	// 由golang程序生成的xid，暴露到外部使用
	VehicleDriverBindingID string `json:"vehicle_driver_binding_id"`
	// 驾驶员id
	DriverID *string `json:"driver_id"`
	// 车辆id
	VehicleID *string `json:"vehicle_id"`
	// 备注
	Remarks *string `json:"remarks"`
	// 租赁合同,,云储存系统返回的完整租赁合同的图片路径
	LeaseContract []*string `json:"lease_contract"`
	// 车牌号
	LicensePlateNumber *string `json:"license_plate_number"`
	// 车牌颜色
	LicensePlateColor *int `json:"license_plate_color"`
	// 号牌种类
	LicensePlateType *int `json:"license_plate_type"`
	// 驾驶员姓名
	DriverName *string `json:"driver_name"`
	// 是否绑定
	IsBinding *bool `json:"is_binding"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 创建时间
	CreateAt time.Time `json:"create_at"`
	// 创建人
	CreateBy string `json:"create_by"`
	// 修改时间
	UpdateAt time.Time `json:"update_at"`
	// 修改人
	UpdateBy string `json:"update_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
}

func (VehicleDriverBinding) IsTimeModel() {}

// 车辆驾驶员绑定信息的输入过滤
type VehicleDriverBindingFilter struct {
	// 车牌号
	LicensePlateNumber *string `json:"license_plate_number"`
	// 经营范围
	BusinessScope *int `json:"business_scope"`
	// 绑定状态
	IsBinding *bool `json:"is_binding"`
	// 企业名称
	EnterpriseName *string `json:"enterprise_name"`
	// 所在县
	District *string `json:"district"`
}

// 车辆信息
type VehicleInfo struct {
	// 车辆外部编码,由golang程序生成的xid，暴露到外部使用
	VehicleID string `json:"vehicle_id"`
	// 车牌号
	LicensePlateNumber *string `json:"license_plate_number"`
	// 车牌颜色
	LicensePlateColor *int `json:"license_plate_color"`
	// 号牌种类
	LicensePlateType *int `json:"license_plate_type"`
	// 所在企业id
	EnterpriseID *string `json:"enterprise_id"`
	// 企业名称
	EnterpriseName *string `json:"enterprise_name"`
	// 所在县
	District *string `json:"district"`
	// 所在部门id
	DepartmentID *string `json:"department_id"`
	// 车架号(后6位)
	VehicleIdentificationNumber *string `json:"vehicle_identification_number"`
	// 道路运输证号
	RoadTransportLicenseNumber *string `json:"road_transport_license_number"`
	// 车辆类型
	VehicleType *int `json:"vehicle_type"`
	// 行业类别
	IndustryCategory *int `json:"industry_category"`
	// 吨位
	Heavy *float64 `json:"heavy"`
	// 座位
	Seats *int `json:"seats"`
	// 营运类型
	OperatingType *int `json:"operating_type"`
	// 营运线路
	OperatingRoute *string `json:"operating_route"`
	// 经营范围
	BusinessScope *int `json:"business_scope"`
	// 机动车管理人
	VehicleManager *string `json:"vehicle_manager"`
	// 机动车管理人联系电话
	VehicleManagerPhone *string `json:"vehicle_manager_phone"`
	// 机动车管理人身份证
	VehicleManagerIDCard *string `json:"vehicle_manager_id_card"`
	// 机动车所有人（六合一）
	Owner *string `json:"owner"`
	// 检验日期（六合一）
	InspectionDate *time.Time `json:"inspection_date"`
	// 报废日期(六合一)
	RetirementDate *time.Time `json:"retirement_date"`
	// 使用性质（六合一）
	UseNature *string `json:"use_nature"`
	// 机动车状态
	VehicleState *int `json:"vehicle_state"`
	// 内网更新时间
	UpdateTimeIn *time.Time `json:"update_time_in"`
	// 车辆信息同步内网反馈信息
	RemarkIn *string `json:"remark_in"`
	// 是否完成
	IsComplete *bool `json:"is_complete"`
	// 行驶证照片,云储存系统返回的路径
	DrivingLicenseePic *string `json:"driving_licensee_pic"`
	// 是否激活
	IsActive *bool `json:"is_active"`
	// 是否录入完成
	IsInput *bool `json:"is_input"`
	// 租车标准价格
	CarRentalPrice *float64 `json:"car_rental_price"`
	// 投保公司
	InsuranceCompany *int `json:"insurance_company"`
	// 投保日期
	InsuranceDate *time.Time `json:"insurance_date"`
	// 维保数据数组
	VehicleMaintenances []*VehicleMaintenance `json:"vehicle_maintenances"`
	// 汽车排量
	VehicleDisplacement *string `json:"vehicle_displacement"`
	// 车辆品牌
	VehicleBrand *int `json:"vehicle_brand"`
	// 准驾车型
	QuasiDrivingModels *int `json:"quasi_driving_models"`
	// 是否上传省厅
	IsUploadProvince *bool `json:"is_upload_province"`
	// 营运状态
	OperatingState *int `json:"operating_state"`
	// 终端ID
	TerminalID *string `json:"terminal_id"`
	// 是否申请安装智能终端
	IsApplyInstallTerminal *bool `json:"is_apply_install_terminal"`
	// 校验状态
	CheckState *int `json:"check_state"`
	// 是否导入
	IsImport *bool `json:"is_import"`
	// 登记时间
	RecordAt *time.Time `json:"record_at"`
	// 登记人
	RecordBy *string `json:"record_by"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 备注
	Remarks *string `json:"remarks"`
	// 创建时间
	CreateAt time.Time `json:"create_at"`
	// 创建人
	CreateBy string `json:"create_by"`
	// 修改时间
	UpdateAt time.Time `json:"update_at"`
	// 修改人
	UpdateBy string `json:"update_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
}

func (VehicleInfo) IsTimeModel() {}

// 车辆信息的输入过滤
type VehicleInfoFilter struct {
	// 车牌号
	LicensePlateNumber *string `json:"license_plate_number"`
	// 经营范围
	BusinessScope *int `json:"business_scope"`
	// 机动车状态
	VehicleState *int `json:"vehicle_state"`
	// 登记时间从
	RegisterBetween *DateRange `json:"registerBetween"`
	// 车辆类型
	VehicleType *int `json:"vehicle_type"`
	// 检验日期（六合一）
	InspectionDate *time.Time `json:"inspection_date"`
	// 使用性质（六合一）
	UseNature *string `json:"use_nature"`
	// 企业名称
	EnterpriseName *string `json:"enterprise_name"`
	// 所在县
	District *string `json:"district"`
}

// 车辆维保数据
type VehicleMaintenance struct {
	// 维保时间
	MaintenanceDate *time.Time `json:"maintenance_date"`
	// 维保公里数
	MaintenanceKilometers *float64 `json:"maintenance_kilometers"`
}

// 排序方向枚举
type SortDirection string

const (
	// 升序
	SortDirectionAscending SortDirection = "ASCENDING"
	// 降序
	SortDirectionDescending SortDirection = "DESCENDING"
)

var AllSortDirection = []SortDirection{
	SortDirectionAscending,
	SortDirectionDescending,
}

func (e SortDirection) IsValid() bool {
	switch e {
	case SortDirectionAscending, SortDirectionDescending:
		return true
	}
	return false
}

func (e SortDirection) String() string {
	return string(e)
}

func (e *SortDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortDirection", str)
	}
	return nil
}

func (e SortDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// 车辆信息可排序字段
type SortableDriverField string

const (
	// 创建时间
	SortableDriverFieldCreated SortableDriverField = "created"
	// 驾驶员身份验证信息ID
	SortableDriverFieldDriverIdentityID SortableDriverField = "driver_identity_id"
)

var AllSortableDriverField = []SortableDriverField{
	SortableDriverFieldCreated,
	SortableDriverFieldDriverIdentityID,
}

func (e SortableDriverField) IsValid() bool {
	switch e {
	case SortableDriverFieldCreated, SortableDriverFieldDriverIdentityID:
		return true
	}
	return false
}

func (e SortableDriverField) String() string {
	return string(e)
}

func (e *SortableDriverField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortableDriverField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortableDriverField", str)
	}
	return nil
}

func (e SortableDriverField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// 车辆信息可排序字段
type SortableVehicleField string

const (
	// 创建时间
	SortableVehicleFieldCreated SortableVehicleField = "created"
	// 车牌号
	SortableVehicleFieldLicensePlateNumber SortableVehicleField = "license_plate_number"
	// 汽车排量
	SortableVehicleFieldVehicleDisplacement SortableVehicleField = "vehicle_displacement"
	// 租车标准价格
	SortableVehicleFieldCarRentalPrice SortableVehicleField = "car_rental_price"
)

var AllSortableVehicleField = []SortableVehicleField{
	SortableVehicleFieldCreated,
	SortableVehicleFieldLicensePlateNumber,
	SortableVehicleFieldVehicleDisplacement,
	SortableVehicleFieldCarRentalPrice,
}

func (e SortableVehicleField) IsValid() bool {
	switch e {
	case SortableVehicleFieldCreated, SortableVehicleFieldLicensePlateNumber, SortableVehicleFieldVehicleDisplacement, SortableVehicleFieldCarRentalPrice:
		return true
	}
	return false
}

func (e SortableVehicleField) String() string {
	return string(e)
}

func (e *SortableVehicleField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortableVehicleField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortableVehicleField", str)
	}
	return nil
}

func (e SortableVehicleField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
