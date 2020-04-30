package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// middleware provides a convenient mechanism for filtering HTTP requests
// entering the application. It returns a new handler which may perform various
// operations and should finish by calling the next HTTP handler.
type middleware func(next http.HandlerFunc) http.HandlerFunc

// chainMiddleware provides syntactic sugar to create a new middleware
// which will be the result of chaining the ones received as parameters.
func chainMiddleware(mw ...middleware) middleware {
	return func(final http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			last := final
			for i := len(mw) - 1; i >= 0; i-- {
				last = mw[i](last)
			}
			last(w, r)
		}
	}
}

// Parse jwt token and decode the user struct
// Attach the user struct in the context of the request
func withUserParsing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Middleware called")
		//Get jwt token from headers
		splitToken := strings.Split(req.Header.Get("Authorization"), "Bearer ")
		if len(splitToken) < 2 {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(w, "No token provided")
			return
		}
		jwtTokenString := splitToken[1]
		token, err := jwt.Parse(jwtTokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte("secret"), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Its an error %s", err)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			next.ServeHTTP(w, req.WithContext(context.WithValue(req.Context(),
				"userID", claims["id"])))
		} else {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Unable to Parse Token")
		}

	}
}
