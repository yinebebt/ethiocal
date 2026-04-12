package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/yinebebt/ethiocal/bahirehasab"
	"github.com/yinebebt/ethiocal/dateconverter"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

// BahireHasab handles GET /api/bahir/{year} and returns festival dates as JSON.
func BahireHasab(w http.ResponseWriter, r *http.Request) {
	yearString := r.PathValue("year")
	if yearString == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "Empty year value",
		})
		return
	}
	year, err := strconv.Atoi(yearString)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "not a valid year",
		})
		return
	}

	festival, err := bahirehasab.BahireHasab(year)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"data": festival,
	})
}

func parseDate(w http.ResponseWriter, r *http.Request) (year, month, day int, ok bool) {
	dateString := r.PathValue("date")
	if dateString == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "empty date",
		})
		return 0, 0, 0, false
	}
	parts := strings.Split(dateString, "-")
	if len(parts) != 3 {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "not a valid date, expected format: year-month-day",
		})
		return 0, 0, 0, false
	}

	year, err1 := strconv.Atoi(parts[0])
	month, err2 := strconv.Atoi(parts[1])
	day, err3 := strconv.Atoi(parts[2])
	if err1 != nil || err2 != nil || err3 != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "date must contain valid integers",
		})
		return 0, 0, 0, false
	}
	return year, month, day, true
}

// Ethiopian handles GET /api/gtoe/{date} and converts a Gregorian date to Ethiopian.
func Ethiopian(w http.ResponseWriter, r *http.Request) {
	year, month, day, ok := parseDate(w, r)
	if !ok {
		return
	}
	etDate, err := dateconverter.Ethiopian(year, month, day)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"ethiopian_date": etDate.Format("2006-01-02"),
	})
}

// Gregorian handles GET /api/etog/{date} and converts an Ethiopian date to Gregorian.
func Gregorian(w http.ResponseWriter, r *http.Request) {
	year, month, day, ok := parseDate(w, r)
	if !ok {
		return
	}
	gregDate, err := dateconverter.Gregorian(year, month, day)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"gregorian_date": gregDate.Format("2006-01-02"),
	})
}

// Init starts the HTTP server with all API routes registered. It blocks until
// a SIGINT or SIGTERM signal is received and then shuts down gracefully.
func Init() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/bahir/{year}", BahireHasab)
	mux.HandleFunc("GET /api/gtoe/{date}", Ethiopian)
	mux.HandleFunc("GET /api/etog/{date}", Gregorian)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("listening on :%s", port)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	log.Println("shutting down gracefully, received signal:", sig)

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}
}
