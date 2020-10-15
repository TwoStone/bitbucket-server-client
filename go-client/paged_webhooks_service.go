package bitbucket

import "context"

func (a WebhookApiService) GetWebhooks(ctx context.Context, projectKey, repositorySlug string) ([]Webhook, error) {
	return a.GetWebhooksPaged(ctx, projectKey, repositorySlug).GetAllPages()
}

func (r apiGetWebhooksPagedRequest) GetAllPages() ([]Webhook, error) {
	webhooks := []Webhook{}

	r.limitOrDefault()

	start := int32(0)
	morePages := true

	for morePages {
		result, _, err := r.Start(start).Execute()
		if err != nil {
			return nil, err
		}

		webhooks = append(webhooks, result.Values...)

		morePages = !result.IsLastPage
		if morePages {
			start = *result.NextPageStart
		}
	}

	return webhooks, nil
}

func (r apiGetWebhooksPagedRequest) limitOrDefault() {
	if r.limit == nil {
		r.limit = PtrInt32(DefaultPageLimit)
	}
}
