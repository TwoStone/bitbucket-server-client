package bitbucket

import "context"

func (a *RepositoriesApiService) SearchRepositories(ctx context.Context) ([]Repository, error) {
	return a.SearchRepositoriesPaged(ctx).GetAllPages()
}

func (r apiSearchRepositoriesPagedRequest) GetAllPages() ([]Repository, error) {
	repositories := []Repository{}

	r.limitOrDefault()

	start := int32(0)
	morePages := true

	for morePages {
		result, _, err := r.Start(start).Execute()
		if err != nil {
			return nil, err
		}

		repositories = append(repositories, result.Values...)

		morePages = !result.IsLastPage
		if morePages {
			start = *result.NextPageStart
		}
	}

	return repositories, nil
}

func (r apiSearchRepositoriesPagedRequest) limitOrDefault() {
	if r.limit == nil {
		r.limit = PtrInt32(DefaultPageLimit)
	}
}

func (a *RepositoriesApiService) GetRepositories(ctx context.Context, projectKey string) ([]Repository, error) {
	return a.GetRepositoriesPaged(ctx, projectKey).GetAllPages()
}

func (r apiGetRepositoriesPagedRequest) GetAllPages() ([]Repository, error) {
	repositories := []Repository{}

	r.limitOrDefault()

	start := int32(0)
	morePages := true

	for morePages {
		result, _, err := r.Start(start).Execute()
		if err != nil {
			return nil, err
		}

		repositories = append(repositories, result.Values...)

		morePages = !result.IsLastPage
		if morePages {
			start = *result.NextPageStart
		}
	}

	return repositories, nil
}

func (r apiGetRepositoriesPagedRequest) limitOrDefault() {
	if r.limit == nil {
		r.limit = PtrInt32(DefaultPageLimit)
	}
}
