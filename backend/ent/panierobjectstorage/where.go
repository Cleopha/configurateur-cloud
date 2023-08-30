// Code generated by ent, DO NOT EDIT.

package panierobjectstorage

import (
	"backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldLTE(FieldID, id))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldEQ(FieldQuantity, v))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(sql.FieldLTE(FieldQuantity, v))
}

// HasPanier applies the HasEdge predicate on the "panier" edge.
func HasPanier() predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PanierTable, PanierColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPanierWith applies the HasEdge predicate on the "panier" edge with a given conditions (other predicates).
func HasPanierWith(preds ...predicate.Panier) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(func(s *sql.Selector) {
		step := newPanierStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasObjectStorage applies the HasEdge predicate on the "objectStorage" edge.
func HasObjectStorage() predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ObjectStorageTable, ObjectStorageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasObjectStorageWith applies the HasEdge predicate on the "objectStorage" edge with a given conditions (other predicates).
func HasObjectStorageWith(preds ...predicate.ObjectStorage) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(func(s *sql.Selector) {
		step := newObjectStorageStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PanierObjectStorage) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PanierObjectStorage) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(func(s *sql.Selector) {
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
func Not(p predicate.PanierObjectStorage) predicate.PanierObjectStorage {
	return predicate.PanierObjectStorage(func(s *sql.Selector) {
		p(s.Not())
	})
}