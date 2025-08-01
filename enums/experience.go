package enums

type ExperienceLevel int

const (
	Junior ExperienceLevel = iota
	Mid
	Senior
	Lead
)

func (e ExperienceLevel) ToString() string {
	switch e {
	case Junior:
		return "Junior"
	case Mid:
		return "Mid"
	case Senior:
		return "Senior"
	case Lead:
		return "Lead"
	default:
		return "Unknown"
	}
}

func ExperienceLevelFromString(s string) ExperienceLevel {
	switch s {
	case "Junior":
		return Junior
	case "Mid":
		return Mid
	case "Senior":
		return Senior
	case "Lead":
		return Lead
	default:
		return Junior
	}
}
