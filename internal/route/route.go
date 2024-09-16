package route

import (
	"net/http"

	bidhandler "tzAvito/internal/api/bids"
	handler "tzAvito/internal/api/tender"
	"tzAvito/internal/db"
	"tzAvito/internal/repository"
	"tzAvito/internal/service"
	"github.com/go-chi/chi"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/rs/cors"
)

func SetupRoutes(router *chi.Mux) {
	tenderRepo := repository.NewRepository(db.DB)
	tenderService := service.NewTenderService(tenderRepo)
	tenderHandler := handler.NewImplementation(tenderService)

	bidRepo := repository.NewBidRepository(db.DB)
	bidService := service.NewBidService(bidRepo)
	bidHandler := bidhandler.NewImplementation(bidService)
	router.Use(cors.AllowAll().Handler)
    router.Use(middleware.Logger) 

    // Ping route
    router.Get("/api/ping", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message": "ok"}`))
    })

    // Tenders routes
    router.Route("/api/tenders", func(r chi.Router) {
        r.Post("/new", tenderHandler.CreateTender)
        r.Get("/{tenderId}/status", tenderHandler.FindTenderByID)
        r.Get("/", tenderHandler.GetTenders)
        r.Get("/my", tenderHandler.GetUserTender)
        r.Put("/{tenderId}/status", tenderHandler.UpdateTenderStatus)
        r.Patch("/{tenderId}/edit", tenderHandler.EditTenderHandler)
        r.Put("/{tenderId}/rollback/{version}", tenderHandler.RollbackToVersion)
    })

    // Bids routes
    router.Route("/api/bids", func(r chi.Router) {
        r.Post("/new", bidHandler.CreateBid)
        r.Get("/my", bidHandler.GetUserBid)
        r.Get("/{bidId}/status", bidHandler.GetBidStatus)
        r.Put("/{bidId}/status", bidHandler.UpdateBidStatus)
        r.Get("/{tenderId}/list", bidHandler.TenderList) 
       // r.Patch("/{bidId}/edit", bidHandler.UpdateBid)
        r.Put("/{bidId}/submit_decision", bidHandler.SubmitDecision)
        // r.Put("/{bidId}/feedback", bidHandler.FeedBack)
        // r.Put("/{bidId}/rollback/{version}", bidHandler.Rollback)
        // r.Get("/{tenderId}/reviews", bidHandler.Reviews)
    })
}
//я не знаю насколько часто возникает проблемы с wildcards но в следующий раз не буду юзать gin для роутера ;< пришлось переписать все на chi 
