package repository

import (
	"fmt"
	"strings"
)

type Service struct {
	ID          int
	Code        string // №1..№11 (для отображения)
	Name        string
	PriceRUB    int    // 0..2500 (как на макете)
	Duration    string // "1 календарный день"
	HomeVisit   bool   // "Доступно с выездом на дом"
	ImageKey    string // ключ объекта в MinIO, латиницей, напр. "n2_wbc.png"
	ImageURL    string // конечный URL (формируется из MINIO_PUBLIC_URL + "/" + ImageKey)
	Description string // большой текст для карточки (деталка)
}

type OrderItem struct {
	Index     int // порядковый № в заявке (1,2,3..)
	ServiceID int
	MMField   string // поле «м-м» (значение/ед.изм. и т.п., например "5128/мкл")
}

type Order struct {
	ID             string
	Items          []OrderItem // (в лабораторной допустимо хранить тут массив)
	ComputedResult string      // «Ваш балл по шкале Рэнсона — ...»
}

type Repository struct {
	Services     []Service
	Order        Order
	MinioBaseURL string // например: http://localhost:9000/pankreatitmed
}

func NewRepository(minioBaseURL string) *Repository {
	r := &Repository{
		MinioBaseURL: minioBaseURL,
	}

	// 11 услуг (реальные названия с твоего макета; цены — как на скринах)
	r.Services = []Service{
		{1, "№1", "Оценка возраста пациента", 0, "1 календарный день", true, "n1_age.png", "", "Возраст > 55 лет — критерий при поступлении."},
		{2, "№2", "Анализ лейкоцитов крови", 400, "1 календарный день", true, "n2_wbc.png", "", "Повышение лейкоцитов может указывать на выраженный воспалительный процесс."},
		{3, "№3", "Измерение уровня глюкозы", 350, "1 календарный день", true, "n3_glucose.png", "", "Гипергликемия — один из ранних критериев."},
		{4, "№4", "Определение уровня ЛДГ", 600, "1 календарный день", true, "n4_ldh.png", "", "ЛДГ > 350 МЕ/л — критерий тяжести."},
		{5, "№5", "Анализ активности АСТ", 500, "1 календарный день", true, "n5_ast.png", "", "АСТ > 250 МЕ/л — критерий тяжести."},
		{6, "№6", "Контроль изменения гематокрита", 300, "1 календарный день через 48 часов", true, "n6_hct.png", "", "Снижение гематокрита в динамике — неблагоприятный признак."},
		{7, "№7", "Измерение уровня мочевины (BUN)", 450, "1 календарный день через 48 часов", true, "n7_bun.png", "", "Рост мочевины указывает на ухудшение."},
		{8, "№8", "Измерение уровня кальция сыворотки", 400, "1 календарный день", true, "n8_ca.png", "", "Гипокальциемия — прогностический критерий."},
		{9, "№9", "Измерение PaO₂", 700, "1 календарный день", true, "n9_pao2.png", "", "PaO₂ < 60 мм рт.ст. — критерий в шкале Рэнсона."},
		{10, "№10", "Оценка кислотно-щелочного состояния", 650, "1 календарный день через 48 часов", true, "n10_acidbase.png", "", "Декомпенсированный ацидоз — неблагоприятен."},
		{11, "№11", "Оценка объёма секвестрированной жидкости", 2500, "1 календарный день через 48 часов", true, "n11_sequestration.png", "", "Большой объём — высокий риск осложнений."},
	}

	// сформировать публичные URL картинок из MinIO
	for i := range r.Services {
		r.Services[i].ImageURL = fmt.Sprintf("%s/%s", strings.TrimRight(minioBaseURL, "/"), r.Services[i].ImageKey)
	}

	// «демо»-заявка заранее заполнена 2 услугами (редактирования в ЛР1 нет)
	r.Order = Order{
		ID: "R-0001",
		Items: []OrderItem{
			{Index: 1, ServiceID: 2, MMField: "5128/мкл"},
			{Index: 2, ServiceID: 9, MMField: "75 мм рт. ст."},
		},
		ComputedResult: "Ваш балл по шкале Рэнсона — 5. Летальный исход — 40%",
	}

	return r
}

// Фильтрация по имени (или точной цене) — серверная
func (r *Repository) ListServices(query string) []Service {
	if strings.TrimSpace(query) == "" {
		return r.Services
	}
	q := strings.ToLower(strings.TrimSpace(query))

	// трактуем как цену
	for _, s := range r.Services {
		if fmt.Sprintf("%d", s.PriceRUB) == q {
			return []Service{s}
		}
	}

	// иначе — поиск по подстроке имени
	res := make([]Service, 0)
	for _, s := range r.Services {
		if strings.Contains(strings.ToLower(s.Name), q) {
			res = append(res, s)
		}
	}
	return res
}

func (r *Repository) GetServiceByID(id int) *Service {
	for i := range r.Services {
		if r.Services[i].ID == id {
			return &r.Services[i]
		}
	}
	return nil
}

func (r *Repository) CartCount() int {
	return len(r.Order.Items)
}
