package domains

import "time"

// ----- AUTH FORM -----
type BasicLoginForm struct {
	Email    string `json:"email" form:"email" valid:"required~this field is required"`
	Password string `json:"password" form:"password" valid:"required~this field is required"`
	Remember bool   `json:"remember" form:"remember"`
}

type BasicRegisterForm struct {
	Remember bool   `json:"remember" form:"remember"`
	Email    string `json:"email" form:"email" valid:"required~this field is required,email~must be a valid email format"`
	Password string `json:"password" form:"password" valid:"required~this field is required,stringlength(6|32)~password length must be 6 to 32 character,"`
}

type ServiceUserRegisterForm struct {
	BasicRegisterForm
	Address  *string `json:"address" form:"address" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Fullname *string `json:"fullname" form:"fullname" valid:"required~this field is required"`
	CardID   *string `json:"identity_card_number" form:"identity_card_number" valid:"required~this field is required,stringlength(16|16)~must contain 16 character,numeric~only numeric character"`
	Phone    *string `json:"phone" form:"phone" valid:"required~this field is required,stringlength(12|13)~must contain minimum 12 character and maximum 13 character,numeric~only numeric character"`
}

type AdminRegisterForm struct {
	BasicRegisterForm
	Address  *string `json:"address" form:"address" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Fullname *string `json:"fullname" form:"fullname" valid:"required~this field is required"`
	Phone    *string `json:"phone" form:"phone" valid:"required~this field is required,stringlength(12|13)~must contain minimum 12 character and maximum 13 character,numeric~only numeric character"`
}

type EmployeeRegisterForm struct {
	BasicRegisterForm
	Address  *string `json:"address" form:"address" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Fullname *string `json:"fullname" form:"fullname" valid:"required~this field is required"`
	CardID   *string `json:"identity_card_number" form:"identity_card_number" valid:"required~this field is required,stringlength(16|16)~must contain 16 character,numeric~only numeric character"`
	Phone    *string `json:"phone" form:"phone" valid:"required~this field is required,stringlength(12|13)~must contain minimum 12 character and maximum 13 character,numeric~only numeric character"`
}

type SupervisorRegisterForm struct {
	BasicRegisterForm
	Address  *string `json:"address" form:"address" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Fullname *string `json:"fullname" form:"fullname" valid:"required~this field is required"`
	CardID   *string `json:"identity_card_number" form:"identity_card_number" valid:"required~this field is required,stringlength(16|16)~must contain 16 character,numeric~only numeric character"`
	Phone    *string `json:"phone" form:"phone" valid:"required~this field is required,stringlength(12|13)~must contain minimum 12 character and maximum 13 character,numeric~only numeric character"`
}

func (s SupervisorRegisterForm) GetCreds() BasicRegisterForm {
	return BasicRegisterForm{
		Email:    s.Email,
		Password: s.Password,
	}
}

type UserEditForm struct {
	Password string `json:"password" form:"password" valid:"required~this field is required,stringlength(6|32)~password length must be 6 to 32 character,"`
}

// ----- END OF AUTH FORM -----
// ----- MASTER DATA FORM -----
type CategoryCreateForm struct {
	CategoryName *string `form:"category_name" json:"category_name,omitempty" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Description  *string `form:"description" json:"description,omitempty" valid:"stringlength(0|1000)~maximum 1000 character"`
	Icon         *string `json:"icon"`
}
type CategoryUpdateForm struct {
	CategoryName *string `form:"category_name" json:"category_name,omitempty" valid:"stringlength(0|255)~maximum 255 character"`
	Description  *string `form:"description" json:"description,omitempty" valid:"stringlength(0|1000)~maximum 1000 character"`
	Icon         *string `json:"icon"`
}
type DistrictCreateForm struct {
	DistrictName *string `json:"district_name,omitempty" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Description  string  `json:"description,omitempty" valid:"stringlength(0|255)~maximum 255 character"`
}
type DistrictUpdateForm struct {
	DistrictName *string `json:"district_name,omitempty" valid:"stringlength(0|255)~maximum 255 character"`
	Description  *string `json:"description,omitempty" valid:"stringlength(0|255)~maximum 255 character"`
}
type SubDistrictCreateForm struct {
	SubDistrictName *string `json:"sub_district_name,omitempty" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Description     *string `json:"description,omitempty" valid:"stringlength(0|255)~maximum 255 character"`
	DistrictID      *uint   `json:"district_id,omitempty" valid:"required~this field is required"`
}
type SubDistrictUpdateForm struct {
	SubDistrictName *string `json:"sub_district_name,omitempty" valid:"stringlength(0|255)~maximum 255 character"`
	Description     *string `json:"description,omitempty" valid:"stringlength(0|255)~maximum 255 character"`
}
type VillageCreateForm struct {
	VillageName   *string `json:"village_name,omitempty" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Description   string  `json:"description,omitempty" valid:"stringlength(0|255)~maximum 255 character"`
	SubDistrictID *uint   `json:"sub_district_id,omitempty" valid:"required~this field is required"`
}
type VillageUpdateForm struct {
	VillageName *string `json:"village_name,omitempty" valid:"stringlength(0|255)~maximum 255 character"`
	Description *string `json:"description,omitempty" valid:"stringlength(0|255)~maximum 255 character"`
}

