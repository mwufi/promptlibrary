// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"prompt-library/backend/ent/conversation"
	"prompt-library/backend/ent/predicate"
	"prompt-library/backend/ent/prompt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PromptUpdate is the builder for updating Prompt entities.
type PromptUpdate struct {
	config
	hooks    []Hook
	mutation *PromptMutation
}

// Where appends a list predicates to the PromptUpdate builder.
func (pu *PromptUpdate) Where(ps ...predicate.Prompt) *PromptUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetTitle sets the "title" field.
func (pu *PromptUpdate) SetTitle(s string) *PromptUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (pu *PromptUpdate) SetNillableTitle(s *string) *PromptUpdate {
	if s != nil {
		pu.SetTitle(*s)
	}
	return pu
}

// SetContent sets the "content" field.
func (pu *PromptUpdate) SetContent(s string) *PromptUpdate {
	pu.mutation.SetContent(s)
	return pu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (pu *PromptUpdate) SetNillableContent(s *string) *PromptUpdate {
	if s != nil {
		pu.SetContent(*s)
	}
	return pu
}

// SetCategory sets the "category" field.
func (pu *PromptUpdate) SetCategory(s string) *PromptUpdate {
	pu.mutation.SetCategory(s)
	return pu
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (pu *PromptUpdate) SetNillableCategory(s *string) *PromptUpdate {
	if s != nil {
		pu.SetCategory(*s)
	}
	return pu
}

// SetVotes sets the "votes" field.
func (pu *PromptUpdate) SetVotes(i int) *PromptUpdate {
	pu.mutation.ResetVotes()
	pu.mutation.SetVotes(i)
	return pu
}

// SetNillableVotes sets the "votes" field if the given value is not nil.
func (pu *PromptUpdate) SetNillableVotes(i *int) *PromptUpdate {
	if i != nil {
		pu.SetVotes(*i)
	}
	return pu
}

// AddVotes adds i to the "votes" field.
func (pu *PromptUpdate) AddVotes(i int) *PromptUpdate {
	pu.mutation.AddVotes(i)
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PromptUpdate) SetCreatedAt(t time.Time) *PromptUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PromptUpdate) SetNillableCreatedAt(t *time.Time) *PromptUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetCreatedBy sets the "created_by" field.
func (pu *PromptUpdate) SetCreatedBy(s string) *PromptUpdate {
	pu.mutation.SetCreatedBy(s)
	return pu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pu *PromptUpdate) SetNillableCreatedBy(s *string) *PromptUpdate {
	if s != nil {
		pu.SetCreatedBy(*s)
	}
	return pu
}

// AddConversationIDs adds the "conversations" edge to the Conversation entity by IDs.
func (pu *PromptUpdate) AddConversationIDs(ids ...int) *PromptUpdate {
	pu.mutation.AddConversationIDs(ids...)
	return pu
}

// AddConversations adds the "conversations" edges to the Conversation entity.
func (pu *PromptUpdate) AddConversations(c ...*Conversation) *PromptUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.AddConversationIDs(ids...)
}

// Mutation returns the PromptMutation object of the builder.
func (pu *PromptUpdate) Mutation() *PromptMutation {
	return pu.mutation
}

// ClearConversations clears all "conversations" edges to the Conversation entity.
func (pu *PromptUpdate) ClearConversations() *PromptUpdate {
	pu.mutation.ClearConversations()
	return pu
}

// RemoveConversationIDs removes the "conversations" edge to Conversation entities by IDs.
func (pu *PromptUpdate) RemoveConversationIDs(ids ...int) *PromptUpdate {
	pu.mutation.RemoveConversationIDs(ids...)
	return pu
}

// RemoveConversations removes "conversations" edges to Conversation entities.
func (pu *PromptUpdate) RemoveConversations(c ...*Conversation) *PromptUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.RemoveConversationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PromptUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PromptUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PromptUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PromptUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PromptUpdate) check() error {
	if v, ok := pu.mutation.Title(); ok {
		if err := prompt.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Prompt.title": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Content(); ok {
		if err := prompt.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Prompt.content": %w`, err)}
		}
	}
	return nil
}

