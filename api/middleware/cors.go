package middleware

import (
	"net/http"
	"regexp"
	"strings"
)

type Cors struct {
	OriginRules []string
}

func NewCors(origin string) *Cors {
	rules := strings.Split(origin, ",")
	for i, s := range rules {
		rules[i] = strings.TrimSpace(s)
	}
	return &Cors{
		OriginRules: rules,
	}
}

func (m *Cors) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.Method == "OPTIONS" {
		if m.allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		return
	}
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	next(w, r)
}

func (m *Cors) allowedOrigin(origin string) bool {
	for _, r := range m.OriginRules {
		if r == "*" {
			return true
		}
		if matched, _ := regexp.MatchString(r, origin); matched {
			return true
		}
	}
	return false
}
