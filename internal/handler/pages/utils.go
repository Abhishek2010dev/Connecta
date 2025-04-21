package pages

import (
	"net/http"

	"github.com/Abhishek2010dev/Connecta/internal/dto"
	"github.com/Abhishek2010dev/Connecta/internal/middleware"
)

func GetAuthPayload(r *http.Request) dto.AuthPaylaod {
	return r.Context().Value(middleware.AuthPayloadKey).(dto.AuthPaylaod)
}
