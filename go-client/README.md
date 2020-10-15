# Go API client for bitbucket

<h1>REST Resources Provided By: Bitbucket Server - REST</h1>
<p>
    This is the reference document for the Atlassian Bitbucket REST API. The REST API is for developers who want to:
</p>
<ul>
    <li>integrate Bitbucket with other applications;</li>
    <li>create scripts that interact with Bitbucket; or</li>
    <li>develop plugins that enhance the Bitbucket UI, using REST to interact with the backend.</li>
</ul>
You can read more about developing Bitbucket plugins in the
<a href=\"https://developer.atlassian.com/server/bitbucket/\">Bitbucket Developer Documentation</a>.
<p></p>
<h2 id=\"gettingstarted\">Getting started</h2>
<p>
    Because the REST API is based on open standards, you can use any web development language or command line tool
    capable of generating an HTTP request to access the API. See the
    <a href=\"https://developer.atlassian.com/server/bitbucket/reference/rest-api/\">developer documentation</a> for a
    basic
    usage example.
</p>
<p>
    If you're already working with the
    <a href=\"https://developer.atlassian.com/server/framework/atlassian-sdk/\">Atlassian SDK</a>,
    the <a href=\"https://developer.atlassian.com/server/framework/atlassian-sdk/using-the-rest-api-browser/\">REST API
        Browser</a> is a great tool for exploring and experimenting with the Bitbucket REST API.
</p>
<h2>
    <a name=\"StructureoftheRESTURIs\"></a>Structure of the REST URIs</h2>
<p>
    Bitbucket's REST APIs provide access to resources (data entities) via URI paths. To use a REST API, your application
    will
    make an HTTP request and parse the response. The Bitbucket REST API uses JSON as its communication format, and the
    standard
    HTTP methods like GET, PUT, POST and DELETE. URIs for Bitbucket's REST API resource have the following structure:
</p>
<pre>    http://host:port/context/rest/api-name/api-version/path/to/resource
</pre>
<p>
    For example, the following URI would retrieve a page of the latest commits to the <strong>jira</strong> repository
    in
    the <strong>Jira</strong> project on <a href=\"https://stash.atlassian.com\">https://stash.atlassian.com</a>.
</p>
<pre>    https://stash.atlassian.com/rest/api/1.0/projects/JIRA/repos/jira/commits
</pre>
<p>
    See the API descriptions below for a full list of available resources.
</p>
<p>
    Alternatively we also publish a list of resources in
    <a href=\"http://en.wikipedia.org/wiki/Web_Application_Description_Language\">WADL</a> format. It is available
    <a href=\"bitbucket-rest.wadl\">here</a>.
</p>
<h2 id=\"paging-params\">Paged APIs</h2>
<p>
    Bitbucket uses paging to conserve server resources and limit response size for resources that return potentially
    large
    collections of items. A request to a paged API will result in a <code>values</code> array wrapped in a JSON object
    with some paging metadata, like this:
</p>
<pre>    {
        \"size\": 3,
        \"limit\": 3,
        \"isLastPage\": false,
        \"values\": [
            { /_* result 0 *_/ },
            { /_* result 1 *_/ },
            { /_* result 2 *_/ }
        ],
        \"start\": 0,
        \"filter\": null,
        \"nextPageStart\": 3
    }
</pre>
<p>
    Clients can use the <code>limit</code> and <code>start</code> query parameters to retrieve the desired number of
    results.
</p>
<p>
    The <code>limit</code> parameter indicates how many results to return per page. Most APIs default to returning
    <code>25</code> if the limit is left unspecified. This number can be increased, but note that a resource-specific
    hard limit will apply. These hard limits can be configured by server administrators, so it's always best practice to
    check the <code>limit</code> attribute on the response to see what limit has been applied.
    The request to get a larger page should look like this:
</p>
<pre>    http://host:port/context/rest/api-name/api-version/path/to/resource?limit={desired size of page}
</pre>
<p>
    For example:
