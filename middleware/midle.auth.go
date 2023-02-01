package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/biFebriansyah/backintro/libs"
)

func AuthMidle(role ...string) Middle {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var header string
			var valid bool

			// get header optional
			if header = r.Header.Get("Authorization"); header == "" {
				libs.NewRes("header not provide, please login", 401, true).Send(w)
				return
			}

			if !strings.Contains(header, "Bearer") {
				libs.NewRes("invalid header type", 401, true).Send(w)
				return
			}

			tokens := strings.Replace(header, "Bearer ", "", -1)

			chekToken, err := libs.CekToken(tokens)
			if err != nil {
				libs.NewRes(err.Error(), 401, true).Send(w)
				return
			}

			for _, rl := range role {
				if rl == chekToken.Role {
					valid = true
				}
			}

			if !valid {
				libs.NewRes("you not have permission to accsess", 401, true).Send(w)
				return
			}

			log.Println("Auth middleware pass")

			// share userid to controller
			ctx := context.WithValue(r.Context(), "user", chekToken.User_Id)

			// serve next middleware
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}
