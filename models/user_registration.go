package models

type UserRegistration struct {
	ID                   string              `json:"id"                   validate:"required"`
	FirstName            string              `json:"firstName"            validate:"required"`
	LastName             string              `json:"lastName"             validate:"required"`
	Email                string              `json:"email"                validate:"required,email"`
	ShirtSize            string              `json:"shirtSize"            validate:"required,oneof=S M L XL"`
	Diet                 string              `json:"diet"                 validate:"required,oneof=NONE VEGAN VEGETARIAN"`
	Age                  int                 `json:"age"                  validate:"required"`
	GraduationYear       int                 `json:"graduationYear"       validate:"required"`
	Transportation       string              `json:"transportation"       validate:"required,oneof=NONE BUS"`
	School               string              `json:"school"               validate:"required"`
	Major                string              `json:"major"                validate:"required"`
	Gender               string              `json:gender                 validate:"required,oneof=MALE FEMALE NONBINARY OTHER"`
	ProfessionalInterest string              `json:"professionalInterest" validate:"required,oneof=INTERNSHIP FULLTIME BOTH"`
	GitHub               string              `json:"github"               validate:"required"`
	Linkedin             string              `json"linkedin"              validate:"required"`
	Interests            string              `json:"interests"            validate:"required"`
	IsNovice             bool                `json:"isNovice"             validate:"required|isdefault"`
	IsPrivate            bool                `json:"isPrivate"            validate:"required|isdefault"`
	PhoneNumber          string              `json:"phoneNumber"          validate:"required"`
	LongForms            []UserLongForm      `json:"longforms"            validate:"required,dive,required"`
	ExtraInfos           []UserExtraInfo     `json:"extraInfos"           validate:"required,dive,required"`
	OsContributors       []UserOsContributor `json:"osContributors"       validate:"required,dive,required"`
	Collaborators        []UserCollaborator  `json:"collaborators"        validate:"required,dive,required"`
}

type UserLongForm struct {
	Response string `json:"response" validate:"required"`
}

type UserExtraInfo struct {
	Response string `json:"response" validate:"required"`
}

type UserOsContributor struct {
	Name        string `json:"name"        validate:"required"`
	ContactInfo string `json:"contactInfo" validate:"required"`
}

type UserCollaborator struct {
	Github string `json:github validate:"required"`
}
