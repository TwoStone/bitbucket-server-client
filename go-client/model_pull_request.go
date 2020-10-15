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

// PullRequest struct for PullRequest
type PullRequest struct {
	Id           *int32            `json:"id,omitempty"`
	Version      *int32            `json:"version,omitempty"`
	Title        string            `json:"title"`
	Description  string            `json:"description"`
	State        *PullRequestState `json:"state,omitempty"`
	Open         *bool             `json:"open,omitempty"`
	Closed       *bool             `json:"closed,omitempty"`
	CreatedDate  *int64            `json:"createdDate,omitempty"`
	UpdatedDate  *int64            `json:"updatedDate,omitempty"`
	FromRef      RepositoryRef     `json:"fromRef"`
	ToRef        RepositoryRef     `json:"toRef"`
	Locked       *bool             `json:"locked,omitempty"`
	Author       *UserRole         `json:"author,omitempty"`
	Reviewers    *[]UserRole       `json:"reviewers,omitempty"`
	Participants *[]UserRole       `json:"participants,omitempty"`
	Links        *ProjectLinks     `json:"links,omitempty"`
}

// NewPullRequest instantiates a new PullRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullRequest(title string, description string, fromRef RepositoryRef, toRef RepositoryRef) *PullRequest {
	this := PullRequest{}
	this.Title = title
	this.Description = description
	this.FromRef = fromRef
	this.ToRef = toRef
	return &this
}

// NewPullRequestWithDefaults instantiates a new PullRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullRequestWithDefaults() *PullRequest {
	this := PullRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PullRequest) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PullRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PullRequest) SetId(v int32) {
	o.Id = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PullRequest) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PullRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *PullRequest) SetVersion(v int32) {
	o.Version = &v
}

