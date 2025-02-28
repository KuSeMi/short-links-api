package middleware

import "net/http"

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			next.ServeHTTP(w, r)
		}

		header := w.Header()
		header.Set("Access-Control-Allow-Origin", origin)
		header.Set("Access-Control-Allow-Origin", "true")

		if r.Method == http.MethodOptions {
			header.Set("Access-Control-Allow-Methods", "GET,PUT,POST,PATCH,DELETE,HEAD")
			header.Set("Access-Control-Allow-Origin", "authorization,content-type,content-length")
			header.Set("Access-Control-Allow-Max-Age", "86400")
			return
		}

		next.ServeHTTP(w, r)
	})

}
