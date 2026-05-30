package controller

import (
	"net/http"
	"strings"
)

type AuthContext struct {
	Username     string   `json:"username"`
	ClientID     string   `json:"client_id"`
	Roles        []string `json:"roles"`
	WorkEntities []string `json:"work_entities"`
}

func ExtractAuthContext(r *http.Request) AuthContext {
	roles := splitNonEmpty(r.Header.Get("X-Auth-Roles"))
	wes := splitNonEmpty(r.Header.Get("X-Auth-Work-Entities"))
	if roles == nil {
		roles = []string{}
	}
	if wes == nil {
		wes = []string{}
	}
	return AuthContext{
		Username:     r.Header.Get("X-Auth-User"),
		ClientID:     r.Header.Get("X-Auth-Client-Id"),
		Roles:        roles,
		WorkEntities: wes,
	}
}

func splitNonEmpty(s string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, ",")
}
