package handlers

import (
	"atlas-clc/packets/inputs"
	"atlas-clc/packets/inputs/models"
	"atlas-clc/packets/inputs/readers"
	"atlas-clc/packets/outputs/writers"
	"atlas-clc/processors"
	"atlas-clc/registries"
	"atlas-clc/sessions"
	"log"
)

type CreateCharacterHandler struct {
}

func (h *CreateCharacterHandler) Handle(l *log.Logger, sessionId int, r *inputs.Reader) {
	s := registries.GetSessionRegistry().GetSession(sessionId)
	if s != nil {
		p := readers.ReadCreateCharacter(r)
		if p != nil {
			h.handle(l, s, p)
		}
	}
}

func (h *CreateCharacterHandler) handle(l *log.Logger, s *sessions.Session, p *models.CreateCharacter) {
	ca, err := processors.SeedCharacter(l, s.AccountId(), s.WorldId(), p.Name(), p.Job(), p.Face(), p.Hair(),
		p.HairColor(), p.SkinColor(), p.Gender(), p.Top(), p.Bottom(), p.Shoes(), p.Weapon())
	if err != nil {
		l.Println("[ERROR] while seeding character")
		return
	}

	c, err := processors.GetCharacterById(l, ca.Id())
	if err != nil {
		l.Println("[ERROR] retrieving newly seeded character")
		return
	}

	s.Announce(writers.WriteNewCharacter(*c))
}
