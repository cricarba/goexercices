package MI
import (
	"fmt"
	"net/http"
	"time"
)

func middlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before handler; start middleware")
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Println("finish middleware", time.Since(start))

	})
}