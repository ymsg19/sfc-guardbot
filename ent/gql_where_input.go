// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"time"

	"github.com/ymsg19/sfc-guardbot/ent/member"
	"github.com/ymsg19/sfc-guardbot/ent/predicate"
)

// MemberWhereInput represents a where input for filtering Member queries.
type MemberWhereInput struct {
	Not *MemberWhereInput   `json:"not,omitempty"`
	Or  []*MemberWhereInput `json:"or,omitempty"`
	And []*MemberWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "discord_uid" field predicates.
	DiscordUID             *string  `json:"discordUID,omitempty"`
	DiscordUIDNEQ          *string  `json:"discordUIDNEQ,omitempty"`
	DiscordUIDIn           []string `json:"discordUIDIn,omitempty"`
	DiscordUIDNotIn        []string `json:"discordUIDNotIn,omitempty"`
	DiscordUIDGT           *string  `json:"discordUIDGT,omitempty"`
	DiscordUIDGTE          *string  `json:"discordUIDGTE,omitempty"`
	DiscordUIDLT           *string  `json:"discordUIDLT,omitempty"`
	DiscordUIDLTE          *string  `json:"discordUIDLTE,omitempty"`
	DiscordUIDContains     *string  `json:"discordUIDContains,omitempty"`
	DiscordUIDHasPrefix    *string  `json:"discordUIDHasPrefix,omitempty"`
	DiscordUIDHasSuffix    *string  `json:"discordUIDHasSuffix,omitempty"`
	DiscordUIDEqualFold    *string  `json:"discordUIDEqualFold,omitempty"`
	DiscordUIDContainsFold *string  `json:"discordUIDContainsFold,omitempty"`

	// "email" field predicates.
	Email             *string  `json:"email,omitempty"`
	EmailNEQ          *string  `json:"emailNEQ,omitempty"`
	EmailIn           []string `json:"emailIn,omitempty"`
	EmailNotIn        []string `json:"emailNotIn,omitempty"`
	EmailGT           *string  `json:"emailGT,omitempty"`
	EmailGTE          *string  `json:"emailGTE,omitempty"`
	EmailLT           *string  `json:"emailLT,omitempty"`
	EmailLTE          *string  `json:"emailLTE,omitempty"`
	EmailContains     *string  `json:"emailContains,omitempty"`
	EmailHasPrefix    *string  `json:"emailHasPrefix,omitempty"`
	EmailHasSuffix    *string  `json:"emailHasSuffix,omitempty"`
	EmailEqualFold    *string  `json:"emailEqualFold,omitempty"`
	EmailContainsFold *string  `json:"emailContainsFold,omitempty"`

	// "is_verified" field predicates.
	IsVerified    *bool `json:"isVerified,omitempty"`
	IsVerifiedNEQ *bool `json:"isVerifiedNEQ,omitempty"`

	// "verification_expiry" field predicates.
	VerificationExpiry      *time.Time  `json:"verificationExpiry,omitempty"`
	VerificationExpiryNEQ   *time.Time  `json:"verificationExpiryNEQ,omitempty"`
	VerificationExpiryIn    []time.Time `json:"verificationExpiryIn,omitempty"`
	VerificationExpiryNotIn []time.Time `json:"verificationExpiryNotIn,omitempty"`
	VerificationExpiryGT    *time.Time  `json:"verificationExpiryGT,omitempty"`
	VerificationExpiryGTE   *time.Time  `json:"verificationExpiryGTE,omitempty"`
	VerificationExpiryLT    *time.Time  `json:"verificationExpiryLT,omitempty"`
	VerificationExpiryLTE   *time.Time  `json:"verificationExpiryLTE,omitempty"`

	// "verification_token_expiry" field predicates.
	VerificationTokenExpiry      *time.Time  `json:"verificationTokenExpiry,omitempty"`
	VerificationTokenExpiryNEQ   *time.Time  `json:"verificationTokenExpiryNEQ,omitempty"`
	VerificationTokenExpiryIn    []time.Time `json:"verificationTokenExpiryIn,omitempty"`
	VerificationTokenExpiryNotIn []time.Time `json:"verificationTokenExpiryNotIn,omitempty"`
	VerificationTokenExpiryGT    *time.Time  `json:"verificationTokenExpiryGT,omitempty"`
	VerificationTokenExpiryGTE   *time.Time  `json:"verificationTokenExpiryGTE,omitempty"`
	VerificationTokenExpiryLT    *time.Time  `json:"verificationTokenExpiryLT,omitempty"`
	VerificationTokenExpiryLTE   *time.Time  `json:"verificationTokenExpiryLTE,omitempty"`
}

