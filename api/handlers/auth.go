package handlers

import (
	"encoding/json"
	"go-project/internal/auth"
	"net/http"
	"time"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
    creds, err := auth.GetCredentials(r)
    if err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    user, err := auth.ValidateCredentials(creds)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    sessionID, err := auth.CreateSession(user)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Set the session ID in a cookie
    http.SetCookie(w, &http.Cookie{
        Name:    "session_token",
        Value:   sessionID,
        Path:    "/",
        Expires: time.Now().Add(24 * time.Hour),
        HttpOnly: true, // Helps mitigate XSS
        Secure:   true, // Only send over HTTPS
        SameSite: http.SameSiteStrictMode, // Helps mitigate CSRF
    })

    // Send a success response (possibly with user data)
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(map[string]string{"message": "Signed in Successfully!"})

    // Write user data to response if needed
}

