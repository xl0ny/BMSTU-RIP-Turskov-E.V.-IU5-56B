package mapper

import (
	"pankreatitmed/internal/app/ds"
	"pankreatitmed/internal/app/dto/request"
	"pankreatitmed/internal/app/dto/response"
)

// client->handler
func CreateCriterionToCriterion(crit request.CreateCriterion) ds.Criterion {
	return ds.Criterion{
		Code:        crit.Code,
		Name:        crit.Name,
		Description: crit.Description,
		Duration:    crit.Duration,
		HomeVisit:   crit.HomeVisit,
		Status:      crit.Status,
		Unit:        crit.Unit,
		RefLow:      &crit.RefLow,
		RefHigh:     &crit.RefLow,
	}
}

// handler->client
func CritertionToSendCriterion(crit ds.Criterion) response.SendCriterion {
	return response.SendCriterion{
		ID:          crit.ID,
		Code:        crit.Code,
		Name:        crit.Name,
		Description: crit.Description,
		Duration:    crit.Duration,
		HomeVisit:   crit.HomeVisit,
		ImageURL:    crit.ImageURL,
		Status:      crit.Status,
		Unit:        crit.Unit,
		RefLow:      crit.RefLow,
		RefHigh:     crit.RefLow,
	}
}

func CriterionsToSendCrtierions(crits []ds.Criterion) []response.SendCriterion {
	list := make([]response.SendCriterion, len(crits))
	for i, c := range crits {
		list[i] = CritertionToSendCriterion(c)
	}
	return list
}