// ----- END OF MASTER DATA FORM -----
// ----- APP SERVICE FORM -----
type ServiceItemCreateForm struct {
	// MinValue         uint   `form:"min_value" json:"min_value" valid:"range(0|1000)~maximum value is 1000 and minimum is 0"`
	ServiceID        uint   `form:"partial_service_id" json:"partial_service_id" valid:"required~this field is required"`
	ItemName         string `form:"item_name" json:"item_name" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Description      string `form:"description" json:"description" valid:"stringlength(0|255)~maximum 255 character"`
	MaxValue         uint   `form:"max_value" json:"max_value" valid:"required~this field is required,range(1|1000)~maximum value is 1000 and minimum is 1"`
	PricePerItem     uint64 `form:"price_per_item" json:"price_per_item" valid:"required~this field is required,range(0|10000000)~maximum value is 10000000 and minimum is 0"`
	IsOptionalChoice bool   `form:"is_optional_choice" json:"is_optional_choice"`
	Unit             string `form:"unit" json:"unit" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
}
type ServiceItemUpdateForm struct {
	// MinValue         uint   `form:"min_value" json:"min_value,omitempty" valid:"range(0|1000)~maximum value is 1000 and minimum is 0"`
	ItemName         string `form:"item_name" json:"item_name,omitempty" valid:"stringlength(0|255)~maximum 255 character"`
	Description      string `form:"description" json:"description,omitempty" valid:"stringlength(0|255)~maximum 255 character"`
	MaxValue         uint   `form:"max_value" json:"max_value,omitempty" valid:"range(1|1000)~maximum value is 1000 and minimum is 1"`
	PricePerItem     uint64 `form:"price_per_item" json:"price_per_item,omitempty" valid:"range(0|10000000)~maximum value is 10000000 and minimum is 0"`
	IsOptionalChoice bool   `form:"is_optional_choice" json:"is_optional_choice"`
	Unit             string `form:"unit" json:"unit,omitempty" valid:"stringlength(0|255)~maximum 255 character"`
}
type PartialServiceCreateForm struct {
	ServiceName string `form:"service_name" json:"service_name" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Description string `form:"description" json:"description" valid:"stringlength(0|1000)~maximum 1000 character"`
	BasePrice   uint64 `form:"base_price" json:"base_price" valid:"required~this field is required,int~must integer,range(0|1000000000)~min value is 0"`
	CategoryID  uint   `form:"category_id" json:"category_id" valid:"required~this field is required,int~must integer"`
}
type PartialServiceUpdateForm struct {
	ServiceName *string `form:"service_name" json:"service_name,omitempty" valid:"stringlength(0|255)~maximum 255 character"`
	Description *string `form:"description" json:"description,omitempty" valid:"stringlength(0|255)~maximum 255 character"`
	BasePrice   *uint64 `form:"base_price" json:"base_price,omitempty" valid:"int~must integer,range(0|1000000000)~min value is 0"`
}

// --> Service Package
type ServicePackageServiceItemDetailCreateForm struct {
	ServiceItemID uint `json:"service_item_id"`
	Value         uint `json:"value"`
}
type ServicePackageServiceDetailCreateForm struct {
	ServiceID    uint                                        `json:"service_id"`
	ServiceItems []ServicePackageServiceItemDetailCreateForm `json:"service_items"`
}
type ServicePackageCreateForm struct {
	PackageName string `json:"package_name" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Description string `json:"description" valid:"stringlength(0|255)~maximum 255 character"`
	BasePrice   uint64 `json:"base_price" valid:"required~this field is required,range(0|)~min value is 0"`
}
type ServicePackageUpdateForm struct {
}

