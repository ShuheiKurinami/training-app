// backend/infrastructure/router/router.go

package router

import (
	"training-app/backend/interface/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(userController *controllers.UserController) *mux.Router {
	router := mux.NewRouter()

	// ユーザー関連のルート
	router.HandleFunc("/api/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", userController.GetUser).Methods("GET")
	router.HandleFunc("/api/users/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", userController.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users", userController.GetAllUsers).Methods("GET")

	return router
}
