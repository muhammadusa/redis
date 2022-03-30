package middle

import(
	"log"
	"time"
	"context"
	"net/http"
	"encoding/json"

	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr: "loclahost:6379",
	Password: "Lv9az31tQPNOlJpQxCYMqX9z3VCT9NhZ4uT1Ksw0gfw=",
	DB: 0,
})  

func RedisCashe(next func(http.ResponseWriter, *http.Request)) http.Handler {
	
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		
		_, err := client.Ping().Result()

		if err != nil {
			
			next(w, r)
			
			return 
		}

		key := r.RequestURI

		response, err := client.Get(key).Result()

		if err == nil && r.Method == "GET" {
		
			log.Println("Executed MiddleWare")
		
			w.Write([]byte(response))
		
			return
		}

		ctx := context.WithValue(r.Context(), "cashe", func (value interface{}, expire int) {
			jsonData, err := json.Marshal(value)

			if err == nil {
			
				client.Set(r.RequestURI, jsonData, time.Duration(expire))
			}
		})

		req := r.WithContext(ctx)

		next(w, req)
	})
}