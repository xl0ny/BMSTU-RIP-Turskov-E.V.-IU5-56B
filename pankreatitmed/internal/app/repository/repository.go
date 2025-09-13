package repository

import (
	"fmt"
	"strings"
)

type Service struct {
	ID          int
	Code        string
	Name        string
	PriceRUB    int
	Duration    string
	HomeVisit   bool
	ImageKey    string
	ImageURL    string
	Description string
}

type OrderItem struct {
	Index     int
	ServiceID int
	MMField   string
}

type Order struct {
	ID             string
	Items          []OrderItem
	ComputedResult string
}

type Repository struct {
	Services     []Service
	Order        Order
	MinioBaseURL string
}

func NewRepository(minioBaseURL string) *Repository {
	r := &Repository{
		MinioBaseURL: minioBaseURL,
	}

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

	for i := range r.Services {
		r.Services[i].ImageURL = fmt.Sprintf("%s/%s", strings.TrimRight(minioBaseURL, "/"), r.Services[i].ImageKey)
	}
	// демо заявка моя
	r.Order = Order{
		ID: "0001",
		Items: []OrderItem{
			{Index: 1, ServiceID: 2, MMField: "5128/мкл"},
			{Index: 2, ServiceID: 9, MMField: "75 мм рт. ст."},
		},
		ComputedResult: "Ваш балл по шкале Рэнсона — 5. Летальный исход — 40%",
	}

	return r
}

func (r *Repository) ListServices(query string) []Service {
	if strings.TrimSpace(query) == "" {
		return r.Services
	}
	q := strings.ToLower(strings.TrimSpace(query))

	//// если трактовать как цену
	//for _, s := range r.Services {
	//	if fmt.Sprintf("%d", s.PriceRUB) == q {
	//		return []Service{s}
	//	}
	//}

	// иначе поиск по подстроке имени
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
