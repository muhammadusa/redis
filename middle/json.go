package middle

import "net/http"

func JSON (next func(http.ResponseWriter, *http.Request)) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		
		w.Header().Set("Content-Type", "application/json")
		
		next(w, r)
	})
}