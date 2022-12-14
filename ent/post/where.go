// Code generated by ent, DO NOT EDIT.

package post

import (
	"moments/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedDate applies equality check predicate on the "created_date" field. It's identical to CreatedDateEQ.
func CreatedDate(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedDate), v))
	})
}

// UpdatedDate applies equality check predicate on the "updated_date" field. It's identical to UpdatedDateEQ.
func UpdatedDate(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedDate), v))
	})
}

// DeletedDate applies equality check predicate on the "deleted_date" field. It's identical to DeletedDateEQ.
func DeletedDate(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedDate), v))
	})
}

// PlainText applies equality check predicate on the "plain_text" field. It's identical to PlainTextEQ.
func PlainText(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlainText), v))
	})
}

// Likes applies equality check predicate on the "likes" field. It's identical to LikesEQ.
func Likes(v uint64) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLikes), v))
	})
}

// LinkURL applies equality check predicate on the "link_url" field. It's identical to LinkURLEQ.
func LinkURL(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLinkURL), v))
	})
}

// CreatedDateEQ applies the EQ predicate on the "created_date" field.
func CreatedDateEQ(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedDate), v))
	})
}

// CreatedDateNEQ applies the NEQ predicate on the "created_date" field.
func CreatedDateNEQ(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedDate), v))
	})
}

// CreatedDateIn applies the In predicate on the "created_date" field.
func CreatedDateIn(vs ...time.Time) predicate.Post {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedDate), v...))
	})
}

// CreatedDateNotIn applies the NotIn predicate on the "created_date" field.
func CreatedDateNotIn(vs ...time.Time) predicate.Post {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedDate), v...))
	})
}

// CreatedDateGT applies the GT predicate on the "created_date" field.
func CreatedDateGT(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedDate), v))
	})
}

// CreatedDateGTE applies the GTE predicate on the "created_date" field.
func CreatedDateGTE(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedDate), v))
	})
}

// CreatedDateLT applies the LT predicate on the "created_date" field.
func CreatedDateLT(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedDate), v))
	})
}

// CreatedDateLTE applies the LTE predicate on the "created_date" field.
func CreatedDateLTE(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedDate), v))
	})
}

// UpdatedDateEQ applies the EQ predicate on the "updated_date" field.
func UpdatedDateEQ(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedDate), v))
	})
}

// UpdatedDateNEQ applies the NEQ predicate on the "updated_date" field.
func UpdatedDateNEQ(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedDate), v))
	})
}

// UpdatedDateIn applies the In predicate on the "updated_date" field.
func UpdatedDateIn(vs ...time.Time) predicate.Post {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedDate), v...))
	})
}

// UpdatedDateNotIn applies the NotIn predicate on the "updated_date" field.
func UpdatedDateNotIn(vs ...time.Time) predicate.Post {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedDate), v...))
	})
}

// UpdatedDateGT applies the GT predicate on the "updated_date" field.
func UpdatedDateGT(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedDate), v))
	})
}

// UpdatedDateGTE applies the GTE predicate on the "updated_date" field.
func UpdatedDateGTE(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedDate), v))
	})
}

// UpdatedDateLT applies the LT predicate on the "updated_date" field.
func UpdatedDateLT(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedDate), v))
	})
}

// UpdatedDateLTE applies the LTE predicate on the "updated_date" field.
func UpdatedDateLTE(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedDate), v))
	})
}

// DeletedDateEQ applies the EQ predicate on the "deleted_date" field.
func DeletedDateEQ(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedDate), v))
	})
}

// DeletedDateNEQ applies the NEQ predicate on the "deleted_date" field.
func DeletedDateNEQ(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedDate), v))
	})
}

// DeletedDateIn applies the In predicate on the "deleted_date" field.
func DeletedDateIn(vs ...time.Time) predicate.Post {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedDate), v...))
	})
}

// DeletedDateNotIn applies the NotIn predicate on the "deleted_date" field.
func DeletedDateNotIn(vs ...time.Time) predicate.Post {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedDate), v...))
	})
}

// DeletedDateGT applies the GT predicate on the "deleted_date" field.
func DeletedDateGT(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedDate), v))
	})
}

// DeletedDateGTE applies the GTE predicate on the "deleted_date" field.
func DeletedDateGTE(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedDate), v))
	})
}

// DeletedDateLT applies the LT predicate on the "deleted_date" field.
func DeletedDateLT(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedDate), v))
	})
}

// DeletedDateLTE applies the LTE predicate on the "deleted_date" field.
func DeletedDateLTE(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedDate), v))
	})
}

// DeletedDateIsNil applies the IsNil predicate on the "deleted_date" field.
func DeletedDateIsNil() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedDate)))
	})
}

// DeletedDateNotNil applies the NotNil predicate on the "deleted_date" field.
func DeletedDateNotNil() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedDate)))
	})
}

// PlainTextEQ applies the EQ predicate on the "plain_text" field.
func PlainTextEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlainText), v))
	})
}

// PlainTextNEQ applies the NEQ predicate on the "plain_text" field.
func PlainTextNEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPlainText), v))
	})
}

