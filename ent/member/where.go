// Code generated by entc, DO NOT EDIT.

package member

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ymsg19/sfc-guardbot/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DiscordUID applies equality check predicate on the "discord_uid" field. It's identical to DiscordUIDEQ.
func DiscordUID(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscordUID), v))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// IsVerified applies equality check predicate on the "is_verified" field. It's identical to IsVerifiedEQ.
func IsVerified(v bool) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsVerified), v))
	})
}

// VerificationExpiry applies equality check predicate on the "verification_expiry" field. It's identical to VerificationExpiryEQ.
func VerificationExpiry(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVerificationExpiry), v))
	})
}

// HashedVerificationToken applies equality check predicate on the "hashed_verification_token" field. It's identical to HashedVerificationTokenEQ.
func HashedVerificationToken(v []byte) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHashedVerificationToken), v))
	})
}

// VerificationTokenExpiry applies equality check predicate on the "verification_token_expiry" field. It's identical to VerificationTokenExpiryEQ.
func VerificationTokenExpiry(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVerificationTokenExpiry), v))
	})
}

// DiscordUIDEQ applies the EQ predicate on the "discord_uid" field.
func DiscordUIDEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscordUID), v))
	})
}

// DiscordUIDNEQ applies the NEQ predicate on the "discord_uid" field.
func DiscordUIDNEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDiscordUID), v))
	})
}

// DiscordUIDIn applies the In predicate on the "discord_uid" field.
func DiscordUIDIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDiscordUID), v...))
	})
}

// DiscordUIDNotIn applies the NotIn predicate on the "discord_uid" field.
func DiscordUIDNotIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDiscordUID), v...))
	})
}

// DiscordUIDGT applies the GT predicate on the "discord_uid" field.
func DiscordUIDGT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDiscordUID), v))
	})
}

// DiscordUIDGTE applies the GTE predicate on the "discord_uid" field.
func DiscordUIDGTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDiscordUID), v))
	})
}

// DiscordUIDLT applies the LT predicate on the "discord_uid" field.
func DiscordUIDLT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDiscordUID), v))
	})
}

// DiscordUIDLTE applies the LTE predicate on the "discord_uid" field.
func DiscordUIDLTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDiscordUID), v))
	})
}

// DiscordUIDContains applies the Contains predicate on the "discord_uid" field.
func DiscordUIDContains(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDiscordUID), v))
	})
}

// DiscordUIDHasPrefix applies the HasPrefix predicate on the "discord_uid" field.
func DiscordUIDHasPrefix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDiscordUID), v))
	})
}

// DiscordUIDHasSuffix applies the HasSuffix predicate on the "discord_uid" field.
func DiscordUIDHasSuffix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDiscordUID), v))
	})
}

// DiscordUIDEqualFold applies the EqualFold predicate on the "discord_uid" field.
func DiscordUIDEqualFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDiscordUID), v))
	})
}

// DiscordUIDContainsFold applies the ContainsFold predicate on the "discord_uid" field.
func DiscordUIDContainsFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDiscordUID), v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// IsVerifiedEQ applies the EQ predicate on the "is_verified" field.
func IsVerifiedEQ(v bool) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsVerified), v))
	})
}

// IsVerifiedNEQ applies the NEQ predicate on the "is_verified" field.
func IsVerifiedNEQ(v bool) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsVerified), v))
	})
}

// VerificationExpiryEQ applies the EQ predicate on the "verification_expiry" field.
func VerificationExpiryEQ(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVerificationExpiry), v))
	})
}

// VerificationExpiryNEQ applies the NEQ predicate on the "verification_expiry" field.
func VerificationExpiryNEQ(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVerificationExpiry), v))
	})
}

// VerificationExpiryIn applies the In predicate on the "verification_expiry" field.
func VerificationExpiryIn(vs ...time.Time) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVerificationExpiry), v...))
	})
}

