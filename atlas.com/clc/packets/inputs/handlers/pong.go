package handlers

import (
	"atlas-clc/packets/inputs"
	"log"
)

type PongHandler struct {
}

func (h *PongHandler) Handle(_ *log.Logger, _ int, _ *inputs.Reader) {
}
