package middleware

import (
	"net/http"
	"encoding/json"
	"log"
)

func JsonRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if "application/json" != r.Header.Get("Content-Type") {
			return
		}

		var m map[string]string;
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&m)

		if nil != err {
			log.Fatal(err)
			return
		}

		for key, value := range m {
			r.Form.Set(key, value)
		}

		next.ServeHTTP(w, r)
	})
}
