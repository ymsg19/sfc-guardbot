package handlers

import (
	"context"
	"log"

	"github.com/bwmarrin/discordgo"

	"github.com/ymsg19/sfc-guardbot/ent"
	"github.com/ymsg19/sfc-guardbot/ent/member"

	"github.com/MakeNowJust/heredoc/v2"
)

type Handlers struct {
	DB *ent.Client
}

func NewHandlers(db *ent.Client) *Handlers {
	return &Handlers{
		DB: db,
	}
}

func (h Handlers) ResigsterAll(s *discordgo.Session) {
	s.AddHandler(h.onMemberJoined)
}

func (h Handlers) onMemberJoined(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	log.Printf("onMemberJoined fired; user id: %s", e.Member.User.ID)
	joined := e.Member.User.ID

	res, err := h.DB.Member.Query().Where(
		member.DiscordUID(joined),
	).Exist(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if res {
		if err := s.GuildMemberRoleAdd(
			e.Member.GuildID,
			e.Member.User.ID,
			"954324951450652684",
		); err != nil {
			log.Fatal(err)
		}
		return
	}

	dm, err := s.UserChannelCreate(e.Member.User.ID)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := s.ChannelMessageSend(dm.ID, heredoc.Doc(`
			あなたは認証が済んでいないため、サーバーからキックされました。
			https://sfc.guardbot.online/ にて認証を行ってください。
			You have been kicked from the server because you are not verified.
			Please verify yourself at https://sfc.guardbot.online/
		`)); err != nil {
		log.Fatal(err)
	}

	if err := s.GuildMemberDeleteWithReason(
		e.Member.GuildID,
		e.Member.User.ID,
		"not verified",
	); err != nil {
		log.Fatal(err)
	}

	if _, err := s.ChannelMessageSend(
		"918127893836628028",
		heredoc.Docf(`
				%s は認証が済んでいないため、サーバーからキックされました。
				%s has been kicked from the server because they is not verified.
			`, e.Member.User.Mention(), e.Member.User.Mention()),
	); err != nil {
		log.Fatal(err)
	}
}
