package api

import (
	"github.com/angdev/chocolat/lib/query"
	"github.com/angdev/chocolat/model"
	"github.com/angdev/chocolat/support/repo"
)

func count(p *model.Project, params *QueryParams) (interface{}, error) {
	r := repo.NewRepository(p.RepoName())
	defer r.Close()

	q := query.New(r.C(params.CollectionName), params.ToQuery().Count())

	presenter := NewPresenter(q)

	if params.Interval.IsGiven() {
		return presenter.PresentInterval(&params.TimeFrame, &params.Interval)
	} else {
		return presenter.Present()
	}
}

func countUnique(p *model.Project, target string, params *QueryParams) (interface{}, error) {
	r := repo.NewRepository(p.RepoName())
	defer r.Close()

	q := query.New(r.C(params.CollectionName), params.ToQuery().CountUnique(target))

	presenter := NewPresenter(q)

	if params.Interval.IsGiven() {
		return presenter.PresentInterval(&params.TimeFrame, &params.Interval)
	} else {
		return presenter.Present()
	}
}

func min(p *model.Project, target string, params *QueryParams) (interface{}, error) {
	r := repo.NewRepository(p.RepoName())
	defer r.Close()

	q := query.New(r.C(params.CollectionName), params.ToQuery().Min(target))

	presenter := NewPresenter(q)

	if params.Interval.IsGiven() {
		return presenter.PresentInterval(&params.TimeFrame, &params.Interval)
	} else {
		return presenter.Present()
	}
}

func max(p *model.Project, target string, params *QueryParams) (interface{}, error) {
	r := repo.NewRepository(p.RepoName())
	defer r.Close()

	q := query.New(r.C(params.CollectionName), params.ToQuery().Max(target))

	presenter := NewPresenter(q)

	if params.Interval.IsGiven() {
		return presenter.PresentInterval(&params.TimeFrame, &params.Interval)
	} else {
		return presenter.Present()
	}
}

func sum(p *model.Project, target string, params *QueryParams) (interface{}, error) {
	r := repo.NewRepository(p.RepoName())
	defer r.Close()

	q := query.New(r.C(params.CollectionName), params.ToQuery().Sum(target))

	presenter := NewPresenter(q)

	if params.Interval.IsGiven() {
		return presenter.PresentInterval(&params.TimeFrame, &params.Interval)
	} else {
		return presenter.Present()
	}
}

func average(p *model.Project, target string, params *QueryParams) (interface{}, error) {
	r := repo.NewRepository(p.RepoName())
	defer r.Close()

	q := query.New(r.C(params.CollectionName), params.ToQuery().Average(target))

	presenter := NewPresenter(q)

	if params.Interval.IsGiven() {
		return presenter.PresentInterval(&params.TimeFrame, &params.Interval)
	} else {
		return presenter.Present()
	}
}

func percentile(p *model.Project, target string, percent float64, params *QueryParams) (interface{}, error) {
	r := repo.NewRepository(p.RepoName())
	defer r.Close()

	q := query.New(r.C(params.CollectionName), params.ToQuery().
		OrderBy(&query.Order{Field: target, Order: query.ASC}).Collect(target))
	var results []queryGroupResult

	if err := q.Execute(&results); err != nil {
		return nil, err
	}

	for i, _ := range results {
		result := results[i].Result.([]interface{})
		offset := int(float64(len(result)) * percent / 100)
		results[i].Result = result[offset]
	}

	return results, nil
}

func median(p *model.Project, target string, params *QueryParams) (interface{}, error) {
	return percentile(p, target, 50, params)
}

func selectUnique(p *model.Project, target string, params *QueryParams) (interface{}, error) {
	// Unique w/ no counting
	return nil, nil
}
