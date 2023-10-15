package services

import (
	"context"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	"go.uber.org/zap"

	"project-management/constants"
	"project-management/mappers"
	"project-management/models"
	"project-management/repositories"
)

const (
	startRow       = 3
	startColumn    = 6
	columnTaskCode = 2
	columnTaskName = 3
	columnLabels   = 5
)

var (
	totalIssue, _ = regexp.Compile(`Total \(\d+ issues\)`)
)

type ExcelMember struct {
	Name   string
	Column int
	Member *models.Member
}

func getMemberByColumn(column int, memberList []ExcelMember) string {
	for _, member := range memberList {
		if member.Column == column {
			return member.Name
		}
	}

	panic("Cannot find member by column")
}

func getColumnByMember(memberName string, memberList []ExcelMember) int {
	for _, member := range memberList {
		if member.Name == memberName {
			return member.Column
		}
	}

	panic("Cannot find column by member")
}

type ExcelLabel struct {
	Name constants.Label
	Row  int
}

func getLabelByRow(row int, labelList []ExcelLabel) constants.Label {
	for _, label := range labelList {
		if label.Row == row {
			return label.Name
		}
	}

	panic("Cannot find label by row")
}

type EffortAllocService struct {
	logger           *zap.Logger
	memberRepository *repositories.MemberRepository
	memberMapper     *mappers.MemberMapper
}

func NewEffortAllocService(
	logger *zap.Logger,
	memberRepository *repositories.MemberRepository,
	memberMapper *mappers.MemberMapper,
) *EffortAllocService {
	return &EffortAllocService{
		logger:           logger,
		memberRepository: memberRepository,
		memberMapper:     memberMapper,
	}
}

func (s *EffortAllocService) ParseWeeklyReport(ctx context.Context, rows [][]string) (*models.UploadEaWeeklyReportResponse, error) {
	memberList, err := s.readMemberList(ctx, rows)

	if err != nil {
		return nil, err
	}

	memberByJiraName := lo.SliceToMap(memberList, func(member ExcelMember) (string, *models.Member) {
		return member.Name, member.Member
	})

	endRow := s.detectEndRow(rows)

	labelList := s.readLabelList(rows)

	firstEffortRow := startRow + 1 // Skip the first row (containing the member list)
	firstEffortCol := startColumn
	lastEffortRow := endRow - 1
	lastEffortCol := len(memberList) + startColumn - 1

	effortAllocation := make(map[string]map[constants.LabelCategory]int)

	// Print effort usage area
	for col := firstEffortCol; col <= lastEffortCol; col++ {
		member := getMemberByColumn(col, memberList)

		for row := firstEffortRow; row <= lastEffortRow; row++ {
			label := getLabelByRow(row, labelList)
			labelCategory := constants.LabelCategoryMap[label]

			timeLogged, _ := strconv.Atoi(strings.TrimSuffix(rows[row][col], "h"))

			if _, ok := effortAllocation[member]; !ok {
				effortAllocation[member] = make(map[constants.LabelCategory]int)
			}

			effortAllocation[member][labelCategory] += timeLogged
		}
	}

	resp := &models.UploadEaWeeklyReportResponse{
		ListResponse: models.ListResponse{
			Total: len(effortAllocation),
		},
		Items: lo.MapToSlice(effortAllocation, func(jiraName string, efforts map[constants.LabelCategory]int) models.EaWeeklyMember {
			return models.EaWeeklyMember{
				Member: memberByJiraName[jiraName],
				Efforts: lo.MapToSlice(efforts, func(category constants.LabelCategory, timeLogged int) *models.EaWeeklyEffort {
					return &models.EaWeeklyEffort{
						Category: category,
						Time:     timeLogged,
					}
				}),
			}
		}),
	}

	return resp, nil
}

func (s *EffortAllocService) readMemberList(ctx context.Context, rows [][]string) ([]ExcelMember, error) {
	memberList := make([]ExcelMember, 0)

	jiraNames := make([]string, 0)

	for row, rowContent := range rows {
		if row == startRow {
			for col, colContent := range rowContent {
				if col >= startColumn {
					colContent = strings.TrimSpace(colContent)

					if colContent != "" && colContent != "Total" {
						memberList = append(memberList, ExcelMember{
							Name:   colContent,
							Column: col,
						})

						jiraNames = append(jiraNames, colContent)
					}
				}
			}

			// Stop reading rows, we have the member list already
			break
		}
	}

	// Get the member list from the database
	members, err := s.memberRepository.FindAllByJiraNameIn(ctx, jiraNames)

	if err != nil {
		return nil, err
	}

	memberMap := make(map[string]*models.Member)

	for _, member := range members {
		memberMap[member.JiraName] = s.memberMapper.ToModel(member)
	}

	for i, member := range memberList {
		if _, ok := memberMap[member.Name]; !ok {
			return nil, errors.New("member " + member.Name + " not found")
		}

		memberList[i].Member = memberMap[member.Name]
	}

	// Sort member list alphabetically
	sort.SliceStable(memberList, func(i, j int) bool {
		return memberList[i].Name < memberList[j].Name
	})

	return memberList, nil
}

func (s *EffortAllocService) readLabelList(rows [][]string) []ExcelLabel {
	labelList := make([]ExcelLabel, 0)

	for row, rowContent := range rows {
		// Skip the first row (containing the member list)
		if row >= startRow+1 {
			if s.isLastRow(rowContent) {
				break
			}

			labels := strings.Split(rowContent[columnLabels], ",")

			if len(labels) > 1 {
				s.logger.Panic("more than one label in one task", zap.String("task", rowContent[columnTaskCode]), zap.String("labels", rowContent[columnLabels]))
			}

			for _, label := range labels {
				label = strings.TrimSpace(label)

				if label == "" {
					s.logger.Panic("empty label", zap.String("task", rowContent[columnTaskCode]))
				}

				if _, ok := constants.LabelCategoryMap[constants.Label(label)]; !ok {
					s.logger.Panic("unknown label", zap.String("task", rowContent[columnTaskCode]), zap.String("label", label))
				}

				labelList = append(labelList, ExcelLabel{
					Name: constants.Label(label),
					Row:  row,
				})
			}
		}
	}

	return labelList
}

func (s *EffortAllocService) detectEndRow(rows [][]string) int {
	for row, rowContent := range rows {
		if s.isLastRow(rowContent) {
			return row
		}
	}

	panic("cannot detect end row")
}

func (s *EffortAllocService) isLastRow(rowContent []string) bool {
	return totalIssue.Match([]byte(rowContent[0]))
}
