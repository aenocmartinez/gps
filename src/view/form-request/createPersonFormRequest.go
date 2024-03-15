package formrequest

type CreatePersonFormRequest struct {
	Name          string `json:"name" binding:"required"`
	Birthdate     string `json:"birthdate" binding:"required"`
	Phone         string `json:"phone" binding:"required"`
	Address       string `json:"address" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	Guardian      string `json:"guardian" binding:"required"`
	GuardianPhone string `json:"guardianPhone" binding:"required"`
}