// PlainTextIn applies the In predicate on the "plain_text" field.
func PlainTextIn(vs ...string) predicate.Post {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPlainText), v...))
	})
}

// PlainTextNotIn applies the NotIn predicate on the "plain_text" field.
func PlainTextNotIn(vs ...string) predicate.Post {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPlainText), v...))
	})
}

// PlainTextGT applies the GT predicate on the "plain_text" field.
func PlainTextGT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPlainText), v))
	})
}

// PlainTextGTE applies the GTE predicate on the "plain_text" field.
func PlainTextGTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPlainText), v))
	})
}

// PlainTextLT applies the LT predicate on the "plain_text" field.
func PlainTextLT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPlainText), v))
	})
}

// PlainTextLTE applies the LTE predicate on the "plain_text" field.
func PlainTextLTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPlainText), v))
	})
}

// PlainTextContains applies the Contains predicate on the "plain_text" field.
func PlainTextContains(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPlainText), v))
	})
}

// PlainTextHasPrefix applies the HasPrefix predicate on the "plain_text" field.
func PlainTextHasPrefix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPlainText), v))
	})
}

// PlainTextHasSuffix applies the HasSuffix predicate on the "plain_text" field.
func PlainTextHasSuffix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPlainText), v))
	})
}

// PlainTextEqualFold applies the EqualFold predicate on the "plain_text" field.
func PlainTextEqualFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPlainText), v))
	})
}

// PlainTextContainsFold applies the ContainsFold predicate on the "plain_text" field.
func PlainTextContainsFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPlainText), v))
	})
}

// LikesEQ applies the EQ predicate on the "likes" field.
func LikesEQ(v uint64) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLikes), v))
	})
}

// LikesNEQ applies the NEQ predicate on the "likes" field.
func LikesNEQ(v uint64) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLikes), v))
	})
}

// LikesIn applies the In predicate on the "likes" field.
func LikesIn(vs ...uint64) predicate.Post {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLikes), v...))
	})
}

// LikesNotIn applies the NotIn predicate on the "likes" field.
func LikesNotIn(vs ...uint64) predicate.Post {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLikes), v...))
	})
}

// LikesGT applies the GT predicate on the "likes" field.
func LikesGT(v uint64) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLikes), v))
	})
}

// LikesGTE applies the GTE predicate on the "likes" field.
func LikesGTE(v uint64) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLikes), v))
	})
}

// LikesLT applies the LT predicate on the "likes" field.
func LikesLT(v uint64) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLikes), v))
	})
}

// LikesLTE applies the LTE predicate on the "likes" field.
func LikesLTE(v uint64) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLikes), v))
	})
}

// LikesIsNil applies the IsNil predicate on the "likes" field.
func LikesIsNil() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLikes)))
	})
}

// LikesNotNil applies the NotNil predicate on the "likes" field.
func LikesNotNil() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLikes)))
	})
}

// LinkURLEQ applies the EQ predicate on the "link_url" field.
func LinkURLEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLinkURL), v))
	})
}

// LinkURLNEQ applies the NEQ predicate on the "link_url" field.
func LinkURLNEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLinkURL), v))
	})
}

// LinkURLIn applies the In predicate on the "link_url" field.
func LinkURLIn(vs ...string) predicate.Post {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLinkURL), v...))
	})
}

// LinkURLNotIn applies the NotIn predicate on the "link_url" field.
func LinkURLNotIn(vs ...string) predicate.Post {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLinkURL), v...))
	})
}

// LinkURLGT applies the GT predicate on the "link_url" field.
func LinkURLGT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLinkURL), v))
	})
}

// LinkURLGTE applies the GTE predicate on the "link_url" field.
func LinkURLGTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLinkURL), v))
	})
}

// LinkURLLT applies the LT predicate on the "link_url" field.
func LinkURLLT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLinkURL), v))
	})
}

// LinkURLLTE applies the LTE predicate on the "link_url" field.
func LinkURLLTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLinkURL), v))
	})
}

// LinkURLContains applies the Contains predicate on the "link_url" field.
func LinkURLContains(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLinkURL), v))
	})
}

// LinkURLHasPrefix applies the HasPrefix predicate on the "link_url" field.
func LinkURLHasPrefix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLinkURL), v))
	})
}

// LinkURLHasSuffix applies the HasSuffix predicate on the "link_url" field.
func LinkURLHasSuffix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLinkURL), v))
	})
}

// LinkURLIsNil applies the IsNil predicate on the "link_url" field.
func LinkURLIsNil() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLinkURL)))
	})
}

// LinkURLNotNil applies the NotNil predicate on the "link_url" field.
func LinkURLNotNil() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLinkURL)))
	})
}

// LinkURLEqualFold applies the EqualFold predicate on the "link_url" field.
func LinkURLEqualFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLinkURL), v))
	})
}

// LinkURLContainsFold applies the ContainsFold predicate on the "link_url" field.
func LinkURLContainsFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLinkURL), v))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
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
func Not(p predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		p(s.Not())
	})
}