// --> END OF SERVICE PACKAGE
// ----- END OF APP SERVICE FORM -----
// ----- USER FORM ------
type ServiceUserUpdateForm struct {
	Address  *string `json:"address,omitempty" form:"address" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Fullname *string `json:"Fullname,omitempty" form:"fullname" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	CardID   *string `json:"identity_card_number,omitempty" form:"identity_card_number" valid:"required~this field is required,stringlength(16|16)~must contain 16 character,numeric~only numeric character"`
	Phone    *string `json:"phone,omitempty" form:"phone" valid:"required~this field is required,stringlength(12|13)~must contain minimum 12 character and maximum 13 character,numeric~only numeric character"`
}
type EmployeeUpdateForm struct {
	Address  *string `json:"address,omitempty" form:"address" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Fullname *string `json:"fullname,omitempty" form:"fullname" valid:"required~this field is required"`
	CardID   *string `json:"identity_card_number,omitempty" form:"identity_card_number" valid:"required~this field is required,stringlength(16|16)~must contain 16 character,numeric~only numeric character"`
	Phone    *string `json:"phone,omitempty" form:"phone" valid:"required~this field is required,stringlength(12|13)~must contain minimum 12 character and maximum 13 character,numeric~only numeric character"`
}
type SupervisorUpdateForm struct {
	Address  *string `json:"address,omitempty" form:"address" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Fullname *string `json:"fullname,omitempty" form:"fullname" valid:"required~this field is required"`
	CardID   *string `json:"identity_card_number,omitempty" form:"identity_card_number" valid:"required~this field is required,stringlength(16|16)~must contain 16 character,numeric~only numeric character"`
	Phone    *string `json:"phone,omitempty" form:"phone" valid:"required~this field is required,stringlength(12|13)~must contain minimum 12 character and maximum 13 character,numeric~only numeric character"`
}
type AdminUpdateForm struct {
	Address  *string `json:"address,omitempty" form:"address" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Fullname *string `json:"fullname,omitempty" form:"fullname" valid:"required~this field is required"`
	Phone    *string `json:"phone,omitempty" form:"phone" valid:"required~this field is required,stringlength(12|13)~must contain minimum 12 character and maximum 13 character,numeric~only numeric character"`
}

// ----- END OF USER FORM -----
// ----- ORDER FORM -----
type ServiceOrderUpdateStatusForm struct {
	Status OrderStatusEnum `json:"status" form:"status" valid:"required~this field is required,in(waiting_for_payment|waiting_for_confirmation|processed|waiting_for_further_payment|ongoing|completed|cancelled)~invalid value"`
}
type ServiceOrderForm struct {
	// ServiceUserID    uint                     `json:"service_user_id" valid:"required~this field is required"`
	ContractDuration uint                     `json:"contract_duration" valid:"required~this field is required"`
	StartDate        time.Time                `json:"start_date" valid:"required~this field is required"`
	Address          string                   `json:"address" valid:"required~this field is required,stringlength(0|2000)~maximum 2000 character"`
	Note             string                   `json:"buyer_note" valid:"stringlength(0|2000)~maximum 2000 character"`
	PaymentType      PaymentTypeEnum          `json:"payment_type" valid:"required~this field is required,in(full|dp|3_termin)~accepted value are full or dp or 3_termin"`
	Details          []ServiceOrderDetailForm `json:"order_details"`
}
type ServiceOrderDetailForm struct {
	PartialServiceID uint                         `json:"partial_service_id" valid:"required~this field is required"`
	Items            []ServiceOrderDetailItemForm `json:"order_detail_items"`
}
type ServiceOrderDetailItemForm struct {
	PartialServiceItemID uint `json:"partial_service_item_id" valid:"required~this field is required"`
	Value                uint `json:"value" valid:"required~this field is required"`
}

// ----- END OF ORDER FORM -----
// ----- PLACEMENT FORM -----
type ServiceOrderPlacementCreateForm struct{}
type ServiceOrderPlacementUpdateForm struct{}
type ServiceOrderPlacementEmployeeCreateForm struct{}
type ServiceOrderPlacementEmployeeUpdateForm struct{}

// ----- END OF PLACEMENT FORM -----
