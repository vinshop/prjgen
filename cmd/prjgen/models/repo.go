package models

type Repo struct {
	Base
	Model *Model
}

func NewRepo(model Model, dir string) Repo {
	return Repo{
		Base: Base{
			Pkg: model.Pkg,
			Dir: model.Dir + "/" + dir,
		},
		Model: &model,
	}
}
