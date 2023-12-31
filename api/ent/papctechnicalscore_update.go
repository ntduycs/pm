// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management/ent/papc"
	"project-management/ent/papctechnicalscore"
	"project-management/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PaPcTechnicalScoreUpdate is the builder for updating PaPcTechnicalScore entities.
type PaPcTechnicalScoreUpdate struct {
	config
	hooks    []Hook
	mutation *PaPcTechnicalScoreMutation
}

// Where appends a list predicates to the PaPcTechnicalScoreUpdate builder.
func (pptsu *PaPcTechnicalScoreUpdate) Where(ps ...predicate.PaPcTechnicalScore) *PaPcTechnicalScoreUpdate {
	pptsu.mutation.Where(ps...)
	return pptsu
}

// SetPaPcID sets the "pa_pc_id" field.
func (pptsu *PaPcTechnicalScoreUpdate) SetPaPcID(i int) *PaPcTechnicalScoreUpdate {
	pptsu.mutation.SetPaPcID(i)
	return pptsu
}

// SetType sets the "type" field.
func (pptsu *PaPcTechnicalScoreUpdate) SetType(pa papctechnicalscore.Type) *PaPcTechnicalScoreUpdate {
	pptsu.mutation.SetType(pa)
	return pptsu
}

// SetSkill sets the "skill" field.
func (pptsu *PaPcTechnicalScoreUpdate) SetSkill(s string) *PaPcTechnicalScoreUpdate {
	pptsu.mutation.SetSkill(s)
	return pptsu
}

// SetSelfScore sets the "self_score" field.
func (pptsu *PaPcTechnicalScoreUpdate) SetSelfScore(f float32) *PaPcTechnicalScoreUpdate {
	pptsu.mutation.ResetSelfScore()
	pptsu.mutation.SetSelfScore(f)
	return pptsu
}

// AddSelfScore adds f to the "self_score" field.
func (pptsu *PaPcTechnicalScoreUpdate) AddSelfScore(f float32) *PaPcTechnicalScoreUpdate {
	pptsu.mutation.AddSelfScore(f)
	return pptsu
}

// SetPeerScore sets the "peer_score" field.
func (pptsu *PaPcTechnicalScoreUpdate) SetPeerScore(f float32) *PaPcTechnicalScoreUpdate {
	pptsu.mutation.ResetPeerScore()
	pptsu.mutation.SetPeerScore(f)
	return pptsu
}

// SetNillablePeerScore sets the "peer_score" field if the given value is not nil.
func (pptsu *PaPcTechnicalScoreUpdate) SetNillablePeerScore(f *float32) *PaPcTechnicalScoreUpdate {
	if f != nil {
		pptsu.SetPeerScore(*f)
	}
	return pptsu
}

// AddPeerScore adds f to the "peer_score" field.
func (pptsu *PaPcTechnicalScoreUpdate) AddPeerScore(f float32) *PaPcTechnicalScoreUpdate {
	pptsu.mutation.AddPeerScore(f)
	return pptsu
}

// ClearPeerScore clears the value of the "peer_score" field.
func (pptsu *PaPcTechnicalScoreUpdate) ClearPeerScore() *PaPcTechnicalScoreUpdate {
	pptsu.mutation.ClearPeerScore()
	return pptsu
}

// SetManagerScore sets the "manager_score" field.
func (pptsu *PaPcTechnicalScoreUpdate) SetManagerScore(f float32) *PaPcTechnicalScoreUpdate {
	pptsu.mutation.ResetManagerScore()
	pptsu.mutation.SetManagerScore(f)
	return pptsu
}

// AddManagerScore adds f to the "manager_score" field.
func (pptsu *PaPcTechnicalScoreUpdate) AddManagerScore(f float32) *PaPcTechnicalScoreUpdate {
	pptsu.mutation.AddManagerScore(f)
	return pptsu
}

// SetPaPc sets the "pa_pc" edge to the PaPc entity.
func (pptsu *PaPcTechnicalScoreUpdate) SetPaPc(p *PaPc) *PaPcTechnicalScoreUpdate {
	return pptsu.SetPaPcID(p.ID)
}

