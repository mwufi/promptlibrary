// Code generated by ent, DO NOT EDIT.

package prompt

import (
	"prompt-library/backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Prompt {
	return predicate.Prompt(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Prompt {
	return predicate.Prompt(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Prompt {
	return predicate.Prompt(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Prompt {
	return predicate.Prompt(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Prompt {
	return predicate.Prompt(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Prompt {
	return predicate.Prompt(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Prompt {
	return predicate.Prompt(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldTitle, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldContent, v))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldCategory, v))
}

// Votes applies equality check predicate on the "votes" field. It's identical to VotesEQ.
func Votes(v int) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldVotes, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldCreatedBy, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Prompt {
	return predicate.Prompt(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Prompt {
	return predicate.Prompt(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldContainsFold(FieldTitle, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Prompt {
	return predicate.Prompt(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Prompt {
	return predicate.Prompt(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldContainsFold(FieldContent, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.Prompt {
	return predicate.Prompt(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.Prompt {
	return predicate.Prompt(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldLTE(FieldCategory, v))
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldContains(FieldCategory, v))
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldHasPrefix(FieldCategory, v))
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldHasSuffix(FieldCategory, v))
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEqualFold(FieldCategory, v))
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldContainsFold(FieldCategory, v))
}

// VotesEQ applies the EQ predicate on the "votes" field.
func VotesEQ(v int) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldVotes, v))
}

// VotesNEQ applies the NEQ predicate on the "votes" field.
func VotesNEQ(v int) predicate.Prompt {
	return predicate.Prompt(sql.FieldNEQ(FieldVotes, v))
}

// VotesIn applies the In predicate on the "votes" field.
func VotesIn(vs ...int) predicate.Prompt {
	return predicate.Prompt(sql.FieldIn(FieldVotes, vs...))
}

// VotesNotIn applies the NotIn predicate on the "votes" field.
func VotesNotIn(vs ...int) predicate.Prompt {
	return predicate.Prompt(sql.FieldNotIn(FieldVotes, vs...))
}

// VotesGT applies the GT predicate on the "votes" field.
func VotesGT(v int) predicate.Prompt {
	return predicate.Prompt(sql.FieldGT(FieldVotes, v))
}

// VotesGTE applies the GTE predicate on the "votes" field.
func VotesGTE(v int) predicate.Prompt {
	return predicate.Prompt(sql.FieldGTE(FieldVotes, v))
}

// VotesLT applies the LT predicate on the "votes" field.
func VotesLT(v int) predicate.Prompt {
	return predicate.Prompt(sql.FieldLT(FieldVotes, v))
}

// VotesLTE applies the LTE predicate on the "votes" field.
func VotesLTE(v int) predicate.Prompt {
	return predicate.Prompt(sql.FieldLTE(FieldVotes, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Prompt {
	return predicate.Prompt(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Prompt {
	return predicate.Prompt(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Prompt {
	return predicate.Prompt(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Prompt {
	return predicate.Prompt(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Prompt {
	return predicate.Prompt(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Prompt {
	return predicate.Prompt(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Prompt {
	return predicate.Prompt(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Prompt {
	return predicate.Prompt(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Prompt {
	return predicate.Prompt(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldContainsFold(FieldCreatedBy, v))
}

// HasConversations applies the HasEdge predicate on the "conversations" edge.
func HasConversations() predicate.Prompt {
	return predicate.Prompt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ConversationsTable, ConversationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConversationsWith applies the HasEdge predicate on the "conversations" edge with a given conditions (other predicates).
func HasConversationsWith(preds ...predicate.Conversation) predicate.Prompt {
	return predicate.Prompt(func(s *sql.Selector) {
		step := newConversationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Prompt) predicate.Prompt {
	return predicate.Prompt(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Prompt) predicate.Prompt {
	return predicate.Prompt(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Prompt) predicate.Prompt {
	return predicate.Prompt(sql.NotPredicates(p))
}
