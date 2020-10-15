package bitbucket

import "context"

func (a *ProjectsApiService) GetProjects(ctx context.Context) ([]Project, error) {
	return a.GetProjectsPaged(ctx).GetAllPages()
}

func (r apiGetProjectsPagedRequest) GetAllPages() ([]Project, error) {
	projects := []Project{}

	r.limitOrDefault()

	start := int32(0)
	morePages := true

	for morePages {
		result, _, err := r.Start(start).Execute()
		if err != nil {
			return nil, err
		}

		projects = append(projects, result.Values...)

		morePages = !result.IsLastPage
		if morePages {
			start = *result.NextPageStart
		}
	}

	return projects, nil
}

func (r apiGetProjectsPagedRequest) limitOrDefault() {
	if r.limit == nil {
		r.limit = PtrInt32(DefaultPageLimit)
	}
}
