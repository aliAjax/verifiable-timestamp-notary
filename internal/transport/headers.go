package transport

import "net/http"

func SetSecurityHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Cache-Control", "no-store")
}
func RequestID(r *http.Request) string {
	if v := r.Header.Get("X-Request-ID"); v != "" {
		return v
	}
	return "generated"
}
func AcceptsJSON(r *http.Request) bool {
	return r.Header.Get("Accept") == "" || r.Header.Get("Accept") == "application/json"
}
