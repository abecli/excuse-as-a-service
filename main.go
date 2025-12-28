package main

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type ExcuseResponse struct {
	Excuse string `json:"excuse"`
}

var excuses []string

func loadExcuses(path string) ([]string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var list []string
	if err := json.Unmarshal(b, &list); err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, fmt.Errorf("excuses list is empty")
	}
	return list, nil
}

// crypto-safe random index (no predictable math/rand seed issues)
func randomIndex(n int) int {
	if n <= 0 {
		return 0
	}
	var buf [8]byte
	if _, err := rand.Read(buf[:]); err != nil {
		// fallback: time-based (should be rare)
		return int(time.Now().UnixNano() % int64(n))
	}
	v := binary.LittleEndian.Uint64(buf[:])
	return int(v % uint64(n))
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// Optional: allow browser calls (handy for testing)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeTxt(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// Optional: allow browser calls (handy for testing)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(msg))
}

func main() {
	var err error
	excuses, err = loadExcuses("excuses.json")
	if err != nil {
		log.Fatalf("failed to load excuses.json: %v", err)
	}

	mux := http.NewServeMux()

	// Health
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, 200, map[string]any{
			"ok":      true,
			"service": "excuse-api",
			"total":   len(excuses),
		})
	})

	// Random excuse (JSON)
	// mux.HandleFunc("/excuse", func(w http.ResponseWriter, r *http.Request) {
	// 	i := randomIndex(len(excuses))
	// 	writeJSON(w, 200, ExcuseResponse{
	// 		Excuse: excuses[i],
	// 	})
	// })

	// Random excuse (plain text) — great for shortcuts/curl
	mux.HandleFunc("/excuse", func(w http.ResponseWriter, r *http.Request) {
		i := randomIndex(len(excuses))
		writeTxt(w, 200, excuses[i])
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Cloud Run uses PORT env; default 8080 works locally too
	}

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("listening on :%s (excuses=%d)", port, len(excuses))
	log.Fatal(server.ListenAndServe())
}