// Mutation returns the PaPcTechnicalScoreMutation object of the builder.
func (pptsu *PaPcTechnicalScoreUpdate) Mutation() *PaPcTechnicalScoreMutation {
	return pptsu.mutation
}

// ClearPaPc clears the "pa_pc" edge to the PaPc entity.
func (pptsu *PaPcTechnicalScoreUpdate) ClearPaPc() *PaPcTechnicalScoreUpdate {
	pptsu.mutation.ClearPaPc()
	return pptsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pptsu *PaPcTechnicalScoreUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pptsu.sqlSave, pptsu.mutation, pptsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pptsu *PaPcTechnicalScoreUpdate) SaveX(ctx context.Context) int {
	affected, err := pptsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pptsu *PaPcTechnicalScoreUpdate) Exec(ctx context.Context) error {
	_, err := pptsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pptsu *PaPcTechnicalScoreUpdate) ExecX(ctx context.Context) {
	if err := pptsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pptsu *PaPcTechnicalScoreUpdate) check() error {
	if v, ok := pptsu.mutation.GetType(); ok {
		if err := papctechnicalscore.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "PaPcTechnicalScore.type": %w`, err)}
		}
	}
	if _, ok := pptsu.mutation.PaPcID(); pptsu.mutation.PaPcCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PaPcTechnicalScore.pa_pc"`)
	}
	return nil
}

