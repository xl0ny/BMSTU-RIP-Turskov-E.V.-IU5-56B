package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"pankreatitmed/internal/app/repository"
)

type Handler struct {
	Repo    *repository.Repository
	tplBase string
}

func NewHandler(repo *repository.Repository, tplDir string) *Handler {
	return &Handler{Repo: repo, tplBase: tplDir}
}

type ServicesPageVM struct {
	Title     string
	Query     string // значение поискового input (сохранится)
	OrderID   string
	CartCount int
	Services  []repository.Service
}

type ServicePageVM struct {
	Title   string
	OrderID string
	Service repository.Service
}

type OrderPageVM struct {
	Title   string
	OrderID string
	Order   repository.Order
	// для удобства можно подтянуть данные услуг по id
	ServicesMap map[int]repository.Service
}

func (h *Handler) ServicesPage(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	oid := r.URL.Query().Get("order_id")
	if oid == "" {
		oid = h.Repo.Order.ID // по умолчанию
	}
	vm := ServicesPageVM{
		Title:     "Услуги — PANKREATITMED",
		Query:     q,
		OrderID:   oid,
		CartCount: h.Repo.CartCount(),
		Services:  h.Repo.ListServices(q),
	}
	h.render(w, "services.html", vm)
}

func (h *Handler) ServicePage(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	oid := r.URL.Query().Get("order_id")
	if oid == "" {
		oid = h.Repo.Order.ID
	}
	id, _ := strconv.Atoi(idStr)
	svc := h.Repo.GetServiceByID(id)
	if svc == nil {
		http.NotFound(w, r)
		return
	}
	vm := ServicePageVM{
		Title:   svc.Name + " — PANKREATITMED",
		OrderID: oid,
		Service: *svc,
	}
	h.render(w, "service.html", vm)
}

func (h *Handler) OrderPage(w http.ResponseWriter, r *http.Request) {
	oid := r.URL.Query().Get("id")
	if oid == "" {
		oid = h.Repo.Order.ID
	}
	// карта услуг по id для быстрых подписей/картинок в шаблоне
	mp := map[int]repository.Service{}
	for _, s := range h.Repo.Services {
		mp[s.ID] = s
	}
	vm := OrderPageVM{
		Title:       "Заявка — PANKREATITMED",
		OrderID:     oid,
		Order:       h.Repo.Order,
		ServicesMap: mp,
	}
	h.render(w, "order.html", vm)
}

func (h *Handler) render(w http.ResponseWriter, name string, data any) {
	fp := filepath.Join(h.tplBase, name)
	tpl := template.Must(template.ParseFiles(fp))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = tpl.Execute(w, data)
}
