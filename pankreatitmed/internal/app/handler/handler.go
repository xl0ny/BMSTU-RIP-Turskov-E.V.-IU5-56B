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

type CriteriaPageVM struct {
	Title     string
	Query     string
	OrderID   string
	CartCount int
	Criteria  []repository.Criterion
}

type CriterionPageVM struct {
	Title     string
	OrderID   string
	Criterion repository.Criterion
}

type OrderPageVM struct {
	Title       string
	OrderID     string
	Order       repository.Order
	CriteriaMap map[int]repository.Criterion
}

func (h *Handler) CriteriaPage(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	oid := r.URL.Query().Get("order_id")
	if oid == "" {
		oid = h.Repo.Order.ID
	}
	vm := CriteriaPageVM{
		Title:     "Услуги — PANKREATITMED",
		Query:     q,
		OrderID:   oid,
		CartCount: h.Repo.CartCount(),
		Criteria:  h.Repo.ListCriteria(q),
	}
	h.render(w, "criteria.html", vm)
}

func (h *Handler) CriterionPage(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	oid := r.URL.Query().Get("order_id")
	if oid == "" {
		oid = h.Repo.Order.ID
	}
	id, _ := strconv.Atoi(idStr)
	svc := h.Repo.GetCriterionByID(id)
	if svc == nil {
		http.NotFound(w, r)
		return
	}
	vm := CriterionPageVM{
		Title:     svc.Name + " — PANKREATITMED",
		OrderID:   oid,
		Criterion: *svc,
	}
	h.render(w, "criterion.html", vm)
}

func (h *Handler) OrderPage(w http.ResponseWriter, r *http.Request) {
	oid := r.URL.Query().Get("id")
	if oid == "" {
		oid = h.Repo.Order.ID
	}
	mp := map[int]repository.Criterion{}
	for _, s := range h.Repo.Criteria {
		mp[s.ID] = s
	}
	vm := OrderPageVM{
		Title:       "Заявка — PANKREATITMED",
		OrderID:     oid,
		Order:       h.Repo.Order,
		CriteriaMap: mp,
	}
	h.render(w, "order.html", vm)
}

func (h *Handler) render(w http.ResponseWriter, name string, data any) {
	fp := filepath.Join(h.tplBase, name)
	tpl := template.Must(template.ParseFiles(fp))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = tpl.Execute(w, data)
}
