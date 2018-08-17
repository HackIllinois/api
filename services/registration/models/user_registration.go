package models

type UserRegistration struct {
	ID                   string   `json:"id"                   validate:"required"`
	FirstName            string   `json:"firstName"            validate:"required"`
	LastName             string   `json:"lastName"             validate:"required"`
	Email                string   `json:"email"                validate:"required,email"`
	GitHub               string   `json:"github"               validate:"required"`
	PhoneNumber          string   `json:"phone"                validate:""`
	Gender               string   `json:"gender"               validate:"required,oneof=MALE FEMALE NONBINARY NODISCLOSE OTHER"`
	StudentType          string   `json:"studentType"          validate:"required,oneof=HIGHSCHOOL UNDERGRAD GRAD POSTGRAD ALUMNI NOSTUDENT NODISCLOSE"`
	Major                string   `json:"major"                validate:"required"`
	School               string   `json:"school"               validate:"required"`
	Transportation       string   `json:"transportation"       validate:"required,oneof=ONCAMPUS BUS DRIVINGREIMBURSEMENT DRIVINGNOREIMBURSEMENT OTHER"`
	ShirtSize            string   `json:"shirtSize"            validate:"required,oneof=S M L XL XXL"`
	Diet                 string   `json:"diet"                 validate:"required,oneof=NONE VEGAN VEGETARIAN GLUTENFREE NOREDMEAT OTHER"`
	GraduationClass      string   `json:"graduationClass"      validate:"required,oneof=FA18 SP19 FA19 SP20 FA20 SP21 FA21 SP22 FA22 SP23 AFTERSP23 NA"`
	JobInterest          []string `json:"jobInterest"          validate:"required,dive,oneof=INTERNSHIP FULLTIME COOP CITYSCHOLAR NA"`
	ProfessionalInterest []string `json:"professionalInterest" validate:"required,dive,oneof=AI SECURITY PARALLEL DATA HCI GFX ALGO THEORY HW NUMANALYSIS ML SWE CV QUANTUM CSO NETWORKING IS VIDEOGAME"`
	HeardFrom            string   `json:"heardFrom"            validate:"required,dive,oneof=FB TWTR POSTERS EMAIL FRIEND ACM WEBSITE OTHER"`
	RPInterest           []string `json:"rpInterest"           validate:"required,dive,oneof=SPEAKERS EXPERIENCES STARTUPFAIR CAREERFAIR MECHMANIA ACMSYMPOSIUM PUZZLEBANG ESCAPEROOM"`
	CreatedAt            int64    `json:"createdAt"            validate:"required"`
	UpdatedAt            int64    `json:"updatedAt"            validate:"required"`
}