// Filter applies the MemberWhereInput filter on the MemberQuery builder.
func (i *MemberWhereInput) Filter(q *MemberQuery) (*MemberQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering members.
// An error is returned if the input is empty or invalid.
func (i *MemberWhereInput) P() (predicate.Member, error) {
	var predicates []predicate.Member
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, member.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Member, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, member.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Member, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, member.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, member.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, member.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, member.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, member.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, member.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, member.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, member.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, member.IDLTE(*i.IDLTE))
	}
	if i.DiscordUID != nil {
		predicates = append(predicates, member.DiscordUIDEQ(*i.DiscordUID))
	}
	if i.DiscordUIDNEQ != nil {
		predicates = append(predicates, member.DiscordUIDNEQ(*i.DiscordUIDNEQ))
	}
	if len(i.DiscordUIDIn) > 0 {
		predicates = append(predicates, member.DiscordUIDIn(i.DiscordUIDIn...))
	}
	if len(i.DiscordUIDNotIn) > 0 {
		predicates = append(predicates, member.DiscordUIDNotIn(i.DiscordUIDNotIn...))
	}
	if i.DiscordUIDGT != nil {
		predicates = append(predicates, member.DiscordUIDGT(*i.DiscordUIDGT))
	}
	if i.DiscordUIDGTE != nil {
		predicates = append(predicates, member.DiscordUIDGTE(*i.DiscordUIDGTE))
	}
	if i.DiscordUIDLT != nil {
		predicates = append(predicates, member.DiscordUIDLT(*i.DiscordUIDLT))
	}
	if i.DiscordUIDLTE != nil {
		predicates = append(predicates, member.DiscordUIDLTE(*i.DiscordUIDLTE))
	}
	if i.DiscordUIDContains != nil {
		predicates = append(predicates, member.DiscordUIDContains(*i.DiscordUIDContains))
	}
	if i.DiscordUIDHasPrefix != nil {
		predicates = append(predicates, member.DiscordUIDHasPrefix(*i.DiscordUIDHasPrefix))
	}
	if i.DiscordUIDHasSuffix != nil {
		predicates = append(predicates, member.DiscordUIDHasSuffix(*i.DiscordUIDHasSuffix))
	}
	if i.DiscordUIDEqualFold != nil {
		predicates = append(predicates, member.DiscordUIDEqualFold(*i.DiscordUIDEqualFold))
	}
	if i.DiscordUIDContainsFold != nil {
		predicates = append(predicates, member.DiscordUIDContainsFold(*i.DiscordUIDContainsFold))
	}
	if i.Email != nil {
		predicates = append(predicates, member.EmailEQ(*i.Email))
	}
	if i.EmailNEQ != nil {
		predicates = append(predicates, member.EmailNEQ(*i.EmailNEQ))
	}
	if len(i.EmailIn) > 0 {
		predicates = append(predicates, member.EmailIn(i.EmailIn...))
	}
	if len(i.EmailNotIn) > 0 {
		predicates = append(predicates, member.EmailNotIn(i.EmailNotIn...))
	}
	if i.EmailGT != nil {
		predicates = append(predicates, member.EmailGT(*i.EmailGT))
	}
	if i.EmailGTE != nil {
		predicates = append(predicates, member.EmailGTE(*i.EmailGTE))
	}
	if i.EmailLT != nil {
		predicates = append(predicates, member.EmailLT(*i.EmailLT))
	}
	if i.EmailLTE != nil {
		predicates = append(predicates, member.EmailLTE(*i.EmailLTE))
	}
	if i.EmailContains != nil {
		predicates = append(predicates, member.EmailContains(*i.EmailContains))
	}
	if i.EmailHasPrefix != nil {
		predicates = append(predicates, member.EmailHasPrefix(*i.EmailHasPrefix))
	}
	if i.EmailHasSuffix != nil {
		predicates = append(predicates, member.EmailHasSuffix(*i.EmailHasSuffix))
	}
	if i.EmailEqualFold != nil {
		predicates = append(predicates, member.EmailEqualFold(*i.EmailEqualFold))
	}
	if i.EmailContainsFold != nil {
		predicates = append(predicates, member.EmailContainsFold(*i.EmailContainsFold))
	}
	if i.IsVerified != nil {
		predicates = append(predicates, member.IsVerifiedEQ(*i.IsVerified))
	}
	if i.IsVerifiedNEQ != nil {
		predicates = append(predicates, member.IsVerifiedNEQ(*i.IsVerifiedNEQ))
	}
	if i.VerificationExpiry != nil {
		predicates = append(predicates, member.VerificationExpiryEQ(*i.VerificationExpiry))
	}
	if i.VerificationExpiryNEQ != nil {
		predicates = append(predicates, member.VerificationExpiryNEQ(*i.VerificationExpiryNEQ))
	}
	if len(i.VerificationExpiryIn) > 0 {
		predicates = append(predicates, member.VerificationExpiryIn(i.VerificationExpiryIn...))
	}
	if len(i.VerificationExpiryNotIn) > 0 {
		predicates = append(predicates, member.VerificationExpiryNotIn(i.VerificationExpiryNotIn...))
	}
	if i.VerificationExpiryGT != nil {
		predicates = append(predicates, member.VerificationExpiryGT(*i.VerificationExpiryGT))
	}
	if i.VerificationExpiryGTE != nil {
		predicates = append(predicates, member.VerificationExpiryGTE(*i.VerificationExpiryGTE))
	}
	if i.VerificationExpiryLT != nil {
		predicates = append(predicates, member.VerificationExpiryLT(*i.VerificationExpiryLT))
	}
	if i.VerificationExpiryLTE != nil {
		predicates = append(predicates, member.VerificationExpiryLTE(*i.VerificationExpiryLTE))
	}
	if i.VerificationTokenExpiry != nil {
		predicates = append(predicates, member.VerificationTokenExpiryEQ(*i.VerificationTokenExpiry))
	}
	if i.VerificationTokenExpiryNEQ != nil {
		predicates = append(predicates, member.VerificationTokenExpiryNEQ(*i.VerificationTokenExpiryNEQ))
	}
	if len(i.VerificationTokenExpiryIn) > 0 {
		predicates = append(predicates, member.VerificationTokenExpiryIn(i.VerificationTokenExpiryIn...))
	}
	if len(i.VerificationTokenExpiryNotIn) > 0 {
		predicates = append(predicates, member.VerificationTokenExpiryNotIn(i.VerificationTokenExpiryNotIn...))
	}
	if i.VerificationTokenExpiryGT != nil {
		predicates = append(predicates, member.VerificationTokenExpiryGT(*i.VerificationTokenExpiryGT))
	}
	if i.VerificationTokenExpiryGTE != nil {
		predicates = append(predicates, member.VerificationTokenExpiryGTE(*i.VerificationTokenExpiryGTE))
	}
	if i.VerificationTokenExpiryLT != nil {
		predicates = append(predicates, member.VerificationTokenExpiryLT(*i.VerificationTokenExpiryLT))
	}
	if i.VerificationTokenExpiryLTE != nil {
		predicates = append(predicates, member.VerificationTokenExpiryLTE(*i.VerificationTokenExpiryLTE))
	}

	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("github.com/ymsg19/sfc-guardbot/ent: empty predicate MemberWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return member.And(predicates...), nil
	}
}