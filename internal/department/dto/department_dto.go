package dto

type DepartmentRequest struct {
	//todo add accesstoken
	Name string `json:"name" validate:"required,min=5,max=30"`
}

type DepartmentResponse struct {
	
}

type UserResponse struct {
	Email           string `json:"email"`
	Username        string `json:"name"`
	UserImageUri    string `json:"userImageUri"`
	CompanyName     string `json:"companyName"`
	CompanyImageUri string `json:"companyImageUri"`
}

type UpdateUserRequest struct {
	Email           *string `json:"email" validate:"omitempty,email,min=1,max=255"`
	Username        *string `json:"name"`
	UserImageUri    *string `json:"userImageUri"`
	CompanyName     *string `json:"companyName"`
	CompanyImageUri *string `json:"companyImageUri"`
}