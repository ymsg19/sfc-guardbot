package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ymsg19/sfc-guardbot/ent"
	"github.com/ymsg19/sfc-guardbot/ent/member"
	"github.com/ymsg19/sfc-guardbot/graph/generated"
	"github.com/ymsg19/sfc-guardbot/graph/model"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) CreateMember(ctx context.Context, input model.MemberInput) (*ent.Member, error) {
	client := ent.FromContext(ctx)

	token, err := generateToken()
	if err != nil {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword(token.ToSlice(), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	res, err := client.Member.Create().
		SetDiscordUID(getDiscordIDFromContext(ctx)).
		SetEmail(input.Email).
		SetHashedVerificationToken(hashed).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	err = sendTokenByMail(input.Email, token.ToString())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *mutationResolver) VerifyMember(ctx context.Context, input model.MemberVerificationInput) (*model.VerificationResult, error) {
	client := ent.FromContext(ctx)

	member, err := client.Member.Get(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if !authorizeDiscordID(ctx, member.DiscordUID) {
		return nil, fmt.Errorf("not authorized")
	}

	if member.IsVerified && member.VerificationExpiry.After(time.Now()) {
		return nil, fmt.Errorf("already verified")
	}

	expiry := member.VerificationTokenExpiry
	if expiry.Before(time.Now()) {
		return nil, fmt.Errorf("verification token expired")
	}

	token, err := base64.URLEncoding.DecodeString(input.VerificationToken)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(member.HashedVerificationToken, token)
	if err != nil {
		return nil, err
	}

	if _, err := member.Update().
		SetIsVerified(true).
		SetVerificationExpiry(time.Now().AddDate(0, 6, 0)).
		Save(ctx); err != nil {
		return nil, err
	}

	invite, err := r.session.ChannelInviteCreate("918127893836628028", discordgo.Invite{
		MaxAge:    0,
		MaxUses:   1,
		Temporary: true,
		Unique:    true,
	})
	if err != nil {
		return nil, err
	}

	return &model.VerificationResult{
		InviteCode: invite.Code,
	}, nil
}

func (r *mutationResolver) RequestMemberVerificationToken(ctx context.Context, id int) (*ent.Member, error) {
	client := ent.FromContext(ctx)

	member, err := client.Member.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if !authorizeDiscordID(ctx, member.DiscordUID) {
		return nil, fmt.Errorf("not authorized")
	}

	token, err := generateToken()
	if err != nil {
		return nil, err
	}

	hashed, _ := bcrypt.GenerateFromPassword(token.ToSlice(), bcrypt.DefaultCost)

	member, err = member.Update().
		SetHashedVerificationToken(hashed).
		SetVerificationTokenExpiry(time.Now().Add(time.Minute * 15)).
		SetVerificationExpiry(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	err = sendTokenByMail(member.Email, token.ToString())
	if err != nil {
		return nil, err
	}

	return member, nil
}

func (r *mutationResolver) DeleteMember(ctx context.Context, id int) (bool, error) {
	client := ent.FromContext(ctx)

	member, err := client.Member.Get(ctx, id)
	if err != nil {
		return false, err
	}

	if !authorizeDiscordID(ctx, member.DiscordUID) {
		return false, fmt.Errorf("not authorized")
	}

	err = client.Member.DeleteOne(member).Exec(ctx)
	return err == nil, err
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

func (r *queryResolver) Member(ctx context.Context) (*ent.Member, error) {
	return r.client.Member.Query().
		Where(member.DiscordUID(getDiscordIDFromContext(ctx))).
		First(ctx)
}

func (r *queryResolver) Members(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.MemberWhereInput) (*ent.MemberConnection, error) {
	// return r.client.Member.Query().
	// 	Paginate(ctx, after, first, before, last, ent.WithMemberFilter(where.Filter))

	// ToDo: Implement later.
	return nil, fmt.Errorf("unimplemented yet")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
