package service

import (
	"github.com/angdev/chocolat/model"
	"github.com/angdev/chocolat/support/repo"
)

type CountParams struct {
	QueryParams
}

func Count(p *model.Project, params *CountParams) (repo.Doc, error) {
	r := repo.NewRepository(p.RepoName())
	defer r.Close()

	pipes := []repo.Doc{}
	pipes = appendPipe(pipes, params.TimeFrame.Pipe())
	pipes = appendPipe(pipes, params.Filters.Pipe())
	pipes = appendPipe(pipes, params.GroupBy.Pipe(countOp()))
	pipes = appendPipe(pipes, countProject(params.GroupBy))

	pipe := r.C(params.CollectionName).Pipe(pipes)
	iter := pipe.Iter()

	var result []repo.Doc
	var d repo.Doc
	for iter.Next(&d) {
		result = append(result, collapseField(d))
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	if params.GroupBy != nil {
		return repo.Doc{"result": result}, nil
	} else if len(result) == 0 {
		return repo.Doc{"result": 0}, nil
	} else {
		return result[0], nil
	}
}

func appendPipe(pipes []repo.Doc, pipe repo.Doc) []repo.Doc {
	if pipe != nil {
		return append(pipes, pipe)
	} else {
		return pipes
	}
}

func countOp() repo.Doc {
	return repo.Doc{
		"count": repo.Doc{"$sum": 1},
	}
}

func countProject(g GroupBy) repo.Doc {
	project := repo.Doc{}
	for _, field := range g {
		project[field] = variablize("_id", field)
	}
	project["_id"] = false
	project["result"] = variablize("count")

	return repo.Doc{
		"$project": project,
	}
}