func (pptsu *PaPcTechnicalScoreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pptsu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(papctechnicalscore.Table, papctechnicalscore.Columns, sqlgraph.NewFieldSpec(papctechnicalscore.FieldID, field.TypeInt))
	if ps := pptsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pptsu.mutation.GetType(); ok {
		_spec.SetField(papctechnicalscore.FieldType, field.TypeEnum, value)
	}
	if value, ok := pptsu.mutation.Skill(); ok {
		_spec.SetField(papctechnicalscore.FieldSkill, field.TypeString, value)
	}
	if value, ok := pptsu.mutation.SelfScore(); ok {
		_spec.SetField(papctechnicalscore.FieldSelfScore, field.TypeFloat32, value)
	}
	if value, ok := pptsu.mutation.AddedSelfScore(); ok {
		_spec.AddField(papctechnicalscore.FieldSelfScore, field.TypeFloat32, value)
	}
	if value, ok := pptsu.mutation.PeerScore(); ok {
		_spec.SetField(papctechnicalscore.FieldPeerScore, field.TypeFloat32, value)
	}
	if value, ok := pptsu.mutation.AddedPeerScore(); ok {
		_spec.AddField(papctechnicalscore.FieldPeerScore, field.TypeFloat32, value)
	}
	if pptsu.mutation.PeerScoreCleared() {
		_spec.ClearField(papctechnicalscore.FieldPeerScore, field.TypeFloat32)
	}
	if value, ok := pptsu.mutation.ManagerScore(); ok {
		_spec.SetField(papctechnicalscore.FieldManagerScore, field.TypeFloat32, value)
	}
	if value, ok := pptsu.mutation.AddedManagerScore(); ok {
		_spec.AddField(papctechnicalscore.FieldManagerScore, field.TypeFloat32, value)
	}
	if pptsu.mutation.PaPcCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   papctechnicalscore.PaPcTable,
			Columns: []string{papctechnicalscore.PaPcColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(papc.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pptsu.mutation.PaPcIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   papctechnicalscore.PaPcTable,
			Columns: []string{papctechnicalscore.PaPcColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(papc.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pptsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{papctechnicalscore.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pptsu.mutation.done = true
	return n, nil
}

// PaPcTechnicalScoreUpdateOne is the builder for updating a single PaPcTechnicalScore entity.
type PaPcTechnicalScoreUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PaPcTechnicalScoreMutation
}

// SetPaPcID sets the "pa_pc_id" field.
func (pptsuo *PaPcTechnicalScoreUpdateOne) SetPaPcID(i int) *PaPcTechnicalScoreUpdateOne {
	pptsuo.mutation.SetPaPcID(i)
	return pptsuo
}

// SetType sets the "type" field.
func (pptsuo *PaPcTechnicalScoreUpdateOne) SetType(pa papctechnicalscore.Type) *PaPcTechnicalScoreUpdateOne {
	pptsuo.mutation.SetType(pa)
	return pptsuo
}

// SetSkill sets the "skill" field.
func (pptsuo *PaPcTechnicalScoreUpdateOne) SetSkill(s string) *PaPcTechnicalScoreUpdateOne {
	pptsuo.mutation.SetSkill(s)
	return pptsuo
}

// SetSelfScore sets the "self_score" field.
func (pptsuo *PaPcTechnicalScoreUpdateOne) SetSelfScore(f float32) *PaPcTechnicalScoreUpdateOne {
	pptsuo.mutation.ResetSelfScore()
	pptsuo.mutation.SetSelfScore(f)
	return pptsuo
}

// AddSelfScore adds f to the "self_score" field.
func (pptsuo *PaPcTechnicalScoreUpdateOne) AddSelfScore(f float32) *PaPcTechnicalScoreUpdateOne {
	pptsuo.mutation.AddSelfScore(f)
	return pptsuo
}

// SetPeerScore sets the "peer_score" field.
func (pptsuo *PaPcTechnicalScoreUpdateOne) SetPeerScore(f float32) *PaPcTechnicalScoreUpdateOne {
	pptsuo.mutation.ResetPeerScore()
	pptsuo.mutation.SetPeerScore(f)
	return pptsuo
}

// SetNillablePeerScore sets the "peer_score" field if the given value is not nil.
func (pptsuo *PaPcTechnicalScoreUpdateOne) SetNillablePeerScore(f *float32) *PaPcTechnicalScoreUpdateOne {
	if f != nil {
		pptsuo.SetPeerScore(*f)
	}
	return pptsuo
}

// AddPeerScore adds f to the "peer_score" field.
func (pptsuo *PaPcTechnicalScoreUpdateOne) AddPeerScore(f float32) *PaPcTechnicalScoreUpdateOne {
	pptsuo.mutation.AddPeerScore(f)
	return pptsuo
}

// ClearPeerScore clears the value of the "peer_score" field.
func (pptsuo *PaPcTechnicalScoreUpdateOne) ClearPeerScore() *PaPcTechnicalScoreUpdateOne {
	pptsuo.mutation.ClearPeerScore()
	return pptsuo
}

// SetManagerScore sets the "manager_score" field.
func (pptsuo *PaPcTechnicalScoreUpdateOne) SetManagerScore(f float32) *PaPcTechnicalScoreUpdateOne {
	pptsuo.mutation.ResetManagerScore()
	pptsuo.mutation.SetManagerScore(f)
	return pptsuo
}

// AddManagerScore adds f to the "manager_score" field.
func (pptsuo *PaPcTechnicalScoreUpdateOne) AddManagerScore(f float32) *PaPcTechnicalScoreUpdateOne {
	pptsuo.mutation.AddManagerScore(f)
	return pptsuo
}

// SetPaPc sets the "pa_pc" edge to the PaPc entity.
func (pptsuo *PaPcTechnicalScoreUpdateOne) SetPaPc(p *PaPc) *PaPcTechnicalScoreUpdateOne {
	return pptsuo.SetPaPcID(p.ID)
}

// Mutation returns the PaPcTechnicalScoreMutation object of the builder.
func (pptsuo *PaPcTechnicalScoreUpdateOne) Mutation() *PaPcTechnicalScoreMutation {
	return pptsuo.mutation
}

// ClearPaPc clears the "pa_pc" edge to the PaPc entity.
func (pptsuo *PaPcTechnicalScoreUpdateOne) ClearPaPc() *PaPcTechnicalScoreUpdateOne {
	pptsuo.mutation.ClearPaPc()
	return pptsuo
}

// Where appends a list predicates to the PaPcTechnicalScoreUpdate builder.
func (pptsuo *PaPcTechnicalScoreUpdateOne) Where(ps ...predicate.PaPcTechnicalScore) *PaPcTechnicalScoreUpdateOne {
	pptsuo.mutation.Where(ps...)
	return pptsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pptsuo *PaPcTechnicalScoreUpdateOne) Select(field string, fields ...string) *PaPcTechnicalScoreUpdateOne {
	pptsuo.fields = append([]string{field}, fields...)
	return pptsuo
}

// Save executes the query and returns the updated PaPcTechnicalScore entity.
func (pptsuo *PaPcTechnicalScoreUpdateOne) Save(ctx context.Context) (*PaPcTechnicalScore, error) {
	return withHooks(ctx, pptsuo.sqlSave, pptsuo.mutation, pptsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pptsuo *PaPcTechnicalScoreUpdateOne) SaveX(ctx context.Context) *PaPcTechnicalScore {
	node, err := pptsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pptsuo *PaPcTechnicalScoreUpdateOne) Exec(ctx context.Context) error {
	_, err := pptsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pptsuo *PaPcTechnicalScoreUpdateOne) ExecX(ctx context.Context) {
	if err := pptsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pptsuo *PaPcTechnicalScoreUpdateOne) check() error {
	if v, ok := pptsuo.mutation.GetType(); ok {
		if err := papctechnicalscore.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "PaPcTechnicalScore.type": %w`, err)}
		}
	}
	if _, ok := pptsuo.mutation.PaPcID(); pptsuo.mutation.PaPcCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PaPcTechnicalScore.pa_pc"`)
	}
	return nil
}

func (pptsuo *PaPcTechnicalScoreUpdateOne) sqlSave(ctx context.Context) (_node *PaPcTechnicalScore, err error) {
	if err := pptsuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(papctechnicalscore.Table, papctechnicalscore.Columns, sqlgraph.NewFieldSpec(papctechnicalscore.FieldID, field.TypeInt))
	id, ok := pptsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PaPcTechnicalScore.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pptsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, papctechnicalscore.FieldID)
		for _, f := range fields {
			if !papctechnicalscore.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != papctechnicalscore.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pptsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pptsuo.mutation.GetType(); ok {
		_spec.SetField(papctechnicalscore.FieldType, field.TypeEnum, value)
	}
	if value, ok := pptsuo.mutation.Skill(); ok {
		_spec.SetField(papctechnicalscore.FieldSkill, field.TypeString, value)
	}
	if value, ok := pptsuo.mutation.SelfScore(); ok {
		_spec.SetField(papctechnicalscore.FieldSelfScore, field.TypeFloat32, value)
	}
	if value, ok := pptsuo.mutation.AddedSelfScore(); ok {
		_spec.AddField(papctechnicalscore.FieldSelfScore, field.TypeFloat32, value)
	}
	if value, ok := pptsuo.mutation.PeerScore(); ok {
		_spec.SetField(papctechnicalscore.FieldPeerScore, field.TypeFloat32, value)
	}
	if value, ok := pptsuo.mutation.AddedPeerScore(); ok {
		_spec.AddField(papctechnicalscore.FieldPeerScore, field.TypeFloat32, value)
	}
	if pptsuo.mutation.PeerScoreCleared() {
		_spec.ClearField(papctechnicalscore.FieldPeerScore, field.TypeFloat32)
	}
	if value, ok := pptsuo.mutation.ManagerScore(); ok {
		_spec.SetField(papctechnicalscore.FieldManagerScore, field.TypeFloat32, value)
	}
	if value, ok := pptsuo.mutation.AddedManagerScore(); ok {
		_spec.AddField(papctechnicalscore.FieldManagerScore, field.TypeFloat32, value)
	}
	if pptsuo.mutation.PaPcCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   papctechnicalscore.PaPcTable,
			Columns: []string{papctechnicalscore.PaPcColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(papc.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pptsuo.mutation.PaPcIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   papctechnicalscore.PaPcTable,
			Columns: []string{papctechnicalscore.PaPcColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(papc.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PaPcTechnicalScore{config: pptsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pptsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{papctechnicalscore.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pptsuo.mutation.done = true
	return _node, nil
}
