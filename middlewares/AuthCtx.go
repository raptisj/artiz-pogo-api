package middlewares

import (
	"artiz-pogo-api/utils"
	"context"
	"fmt"
	"net/http"
)

func AuthCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := utils.VerifyToken(r.Header.Get("Authorization"))
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		userId := token.UID
		ctx = context.WithValue(ctx, "userId", userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
