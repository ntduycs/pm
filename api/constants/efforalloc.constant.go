package constants

type LabelCategory string
type Label string

const (
	LabelCategoryManagement        LabelCategory = "Management"
	LabelCategoryCompanyActivities LabelCategory = "Company Activities"
	LabelCategoryLeaveHoliday      LabelCategory = "Leave & Holiday"
	LabelCategoryTraining          LabelCategory = "Training"
	LabelCategorySupport           LabelCategory = "Support"
	LabelCategoryIdle              LabelCategory = "Idle"
	LabelCategoryDevelopment       LabelCategory = "Development"
	LabelCategoryDocumentation     LabelCategory = "Documentation"
	LabelCategoryDesign            LabelCategory = "Design"
	LabelCategoryBA                LabelCategory = "Business Analyst"
	LabelCategoryTestCases         LabelCategory = "Test Cases"
	LabelCategoryTestExecution     LabelCategory = "Test Execution"
)

const (
	LabelPaPc              Label = "m_PA/PC"
	LabelTechAssessment    Label = "m_tech_assessment"
	LabelInternalMeeting   Label = "m_internal_meeting"
	LabelExternalMeeting   Label = "m_external_meeting"
	LabelCompanyActivities Label = "m_company_activities"
	LabelLeave             Label = "m_leave"
	LabelHoliday           Label = "m_holiday_leave"
	LabelTraining          Label = "m_training"
	LabelSupport           Label = "m_support"
	LabelIdle              Label = "m_idle"
	LabelPlanning          Label = "m_planning"

	LabelDesignSystem       Label = "t_dev_design_system"
	LabelDesignDB           Label = "t_dev_design_db"
	LabelInitProject        Label = "t_dev_design_prj_init"
	LabelDesignDetail       Label = "t_dev_design_detail"
	LabelCoding             Label = "t_dev_coding"
	LabelBugFixing          Label = "t_dev_bugfixing"
	LabelReviewCode         Label = "t_dev_review_code"
	LabelReviewDesignSystem Label = "t_dev_review_design_system"
	LabelReviewDesignDB     Label = "t_dev_review_design_db"
	LabelReviewDesignDetail Label = "t_dev_review_design_detail"
	LabelDeploy             Label = "t_dev_deploy"
	LabelResearch           Label = "t_dev_research"
	LabelDocument           Label = "t_dev_document"
	LabelUnitTest           Label = "t_dev_unit_test"

	LabelDesUx           Label = "t_des_UX"
	LabelDesUiSystem     Label = "t_des_UI_system"
	LabelDesUiMaster     Label = "t_des_UI_master"
	LabelDesUiNormal     Label = "t_des_UI_normal"
	LabelDesUiAdapt      Label = "t_des_UI_adapt"
	LabelDesReviewUx     Label = "t_des_review_UX"
	LabelDesReviewUi     Label = "t_des_review_UI"
	LabelDesResearch     Label = "t_des_research"
	LabelDesReviewOutput Label = "t_des_review_output"

	LabelReqInvestigate   Label = "t_req_invest"
	LabelReqBiz           Label = "t_req_biz"
	LabelReqSpecFunc      Label = "t_req_spec_func"
	LabelReqSpecUseCase   Label = "t_req_spec_usecase"
	LabelReqSpecWireframe Label = "t_req_spec_wireframe"
	LabelReqSpecFlowchart Label = "t_req_spec_flowchart"
	LabelReqSpecUserStory Label = "t_req_spec_userstory"
	LabelReqReview        Label = "t_req_review"
	LabelUatPlanning      Label = "t_UAT_plan"
	LabelUatExecution     Label = "t_UAT_execute"
	LabelDocUserManual    Label = "t_doc_user_manual"
	LabelDocOpGuide       Label = "t_doc_operation_guide"
	LabelDocReview        Label = "t_doc_review"
	LabelReqResearch      Label = "t_req_research"

	LabelTestFuncBreakdown Label = "t_testfunction_breakdown"
	LabelTestCreate        Label = "t_testcase_create"
	LabelTestReview        Label = "t_testcase_review"
	LabelTestExec          Label = "t_execute_testcase"
	LabelTestSuiteCreate   Label = "t_test_suite"
	LabelTestManagement    Label = "t_test_mngt"
	LabelContentVerify     Label = "t_content_verify"
	LabelBugVerify         Label = "t_verify_bug"
	LabelRegressionTest    Label = "t_regression_test"
	LabelSmokeTest         Label = "t_smoke_test"
	LabelSanityTest        Label = "t_sanity_test"
	LabelTestScriptImpl    Label = "t_testscript_implement"
	LabelTestScriptReview  Label = "t_testscript_review"
	LabelTestScriptExec    Label = "t_testscript_execute"
	LabelApiTestCreate     Label = "t_apitest_create"
	LabelApiTestExec       Label = "t_apitest_execute"
	LabelApiTestReview     Label = "t_apitest_review"
	LabelTestReport        Label = "t_test_report"
	LabelBugReport         Label = "t_bug_report"
)

