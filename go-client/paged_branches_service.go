package bitbucket

import "context"

func (a *BranchesApiService) GetBranches(ctx context.Context, projectKey string, repositorySlug string) ([]Branch, error) {
	return a.GetBranchesPaged(ctx, projectKey, repositorySlug).GetAllPages()
}

func (r apiGetBranchesPagedRequest) GetAllPages() ([]Branch, error) {
	branches := []Branch{}

	r.limitOrDefault()

	start := int32(0)
	morePages := true

	for morePages {
		result, _, err := r.Start(start).Execute()
		if err != nil {
			return nil, err
		}

		branches = append(branches, result.Values...)

		morePages = !result.IsLastPage
		if morePages {
			start = *result.NextPageStart
		}
	}

	return branches, nil
}

func (r apiGetBranchesPagedRequest) limitOrDefault() {
	if r.limit == nil {
		r.limit = PtrInt32(DefaultPageLimit)
	}
}
