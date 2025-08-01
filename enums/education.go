package enums

type EducationLevel int

const (
	Secondary EducationLevel = iota
	Higher
	Incomplete
	Bachelor
	Master
	Doctor
)

func (e EducationLevel) ToString() string {
	switch e {
	case Secondary:
		return "SECONDARY"
	case Higher:
		return "HIGHER"
	case Incomplete:
		return "INCOMPLETE"
	case Bachelor:
		return "BACHELOR"
	case Master:
		return "MASTER"
	case Doctor:
		return "DOCTOR"
	default:
		return ""
	}
}

func EducationLevelFromString(name string) EducationLevel {
	switch name {
	case "SECONDARY":
		return Secondary
	case "HIGHER":
		return Higher
	case "INCOMPLETE":
		return Incomplete
	case "BACHELOR":
		return Bachelor
	case "MASTER":
		return Master
	case "DOCTOR":
		return Doctor
	default:
		return Secondary
	}
}
