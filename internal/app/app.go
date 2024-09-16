package app

import (
	"context"
	"log"
	"net/http"
	"path/filepath"
	"tzAvito/internal/config"
	"tzAvito/internal/db"
	"tzAvito/internal/route"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//моя попытка сделать singleton одиночка паттерн не знаю насколько правильно получился
type App struct {
	serviceProvider *ServiceProvider
	router          *chi.Mux
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initConfig(_ context.Context) error {
	envPath := filepath.Join("..", "pkg", ".env")
	err := config.Load(envPath)
	if err != nil {
		return err
	}
	db.ConnectDB()

	return nil

}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initRouter,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
func (a *App) Run() error {
	address := "0.0.0.0:8080"
	//address := "localhost:8080"
	log.Printf("HTTP server is running on %s", address)

	err := http.ListenAndServe(address, a.router)
	if err != nil {
		return err
	}

	return nil
}
func (a *App) initRouter(_ context.Context) error {
	a.router = chi.NewRouter()
	a.router.Use(CORSMiddleware)
	a.router.Use(middleware.Logger) 
	route.SetupRoutes(a.router)

	return nil
}
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
