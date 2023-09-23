// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management/ent/member"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Member is the model entity for the Member schema.
type Member struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Level holds the value of the "level" field.
	Level member.Level `json:"level,omitempty"`
	// comma separated values of positions
	Positions string `json:"positions,omitempty"`
	// key performance indicator
	Kpi float32 `json:"kpi,omitempty"`
	// Category holds the value of the "category" field.
	Category member.Category `json:"category,omitempty"`
	// total effort in percentage
	TotalEffort float32 `json:"total_effort,omitempty"`
	// start date of member
	StartDate *time.Time `json:"start_date,omitempty"`
	// end date of member
	EndDate *time.Time `json:"end_date,omitempty"`
	// Status holds the value of the "status" field.
	Status       member.Status `json:"status,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Member) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case member.FieldKpi, member.FieldTotalEffort:
			values[i] = new(sql.NullFloat64)
		case member.FieldID:
			values[i] = new(sql.NullInt64)
		case member.FieldName, member.FieldLevel, member.FieldPositions, member.FieldCategory, member.FieldStatus:
			values[i] = new(sql.NullString)
		case member.FieldStartDate, member.FieldEndDate:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Member fields.
func (m *Member) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case member.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case member.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case member.FieldLevel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				m.Level = member.Level(value.String)
			}
		case member.FieldPositions:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field positions", values[i])
			} else if value.Valid {
				m.Positions = value.String
			}
		case member.FieldKpi:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field kpi", values[i])
			} else if value.Valid {
				m.Kpi = float32(value.Float64)
			}
		case member.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				m.Category = member.Category(value.String)
			}
		case member.FieldTotalEffort:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field total_effort", values[i])
			} else if value.Valid {
				m.TotalEffort = float32(value.Float64)
			}
		case member.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_date", values[i])
			} else if value.Valid {
				m.StartDate = new(time.Time)
				*m.StartDate = value.Time
			}
		case member.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_date", values[i])
			} else if value.Valid {
				m.EndDate = new(time.Time)
				*m.EndDate = value.Time
			}
		case member.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				m.Status = member.Status(value.String)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Member.
// This includes values selected through modifiers, order, etc.
func (m *Member) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// Update returns a builder for updating this Member.
// Note that you need to call Member.Unwrap() before calling this method if this Member
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Member) Update() *MemberUpdateOne {
	return NewMemberClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Member entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Member) Unwrap() *Member {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Member is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Member) String() string {
	var builder strings.Builder
	builder.WriteString("Member(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("level=")
	builder.WriteString(fmt.Sprintf("%v", m.Level))
	builder.WriteString(", ")
	builder.WriteString("positions=")
	builder.WriteString(m.Positions)
	builder.WriteString(", ")
	builder.WriteString("kpi=")
	builder.WriteString(fmt.Sprintf("%v", m.Kpi))
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(fmt.Sprintf("%v", m.Category))
	builder.WriteString(", ")
	builder.WriteString("total_effort=")
	builder.WriteString(fmt.Sprintf("%v", m.TotalEffort))
	builder.WriteString(", ")
	if v := m.StartDate; v != nil {
		builder.WriteString("start_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := m.EndDate; v != nil {
		builder.WriteString("end_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", m.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Members is a parsable slice of Member.
type Members []*Member
