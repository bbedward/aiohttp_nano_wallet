// Code generated by ent, DO NOT EDIT.

package block

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/appditto/pippin_nano_wallet/libs/database/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AccountID applies equality check predicate on the "account_id" field. It's identical to AccountIDEQ.
func AccountID(v uuid.UUID) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountID), v))
	})
}

// AdhocAccountID applies equality check predicate on the "adhoc_account_id" field. It's identical to AdhocAccountIDEQ.
func AdhocAccountID(v uuid.UUID) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdhocAccountID), v))
	})
}

// BlockHash applies equality check predicate on the "block_hash" field. It's identical to BlockHashEQ.
func BlockHash(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBlockHash), v))
	})
}

// SendID applies equality check predicate on the "send_id" field. It's identical to SendIDEQ.
func SendID(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSendID), v))
	})
}

// Subtype applies equality check predicate on the "subtype" field. It's identical to SubtypeEQ.
func Subtype(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubtype), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// AccountIDEQ applies the EQ predicate on the "account_id" field.
func AccountIDEQ(v uuid.UUID) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountID), v))
	})
}

// AccountIDNEQ applies the NEQ predicate on the "account_id" field.
func AccountIDNEQ(v uuid.UUID) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccountID), v))
	})
}

// AccountIDIn applies the In predicate on the "account_id" field.
func AccountIDIn(vs ...uuid.UUID) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAccountID), v...))
	})
}

// AccountIDNotIn applies the NotIn predicate on the "account_id" field.
func AccountIDNotIn(vs ...uuid.UUID) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAccountID), v...))
	})
}

// AccountIDIsNil applies the IsNil predicate on the "account_id" field.
func AccountIDIsNil() predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAccountID)))
	})
}

// AccountIDNotNil applies the NotNil predicate on the "account_id" field.
func AccountIDNotNil() predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAccountID)))
	})
}

// AdhocAccountIDEQ applies the EQ predicate on the "adhoc_account_id" field.
func AdhocAccountIDEQ(v uuid.UUID) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdhocAccountID), v))
	})
}

// AdhocAccountIDNEQ applies the NEQ predicate on the "adhoc_account_id" field.
func AdhocAccountIDNEQ(v uuid.UUID) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdhocAccountID), v))
	})
}

// AdhocAccountIDIn applies the In predicate on the "adhoc_account_id" field.
func AdhocAccountIDIn(vs ...uuid.UUID) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAdhocAccountID), v...))
	})
}

// AdhocAccountIDNotIn applies the NotIn predicate on the "adhoc_account_id" field.
func AdhocAccountIDNotIn(vs ...uuid.UUID) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAdhocAccountID), v...))
	})
}

// AdhocAccountIDIsNil applies the IsNil predicate on the "adhoc_account_id" field.
func AdhocAccountIDIsNil() predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAdhocAccountID)))
	})
}

// AdhocAccountIDNotNil applies the NotNil predicate on the "adhoc_account_id" field.
func AdhocAccountIDNotNil() predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAdhocAccountID)))
	})
}

// BlockHashEQ applies the EQ predicate on the "block_hash" field.
func BlockHashEQ(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBlockHash), v))
	})
}

// BlockHashNEQ applies the NEQ predicate on the "block_hash" field.
func BlockHashNEQ(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBlockHash), v))
	})
}

// BlockHashIn applies the In predicate on the "block_hash" field.
func BlockHashIn(vs ...string) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBlockHash), v...))
	})
}

// BlockHashNotIn applies the NotIn predicate on the "block_hash" field.
func BlockHashNotIn(vs ...string) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBlockHash), v...))
	})
}

// BlockHashGT applies the GT predicate on the "block_hash" field.
func BlockHashGT(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBlockHash), v))
	})
}

// BlockHashGTE applies the GTE predicate on the "block_hash" field.
func BlockHashGTE(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBlockHash), v))
	})
}

// BlockHashLT applies the LT predicate on the "block_hash" field.
func BlockHashLT(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBlockHash), v))
	})
}

// BlockHashLTE applies the LTE predicate on the "block_hash" field.
func BlockHashLTE(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBlockHash), v))
	})
}

// BlockHashContains applies the Contains predicate on the "block_hash" field.
func BlockHashContains(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBlockHash), v))
	})
}

// BlockHashHasPrefix applies the HasPrefix predicate on the "block_hash" field.
func BlockHashHasPrefix(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBlockHash), v))
	})
}

// BlockHashHasSuffix applies the HasSuffix predicate on the "block_hash" field.
func BlockHashHasSuffix(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBlockHash), v))
	})
}

// BlockHashEqualFold applies the EqualFold predicate on the "block_hash" field.
func BlockHashEqualFold(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBlockHash), v))
	})
}