// VerificationExpiryNotIn applies the NotIn predicate on the "verification_expiry" field.
func VerificationExpiryNotIn(vs ...time.Time) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVerificationExpiry), v...))
	})
}

// VerificationExpiryGT applies the GT predicate on the "verification_expiry" field.
func VerificationExpiryGT(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVerificationExpiry), v))
	})
}

// VerificationExpiryGTE applies the GTE predicate on the "verification_expiry" field.
func VerificationExpiryGTE(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVerificationExpiry), v))
	})
}

// VerificationExpiryLT applies the LT predicate on the "verification_expiry" field.
func VerificationExpiryLT(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVerificationExpiry), v))
	})
}

// VerificationExpiryLTE applies the LTE predicate on the "verification_expiry" field.
func VerificationExpiryLTE(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVerificationExpiry), v))
	})
}

// HashedVerificationTokenEQ applies the EQ predicate on the "hashed_verification_token" field.
func HashedVerificationTokenEQ(v []byte) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHashedVerificationToken), v))
	})
}

// HashedVerificationTokenNEQ applies the NEQ predicate on the "hashed_verification_token" field.
func HashedVerificationTokenNEQ(v []byte) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHashedVerificationToken), v))
	})
}

// HashedVerificationTokenIn applies the In predicate on the "hashed_verification_token" field.
func HashedVerificationTokenIn(vs ...[]byte) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHashedVerificationToken), v...))
	})
}

// HashedVerificationTokenNotIn applies the NotIn predicate on the "hashed_verification_token" field.
func HashedVerificationTokenNotIn(vs ...[]byte) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHashedVerificationToken), v...))
	})
}

// HashedVerificationTokenGT applies the GT predicate on the "hashed_verification_token" field.
func HashedVerificationTokenGT(v []byte) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHashedVerificationToken), v))
	})
}

// HashedVerificationTokenGTE applies the GTE predicate on the "hashed_verification_token" field.
func HashedVerificationTokenGTE(v []byte) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHashedVerificationToken), v))
	})
}

// HashedVerificationTokenLT applies the LT predicate on the "hashed_verification_token" field.
func HashedVerificationTokenLT(v []byte) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHashedVerificationToken), v))
	})
}

// HashedVerificationTokenLTE applies the LTE predicate on the "hashed_verification_token" field.
func HashedVerificationTokenLTE(v []byte) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHashedVerificationToken), v))
	})
}

// VerificationTokenExpiryEQ applies the EQ predicate on the "verification_token_expiry" field.
func VerificationTokenExpiryEQ(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVerificationTokenExpiry), v))
	})
}

// VerificationTokenExpiryNEQ applies the NEQ predicate on the "verification_token_expiry" field.
func VerificationTokenExpiryNEQ(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVerificationTokenExpiry), v))
	})
}

// VerificationTokenExpiryIn applies the In predicate on the "verification_token_expiry" field.
func VerificationTokenExpiryIn(vs ...time.Time) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVerificationTokenExpiry), v...))
	})
}

// VerificationTokenExpiryNotIn applies the NotIn predicate on the "verification_token_expiry" field.
func VerificationTokenExpiryNotIn(vs ...time.Time) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVerificationTokenExpiry), v...))
	})
}

// VerificationTokenExpiryGT applies the GT predicate on the "verification_token_expiry" field.
func VerificationTokenExpiryGT(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVerificationTokenExpiry), v))
	})
}

// VerificationTokenExpiryGTE applies the GTE predicate on the "verification_token_expiry" field.
func VerificationTokenExpiryGTE(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVerificationTokenExpiry), v))
	})
}

// VerificationTokenExpiryLT applies the LT predicate on the "verification_token_expiry" field.
func VerificationTokenExpiryLT(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVerificationTokenExpiry), v))
	})
}

// VerificationTokenExpiryLTE applies the LTE predicate on the "verification_token_expiry" field.
func VerificationTokenExpiryLTE(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVerificationTokenExpiry), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Member) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Member) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
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
func Not(p predicate.Member) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		p(s.Not())
	})
}