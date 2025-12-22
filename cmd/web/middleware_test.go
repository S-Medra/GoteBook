package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/S-Medra/GoteBook/internal/assert"
)

func TestSecureHeaders(t *testing.T) {
	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	secureHeaders(next).ServeHTTP(rr, r)

	rs := rr.Result()
	defer rs.Body.Close()

	assert.Equal(t, "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com", rs.Header.Get("Content-Security-Policy"))
	assert.Equal(t, "origin-when-cross-origin", rs.Header.Get("Referrer-Policy"))
	assert.Equal(t, "nosniff", rs.Header.Get("X-Content-Type-Options"))
	assert.Equal(t, "deny", rs.Header.Get("X-Frame-Options"))
	assert.Equal(t, "0", rs.Header.Get("X-XSS-Protection"))
	assert.Equal(t, http.StatusOK, rs.StatusCode)

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	body = bytes.TrimSpace(body)
	assert.Equal(t, "OK", string(body))
}
