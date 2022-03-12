package schema

import (
	"errors"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// EmailValidator returns the validator that check the passed email ends with @sfc.keio.ac.jp.
func EmailValidator() func(email string) error {
	return func(email string) error {
		// ToDo: Replace "@gmail.com" to "@sfc.keio.ac.jp" when it is avaliable.
		if !strings.HasSuffix(email, "@gmail.com") {
			return errors.New("invalid email")
		}
		return nil
	}
}

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		// Basic information.
		field.String("discord_uid").Unique().NotEmpty().Immutable(),
		field.String("email").Unique().
			NotEmpty().Validate(EmailValidator()).Immutable(),

		// Data for verification.
		field.Bool("is_verified").Default(false),
		field.Time("verification_expiry").Default(func() time.Time {
			return time.Now()
		}),
		field.Bytes("hashed_verification_token").NotEmpty().Sensitive(),
		field.Time("verification_token_expiry").Default(func() time.Time {
			return time.Now().Add(time.Minute * 15)
		}),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return nil
}
