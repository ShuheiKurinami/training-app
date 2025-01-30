// backend/infrastructure/router/router.go

package router

import (
	"net/http"

	"github.com/ShuheiKurinami/training-app/backend/interface/controllers"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func SetupRoutes(userController *controllers.UserController) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", userController.GetUser).Methods("GET")
	router.HandleFunc("/api/users/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", userController.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users", userController.GetAllUsers).Methods("GET")

	// CORS を許可する
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(router)

	return corsHandler
}
