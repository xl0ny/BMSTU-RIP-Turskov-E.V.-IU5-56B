package repository

import (
	"fmt"
	"strings"
)

type Criterion struct {
	ID          int
	Code        string
	Name        string
	Indicator   string
	Duration    string
	HomeVisit   bool
	ImageKey    string
	ImageURL    string
	Description string
}

type OrderItem struct {
	Index       int
	CriterionID int
	MMField     string
}

type Order struct {
	ID             string
	Items          []OrderItem
	ComputedResult string
}

type Repository struct {
	Criteria     []Criterion
	Order        Order
	MinioBaseURL string
}

func NewRepository(minioBaseURL string) *Repository {
	r := &Repository{
		MinioBaseURL: minioBaseURL,
	}

	r.Criteria = []Criterion{
		{1, "№1", "Оценка возраста пациента", "> 55 лет", "1 календарный день", true, "n1_age.png", "", "Возраст > 55 лет — критерий при поступлении."},
		{2, "№2", "Анализ лейкоцитов крови", "> 16 000/мм³", "1 календарный день", true, "n2_wbc.png", "", "Повышение лейкоцитов может указывать на выраженный воспалительный процесс."},
		{3, "№3", "Измерение уровня глюкозы", "> 200 мг/дл (11,1 ммоль/л)", "1 календарный день", true, "n3_glucose.png", "", "Гипергликемия — один из ранних критериев."},
		{4, "№4", "Определение уровня ЛДГ", "> 350 Ед/л", "1 календарный день", true, "n4_ldh.png", "", "ЛДГ > 350 МЕ/л — критерий тяжести."},
		{5, "№5", "Анализ активности АСТ", "> 250 Ед/л", "1 календарный день", true, "n5_ast.png", "", "АСТ > 250 МЕ/л — критерий тяжести."},
		{6, "№6", "Контроль изменения гематокрита", "Падение > 10% (за 48 ч)", "1 календарный день через 48 часов", true, "n6_hct.png", "", "Снижение гематокрита в динамике — неблагоприятный признак."},
		{7, "№7", "Измерение уровня мочевины (BUN)", "Повышение > 5 мг/дл", "1 календарный день через 48 часов", true, "n7_bun.png", "", "Рост мочевины указывает на ухудшение."},
		{8, "№8", "Измерение уровня кальция сыворотки", "< 8,0 мг/дл (2,0 ммоль/л)", "1 календарный день", true, "n8_ca.png", "", "Гипокальциемия — прогностический критерий."},
		{9, "№9", "Измерение PaO₂", "< 60 мм рт.ст.", "1 календарный день", true, "n9_pao2.png", "", "PaO₂ < 60 мм рт.ст. — критерий в шкале Рэнсона."},
		{10, "№10", "Оценка кислотно-щелочного состояния", "Дефицит оснований > 4 мЭкв/л", "1 календарный день через 48 часов", true, "n10_acidbase.png", "", "Декомпенсированный ацидоз — неблагоприятен."},
		{11, "№11", "Оценка объёма секвестрированной жидкости", "> 6 л", "1 календарный день через 48 часов", true, "n11_sequestration.png", "", "Большой объём — высокий риск осложнений."},
	}
	for i := range r.Criteria {
		r.Criteria[i].ImageURL = fmt.Sprintf("%s/%s", strings.TrimRight(minioBaseURL, "/"), r.Criteria[i].ImageKey)
	}
	// демо заявка моя
	r.Order = Order{
		ID: "0001",
		Items: []OrderItem{
			{Index: 1, CriterionID: 2, MMField: "5128/мкл"},
			{Index: 2, CriterionID: 9, MMField: "75 мм рт. ст."},
		},
		ComputedResult: "Ваш балл по шкале Рэнсона — 5. Летальный исход — 40%",
	}

	return r
}

func (r *Repository) ListCriteria(query string) []Criterion {
	if strings.TrimSpace(query) == "" {
		return r.Criteria
	}
	q := strings.ToLower(strings.TrimSpace(query))

	//// если трактовать как цену
	//for _, s := range r.CriteriaMap {
	//	if fmt.Sprintf("%d", s.Indicator) == q {
	//		return []Indicator{s}
	//	}
	//}

	// иначе поиск по подстроке имени
	res := make([]Criterion, 0)
	for _, s := range r.Criteria {
		if strings.Contains(strings.ToLower(s.Name), q) {
			res = append(res, s)
		}
	}
	return res
}

func (r *Repository) GetCriterionByID(id int) *Criterion {
	for i := range r.Criteria {
		if r.Criteria[i].ID == id {
			return &r.Criteria[i]
		}
	}
	return nil
}

func (r *Repository) CartCount() int {
	return len(r.Order.Items)
}
