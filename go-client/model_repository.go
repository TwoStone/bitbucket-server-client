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

// Repository struct for Repository
type Repository struct {
	Id            *int32           `json:"id,omitempty"`
	Slug          string           `json:"slug"`
	Name          string           `json:"name"`
	Description   string           `json:"description"`
	HierarchyId   *string          `json:"hierarchyId,omitempty"`
	ScmId         *string          `json:"scmId,omitempty"`
	State         *string          `json:"state,omitempty"`
	StatusMessage *string          `json:"statusMessage,omitempty"`
	Forkable      bool             `json:"forkable"`
	Public        bool             `json:"public"`
	Project       Project          `json:"project"`
	Links         *RepositoryLinks `json:"links,omitempty"`
}

// NewRepository instantiates a new Repository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepository(slug string, name string, description string, forkable bool, public bool, project Project) *Repository {
	this := Repository{}
	this.Slug = slug
	this.Name = name
	this.Description = description
	this.Forkable = forkable
	this.Public = public
	this.Project = project
	return &this
}

// NewRepositoryWithDefaults instantiates a new Repository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryWithDefaults() *Repository {
	this := Repository{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Repository) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Repository) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Repository) SetId(v int32) {
	o.Id = &v
}

// GetSlug returns the Slug field value
func (o *Repository) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Repository) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Repository) SetSlug(v string) {
	o.Slug = v
}

// GetName returns the Name field value
func (o *Repository) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Repository) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Repository) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Repository) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Repository) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Repository) SetDescription(v string) {
	o.Description = v
}

// GetHierarchyId returns the HierarchyId field value if set, zero value otherwise.
func (o *Repository) GetHierarchyId() string {
	if o == nil || o.HierarchyId == nil {
		var ret string
		return ret
	}
	return *o.HierarchyId
}

// GetHierarchyIdOk returns a tuple with the HierarchyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetHierarchyIdOk() (*string, bool) {
	if o == nil || o.HierarchyId == nil {
		return nil, false
	}
	return o.HierarchyId, true
}

// HasHierarchyId returns a boolean if a field has been set.
func (o *Repository) HasHierarchyId() bool {
	if o != nil && o.HierarchyId != nil {
		return true
	}

	return false
}

// SetHierarchyId gets a reference to the given string and assigns it to the HierarchyId field.
func (o *Repository) SetHierarchyId(v string) {
	o.HierarchyId = &v
}

// GetScmId returns the ScmId field value if set, zero value otherwise.
func (o *Repository) GetScmId() string {
	if o == nil || o.ScmId == nil {
		var ret string
		return ret
	}
	return *o.ScmId
}

// GetScmIdOk returns a tuple with the ScmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetScmIdOk() (*string, bool) {
	if o == nil || o.ScmId == nil {
		return nil, false
	}
	return o.ScmId, true
}

// HasScmId returns a boolean if a field has been set.
func (o *Repository) HasScmId() bool {
	if o != nil && o.ScmId != nil {
		return true
	}

	return false
}

// SetScmId gets a reference to the given string and assigns it to the ScmId field.
func (o *Repository) SetScmId(v string) {
	o.ScmId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Repository) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Repository) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Repository) SetState(v string) {
	o.State = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *Repository) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *Repository) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *Repository) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetForkable returns the Forkable field value
func (o *Repository) GetForkable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Forkable
}

// GetForkableOk returns a tuple with the Forkable field value
// and a boolean to check if the value has been set.
func (o *Repository) GetForkableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Forkable, true
}

// SetForkable sets field value
func (o *Repository) SetForkable(v bool) {
	o.Forkable = v
}

// GetPublic returns the Public field value
func (o *Repository) GetPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Public
}

// GetPublicOk returns a tuple with the Public field value
// and a boolean to check if the value has been set.
func (o *Repository) GetPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Public, true
}

// SetPublic sets field value
func (o *Repository) SetPublic(v bool) {
	o.Public = v
}

// GetProject returns the Project field value
func (o *Repository) GetProject() Project {
	if o == nil {
		var ret Project
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *Repository) GetProjectOk() (*Project, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *Repository) SetProject(v Project) {
	o.Project = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Repository) GetLinks() RepositoryLinks {
	if o == nil || o.Links == nil {
		var ret RepositoryLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetLinksOk() (*RepositoryLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Repository) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RepositoryLinks and assigns it to the Links field.
func (o *Repository) SetLinks(v RepositoryLinks) {
	o.Links = &v
}

func (o Repository) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.HierarchyId != nil {
		toSerialize["hierarchyId"] = o.HierarchyId
	}
	if o.ScmId != nil {
		toSerialize["scmId"] = o.ScmId
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.StatusMessage != nil {
		toSerialize["statusMessage"] = o.StatusMessage
	}
	if true {
		toSerialize["forkable"] = o.Forkable
	}
	if true {
		toSerialize["public"] = o.Public
	}
	if true {
		toSerialize["project"] = o.Project
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableRepository struct {
	value *Repository
	isSet bool
}

func (v NullableRepository) Get() *Repository {
	return v.value
}

func (v *NullableRepository) Set(val *Repository) {
	v.value = val
	v.isSet = true
}

func (v NullableRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepository(val *Repository) *NullableRepository {
	return &NullableRepository{value: val, isSet: true}
}

func (v NullableRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