var (
	LabelCategories = []LabelCategory{
		LabelCategoryManagement,
		LabelCategoryCompanyActivities,
		LabelCategoryLeaveHoliday,
		LabelCategoryTraining,
		LabelCategorySupport,
		LabelCategoryIdle,
		LabelCategoryDevelopment,
		LabelCategoryDocumentation,
		LabelCategoryDesign,
		LabelCategoryBA,
		LabelCategoryTestCases,
		LabelCategoryTestExecution,
	}

	Labels = []Label{
		LabelPaPc,
		LabelTechAssessment,
		LabelExternalMeeting,
		LabelCompanyActivities,
		LabelLeave,
		LabelHoliday,
		LabelTraining,
		LabelSupport,
		LabelIdle,
		LabelPlanning,
		LabelDesignSystem,
		LabelDesignDB,
		LabelInitProject,
		LabelDesignDetail,
		LabelCoding,
		LabelBugFixing,
		LabelReviewCode,
		LabelReviewDesignSystem,
		LabelReviewDesignDB,
		LabelReviewDesignDetail,
		LabelDeploy,
		LabelResearch,
		LabelDocument,
		LabelUnitTest,
		LabelDesUx,
		LabelDesUiSystem,
		LabelDesUiMaster,
		LabelDesUiNormal,
		LabelDesUiAdapt,
		LabelDesReviewUx,
		LabelDesReviewUi,
		LabelDesResearch,
		LabelDesReviewOutput,
		LabelReqInvestigate,
		LabelReqBiz,
		LabelReqSpecFunc,
		LabelReqSpecUseCase,
		LabelReqSpecWireframe,
		LabelReqSpecFlowchart,
		LabelReqSpecUserStory,
		LabelReqReview,
		LabelUatPlanning,
		LabelUatExecution,
		LabelDocUserManual,
		LabelDocOpGuide,
		LabelDocReview,
		LabelReqResearch,
		LabelTestFuncBreakdown,
		LabelTestCreate,
		LabelTestReview,
		LabelTestExec,
		LabelTestSuiteCreate,
		LabelTestManagement,
		LabelContentVerify,
		LabelBugVerify,
		LabelRegressionTest,
		LabelSmokeTest,
		LabelSanityTest,
		LabelTestScriptImpl,
		LabelTestScriptReview,
		LabelTestScriptExec,
		LabelApiTestCreate,
		LabelApiTestExec,
		LabelApiTestReview,
		LabelTestReport,
		LabelBugReport,
	}

	LabelCategoryMap = map[Label]LabelCategory{
		LabelPaPc:            LabelCategoryManagement,
		LabelTechAssessment:  LabelCategoryManagement,
		LabelInternalMeeting: LabelCategoryManagement,
		LabelExternalMeeting: LabelCategoryManagement,
		LabelPlanning:        LabelCategoryManagement,
		LabelTestManagement:  LabelCategoryManagement,
		LabelTestReport:      LabelCategoryManagement,

		LabelCompanyActivities: LabelCategoryCompanyActivities,

		LabelLeave:   LabelCategoryLeaveHoliday,
		LabelHoliday: LabelCategoryLeaveHoliday,

		LabelTraining:    LabelCategoryTraining,
		LabelResearch:    LabelCategoryTraining,
		LabelDesResearch: LabelCategoryTraining,
		LabelReqResearch: LabelCategoryTraining,

		LabelSupport: LabelCategorySupport,

		LabelIdle: LabelCategoryIdle,

		LabelDesignSystem:       LabelCategoryDevelopment,
		LabelDesignDB:           LabelCategoryDevelopment,
		LabelInitProject:        LabelCategoryDevelopment,
		LabelDesignDetail:       LabelCategoryDevelopment,
		LabelCoding:             LabelCategoryDevelopment,
		LabelBugFixing:          LabelCategoryDevelopment,
		LabelReviewCode:         LabelCategoryDevelopment,
		LabelReviewDesignSystem: LabelCategoryDevelopment,
		LabelReviewDesignDB:     LabelCategoryDevelopment,
		LabelReviewDesignDetail: LabelCategoryDevelopment,
		LabelDeploy:             LabelCategoryDevelopment,
		LabelUnitTest:           LabelCategoryDevelopment,

		LabelDocument: LabelCategoryDocumentation,

		LabelDesUx:           LabelCategoryDesign,
		LabelDesUiSystem:     LabelCategoryDesign,
		LabelDesUiMaster:     LabelCategoryDesign,
		LabelDesUiNormal:     LabelCategoryDesign,
		LabelDesUiAdapt:      LabelCategoryDesign,
		LabelDesReviewUx:     LabelCategoryDesign,
		LabelDesReviewUi:     LabelCategoryDesign,
		LabelDesReviewOutput: LabelCategoryDesign,

		LabelReqInvestigate:   LabelCategoryBA,
		LabelReqBiz:           LabelCategoryBA,
		LabelReqSpecFunc:      LabelCategoryBA,
		LabelReqSpecUseCase:   LabelCategoryBA,
		LabelReqSpecWireframe: LabelCategoryBA,
		LabelReqSpecFlowchart: LabelCategoryBA,
		LabelReqSpecUserStory: LabelCategoryBA,
		LabelReqReview:        LabelCategoryBA,
		LabelUatPlanning:      LabelCategoryBA,
		LabelUatExecution:     LabelCategoryBA,
		LabelDocUserManual:    LabelCategoryBA,
		LabelDocOpGuide:       LabelCategoryBA,
		LabelDocReview:        LabelCategoryBA,

		LabelTestFuncBreakdown: LabelCategoryTestCases,
		LabelTestCreate:        LabelCategoryTestCases,
		LabelTestReview:        LabelCategoryTestCases,
		LabelTestSuiteCreate:   LabelCategoryTestCases,

		LabelTestExec:         LabelCategoryTestExecution,
		LabelContentVerify:    LabelCategoryTestExecution,
		LabelBugVerify:        LabelCategoryTestExecution,
		LabelRegressionTest:   LabelCategoryTestExecution,
		LabelSmokeTest:        LabelCategoryTestExecution,
		LabelSanityTest:       LabelCategoryTestExecution,
		LabelTestScriptImpl:   LabelCategoryTestExecution,
		LabelTestScriptReview: LabelCategoryTestExecution,
		LabelTestScriptExec:   LabelCategoryTestExecution,
		LabelApiTestCreate:    LabelCategoryTestExecution,
		LabelApiTestExec:      LabelCategoryTestExecution,
		LabelApiTestReview:    LabelCategoryTestExecution,
		LabelBugReport:        LabelCategoryTestExecution,
	}

	ProductiveLabelCategories = []LabelCategory{
		LabelCategoryManagement,
		LabelCategoryDevelopment,
		LabelCategoryTestCases,
		LabelCategoryTestExecution,
		LabelCategoryBA,
		LabelCategoryDesign,
		LabelCategoryTraining,
		LabelCategoryDocumentation,
		LabelCategorySupport,
		LabelCategoryCompanyActivities,
	}
)
