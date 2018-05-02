package middleware

import (
    "net/http"
    "../errors"
    "encoding/json"
    "fmt"
)

func ErrorMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if r := recover(); r != nil {
                w.Header().Set("Content-Type", "application/json")
                switch panic_error := r.(type) {
                case errors.APIError:
                    w.WriteHeader(panic_error.Status)
                    json.NewEncoder(w).Encode(panic_error)
                default:
                    fmt.Println(r)
                    w.WriteHeader(http.StatusInternalServerError)
                    json.NewEncoder(w).Encode(errors.APIError{Status: 500, Message: "Unknown Error"})
                }
            }
        }()

    	next.ServeHTTP(w, r)
    })
}
