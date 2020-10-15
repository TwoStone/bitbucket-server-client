package bitbucket

import "context"

func (a CommitsApiService) GetCommits(ctx context.Context, projectKey string, repositorySlug string) ([]Commit, error) {
	return a.GetCommitsPaged(ctx, projectKey, repositorySlug).GetAllPages()
}

func (r apiGetCommitsPagedRequest) GetAllPages() ([]Commit, error) {
	commits := []Commit{}

	r.limitOrDefault()

	start := int32(0)
	morePages := true

	for morePages {
		result, _, err := r.Start(start).Execute()
		if err != nil {
			return nil, err
		}

		commits = append(commits, result.Values...)

		morePages = !result.IsLastPage
		if morePages {
			start = *result.NextPageStart
		}
	}

	return commits, nil
}

func (r apiGetCommitsPagedRequest) limitOrDefault() {
	if r.limit == nil {
		r.limit = PtrInt32(DefaultPageLimit)
	}
}