// BlockHashContainsFold applies the ContainsFold predicate on the "block_hash" field.
func BlockHashContainsFold(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBlockHash), v))
	})
}

// SendIDEQ applies the EQ predicate on the "send_id" field.
func SendIDEQ(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSendID), v))
	})
}

// SendIDNEQ applies the NEQ predicate on the "send_id" field.
func SendIDNEQ(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSendID), v))
	})
}

// SendIDIn applies the In predicate on the "send_id" field.
func SendIDIn(vs ...string) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSendID), v...))
	})
}

// SendIDNotIn applies the NotIn predicate on the "send_id" field.
func SendIDNotIn(vs ...string) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSendID), v...))
	})
}

// SendIDGT applies the GT predicate on the "send_id" field.
func SendIDGT(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSendID), v))
	})
}

// SendIDGTE applies the GTE predicate on the "send_id" field.
func SendIDGTE(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSendID), v))
	})
}

// SendIDLT applies the LT predicate on the "send_id" field.
func SendIDLT(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSendID), v))
	})
}

// SendIDLTE applies the LTE predicate on the "send_id" field.
func SendIDLTE(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSendID), v))
	})
}

// SendIDContains applies the Contains predicate on the "send_id" field.
func SendIDContains(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSendID), v))
	})
}

// SendIDHasPrefix applies the HasPrefix predicate on the "send_id" field.
func SendIDHasPrefix(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSendID), v))
	})
}

// SendIDHasSuffix applies the HasSuffix predicate on the "send_id" field.
func SendIDHasSuffix(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSendID), v))
	})
}

// SendIDIsNil applies the IsNil predicate on the "send_id" field.
func SendIDIsNil() predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSendID)))
	})
}

// SendIDNotNil applies the NotNil predicate on the "send_id" field.
func SendIDNotNil() predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSendID)))
	})
}

// SendIDEqualFold applies the EqualFold predicate on the "send_id" field.
func SendIDEqualFold(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSendID), v))
	})
}

// SendIDContainsFold applies the ContainsFold predicate on the "send_id" field.
func SendIDContainsFold(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSendID), v))
	})
}

// SubtypeEQ applies the EQ predicate on the "subtype" field.
func SubtypeEQ(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubtype), v))
	})
}

// SubtypeNEQ applies the NEQ predicate on the "subtype" field.
func SubtypeNEQ(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSubtype), v))
	})
}

// SubtypeIn applies the In predicate on the "subtype" field.
func SubtypeIn(vs ...string) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSubtype), v...))
	})
}

// SubtypeNotIn applies the NotIn predicate on the "subtype" field.
func SubtypeNotIn(vs ...string) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSubtype), v...))
	})
}

// SubtypeGT applies the GT predicate on the "subtype" field.
func SubtypeGT(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSubtype), v))
	})
}

// SubtypeGTE applies the GTE predicate on the "subtype" field.
func SubtypeGTE(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSubtype), v))
	})
}

// SubtypeLT applies the LT predicate on the "subtype" field.
func SubtypeLT(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSubtype), v))
	})
}

// SubtypeLTE applies the LTE predicate on the "subtype" field.
func SubtypeLTE(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSubtype), v))
	})
}

// SubtypeContains applies the Contains predicate on the "subtype" field.
func SubtypeContains(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSubtype), v))
	})
}

// SubtypeHasPrefix applies the HasPrefix predicate on the "subtype" field.
func SubtypeHasPrefix(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSubtype), v))
	})
}

// SubtypeHasSuffix applies the HasSuffix predicate on the "subtype" field.
func SubtypeHasSuffix(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSubtype), v))
	})
}

// SubtypeEqualFold applies the EqualFold predicate on the "subtype" field.
func SubtypeEqualFold(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSubtype), v))
	})
}

// SubtypeContainsFold applies the ContainsFold predicate on the "subtype" field.
func SubtypeContainsFold(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSubtype), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasAccount applies the HasEdge predicate on the "account" edge.
func HasAccount() predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AccountTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountWith applies the HasEdge predicate on the "account" edge with a given conditions (other predicates).
func HasAccountWith(preds ...predicate.Account) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AccountInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAdhocAccount applies the HasEdge predicate on the "adhoc_account" edge.
func HasAdhocAccount() predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AdhocAccountTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AdhocAccountTable, AdhocAccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAdhocAccountWith applies the HasEdge predicate on the "adhoc_account" edge with a given conditions (other predicates).
func HasAdhocAccountWith(preds ...predicate.AdhocAccount) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AdhocAccountInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AdhocAccountTable, AdhocAccountColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Block) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Block) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Block) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		p(s.Not())
	})
}
