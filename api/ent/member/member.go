// Code generated by ent, DO NOT EDIT.

package member

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the member type in the database.
	Label = "member"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// FieldPositions holds the string denoting the positions field in the database.
	FieldPositions = "positions"
	// FieldKpi holds the string denoting the kpi field in the database.
	FieldKpi = "kpi"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldTotalEffort holds the string denoting the total_effort field in the database.
	FieldTotalEffort = "total_effort"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the member in the database.
	Table = "members"
)

// Columns holds all SQL columns for member fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldLevel,
	FieldPositions,
	FieldKpi,
	FieldCategory,
	FieldTotalEffort,
	FieldStartDate,
	FieldEndDate,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// KpiValidator is a validator for the "kpi" field. It is called by the builders before save.
	KpiValidator func(float32) error
	// TotalEffortValidator is a validator for the "total_effort" field. It is called by the builders before save.
	TotalEffortValidator func(float32) error
)

// Level defines the type for the "level" enum field.
type Level string

// Level values.
const (
	LevelLV1  Level = "LV1"
	LevelLV2  Level = "LV2"
	LevelLV3  Level = "LV3"
	LevelLV4  Level = "LV4"
	LevelLV5  Level = "LV5"
	LevelLV6  Level = "LV6"
	LevelLV7  Level = "LV7"
	LevelLV8  Level = "LV8"
	LevelLV9  Level = "LV9"
	LevelLV10 Level = "LV10"
)

func (l Level) String() string {
	return string(l)
}

// LevelValidator is a validator for the "level" field enum values. It is called by the builders before save.
func LevelValidator(l Level) error {
	switch l {
	case LevelLV1, LevelLV2, LevelLV3, LevelLV4, LevelLV5, LevelLV6, LevelLV7, LevelLV8, LevelLV9, LevelLV10:
		return nil
	default:
		return fmt.Errorf("member: invalid enum value for level field: %q", l)
	}
}

// Category defines the type for the "category" enum field.
type Category string

// Category values.
const (
	CategoryOfficial Category = "Official"
	CategoryBuffer   Category = "Buffer"
)

func (c Category) String() string {
	return string(c)
}

// CategoryValidator is a validator for the "category" field enum values. It is called by the builders before save.
func CategoryValidator(c Category) error {
	switch c {
	case CategoryOfficial, CategoryBuffer:
		return nil
	default:
		return fmt.Errorf("member: invalid enum value for category field: %q", c)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusPending  Status = "pending"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusActive, StatusInactive, StatusPending:
		return nil
	default:
		return fmt.Errorf("member: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Member queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByLevel orders the results by the level field.
func ByLevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLevel, opts...).ToFunc()
}

// ByPositions orders the results by the positions field.
func ByPositions(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPositions, opts...).ToFunc()
}

// ByKpi orders the results by the kpi field.
func ByKpi(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKpi, opts...).ToFunc()
}

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
}

// ByTotalEffort orders the results by the total_effort field.
func ByTotalEffort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalEffort, opts...).ToFunc()
}

// ByStartDate orders the results by the start_date field.
func ByStartDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartDate, opts...).ToFunc()
}

// ByEndDate orders the results by the end_date field.
func ByEndDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndDate, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}
