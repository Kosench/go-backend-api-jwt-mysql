package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go-backend-api-jwt-mysql/config"
	"go-backend-api-jwt-mysql/types"
	"go-backend-api-jwt-mysql/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

type contextKey string

const UserKey contextKey = "userID"

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func WithJWTAuth(handleFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the token from the user request
		tokenString := getTokenFromRequest(r)

		// validate the JWT
		token, err := validateJWT(tokenString)
		if err != nil {
			log.Printf("failed to validate token %v ", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}
		// if is we need to fetch the userID from DB (from the token)
		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		userID, err := strconv.Atoi(str)

		u, err := store.GetUserByID(userID)
		if err != nil {
			log.Printf("failed to get user by id %v", err)
			permissionDenied(w)
			return
		}

		//set context user_ID to the user
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		handleFunc(w, r)
	}
}

func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")

	if tokenAuth != "" {
		return tokenAuth
	}

	return ""
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permissin denied"))
}
