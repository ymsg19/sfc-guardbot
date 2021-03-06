// Code generated by entc, DO NOT EDIT.

package member

import (
	"time"
)

const (
	// Label holds the string label denoting the member type in the database.
	Label = "member"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDiscordUID holds the string denoting the discord_uid field in the database.
	FieldDiscordUID = "discord_uid"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldIsVerified holds the string denoting the is_verified field in the database.
	FieldIsVerified = "is_verified"
	// FieldVerificationExpiry holds the string denoting the verification_expiry field in the database.
	FieldVerificationExpiry = "verification_expiry"
	// FieldHashedVerificationToken holds the string denoting the hashed_verification_token field in the database.
	FieldHashedVerificationToken = "hashed_verification_token"
	// FieldVerificationTokenExpiry holds the string denoting the verification_token_expiry field in the database.
	FieldVerificationTokenExpiry = "verification_token_expiry"
	// Table holds the table name of the member in the database.
	Table = "members"
)

// Columns holds all SQL columns for member fields.
var Columns = []string{
	FieldID,
	FieldDiscordUID,
	FieldEmail,
	FieldIsVerified,
	FieldVerificationExpiry,
	FieldHashedVerificationToken,
	FieldVerificationTokenExpiry,
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
	// DiscordUIDValidator is a validator for the "discord_uid" field. It is called by the builders before save.
	DiscordUIDValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultIsVerified holds the default value on creation for the "is_verified" field.
	DefaultIsVerified bool
	// DefaultVerificationExpiry holds the default value on creation for the "verification_expiry" field.
	DefaultVerificationExpiry func() time.Time
	// HashedVerificationTokenValidator is a validator for the "hashed_verification_token" field. It is called by the builders before save.
	HashedVerificationTokenValidator func([]byte) error
	// DefaultVerificationTokenExpiry holds the default value on creation for the "verification_token_expiry" field.
	DefaultVerificationTokenExpiry func() time.Time
)
