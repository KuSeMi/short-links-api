package stat

import (
	"demo/configs"
	"demo/pkg/event"
	"demo/pkg/middleware"
	"demo/pkg/res"
	"net/http"
	"time"
)

const (
	GroupByDate  = "day"
	GroupByMonth = "month"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

type StatHandler struct {
	StatRepository *StatRepository
	EventBus       *event.EventBus
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}
	router.Handle("GET /stat", middleware.IsAuthed(handler.GetStat(), deps.Config))

}

func (handler *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "invalid from params", http.StatusBadRequest)
			return
		}
		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "invalid to params", http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by != GroupByDate && by != GroupByMonth {
			http.Error(w, "invalid by params", http.StatusBadRequest)
			return
		}
		stats := handler.StatRepository.GetStats(by, from, to)
		res.Json(w, stats, 200)
	}
}