</p>
<pre>    https://stash.atlassian.com/rest/api/1.0/projects/JIRA/repos/jira/commits?limit=1000
</pre>
<p>
    The <code>start</code> parameter indicates which item should be used as the first item in the page of results. All
    paged responses contain an <code>isLastPage</code> attribute indicating whether another page of items exists.
</p>
<p><strong>Important:</strong> If more than one page exists (i.e. the response contains
    <code>\"isLastPage\": false</code>), the response object will also contain a <code>nextPageStart</code> attribute
    which <strong><em>must</em></strong> be used by the client as the <code>start</code> parameter on the next request.
    Identifiers of adjacent objects in a page may not be contiguous, so the start of the next page is <em>not</em>
    necessarily the start of the last page plus the last page's size. A client should always use
    <code>nextPageStart</code> to avoid unexpected results from a paged API.
    The request to get a subsequent page should look like this:
</p>
<pre>    http://host:port/context/rest/api-name/api-version/path/to/resource?start={nextPageStart from previous response}
</pre>
<p>
    For example:
</p>
<pre>    https://stash.atlassian.com/rest/api/1.0/projects/JIRA/repos/jira/commits?start=25
</pre>
<h2 id=\"authentication\">Authentication</h2>
<p>
    Any authentication that works against Bitbucket will work against the REST API. <b>The preferred authentication
        methods
        are HTTP Basic (when using SSL) and OAuth</b>. Other supported methods include: HTTP Cookies and Trusted
    Applications.
</p>
<p>
    You can find OAuth code samples in several programming languages at
    <a
        href=\"https://bitbucket.org/atlassian_tutorial/atlassian-oauth-examples\">bitbucket.org/atlassian_tutorial/atlassian-oauth-examples</a>.
</p>
<p>
    The log-in page uses cookie-based authentication, so if you are using Bitbucket in a browser you can call REST from
    JavaScript on the page and rely on the authentication that the browser has established.
</p>
<h2 id=\"errors-and-validation\">Errors &amp; Validation</h2>
<p>
    If a request fails due to client error, the resource will return an HTTP response code in the 40x range. These can
    be broadly categorised into:
</p>
<table>
    <tbody>
        <tr>
            <th>HTTP Code</th>
            <th>Description</th>
        </tr>
        <tr>
            <td>400 (Bad Request)</td>
            <td>
                One or more of the required parameters or attributes:
                <ul>
                    <li>were missing from the request;</li>
                    <li>incorrectly formatted; or</li>
                    <li>inappropriate in the given context.</li>
                </ul>
            </td>
        </tr>
        <tr>
            <td>401 (Unauthorized)</td>
            <td>
                Either:
                <ul>
                    <li>Authentication is required but was not attempted.</li>
                    <li>Authentication was attempted but failed.</li>
                    <li>Authentication was successful but the authenticated user does not have the requisite permission
                        for the resource.</li>
                </ul>
                See the individual resource documentation for details of required permissions.
            </td>
        </tr>
        <tr>
            <td>403 (Forbidden)</td>
            <td>
                Actions are usually \"forbidden\" if they involve breaching the licensed user limit of the server, or
                degrading the authenticated user's permission level. See the individual resource documentation for more
                details.
            </td>
        </tr>
        <tr>
            <td>404 (Not Found)</td>
            <td>
                The entity you are attempting to access, or the project or repository containing it, does not exist.
            </td>
        </tr>
        <tr>
            <td>405 (Method Not Allowed)</td>
            <td>
                The request HTTP method is not appropriate for the targeted resource. For example an HTTP GET to a
                resource that only accepts an HTTP POST will result in a 405.
            </td>
        </tr>
        <tr>
            <td>409 (Conflict)</td>
            <td>
                The attempted update failed due to some conflict with an existing resource. For example:
                <ul>
                    <li>Creating a project with a key that already exists</li>
                    <li>Merging an out-of-date pull request</li>
                    <li>Deleting a comment that has replies</li>
                    <li>etc.</li>
                </ul>
                See the individual resource documentation for more details.
            </td>
        </tr>
        <tr>
            <td>415 (Unsupported Media Type)</td>
            <td>
                The request entity has a <code>Content-Type</code> that the server does not support. Almost all of the
                Bitbucket REST API accepts <code>application/json</code> format, but check the individual resource
                documentation for more details. Additionally, double-check that you are setting the
                <code>Content-Type</code> header correctly on your request (e.g. using
                <code>-H \"Content-Type: application/json\"</code> in cURL).
            </td>
        </tr>
    </tbody>
</table>
<p>
    For <strong>400</strong> HTTP codes the response will typically contain one or more validation error messages,
    for example:
</p>
<pre>    {
        \"errors\": [
            {
                \"context\": \"name\",
                \"message\": \"The name should be between 1 and 255 characters.\",
                \"exceptionName\": null
            },
            {
                \"context\": \"email\",
                \"message\": \"The email should be a valid email address.\",
                \"exceptionName\": null
            }
        ]
    }
    </pre>
<p>
    The <code>context</code> attribute indicates which parameter or request entity attribute failed validation. Note
    that the <code>context</code> may be null.
</p>
<p>
    For <strong>401</strong>, <strong>403</strong>, <strong>404</strong> and <strong>409</strong>
    HTTP codes, the response will contain one or more descriptive error messages:
</p>
<pre>    {
        \"errors\": [
            {
                \"context\": null,
                \"message\": \"A detailed error message.\",
                \"exceptionName\": null
            }
        ]
    }
    </pre>