// GetTitle returns the Title field value
func (o *PullRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PullRequest) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *PullRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PullRequest) SetDescription(v string) {
	o.Description = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PullRequest) GetState() PullRequestState {
	if o == nil || o.State == nil {
		var ret PullRequestState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetStateOk() (*PullRequestState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PullRequest) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given PullRequestState and assigns it to the State field.
func (o *PullRequest) SetState(v PullRequestState) {
	o.State = &v
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *PullRequest) GetOpen() bool {
	if o == nil || o.Open == nil {
		var ret bool
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetOpenOk() (*bool, bool) {
	if o == nil || o.Open == nil {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *PullRequest) HasOpen() bool {
	if o != nil && o.Open != nil {
		return true
	}

	return false
}

// SetOpen gets a reference to the given bool and assigns it to the Open field.
func (o *PullRequest) SetOpen(v bool) {
	o.Open = &v
}

// GetClosed returns the Closed field value if set, zero value otherwise.
func (o *PullRequest) GetClosed() bool {
	if o == nil || o.Closed == nil {
		var ret bool
		return ret
	}
	return *o.Closed
}

// GetClosedOk returns a tuple with the Closed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetClosedOk() (*bool, bool) {
	if o == nil || o.Closed == nil {
		return nil, false
	}
	return o.Closed, true
}

// HasClosed returns a boolean if a field has been set.
func (o *PullRequest) HasClosed() bool {
	if o != nil && o.Closed != nil {
		return true
	}

	return false
}

// SetClosed gets a reference to the given bool and assigns it to the Closed field.
func (o *PullRequest) SetClosed(v bool) {
	o.Closed = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *PullRequest) GetCreatedDate() int64 {
	if o == nil || o.CreatedDate == nil {
		var ret int64
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetCreatedDateOk() (*int64, bool) {
	if o == nil || o.CreatedDate == nil {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *PullRequest) HasCreatedDate() bool {
	if o != nil && o.CreatedDate != nil {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given int64 and assigns it to the CreatedDate field.
func (o *PullRequest) SetCreatedDate(v int64) {
	o.CreatedDate = &v
}

// GetUpdatedDate returns the UpdatedDate field value if set, zero value otherwise.
func (o *PullRequest) GetUpdatedDate() int64 {
	if o == nil || o.UpdatedDate == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetUpdatedDateOk() (*int64, bool) {
	if o == nil || o.UpdatedDate == nil {
		return nil, false
	}
	return o.UpdatedDate, true
}

// HasUpdatedDate returns a boolean if a field has been set.
func (o *PullRequest) HasUpdatedDate() bool {
	if o != nil && o.UpdatedDate != nil {
		return true
	}

	return false
}

// SetUpdatedDate gets a reference to the given int64 and assigns it to the UpdatedDate field.
func (o *PullRequest) SetUpdatedDate(v int64) {
	o.UpdatedDate = &v
}

// GetFromRef returns the FromRef field value
func (o *PullRequest) GetFromRef() RepositoryRef {
	if o == nil {
		var ret RepositoryRef
		return ret
	}

	return o.FromRef
}

// GetFromRefOk returns a tuple with the FromRef field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetFromRefOk() (*RepositoryRef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromRef, true
}

// SetFromRef sets field value
func (o *PullRequest) SetFromRef(v RepositoryRef) {
	o.FromRef = v
}

// GetToRef returns the ToRef field value
func (o *PullRequest) GetToRef() RepositoryRef {
	if o == nil {
		var ret RepositoryRef
		return ret
	}

	return o.ToRef
}

// GetToRefOk returns a tuple with the ToRef field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetToRefOk() (*RepositoryRef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToRef, true
}

// SetToRef sets field value
func (o *PullRequest) SetToRef(v RepositoryRef) {
	o.ToRef = v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *PullRequest) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *PullRequest) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *PullRequest) SetLocked(v bool) {
	o.Locked = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *PullRequest) GetAuthor() UserRole {
	if o == nil || o.Author == nil {
		var ret UserRole
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetAuthorOk() (*UserRole, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *PullRequest) HasAuthor() bool {
	if o != nil && o.Author != nil {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given UserRole and assigns it to the Author field.
func (o *PullRequest) SetAuthor(v UserRole) {
	o.Author = &v
}

// GetReviewers returns the Reviewers field value if set, zero value otherwise.
func (o *PullRequest) GetReviewers() []UserRole {
	if o == nil || o.Reviewers == nil {
		var ret []UserRole
		return ret
	}
	return *o.Reviewers
}

// GetReviewersOk returns a tuple with the Reviewers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetReviewersOk() (*[]UserRole, bool) {
	if o == nil || o.Reviewers == nil {
		return nil, false
	}
	return o.Reviewers, true
}

// HasReviewers returns a boolean if a field has been set.
func (o *PullRequest) HasReviewers() bool {
	if o != nil && o.Reviewers != nil {
		return true
	}

	return false
}

// SetReviewers gets a reference to the given []UserRole and assigns it to the Reviewers field.
func (o *PullRequest) SetReviewers(v []UserRole) {
	o.Reviewers = &v
}

// GetParticipants returns the Participants field value if set, zero value otherwise.
func (o *PullRequest) GetParticipants() []UserRole {
	if o == nil || o.Participants == nil {
		var ret []UserRole
		return ret
	}
	return *o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetParticipantsOk() (*[]UserRole, bool) {
	if o == nil || o.Participants == nil {
		return nil, false
	}
	return o.Participants, true
}

// HasParticipants returns a boolean if a field has been set.
func (o *PullRequest) HasParticipants() bool {
	if o != nil && o.Participants != nil {
		return true
	}

	return false
}

// SetParticipants gets a reference to the given []UserRole and assigns it to the Participants field.
func (o *PullRequest) SetParticipants(v []UserRole) {
	o.Participants = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PullRequest) GetLinks() ProjectLinks {
	if o == nil || o.Links == nil {
		var ret ProjectLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetLinksOk() (*ProjectLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PullRequest) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ProjectLinks and assigns it to the Links field.
func (o *PullRequest) SetLinks(v ProjectLinks) {
	o.Links = &v
}

func (o PullRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Open != nil {
		toSerialize["open"] = o.Open
	}
	if o.Closed != nil {
		toSerialize["closed"] = o.Closed
	}
	if o.CreatedDate != nil {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if o.UpdatedDate != nil {
		toSerialize["updatedDate"] = o.UpdatedDate
	}
	if true {
		toSerialize["fromRef"] = o.FromRef
	}
	if true {
		toSerialize["toRef"] = o.ToRef
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.Reviewers != nil {
		toSerialize["reviewers"] = o.Reviewers
	}
	if o.Participants != nil {
		toSerialize["participants"] = o.Participants
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullablePullRequest struct {
	value *PullRequest
	isSet bool
}

func (v NullablePullRequest) Get() *PullRequest {
	return v.value
}

func (v *NullablePullRequest) Set(val *PullRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePullRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePullRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullRequest(val *PullRequest) *NullablePullRequest {
	return &NullablePullRequest{value: val, isSet: true}
}

func (v NullablePullRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
