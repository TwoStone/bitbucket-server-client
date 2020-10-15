/*
 * bitbucket-server-api
 *
 * <h1>REST Resources Provided By: Bitbucket Server - REST</h1> <p>     This is the reference document for the Atlassian Bitbucket REST API. The REST API is for developers who want to: </p> <ul>     <li>integrate Bitbucket with other applications;</li>     <li>create scripts that interact with Bitbucket; or</li>     <li>develop plugins that enhance the Bitbucket UI, using REST to interact with the backend.</li> </ul> You can read more about developing Bitbucket plugins in the <a href=\"https://developer.atlassian.com/server/bitbucket/\">Bitbucket Developer Documentation</a>. <p></p> <h2 id=\"gettingstarted\">Getting started</h2> <p>     Because the REST API is based on open standards, you can use any web development language or command line tool     capable of generating an HTTP request to access the API. See the     <a href=\"https://developer.atlassian.com/server/bitbucket/reference/rest-api/\">developer documentation</a> for a     basic     usage example. </p> <p>     If you're already working with the     <a href=\"https://developer.atlassian.com/server/framework/atlassian-sdk/\">Atlassian SDK</a>,     the <a href=\"https://developer.atlassian.com/server/framework/atlassian-sdk/using-the-rest-api-browser/\">REST API         Browser</a> is a great tool for exploring and experimenting with the Bitbucket REST API. </p> <h2>     <a name=\"StructureoftheRESTURIs\"></a>Structure of the REST URIs</h2> <p>     Bitbucket's REST APIs provide access to resources (data entities) via URI paths. To use a REST API, your application     will     make an HTTP request and parse the response. The Bitbucket REST API uses JSON as its communication format, and the     standard     HTTP methods like GET, PUT, POST and DELETE. URIs for Bitbucket's REST API resource have the following structure: </p> <pre>    http://host:port/context/rest/api-name/api-version/path/to/resource </pre> <p>     For example, the following URI would retrieve a page of the latest commits to the <strong>jira</strong> repository     in     the <strong>Jira</strong> project on <a href=\"https://stash.atlassian.com\">https://stash.atlassian.com</a>. </p> <pre>    https://stash.atlassian.com/rest/api/1.0/projects/JIRA/repos/jira/commits </pre> <p>     See the API descriptions below for a full list of available resources. </p> <p>     Alternatively we also publish a list of resources in     <a href=\"http://en.wikipedia.org/wiki/Web_Application_Description_Language\">WADL</a> format. It is available     <a href=\"bitbucket-rest.wadl\">here</a>. </p> <h2 id=\"paging-params\">Paged APIs</h2> <p>     Bitbucket uses paging to conserve server resources and limit response size for resources that return potentially     large     collections of items. A request to a paged API will result in a <code>values</code> array wrapped in a JSON object     with some paging metadata, like this: </p> <pre>    {         \"size\": 3,         \"limit\": 3,         \"isLastPage\": false,         \"values\": [             { /_* result 0 *_/ },             { /_* result 1 *_/ },             { /_* result 2 *_/ }         ],         \"start\": 0,         \"filter\": null,         \"nextPageStart\": 3     } </pre> <p>     Clients can use the <code>limit</code> and <code>start</code> query parameters to retrieve the desired number of     results. </p> <p>     The <code>limit</code> parameter indicates how many results to return per page. Most APIs default to returning     <code>25</code> if the limit is left unspecified. This number can be increased, but note that a resource-specific     hard limit will apply. These hard limits can be configured by server administrators, so it's always best practice to     check the <code>limit</code> attribute on the response to see what limit has been applied.     The request to get a larger page should look like this: </p> <pre>    http://host:port/context/rest/api-name/api-version/path/to/resource?limit={desired size of page} </pre> <p>     For example: </p> <pre>    https://stash.atlassian.com/rest/api/1.0/projects/JIRA/repos/jira/commits?limit=1000 </pre> <p>     The <code>start</code> parameter indicates which item should be used as the first item in the page of results. All     paged responses contain an <code>isLastPage</code> attribute indicating whether another page of items exists. </p> <p><strong>Important:</strong> If more than one page exists (i.e. the response contains     <code>\"isLastPage\": false</code>), the response object will also contain a <code>nextPageStart</code> attribute     which <strong><em>must</em></strong> be used by the client as the <code>start</code> parameter on the next request.     Identifiers of adjacent objects in a page may not be contiguous, so the start of the next page is <em>not</em>     necessarily the start of the last page plus the last page's size. A client should always use     <code>nextPageStart</code> to avoid unexpected results from a paged API.     The request to get a subsequent page should look like this: </p> <pre>    http://host:port/context/rest/api-name/api-version/path/to/resource?start={nextPageStart from previous response} </pre> <p>     For example: </p> <pre>    https://stash.atlassian.com/rest/api/1.0/projects/JIRA/repos/jira/commits?start=25 </pre> <h2 id=\"authentication\">Authentication</h2> <p>     Any authentication that works against Bitbucket will work against the REST API. <b>The preferred authentication         methods         are HTTP Basic (when using SSL) and OAuth</b>. Other supported methods include: HTTP Cookies and Trusted     Applications. </p> <p>     You can find OAuth code samples in several programming languages at     <a         href=\"https://bitbucket.org/atlassian_tutorial/atlassian-oauth-examples\">bitbucket.org/atlassian_tutorial/atlassian-oauth-examples</a>. </p> <p>     The log-in page uses cookie-based authentication, so if you are using Bitbucket in a browser you can call REST from     JavaScript on the page and rely on the authentication that the browser has established. </p> <h2 id=\"errors-and-validation\">Errors &amp; Validation</h2> <p>     If a request fails due to client error, the resource will return an HTTP response code in the 40x range. These can     be broadly categorised into: </p> <table>     <tbody>         <tr>             <th>HTTP Code</th>             <th>Description</th>         </tr>         <tr>             <td>400 (Bad Request)</td>             <td>                 One or more of the required parameters or attributes:                 <ul>                     <li>were missing from the request;</li>                     <li>incorrectly formatted; or</li>                     <li>inappropriate in the given context.</li>                 </ul>             </td>         </tr>         <tr>             <td>401 (Unauthorized)</td>             <td>                 Either:                 <ul>                     <li>Authentication is required but was not attempted.</li>                     <li>Authentication was attempted but failed.</li>                     <li>Authentication was successful but the authenticated user does not have the requisite permission                         for the resource.</li>                 </ul>                 See the individual resource documentation for details of required permissions.             </td>         </tr>         <tr>             <td>403 (Forbidden)</td>             <td>                 Actions are usually \"forbidden\" if they involve breaching the licensed user limit of the server, or                 degrading the authenticated user's permission level. See the individual resource documentation for more                 details.             </td>         </tr>         <tr>             <td>404 (Not Found)</td>             <td>                 The entity you are attempting to access, or the project or repository containing it, does not exist.             </td>         </tr>         <tr>             <td>405 (Method Not Allowed)</td>             <td>                 The request HTTP method is not appropriate for the targeted resource. For example an HTTP GET to a                 resource that only accepts an HTTP POST will result in a 405.             </td>         </tr>         <tr>             <td>409 (Conflict)</td>             <td>                 The attempted update failed due to some conflict with an existing resource. For example:                 <ul>                     <li>Creating a project with a key that already exists</li>                     <li>Merging an out-of-date pull request</li>                     <li>Deleting a comment that has replies</li>                     <li>etc.</li>                 </ul>                 See the individual resource documentation for more details.             </td>         </tr>         <tr>             <td>415 (Unsupported Media Type)</td>             <td>                 The request entity has a <code>Content-Type</code> that the server does not support. Almost all of the                 Bitbucket REST API accepts <code>application/json</code> format, but check the individual resource                 documentation for more details. Additionally, double-check that you are setting the                 <code>Content-Type</code> header correctly on your request (e.g. using                 <code>-H \"Content-Type: application/json\"</code> in cURL).             </td>         </tr>     </tbody> </table> <p>     For <strong>400</strong> HTTP codes the response will typically contain one or more validation error messages,     for example: </p> <pre>    {         \"errors\": [             {                 \"context\": \"name\",                 \"message\": \"The name should be between 1 and 255 characters.\",                 \"exceptionName\": null             },             {                 \"context\": \"email\",                 \"message\": \"The email should be a valid email address.\",                 \"exceptionName\": null             }         ]     }     </pre> <p>     The <code>context</code> attribute indicates which parameter or request entity attribute failed validation. Note     that the <code>context</code> may be null. </p> <p>     For <strong>401</strong>, <strong>403</strong>, <strong>404</strong> and <strong>409</strong>     HTTP codes, the response will contain one or more descriptive error messages: </p> <pre>    {         \"errors\": [             {                 \"context\": null,                 \"message\": \"A detailed error message.\",                 \"exceptionName\": null             }         ]     }     </pre> <p>     A <strong>500</strong> (Server Error) HTTP code indicates an incorrect resource url or an unexpected server error.     Double-check the URL you are trying to access, then report the issue to your server administrator or     <a href=\"https://support.atlassian.com/\">Atlassian Support</a> if problems persist. </p> <h2 id=\"personal-repositories\">Personal Repositories</h2> <p>     Bitbucket allows users to manage their own repositories, called personal repositories. These are repositories     associated     with the user and to which they always have REPO_ADMIN permission. </p> <p>     Accessing personal repositories via REST is achieved through the normal project-centric REST URLs     using the user's slug prefixed by tilde as the project key. E.g. to list personal repositories for a     user with slug \"johnsmith\" you would make a GET to: </p> <pre>http://example.com/rest/api/1.0/projects/~johnsmith/repos</pre> <p></p> <p>     In addition to this, Bitbucket allows access to these repositories through an alternate set of user-centric REST     URLs     beginning with: </p> <pre>http://example.com/rest/api/1.0/users/~{userSlug}/repos</pre> E.g. to list the forks of the repository with slug \"nodejs\" in the personal project of user with slug \"johnsmith\" using the regular REST URL you would make a GET to: <pre>http://example.com/rest/api/1.0/projects/~johnsmith/repos/nodejs/forks</pre> Using the alternate URL, you would make a GET to: <pre>http://example.com/rest/api/1.0/users/johnsmith/repos/nodejs/forks</pre> <p></p>
 *
 * API version: 7.3.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bitbucket

import (
	"encoding/json"
)

// PostWebhook struct for PostWebhook
type PostWebhook struct {
	BranchCreated      *bool  `json:"branchCreated,omitempty"`
	BranchDeleted      *bool  `json:"branchDeleted,omitempty"`
	BranchesToIgnore   string `json:"branchesToIgnore"`
	BuildStatus        *bool  `json:"buildStatus,omitempty"`
	CommittersToIgnore string `json:"committersToIgnore"`
	Enabled            bool   `json:"enabled"`
	Id                 int32  `json:"id"`
	PrCommented        *bool  `json:"prCommented,omitempty"`
	PrCreated          *bool  `json:"prCreated,omitempty"`
	PrDeclined         *bool  `json:"prDeclined,omitempty"`
	PrMerged           *bool  `json:"prMerged,omitempty"`
	PrReopened         *bool  `json:"prReopened,omitempty"`
	PrRescoped         *bool  `json:"prRescoped,omitempty"`
	PrUpdated          *bool  `json:"prUpdated,omitempty"`
	RepoPush           *bool  `json:"repoPush,omitempty"`
	TagCreated         *bool  `json:"tagCreated,omitempty"`
	Title              string `json:"title"`
	Url                string `json:"url"`
}

// NewPostWebhook instantiates a new PostWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostWebhook(branchesToIgnore string, committersToIgnore string, enabled bool, id int32, title string, url string) *PostWebhook {
	this := PostWebhook{}
	this.BranchesToIgnore = branchesToIgnore
	this.CommittersToIgnore = committersToIgnore
	this.Enabled = enabled
	this.Id = id
	this.Title = title
	this.Url = url
	return &this
}

// NewPostWebhookWithDefaults instantiates a new PostWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostWebhookWithDefaults() *PostWebhook {
	this := PostWebhook{}
	return &this
}

// GetBranchCreated returns the BranchCreated field value if set, zero value otherwise.
func (o *PostWebhook) GetBranchCreated() bool {
	if o == nil || o.BranchCreated == nil {
		var ret bool
		return ret
	}
	return *o.BranchCreated
}

// GetBranchCreatedOk returns a tuple with the BranchCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetBranchCreatedOk() (*bool, bool) {
	if o == nil || o.BranchCreated == nil {
		return nil, false
	}
	return o.BranchCreated, true
}

// HasBranchCreated returns a boolean if a field has been set.
func (o *PostWebhook) HasBranchCreated() bool {
	if o != nil && o.BranchCreated != nil {
		return true
	}

	return false
}

// SetBranchCreated gets a reference to the given bool and assigns it to the BranchCreated field.
func (o *PostWebhook) SetBranchCreated(v bool) {
	o.BranchCreated = &v
}

// GetBranchDeleted returns the BranchDeleted field value if set, zero value otherwise.
func (o *PostWebhook) GetBranchDeleted() bool {
	if o == nil || o.BranchDeleted == nil {
		var ret bool
		return ret
	}
	return *o.BranchDeleted
}

// GetBranchDeletedOk returns a tuple with the BranchDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetBranchDeletedOk() (*bool, bool) {
	if o == nil || o.BranchDeleted == nil {
		return nil, false
	}
	return o.BranchDeleted, true
}

// HasBranchDeleted returns a boolean if a field has been set.
func (o *PostWebhook) HasBranchDeleted() bool {
	if o != nil && o.BranchDeleted != nil {
		return true
	}

	return false
}

// SetBranchDeleted gets a reference to the given bool and assigns it to the BranchDeleted field.
func (o *PostWebhook) SetBranchDeleted(v bool) {
	o.BranchDeleted = &v
}

// GetBranchesToIgnore returns the BranchesToIgnore field value
func (o *PostWebhook) GetBranchesToIgnore() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BranchesToIgnore
}

// GetBranchesToIgnoreOk returns a tuple with the BranchesToIgnore field value
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetBranchesToIgnoreOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchesToIgnore, true
}

// SetBranchesToIgnore sets field value
func (o *PostWebhook) SetBranchesToIgnore(v string) {
	o.BranchesToIgnore = v
}

// GetBuildStatus returns the BuildStatus field value if set, zero value otherwise.
func (o *PostWebhook) GetBuildStatus() bool {
	if o == nil || o.BuildStatus == nil {
		var ret bool
		return ret
	}
	return *o.BuildStatus
}

// GetBuildStatusOk returns a tuple with the BuildStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetBuildStatusOk() (*bool, bool) {
	if o == nil || o.BuildStatus == nil {
		return nil, false
	}
	return o.BuildStatus, true
}

// HasBuildStatus returns a boolean if a field has been set.
func (o *PostWebhook) HasBuildStatus() bool {
	if o != nil && o.BuildStatus != nil {
		return true
	}

	return false
}

// SetBuildStatus gets a reference to the given bool and assigns it to the BuildStatus field.
func (o *PostWebhook) SetBuildStatus(v bool) {
	o.BuildStatus = &v
}

// GetCommittersToIgnore returns the CommittersToIgnore field value
func (o *PostWebhook) GetCommittersToIgnore() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommittersToIgnore
}

// GetCommittersToIgnoreOk returns a tuple with the CommittersToIgnore field value
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetCommittersToIgnoreOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommittersToIgnore, true
}

// SetCommittersToIgnore sets field value
func (o *PostWebhook) SetCommittersToIgnore(v string) {
	o.CommittersToIgnore = v
}

// GetEnabled returns the Enabled field value
func (o *PostWebhook) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PostWebhook) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value
func (o *PostWebhook) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostWebhook) SetId(v int32) {
	o.Id = v
}

// GetPrCommented returns the PrCommented field value if set, zero value otherwise.
func (o *PostWebhook) GetPrCommented() bool {
	if o == nil || o.PrCommented == nil {
		var ret bool
		return ret
	}
	return *o.PrCommented
}

// GetPrCommentedOk returns a tuple with the PrCommented field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetPrCommentedOk() (*bool, bool) {
	if o == nil || o.PrCommented == nil {
		return nil, false
	}
	return o.PrCommented, true
}

// HasPrCommented returns a boolean if a field has been set.
func (o *PostWebhook) HasPrCommented() bool {
	if o != nil && o.PrCommented != nil {
		return true
	}

	return false
}

// SetPrCommented gets a reference to the given bool and assigns it to the PrCommented field.
func (o *PostWebhook) SetPrCommented(v bool) {
	o.PrCommented = &v
}

// GetPrCreated returns the PrCreated field value if set, zero value otherwise.
func (o *PostWebhook) GetPrCreated() bool {
	if o == nil || o.PrCreated == nil {
		var ret bool
		return ret
	}
	return *o.PrCreated
}

// GetPrCreatedOk returns a tuple with the PrCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetPrCreatedOk() (*bool, bool) {
	if o == nil || o.PrCreated == nil {
		return nil, false
	}
	return o.PrCreated, true
}

// HasPrCreated returns a boolean if a field has been set.
func (o *PostWebhook) HasPrCreated() bool {
	if o != nil && o.PrCreated != nil {
		return true
	}

	return false
}

// SetPrCreated gets a reference to the given bool and assigns it to the PrCreated field.
func (o *PostWebhook) SetPrCreated(v bool) {
	o.PrCreated = &v
}

// GetPrDeclined returns the PrDeclined field value if set, zero value otherwise.
func (o *PostWebhook) GetPrDeclined() bool {
	if o == nil || o.PrDeclined == nil {
		var ret bool
		return ret
	}
	return *o.PrDeclined
}

// GetPrDeclinedOk returns a tuple with the PrDeclined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetPrDeclinedOk() (*bool, bool) {
	if o == nil || o.PrDeclined == nil {
		return nil, false
	}
	return o.PrDeclined, true
}

// HasPrDeclined returns a boolean if a field has been set.
func (o *PostWebhook) HasPrDeclined() bool {
	if o != nil && o.PrDeclined != nil {
		return true
	}

	return false
}

// SetPrDeclined gets a reference to the given bool and assigns it to the PrDeclined field.
func (o *PostWebhook) SetPrDeclined(v bool) {
	o.PrDeclined = &v
}

// GetPrMerged returns the PrMerged field value if set, zero value otherwise.
func (o *PostWebhook) GetPrMerged() bool {
	if o == nil || o.PrMerged == nil {
		var ret bool
		return ret
	}
	return *o.PrMerged
}

// GetPrMergedOk returns a tuple with the PrMerged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetPrMergedOk() (*bool, bool) {
	if o == nil || o.PrMerged == nil {
		return nil, false
	}
	return o.PrMerged, true
}

// HasPrMerged returns a boolean if a field has been set.
func (o *PostWebhook) HasPrMerged() bool {
	if o != nil && o.PrMerged != nil {
		return true
	}

	return false
}

// SetPrMerged gets a reference to the given bool and assigns it to the PrMerged field.
func (o *PostWebhook) SetPrMerged(v bool) {
	o.PrMerged = &v
}

// GetPrReopened returns the PrReopened field value if set, zero value otherwise.
func (o *PostWebhook) GetPrReopened() bool {
	if o == nil || o.PrReopened == nil {
		var ret bool
		return ret
	}
	return *o.PrReopened
}

// GetPrReopenedOk returns a tuple with the PrReopened field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetPrReopenedOk() (*bool, bool) {
	if o == nil || o.PrReopened == nil {
		return nil, false
	}
	return o.PrReopened, true
}

// HasPrReopened returns a boolean if a field has been set.
func (o *PostWebhook) HasPrReopened() bool {
	if o != nil && o.PrReopened != nil {
		return true
	}

	return false
}

// SetPrReopened gets a reference to the given bool and assigns it to the PrReopened field.
func (o *PostWebhook) SetPrReopened(v bool) {
	o.PrReopened = &v
}

// GetPrRescoped returns the PrRescoped field value if set, zero value otherwise.
func (o *PostWebhook) GetPrRescoped() bool {
	if o == nil || o.PrRescoped == nil {
		var ret bool
		return ret
	}
	return *o.PrRescoped
}

// GetPrRescopedOk returns a tuple with the PrRescoped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetPrRescopedOk() (*bool, bool) {
	if o == nil || o.PrRescoped == nil {
		return nil, false
	}
	return o.PrRescoped, true
}

// HasPrRescoped returns a boolean if a field has been set.
func (o *PostWebhook) HasPrRescoped() bool {
	if o != nil && o.PrRescoped != nil {
		return true
	}

	return false
}

// SetPrRescoped gets a reference to the given bool and assigns it to the PrRescoped field.
func (o *PostWebhook) SetPrRescoped(v bool) {
	o.PrRescoped = &v
}

// GetPrUpdated returns the PrUpdated field value if set, zero value otherwise.
func (o *PostWebhook) GetPrUpdated() bool {
	if o == nil || o.PrUpdated == nil {
		var ret bool
		return ret
	}
	return *o.PrUpdated
}

// GetPrUpdatedOk returns a tuple with the PrUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetPrUpdatedOk() (*bool, bool) {
	if o == nil || o.PrUpdated == nil {
		return nil, false
	}
	return o.PrUpdated, true
}

// HasPrUpdated returns a boolean if a field has been set.
func (o *PostWebhook) HasPrUpdated() bool {
	if o != nil && o.PrUpdated != nil {
		return true
	}

	return false
}

// SetPrUpdated gets a reference to the given bool and assigns it to the PrUpdated field.
func (o *PostWebhook) SetPrUpdated(v bool) {
	o.PrUpdated = &v
}

// GetRepoPush returns the RepoPush field value if set, zero value otherwise.
func (o *PostWebhook) GetRepoPush() bool {
	if o == nil || o.RepoPush == nil {
		var ret bool
		return ret
	}
	return *o.RepoPush
}

// GetRepoPushOk returns a tuple with the RepoPush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetRepoPushOk() (*bool, bool) {
	if o == nil || o.RepoPush == nil {
		return nil, false
	}
	return o.RepoPush, true
}

// HasRepoPush returns a boolean if a field has been set.
func (o *PostWebhook) HasRepoPush() bool {
	if o != nil && o.RepoPush != nil {
		return true
	}

	return false
}

// SetRepoPush gets a reference to the given bool and assigns it to the RepoPush field.
func (o *PostWebhook) SetRepoPush(v bool) {
	o.RepoPush = &v
}

// GetTagCreated returns the TagCreated field value if set, zero value otherwise.
func (o *PostWebhook) GetTagCreated() bool {
	if o == nil || o.TagCreated == nil {
		var ret bool
		return ret
	}
	return *o.TagCreated
}

// GetTagCreatedOk returns a tuple with the TagCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetTagCreatedOk() (*bool, bool) {
	if o == nil || o.TagCreated == nil {
		return nil, false
	}
	return o.TagCreated, true
}

// HasTagCreated returns a boolean if a field has been set.
func (o *PostWebhook) HasTagCreated() bool {
	if o != nil && o.TagCreated != nil {
		return true
	}

	return false
}

// SetTagCreated gets a reference to the given bool and assigns it to the TagCreated field.
func (o *PostWebhook) SetTagCreated(v bool) {
	o.TagCreated = &v
}

// GetTitle returns the Title field value
func (o *PostWebhook) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PostWebhook) SetTitle(v string) {
	o.Title = v
}

// GetUrl returns the Url field value
func (o *PostWebhook) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PostWebhook) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PostWebhook) SetUrl(v string) {
	o.Url = v
}

func (o PostWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BranchCreated != nil {
		toSerialize["branchCreated"] = o.BranchCreated
	}
	if o.BranchDeleted != nil {
		toSerialize["branchDeleted"] = o.BranchDeleted
	}
	if true {
		toSerialize["branchesToIgnore"] = o.BranchesToIgnore
	}
	if o.BuildStatus != nil {
		toSerialize["buildStatus"] = o.BuildStatus
	}
	if true {
		toSerialize["committersToIgnore"] = o.CommittersToIgnore
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.PrCommented != nil {
		toSerialize["prCommented"] = o.PrCommented
	}
	if o.PrCreated != nil {
		toSerialize["prCreated"] = o.PrCreated
	}
	if o.PrDeclined != nil {
		toSerialize["prDeclined"] = o.PrDeclined
	}
	if o.PrMerged != nil {
		toSerialize["prMerged"] = o.PrMerged
	}
	if o.PrReopened != nil {
		toSerialize["prReopened"] = o.PrReopened
	}
	if o.PrRescoped != nil {
		toSerialize["prRescoped"] = o.PrRescoped
	}
	if o.PrUpdated != nil {
		toSerialize["prUpdated"] = o.PrUpdated
	}
	if o.RepoPush != nil {
		toSerialize["repoPush"] = o.RepoPush
	}
	if o.TagCreated != nil {
		toSerialize["tagCreated"] = o.TagCreated
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullablePostWebhook struct {
	value *PostWebhook
	isSet bool
}

func (v NullablePostWebhook) Get() *PostWebhook {
	return v.value
}

func (v *NullablePostWebhook) Set(val *PostWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullablePostWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullablePostWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostWebhook(val *PostWebhook) *NullablePostWebhook {
	return &NullablePostWebhook{value: val, isSet: true}
}

func (v NullablePostWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