<p>
    A <strong>500</strong> (Server Error) HTTP code indicates an incorrect resource url or an unexpected server error.
    Double-check the URL you are trying to access, then report the issue to your server administrator or
    <a href=\"https://support.atlassian.com/\">Atlassian Support</a> if problems persist.
</p>
<h2 id=\"personal-repositories\">Personal Repositories</h2>
<p>
    Bitbucket allows users to manage their own repositories, called personal repositories. These are repositories
    associated
    with the user and to which they always have REPO_ADMIN permission.
</p>
<p>
    Accessing personal repositories via REST is achieved through the normal project-centric REST URLs
    using the user's slug prefixed by tilde as the project key. E.g. to list personal repositories for a
    user with slug \"johnsmith\" you would make a GET to:
</p>
<pre>http://example.com/rest/api/1.0/projects/~johnsmith/repos</pre>
<p></p>
<p>
    In addition to this, Bitbucket allows access to these repositories through an alternate set of user-centric REST
    URLs
    beginning with:
</p>
<pre>http://example.com/rest/api/1.0/users/~{userSlug}/repos</pre>
E.g. to list the forks of the repository with slug
\"nodejs\" in the personal project of user with slug \"johnsmith\" using the regular REST URL you would make a GET to:
<pre>http://example.com/rest/api/1.0/projects/~johnsmith/repos/nodejs/forks</pre>
Using the alternate URL, you would make a GET to:
<pre>http://example.com/rest/api/1.0/users/johnsmith/repos/nodejs/forks</pre>
<p></p>

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 7.3.1
- Package version: 1.3.3
- Build package: org.openapitools.codegen.languages.GoClientExperimentalCodegen
For more information, please visit [https://github.com/TwoStone/bitbucket-server-api](https://github.com/TwoStone/bitbucket-server-api)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./bitbucket"
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BranchesApi* | [**CreateBranch**](docs/BranchesApi.md#createbranch) | **Post** /rest/branch-utils/1.0/projects/{projectKey}/repos/{repositorySlug}/branches | Delete branch
*BranchesApi* | [**DeleteBranch**](docs/BranchesApi.md#deletebranch) | **Delete** /rest/branch-utils/1.0/projects/{projectKey}/repos/{repositorySlug}/branches | 
*BranchesApi* | [**GetBranchesPaged**](docs/BranchesApi.md#getbranchespaged) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/branches | Get Branches
*BranchesApi* | [**GetDefaultBranch**](docs/BranchesApi.md#getdefaultbranch) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/branches/default | Get default branch
*BuildStatusApi* | [**GetBuildStatusesPaged**](docs/BuildStatusApi.md#getbuildstatusespaged) | **Get** /rest/build-status/1.0/commits/{commitHash} | Get build statuses
*BuildStatusApi* | [**PostBuildResult**](docs/BuildStatusApi.md#postbuildresult) | **Post** /rest/build-status/1.0/commits/{commitHash} | Post build-result
*CommitsApi* | [**GetCommit**](docs/CommitsApi.md#getcommit) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId} | Get commit
*CommitsApi* | [**GetCommitsPaged**](docs/CommitsApi.md#getcommitspaged) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits | Get commits
*PostWebhookApi* | [**DeletePostWebhook**](docs/PostWebhookApi.md#deletepostwebhook) | **Delete** /rest/webhook/1.0/projects/{projectKey}/repos/{repositorySlug}/configurations/{ID} | Delete post webhook
*PostWebhookApi* | [**GetPostWebhooks**](docs/PostWebhookApi.md#getpostwebhooks) | **Get** /rest/webhook/1.0/projects/{projectKey}/repos/{repositorySlug}/configurations | Get Post Webhooks
*ProjectsApi* | [**GetProject**](docs/ProjectsApi.md#getproject) | **Get** /rest/api/1.0/projects/{projectKey} | REST resource for working with projects
*ProjectsApi* | [**GetProjectsPaged**](docs/ProjectsApi.md#getprojectspaged) | **Get** /rest/api/1.0/projects | Get Projects
*PullRequestsApi* | [**CreatePullRequest**](docs/PullRequestsApi.md#createpullrequest) | **Post** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests | Create Pull Request
*PullRequestsApi* | [**DeletePullRequest**](docs/PullRequestsApi.md#deletepullrequest) | **Delete** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId} | Delete pull request
*PullRequestsApi* | [**GetDiff**](docs/PullRequestsApi.md#getdiff) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}.diff | Get PR Diff
*PullRequestsApi* | [**GetPullRequest**](docs/PullRequestsApi.md#getpullrequest) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId} | Get pull request
*PullRequestsApi* | [**GetPullRequestsPaged**](docs/PullRequestsApi.md#getpullrequestspaged) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests | Get Pull Request Page
*PullRequestsApi* | [**GetRestApi10ProjectsProjectKeyReposRepositorySlugPullRequestsPullRequestIdPatch**](docs/PullRequestsApi.md#getrestapi10projectsprojectkeyreposrepositoryslugpullrequestspullrequestidpatch) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}.patch | Get PR Patch
*PullRequestsApi* | [**UpdatePullRequest**](docs/PullRequestsApi.md#updatepullrequest) | **Put** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId} | Update pull request
*RepositoriesApi* | [**BrowseRepositoryPaged**](docs/RepositoriesApi.md#browserepositorypaged) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/browse | browseRepository
*RepositoriesApi* | [**BrowseRepositoryPathPaged**](docs/RepositoriesApi.md#browserepositorypathpaged) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/browse/{path} | browseRepositoryPath
*RepositoriesApi* | [**CreateRepository**](docs/RepositoriesApi.md#createrepository) | **Post** /rest/api/1.0/projects/{projectKey}/repos | Create repository
*RepositoriesApi* | [**GetConfiguredDefaultBranch**](docs/RepositoriesApi.md#getconfigureddefaultbranch) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/default-branch | Get default branch
*RepositoriesApi* | [**GetRepositoriesPaged**](docs/RepositoriesApi.md#getrepositoriespaged) | **Get** /rest/api/1.0/projects/{projectKey}/repos | Get Repositories
*RepositoriesApi* | [**GetRepository**](docs/RepositoriesApi.md#getrepository) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug} | Get Repository
*RepositoriesApi* | [**SearchRepositoriesPaged**](docs/RepositoriesApi.md#searchrepositoriespaged) | **Get** /rest/api/1.0/repos | Search repositories
*WebhookApi* | [**CreateWebhook**](docs/WebhookApi.md#createwebhook) | **Post** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks | Create webhook
*WebhookApi* | [**DeleteWebhook**](docs/WebhookApi.md#deletewebhook) | **Delete** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId} | Delete Webhook
*WebhookApi* | [**GetWebhook**](docs/WebhookApi.md#getwebhook) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId} | Get Webhook
*WebhookApi* | [**GetWebhooksPaged**](docs/WebhookApi.md#getwebhookspaged) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks | Get webhooks
*WebhookApi* | [**UpdateWebhook**](docs/WebhookApi.md#updatewebhook) | **Put** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId} | Update webhook


## Documentation For Models

 - [Branch](docs/Branch.md)
 - [BranchMetadata](docs/BranchMetadata.md)
 - [BranchMetadataAheadBehind](docs/BranchMetadataAheadBehind.md)
 - [BranchMetadataBuildStatus](docs/BranchMetadataBuildStatus.md)
 - [BranchMetadataJiraIssue](docs/BranchMetadataJiraIssue.md)
 - [BranchMetadataOutgoingPullRequest](docs/BranchMetadataOutgoingPullRequest.md)
 - [BranchMetadataOutgoingPullRequestOneOf](docs/BranchMetadataOutgoingPullRequestOneOf.md)
 - [BranchMetadataOutgoingPullRequestOneOf1](docs/BranchMetadataOutgoingPullRequestOneOf1.md)
 - [BranchesPage](docs/BranchesPage.md)
 - [BranchesPageAllOf](docs/BranchesPageAllOf.md)
 - [BuildStatus](docs/BuildStatus.md)
 - [BuildStatusPage](docs/BuildStatusPage.md)
 - [BuildStatusPageAllOf](docs/BuildStatusPageAllOf.md)
 - [Children](docs/Children.md)
 - [ChildrenAllOf](docs/ChildrenAllOf.md)
 - [ChildrenAllOfValues](docs/ChildrenAllOfValues.md)
 - [Commit](docs/Commit.md)
 - [CommitParents](docs/CommitParents.md)
 - [CommitsPage](docs/CommitsPage.md)
 - [CommitsPageAllOf](docs/CommitsPageAllOf.md)
 - [CreateBranch](docs/CreateBranch.md)
 - [CreateRepository](docs/CreateRepository.md)
 - [DefaultBranch](docs/DefaultBranch.md)
 - [DeleteBranch](docs/DeleteBranch.md)
 - [Directory](docs/Directory.md)
 - [Errors](docs/Errors.md)
 - [ErrorsErrors](docs/ErrorsErrors.md)
 - [File](docs/File.md)
 - [FileLines](docs/FileLines.md)
 - [FileOrDirectory](docs/FileOrDirectory.md)
 - [Link](docs/Link.md)
 - [Page](docs/Page.md)
 - [Path](docs/Path.md)
 - [PostWebhook](docs/PostWebhook.md)
 - [Project](docs/Project.md)
 - [ProjectLinks](docs/ProjectLinks.md)
 - [ProjectsPage](docs/ProjectsPage.md)
 - [ProjectsPageAllOf](docs/ProjectsPageAllOf.md)
 - [PullRequest](docs/PullRequest.md)
 - [PullRequestDelete](docs/PullRequestDelete.md)
 - [PullRequestState](docs/PullRequestState.md)
 - [PullRequestUpdate](docs/PullRequestUpdate.md)
 - [PullRequestsPage](docs/PullRequestsPage.md)
 - [PullRequestsPageAllOf](docs/PullRequestsPageAllOf.md)
 - [RepositoriesPage](docs/RepositoriesPage.md)
 - [RepositoriesPageAllOf](docs/RepositoriesPageAllOf.md)
 - [Repository](docs/Repository.md)
 - [RepositoryLinks](docs/RepositoryLinks.md)
 - [RepositoryRef](docs/RepositoryRef.md)
 - [RepositoryRefRepository](docs/RepositoryRefRepository.md)
 - [RepositoryRefRepositoryProject](docs/RepositoryRefRepositoryProject.md)
 - [User](docs/User.md)
 - [UserRole](docs/UserRole.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookConfiguration](docs/WebhookConfiguration.md)
 - [WebhookEvent](docs/WebhookEvent.md)
 - [WebhooksPage](docs/WebhooksPage.md)
 - [WebhooksPageAllOf](docs/WebhooksPageAllOf.md)


## Documentation For Authorization



### usernamePassword

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