func (pu *PromptUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(prompt.Table, prompt.Columns, sqlgraph.NewFieldSpec(prompt.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.SetField(prompt.FieldTitle, field.TypeString, value)
	}
	if value, ok := pu.mutation.Content(); ok {
		_spec.SetField(prompt.FieldContent, field.TypeString, value)
	}
	if value, ok := pu.mutation.Category(); ok {
		_spec.SetField(prompt.FieldCategory, field.TypeString, value)
	}
	if value, ok := pu.mutation.Votes(); ok {
		_spec.SetField(prompt.FieldVotes, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedVotes(); ok {
		_spec.AddField(prompt.FieldVotes, field.TypeInt, value)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(prompt.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.CreatedBy(); ok {
		_spec.SetField(prompt.FieldCreatedBy, field.TypeString, value)
	}
	if pu.mutation.ConversationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prompt.ConversationsTable,
			Columns: []string{prompt.ConversationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(conversation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedConversationsIDs(); len(nodes) > 0 && !pu.mutation.ConversationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prompt.ConversationsTable,
			Columns: []string{prompt.ConversationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(conversation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ConversationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prompt.ConversationsTable,
			Columns: []string{prompt.ConversationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(conversation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prompt.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PromptUpdateOne is the builder for updating a single Prompt entity.
type PromptUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PromptMutation
}

// SetTitle sets the "title" field.
func (puo *PromptUpdateOne) SetTitle(s string) *PromptUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (puo *PromptUpdateOne) SetNillableTitle(s *string) *PromptUpdateOne {
	if s != nil {
		puo.SetTitle(*s)
	}
	return puo
}

// SetContent sets the "content" field.
func (puo *PromptUpdateOne) SetContent(s string) *PromptUpdateOne {
	puo.mutation.SetContent(s)
	return puo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (puo *PromptUpdateOne) SetNillableContent(s *string) *PromptUpdateOne {
	if s != nil {
		puo.SetContent(*s)
	}
	return puo
}

// SetCategory sets the "category" field.
func (puo *PromptUpdateOne) SetCategory(s string) *PromptUpdateOne {
	puo.mutation.SetCategory(s)
	return puo
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (puo *PromptUpdateOne) SetNillableCategory(s *string) *PromptUpdateOne {
	if s != nil {
		puo.SetCategory(*s)
	}
	return puo
}

// SetVotes sets the "votes" field.
func (puo *PromptUpdateOne) SetVotes(i int) *PromptUpdateOne {
	puo.mutation.ResetVotes()
	puo.mutation.SetVotes(i)
	return puo
}

// SetNillableVotes sets the "votes" field if the given value is not nil.
func (puo *PromptUpdateOne) SetNillableVotes(i *int) *PromptUpdateOne {
	if i != nil {
		puo.SetVotes(*i)
	}
	return puo
}

// AddVotes adds i to the "votes" field.
func (puo *PromptUpdateOne) AddVotes(i int) *PromptUpdateOne {
	puo.mutation.AddVotes(i)
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *PromptUpdateOne) SetCreatedAt(t time.Time) *PromptUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PromptUpdateOne) SetNillableCreatedAt(t *time.Time) *PromptUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetCreatedBy sets the "created_by" field.
func (puo *PromptUpdateOne) SetCreatedBy(s string) *PromptUpdateOne {
	puo.mutation.SetCreatedBy(s)
	return puo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (puo *PromptUpdateOne) SetNillableCreatedBy(s *string) *PromptUpdateOne {
	if s != nil {
		puo.SetCreatedBy(*s)
	}
	return puo
}

// AddConversationIDs adds the "conversations" edge to the Conversation entity by IDs.
func (puo *PromptUpdateOne) AddConversationIDs(ids ...int) *PromptUpdateOne {
	puo.mutation.AddConversationIDs(ids...)
	return puo
}

// AddConversations adds the "conversations" edges to the Conversation entity.
func (puo *PromptUpdateOne) AddConversations(c ...*Conversation) *PromptUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.AddConversationIDs(ids...)
}

// Mutation returns the PromptMutation object of the builder.
func (puo *PromptUpdateOne) Mutation() *PromptMutation {
	return puo.mutation
}

// ClearConversations clears all "conversations" edges to the Conversation entity.
func (puo *PromptUpdateOne) ClearConversations() *PromptUpdateOne {
	puo.mutation.ClearConversations()
	return puo
}

// RemoveConversationIDs removes the "conversations" edge to Conversation entities by IDs.
func (puo *PromptUpdateOne) RemoveConversationIDs(ids ...int) *PromptUpdateOne {
	puo.mutation.RemoveConversationIDs(ids...)
	return puo
}

// RemoveConversations removes "conversations" edges to Conversation entities.
func (puo *PromptUpdateOne) RemoveConversations(c ...*Conversation) *PromptUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.RemoveConversationIDs(ids...)
}

// Where appends a list predicates to the PromptUpdate builder.
func (puo *PromptUpdateOne) Where(ps ...predicate.Prompt) *PromptUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PromptUpdateOne) Select(field string, fields ...string) *PromptUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Prompt entity.
func (puo *PromptUpdateOne) Save(ctx context.Context) (*Prompt, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PromptUpdateOne) SaveX(ctx context.Context) *Prompt {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PromptUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PromptUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PromptUpdateOne) check() error {
	if v, ok := puo.mutation.Title(); ok {
		if err := prompt.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Prompt.title": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Content(); ok {
		if err := prompt.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Prompt.content": %w`, err)}
		}
	}
	return nil
}

func (puo *PromptUpdateOne) sqlSave(ctx context.Context) (_node *Prompt, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(prompt.Table, prompt.Columns, sqlgraph.NewFieldSpec(prompt.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Prompt.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, prompt.FieldID)
		for _, f := range fields {
			if !prompt.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != prompt.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Title(); ok {
		_spec.SetField(prompt.FieldTitle, field.TypeString, value)
	}
	if value, ok := puo.mutation.Content(); ok {
		_spec.SetField(prompt.FieldContent, field.TypeString, value)
	}
	if value, ok := puo.mutation.Category(); ok {
		_spec.SetField(prompt.FieldCategory, field.TypeString, value)
	}
	if value, ok := puo.mutation.Votes(); ok {
		_spec.SetField(prompt.FieldVotes, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedVotes(); ok {
		_spec.AddField(prompt.FieldVotes, field.TypeInt, value)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(prompt.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.CreatedBy(); ok {
		_spec.SetField(prompt.FieldCreatedBy, field.TypeString, value)
	}
	if puo.mutation.ConversationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prompt.ConversationsTable,
			Columns: []string{prompt.ConversationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(conversation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedConversationsIDs(); len(nodes) > 0 && !puo.mutation.ConversationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prompt.ConversationsTable,
			Columns: []string{prompt.ConversationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(conversation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ConversationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prompt.ConversationsTable,
			Columns: []string{prompt.ConversationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(conversation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Prompt{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prompt.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
