package constants

type MemberLevel int

const (
	MemberLevel0 MemberLevel = iota
	MemberLevel1
	MemberLevel2
	MemberLevel3
	MemberLevel4
	MemberLevel5
	MemberLevel6
	MemberLevel7
	MemberLevel8
	MemberLevel9
	MemberLevel10
	MemberLevel11
	MemberLevel12
)

type MemberPosition string

type MemberCategory string

const (
	MemberPositionPM       MemberPosition = "Project Manager"
	MemberPositionBE       MemberPosition = "Backend Engineer"
	MemberPositionFE       MemberPosition = "Frontend Engineer"
	MemberPositionQA       MemberPosition = "QA Engineer"
	MemberPositionDesigner MemberPosition = "UI/UX Designer"
	MemberPositionBA       MemberPosition = "Business Analyst"
	MemberPositionMobile   MemberPosition = "Mobile Developer"

	MemberCategoryOfficial MemberCategory = "official"
	MemberCategoryBuffer   MemberCategory = "buffer"
)

var (
	MemberPositions = []string{
		string(MemberPositionPM),
		string(MemberPositionBE),
		string(MemberPositionFE),
		string(MemberPositionQA),
		string(MemberPositionDesigner),
		string(MemberPositionBA),
		string(MemberPositionMobile),
	}

	MemberCategories = []string{
		string(MemberCategoryOfficial),
		string(MemberCategoryBuffer),
	}
)
