package consts

type UserGender int

const (
	UserGenderNone   UserGender = 0 //未知
	UserGenderMale              = 1 //男性
	UserGenderFemale            = 2 //女性
)

type EnumOption struct {
	Key  string `json:"key"`
	Text string `json:"text"`
}
