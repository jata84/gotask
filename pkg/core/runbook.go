package core

type RunBook struct {
	project       *Project
	run_book_list map[string][]string
}

func NewRunBook(project *Project) *RunBook {
	return &RunBook{
		project:       project,
		run_book_list: make(map[string][]string),
	}
}

func (rb *RunBook) AddPipeLine(name string, pipeline_list []string) {
	rb.run_book_list[name] = pipeline_list
}

func (rb *RunBook) Run() {
	for _, pipeline := range rb.run_book_list {
		for _, p := range pipeline {
			rb.project.pipelines.pipeline_list.Get(p).Run()
		}

	}
}
