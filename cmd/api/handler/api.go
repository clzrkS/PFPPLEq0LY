package handler

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func ExternalRoutes(app *ApplicationConfig) http.Handler {
	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))


	// router.Use(middleware.AuthMiddleware(app.AuthService))

	router.Route("/v1", func(router chi.Router) {
		router.Get("/health", ApplicationHealthHandler)

		router.Route("/student", func(router chi.Router) {
			router.Get("/", GetStudentHandler)
			router.Post("/create", app.CreateStudentHandler)
			router.Get("/get-by-id/{studentId}", app.GetStudentByIdHandler)
			router.Delete("/delete-by-id/{studentId}", app.DeleteStudentByIdHandler)
			router.Put("/deactivate-by-id", app.DeactivateStudentByIdHandler)
			router.Get("/get-all", app.GetAllStudentsHandler)
			router.Get("/get-by-email", app.GetStudentByEmailHandler)
			// router.Get("/login", app.LoginHandler)
		})
	})

	return router
}
