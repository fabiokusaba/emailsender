package endpoints

import (
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/render"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "request doesn't contain an authorization header"})
			return
		}

		token = strings.Replace(token, "Bearer ", "", 1)
		provider, err := oidc.NewProvider(r.Context(), "http://localhost:8080/realms/emailsenderprovider")
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, map[string]string{"error": "error to connect to the provider"})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: "emailsenderclient"})
		// verifier := provider.Verifier(&oidc.Config{SkipClientIDCheck: true})
		_, err = verifier.Verify(r.Context(), token)
		if err != nil {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "invalid token"})
			return
		}

		next.ServeHTTP(w, r)
	})
}
