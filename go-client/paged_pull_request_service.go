package bitbucket

func (r apiGetPullRequestsPagedRequest) GetAllPages() ([]PullRequest, error) {
	pullRequests := []PullRequest{}

	if r.limit != nil {
		r.limit = PtrInt32(DefaultPageLimit)
	}

	start := int32(0)
	morePages := true

	for morePages {
		page, _, err := r.Start(start).Execute()
		if err != nil {
			return nil, err
		}

		pullRequests = append(pullRequests, page.Values...)

		morePages = !page.IsLastPage
		if morePages {
			start = *page.NextPageStart
		}
	}

	return pullRequests, nil
}
