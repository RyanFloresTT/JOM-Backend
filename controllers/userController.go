package controllers

import (
	"encoding/json"
	"go-backend/initializers"
	"go-backend/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	// get email and pass from context
	var body struct {
		Email    string `json:"Email"`
		Password string `json:"Password"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusBadRequest)
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		http.Error(w, "Failed to create user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// get email and pass from context
	var body struct {
		Email    string `json:"Email"`
		Password string `json:"Password"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Look up user
	var user models.User
	initializers.DB.First(&user, "email=?", body.Email)

	if user.ID == 0 {
		http.Error(w, "Invalid Email or Password", http.StatusBadRequest)
		return
	}

	// Compare pass hashes
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		http.Error(w, "Invalid Email or Password", http.StatusBadRequest)
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("TokenSecret")))
	if err != nil {
		http.Error(w, "Problem generating token", http.StatusBadRequest)
		return
	}

	// Instead of setting the cookie, send the details in a JSON response
	response := map[string]string{
		"cookieName":  "Authorization",
		"cookieValue": tokenString,
		"expires":     time.Now().Add(time.Hour * 24 * 30).Format(time.RFC3339),
	}

	// Set content type and respond with the cookie details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully logged out",
	})
}

func Validate(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*models.User)
	if !ok || user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"validated":     true,
		"wants emails?": user.EmailPromotions,
	})
}
