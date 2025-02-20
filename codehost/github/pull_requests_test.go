// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package github_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-github/v45/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
	host "github.com/reviewpad/reviewpad/v3/codehost/github"
	"github.com/reviewpad/reviewpad/v3/lang/aladino"
	"github.com/stretchr/testify/assert"
)

type paginatedRequestResult struct {
	pageNum int
}

func TestGetPullRequestHeadOwnerName(t *testing.T) {
	mockedHeadOwnerName := "reviewpad"
	mockedPullRequest := aladino.GetDefaultMockPullRequestDetailsWith(&github.PullRequest{
		Head: &github.PullRequestBranch{
			Repo: &github.Repository{
				Owner: &github.User{
					Login: github.String(mockedHeadOwnerName),
				},
			},
		},
	})
	wantOwnerName := mockedPullRequest.Head.Repo.Owner.GetLogin()
	gotOwnerName := host.GetPullRequestHeadOwnerName(mockedPullRequest)

	assert.Equal(t, wantOwnerName, gotOwnerName)
	assert.Equal(t, mockedHeadOwnerName, gotOwnerName)
}

func TestGetPullRequestHeadRepoName(t *testing.T) {
	mockedHeadRepoName := "mocks-test"
	mockedPullRequest := aladino.GetDefaultMockPullRequestDetailsWith(&github.PullRequest{
		Head: &github.PullRequestBranch{
			Repo: &github.Repository{
				Name: &mockedHeadRepoName,
			},
		},
	})
	wantRepoName := mockedPullRequest.Head.Repo.GetName()
	gotRepoName := host.GetPullRequestHeadRepoName(mockedPullRequest)

	assert.Equal(t, wantRepoName, gotRepoName)
	assert.Equal(t, mockedHeadRepoName, gotRepoName)
}

func TestGetPullRequestBaseOwnerName(t *testing.T) {
	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	wantOwnerName := mockedPullRequest.Base.Repo.Owner.GetLogin()
	gotOwnerName := host.GetPullRequestBaseOwnerName(mockedPullRequest)

	assert.Equal(t, wantOwnerName, gotOwnerName)
}

func TestGetPullRequestBaseRepoName(t *testing.T) {
	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	wantRepoName := mockedPullRequest.Base.Repo.GetName()
	gotRepoName := host.GetPullRequestBaseRepoName(mockedPullRequest)

	assert.Equal(t, wantRepoName, gotRepoName)
}

func TestGetPullRequestNumber(t *testing.T) {
	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	wantPullRequestNumber := mockedPullRequest.GetNumber()
	gotPullRequestNumber := host.GetPullRequestNumber(mockedPullRequest)

	assert.Equal(t, wantPullRequestNumber, gotPullRequestNumber)
}

func TestPaginatedRequest_WhenFirstRequestFails(t *testing.T) {
	failMessage := "PaginatedRequestFail"
	initFn := func() interface{} {
		return paginatedRequestResult{}
	}
	reqFn := func(i interface{}, page int) (interface{}, *github.Response, error) {
		return nil, nil, errors.New(failMessage)
	}

	res, err := host.PaginatedRequest(initFn, reqFn)

	assert.Nil(t, res)
	assert.EqualError(t, err, failMessage)
}

func TestPaginatedRequest_WhenFurtherRequestsFail(t *testing.T) {
	failMessage := "PaginatedRequestFail"
	initFn := func() interface{} {
		return paginatedRequestResult{
			pageNum: 1,
		}
	}
	reqFn := func(i interface{}, page int) (interface{}, *github.Response, error) {
		if page == 1 {
			respHeader := make(http.Header)
			respHeader.Add("Link", "<https://api.github.com/user/58276/repos?page=3>; rel=\"last\"")
			resp := &github.Response{
				Response: &http.Response{
					Header: respHeader,
				},
				NextPage: 3,
			}

			return paginatedRequestResult{pageNum: 1}, resp, nil
		}

		return nil, nil, errors.New(failMessage)
	}

	res, err := host.PaginatedRequest(initFn, reqFn)

	assert.Nil(t, res)
	assert.EqualError(t, err, failMessage)
}

func TestPaginatedRequest(t *testing.T) {
	initFn := func() interface{} {
		return []*paginatedRequestResult{
			{pageNum: 1},
		}
	}
	reqFn := func(i interface{}, page int) (interface{}, *github.Response, error) {
		results := i.([]*paginatedRequestResult)
		if page == 1 {
			respHeader := make(http.Header)
			respHeader.Add("Link", "<https://api.github.com/user/58276/repos?page=3>; rel=\"last\"")
			resp := &github.Response{
				Response: &http.Response{
					Header: respHeader,
				},
			}

			return results, resp, nil
		}

		return results, nil, nil
	}

	wantRes := []*paginatedRequestResult{{pageNum: 1}}
	gotRes, err := host.PaginatedRequest(initFn, reqFn)

	assert.Nil(t, err)
	assert.Equal(t, gotRes, wantRes)
}

func TestParseNumPagesFromLink_WhenHTTPLinkHeaderHasNoRel(t *testing.T) {
	link := "<https://api.github.com/user/58276/repos?page=1>"

	wantNumPages := 0

	gotNumPages := host.ParseNumPagesFromLink(link)

	assert.Equal(t, wantNumPages, gotNumPages)
}

func TestParseNumPagesFromLink_WhenHTTPLinkHeaderIsInvalid(t *testing.T) {
	link := "<invalid%+url>; rel=\"last\""

	wantNumPages := 0

	gotNumPages := host.ParseNumPagesFromLink(link)

	assert.Equal(t, wantNumPages, gotNumPages)
}

func TestParseNumPagesFromLink_WhenHTTPLinkHeaderHasNoQueryParamPage(t *testing.T) {
	link := "<https://api.github.com/user/58276/repos>; rel=\"last\""

	wantNumPages := 0

	gotNumPages := host.ParseNumPagesFromLink(link)

	assert.Equal(t, wantNumPages, gotNumPages)
}

func TestParseNumPagesFromLink_WhenHTTPLinkHeaderHasInvalidQueryParamPage(t *testing.T) {
	link := "<https://api.github.com/user/58276/repos?page=7B316>; rel=\"last\""

	wantNumPages := 0

	gotNumPages := host.ParseNumPagesFromLink(link)

	assert.Equal(t, wantNumPages, gotNumPages)
}

func TestParseNumPagesFromLink(t *testing.T) {
	link := "<https://api.github.com/user/58276/repos?page=3>; rel=\"last\""

	// The number of pages is provided in the url query parameter "page"
	wantNumPages := 3

	gotNumPages := host.ParseNumPagesFromLink(link)

	assert.Equal(t, wantNumPages, gotNumPages)
}

func TestParseNumPages_WhenHTTPLinkHeaderIsNotProvided(t *testing.T) {
	respHeader := make(http.Header)
	respHeader.Add("Link", " ")
	resp := &github.Response{
		Response: &http.Response{
			Header: respHeader,
		},
	}

	wantNumPages := 0

	gotNumPages := host.ParseNumPages(resp)

	assert.Equal(t, wantNumPages, gotNumPages)
}

func TestParseNumPages(t *testing.T) {
	respHeader := make(http.Header)
	respHeader.Add("Link", "<https://api.github.com/user/58276/repos?page=3>; rel=\"last\"")
	resp := &github.Response{
		Response: &http.Response{
			Header: respHeader,
		},
	}

	// The number of pages is provided in the url query parameter "page"
	wantNumPages := 3

	gotNumPages := host.ParseNumPages(resp)

	assert.Equal(t, wantNumPages, gotNumPages)
}

func TestGetPullRequestComments_WhenListCommentsRequestFails(t *testing.T) {
	failMessage := "ListCommentsRequestFail"
	mockedGithubClient := aladino.MockDefaultGithubClient(
		[]mock.MockBackendOption{
			mock.WithRequestMatchHandler(
				mock.GetReposIssuesCommentsByOwnerByRepoByIssueNumber,
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					mock.WriteError(
						w,
						http.StatusInternalServerError,
						failMessage,
					)
				}),
			),
		},
		nil,
	)

	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	comments, err := mockedGithubClient.GetComments(
		context.Background(),
		mockedPullRequest.Base.Repo.Owner.GetLogin(),
		mockedPullRequest.Base.Repo.GetName(),
		mockedPullRequest.GetNumber(),
		&github.IssueListCommentsOptions{},
	)

	assert.Nil(t, comments)
	assert.Equal(t, err.(*github.ErrorResponse).Message, failMessage)
}

func TestGetPullRequestComments(t *testing.T) {
	wantComments := []*github.IssueComment{
		{Body: github.String("Lorem Ipsum")},
	}
	mockedGithubClient := aladino.MockDefaultGithubClient(
		[]mock.MockBackendOption{
			mock.WithRequestMatch(
				mock.GetReposIssuesCommentsByOwnerByRepoByIssueNumber,
				wantComments,
			),
		},
		nil,
	)

	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	gotComments, err := mockedGithubClient.GetComments(
		context.Background(),
		mockedPullRequest.Base.Repo.Owner.GetLogin(),
		mockedPullRequest.Base.Repo.GetName(),
		mockedPullRequest.GetNumber(),
		&github.IssueListCommentsOptions{},
	)

	assert.Nil(t, err)
	assert.Equal(t, wantComments, gotComments)
}

func TestGetPullRequestFiles(t *testing.T) {
	wantFiles := []*github.CommitFile{
		{
			Filename: github.String("default-mock-repo/file1.ts"),
			Patch:    nil,
		},
	}
	mockedGithubClient := aladino.MockDefaultGithubClient(
		[]mock.MockBackendOption{
			mock.WithRequestMatchHandler(
				mock.GetReposPullsFilesByOwnerByRepoByPullNumber,
				http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
					w.Write(mock.MustMarshal(wantFiles))
				}),
			),
		},
		nil,
	)

	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	gotFiles, err := mockedGithubClient.GetPullRequestFiles(
		context.Background(),
		mockedPullRequest.Base.Repo.Owner.GetLogin(),
		mockedPullRequest.Base.Repo.GetName(),
		mockedPullRequest.GetNumber(),
	)

	assert.Nil(t, err)
	assert.Equal(t, wantFiles, gotFiles)
}

func TestGetPullRequestReviewers_WhenListReviewersRequestFails(t *testing.T) {
	failMessage := "ListReviewersRequestFail"
	mockedGithubClient := aladino.MockDefaultGithubClient(
		[]mock.MockBackendOption{
			mock.WithRequestMatchHandler(
				mock.GetReposPullsRequestedReviewersByOwnerByRepoByPullNumber,
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					mock.WriteError(
						w,
						http.StatusInternalServerError,
						failMessage,
					)
				}),
			),
		},
		nil,
	)

	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	reviewers, err := mockedGithubClient.GetPullRequestReviewers(
		context.Background(),
		mockedPullRequest.Base.Repo.Owner.GetLogin(),
		mockedPullRequest.Base.Repo.GetName(),
		mockedPullRequest.GetNumber(),
		nil,
	)

	assert.Nil(t, reviewers)
	assert.Equal(t, err.(*github.ErrorResponse).Message, failMessage)
}

func TestGetPullRequestReviewers(t *testing.T) {
	wantReviewers := &github.Reviewers{
		Users: []*github.User{
			{Login: github.String("mary")},
		},
		Teams: []*github.Team{
			{Slug: github.String("reviewpad-team")},
		},
	}
	mockedGithubClient := aladino.MockDefaultGithubClient(
		[]mock.MockBackendOption{
			mock.WithRequestMatch(
				mock.GetReposPullsRequestedReviewersByOwnerByRepoByPullNumber,
				wantReviewers,
			),
		},
		nil,
	)

	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	gotReviewers, err := mockedGithubClient.GetPullRequestReviewers(
		context.Background(),
		mockedPullRequest.Base.Repo.Owner.GetLogin(),
		mockedPullRequest.Base.Repo.GetName(),
		mockedPullRequest.GetNumber(),
		nil,
	)

	assert.Nil(t, err)
	assert.Equal(t, wantReviewers, gotReviewers)
}

func TestGetRepoCollaborators_WhenListCollaboratorsRequestFails(t *testing.T) {
	failMessage := "ListCollaboratorsRequestFail"
	mockedGithubClient := aladino.MockDefaultGithubClient(
		[]mock.MockBackendOption{
			mock.WithRequestMatchHandler(
				mock.GetReposCollaboratorsByOwnerByRepo,
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					mock.WriteError(
						w,
						http.StatusInternalServerError,
						failMessage,
					)
				}),
			),
		},
		nil,
	)

	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	collaborators, err := mockedGithubClient.GetRepoCollaborators(
		context.Background(),
		mockedPullRequest.Base.Repo.Owner.GetLogin(),
		mockedPullRequest.Base.Repo.GetName(),
	)

	assert.Nil(t, collaborators)
	assert.Equal(t, err.(*github.ErrorResponse).Message, failMessage)
}

func TestGetRepoCollaborators(t *testing.T) {
	wantCollaborators := []*github.User{
		{Login: github.String("mary")},
	}
	mockedGithubClient := aladino.MockDefaultGithubClient(
		[]mock.MockBackendOption{
			mock.WithRequestMatch(
				mock.GetReposCollaboratorsByOwnerByRepo,
				wantCollaborators,
			),
		},
		nil,
	)

	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	gotCollaborators, err := mockedGithubClient.GetRepoCollaborators(
		context.Background(),
		mockedPullRequest.Base.Repo.Owner.GetLogin(),
		mockedPullRequest.Base.Repo.GetName(),
	)

	assert.Nil(t, err)
	assert.Equal(t, wantCollaborators, gotCollaborators)
}

func TestGetIssuesAvailableAssignees_WhenListAssigneesRequestFails(t *testing.T) {
	failMessage := "ListAssigneesRequestFail"
	mockedGithubClient := aladino.MockDefaultGithubClient(
		[]mock.MockBackendOption{
			mock.WithRequestMatchHandler(
				mock.GetReposAssigneesByOwnerByRepo,
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					mock.WriteError(
						w,
						http.StatusInternalServerError,
						failMessage,
					)
				}),
			),
		},
		nil,
	)

	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	gotAssignees, err := mockedGithubClient.GetIssuesAvailableAssignees(
		context.Background(),
		mockedPullRequest.GetUser().GetLogin(),
		mockedPullRequest.GetBase().GetRepo().GetName(),
	)

	assert.Nil(t, gotAssignees)
	assert.Equal(t, err.(*github.ErrorResponse).Message, failMessage)
}

func TestGetIssuesAvailableAssignees(t *testing.T) {
	wantAssignees := []*github.User{
		{Login: github.String("jane")},
	}
	mockedGithubClient := aladino.MockDefaultGithubClient(
		[]mock.MockBackendOption{
			mock.WithRequestMatch(
				mock.GetReposAssigneesByOwnerByRepo,
				wantAssignees,
			),
		},
		nil,
	)

	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	gotAssignees, err := mockedGithubClient.GetIssuesAvailableAssignees(
		context.Background(),
		mockedPullRequest.GetUser().GetLogin(),
		mockedPullRequest.GetBase().GetRepo().GetName(),
	)

	assert.Nil(t, err)
	assert.Equal(t, wantAssignees, gotAssignees)
}

func TestGetPullRequestCommits_WhenListCommistsRequestFails(t *testing.T) {
	failMessage := "ListCommitsRequestFail"
	mockedGithubClient := aladino.MockDefaultGithubClient(
		[]mock.MockBackendOption{
			mock.WithRequestMatchHandler(
				mock.GetReposPullsCommitsByOwnerByRepoByPullNumber,
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					mock.WriteError(
						w,
						http.StatusInternalServerError,
						failMessage,
					)
				}),
			),
		},
		nil,
	)

	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	gotCommits, err := mockedGithubClient.GetPullRequestCommits(
		context.Background(),
		mockedPullRequest.GetUser().GetLogin(),
		mockedPullRequest.GetBase().GetRepo().GetName(),
		mockedPullRequest.GetNumber(),
	)

	assert.Nil(t, gotCommits)
	assert.Equal(t, err.(*github.ErrorResponse).Message, failMessage)
}

func TestGetPullRequestCommits(t *testing.T) {
	wantCommits := []*github.RepositoryCommit{
		{
			Commit: &github.Commit{
				Message: github.String("Lorem Ipsum"),
			},
		},
	}
	mockedGithubClient := aladino.MockDefaultGithubClient(
		[]mock.MockBackendOption{
			mock.WithRequestMatch(
				mock.GetReposPullsCommitsByOwnerByRepoByPullNumber,
				wantCommits,
			),
		},
		nil,
	)

	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	gotCommits, err := mockedGithubClient.GetPullRequestCommits(
		context.Background(),
		mockedPullRequest.GetUser().GetLogin(),
		mockedPullRequest.GetBase().GetRepo().GetName(),
		mockedPullRequest.GetNumber(),
	)

	assert.Nil(t, err)
	assert.Equal(t, wantCommits, gotCommits)
}

func TestGetPullRequestReviews(t *testing.T) {
	wantReviews := []*github.PullRequestReview{
		{
			State: github.String("COMMENTED"),
		},
		{
			State: github.String("COMMENTED"),
		},
	}
	mockedGithubClient := aladino.MockDefaultGithubClient(
		[]mock.MockBackendOption{
			mock.WithRequestMatch(
				mock.GetReposPullsReviewsByOwnerByRepoByPullNumber,
				wantReviews,
			),
		},
		nil,
	)

	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	gotReviews, err := mockedGithubClient.GetPullRequestReviews(
		context.Background(),
		mockedPullRequest.GetUser().GetLogin(),
		mockedPullRequest.GetBase().GetRepo().GetName(),
		mockedPullRequest.GetNumber(),
	)

	assert.Nil(t, err)
	assert.Equal(t, wantReviews, gotReviews)
}

func TestGetPullRequestReviews_WhenRequestFails(t *testing.T) {
	failMessage := "ListPullRequestReviewsFail"
	mockedGithubClient := aladino.MockDefaultGithubClient(
		[]mock.MockBackendOption{
			mock.WithRequestMatchHandler(
				mock.GetReposPullsReviewsByOwnerByRepoByPullNumber,
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					mock.WriteError(
						w,
						http.StatusInternalServerError,
						failMessage,
					)
				}),
			),
		},
		nil,
	)

	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	gotReviews, err := mockedGithubClient.GetPullRequestReviews(
		context.Background(),
		mockedPullRequest.GetUser().GetLogin(),
		mockedPullRequest.GetBase().GetRepo().GetName(),
		mockedPullRequest.GetNumber(),
	)

	assert.Nil(t, gotReviews)
	assert.Equal(t, err.(*github.ErrorResponse).Message, failMessage)
}

func TestGetPullRequests(t *testing.T) {
	ownerName := "testOrg"
	repoName := "testRepo"

	wantPullRequests := []*github.PullRequest{}

	mockedGithubClient := aladino.MockDefaultGithubClient(
		[]mock.MockBackendOption{
			mock.WithRequestMatch(
				mock.GetReposPullsByOwnerByRepo,
				wantPullRequests,
			),
		},
		nil,
	)

	gotReviews, err := mockedGithubClient.GetPullRequests(
		context.Background(),
		ownerName,
		repoName,
	)

	assert.Nil(t, err)
	assert.Equal(t, wantPullRequests, gotReviews)
}

func TestGetPullRequests_WhenRequestFails(t *testing.T) {
	failMessage := "ListPullRequests"

	ownerName := "testOrg"
	repoName := "testRepo"

	mockedGithubClient := aladino.MockDefaultGithubClient(
		[]mock.MockBackendOption{
			mock.WithRequestMatchHandler(
				mock.GetReposPullsByOwnerByRepo,
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					mock.WriteError(
						w,
						http.StatusInternalServerError,
						failMessage,
					)
				}),
			),
		},
		nil,
	)

	gotReviews, err := mockedGithubClient.GetPullRequests(
		context.Background(),
		ownerName,
		repoName,
	)

	assert.Nil(t, gotReviews)
	assert.Equal(t, err.(*github.ErrorResponse).Message, failMessage)
}

func TestGetReviewThreads_WhenRequestFails(t *testing.T) {
	failMessage := "GetReviewThreads"
	mockedGithubClient := aladino.MockDefaultGithubClient(
		nil,
		func(w http.ResponseWriter, req *http.Request) {
			http.Error(w, failMessage, http.StatusNotFound)
		},
	)

	gotThreads, err := mockedGithubClient.GetReviewThreads(
		context.Background(),
		aladino.DefaultMockPrOwner,
		aladino.DefaultMockPrRepoName,
		aladino.DefaultMockPrNum,
		2,
	)

	assert.Nil(t, gotThreads)
	assert.Equal(t, err.Error(), fmt.Sprintf("non-200 OK status code: 404 Not Found body: \"%s\\n\"", failMessage))
}

func TestGetReviewThreads(t *testing.T) {
	mockedGraphQLQuery := fmt.Sprintf(
		"{\"query\":\"query($pullRequestNumber:Int!$repositoryName:String!$repositoryOwner:String!$reviewThreadsCursor:String){repository(owner: $repositoryOwner, name: $repositoryName){pullRequest(number: $pullRequestNumber){reviewThreads(first: 10, after: $reviewThreadsCursor){nodes{isResolved,isOutdated},pageInfo{endCursor,hasNextPage}}}}}\",\"variables\":{\"pullRequestNumber\":%d,\"repositoryName\":\"%s\",\"repositoryOwner\":\"%s\",\"reviewThreadsCursor\":null}}\n",
		aladino.DefaultMockPrNum,
		aladino.DefaultMockPrRepoName,
		aladino.DefaultMockPrOwner,
	)

	mockedGithubClient := aladino.MockDefaultGithubClient(
		nil,
		func(w http.ResponseWriter, req *http.Request) {
			query := aladino.MustRead(req.Body)
			switch query {
			case mockedGraphQLQuery:
				aladino.MustWrite(
					w,
					`{"data": {
                        "repository": {
                            "pullRequest": {
                                "reviewThreads": {
                                    "nodes": [{
                                        "isResolved": true,
                                        "isOutdated": false
                                    }]
                                }
                            }
                        }
                    }}`,
				)
			}
		},
	)

	wantReviewThreads := []host.GQLReviewThread{{
		IsResolved: true,
		IsOutdated: false,
	}}
	gotReviewThreads, err := mockedGithubClient.GetReviewThreads(
		context.Background(),
		aladino.DefaultMockPrOwner,
		aladino.DefaultMockPrRepoName,
		aladino.DefaultMockPrNum,
		2,
	)

	assert.Nil(t, err)
	assert.Equal(t, gotReviewThreads, wantReviewThreads)

}

func TestGetIssueTimeline_WhenListIssueTimelineRequestFails(t *testing.T) {
	failMessage := "ListIssueTimelineRequestFail"
	mockedGithubClient := aladino.MockDefaultGithubClient(
		[]mock.MockBackendOption{
			mock.WithRequestMatchHandler(
				mock.GetReposIssuesTimelineByOwnerByRepoByIssueNumber,
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					mock.WriteError(
						w,
						http.StatusInternalServerError,
						failMessage,
					)
				}),
			),
		},
		nil,
	)

	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	gotTimeline, err := mockedGithubClient.GetIssueTimeline(
		context.Background(),
		mockedPullRequest.GetBase().GetRepo().GetOwner().GetLogin(),
		mockedPullRequest.GetBase().GetRepo().GetName(),
		mockedPullRequest.GetNumber(),
	)

	assert.Nil(t, gotTimeline)
	assert.Equal(t, err.(*github.ErrorResponse).Message, failMessage)
}

func TestGetIssueTimeLine(t *testing.T) {
	nonLastEventDate := time.Date(2022, 04, 13, 20, 49, 13, 651387237, time.UTC)
	lastEventDate := time.Date(2022, 04, 16, 20, 49, 34, 0, time.UTC)
	wantTimeline := []*github.Timeline{
		{
			ID:        github.Int64(6430295168),
			Event:     github.String("locked"),
			CreatedAt: &nonLastEventDate,
		},
		{
			ID:        github.Int64(6430296748),
			Event:     github.String("labeled"),
			CreatedAt: &lastEventDate,
		},
	}
	mockedGithubClient := aladino.MockDefaultGithubClient(
		[]mock.MockBackendOption{
			mock.WithRequestMatch(
				mock.GetReposIssuesTimelineByOwnerByRepoByIssueNumber,
				wantTimeline,
			),
		},
		nil,
	)

	mockedPullRequest := aladino.GetDefaultMockPullRequestDetails()
	gotTimeline, err := mockedGithubClient.GetIssueTimeline(
		context.Background(),
		mockedPullRequest.GetBase().GetRepo().GetOwner().GetLogin(),
		mockedPullRequest.GetBase().GetRepo().GetName(),
		mockedPullRequest.GetNumber(),
	)

	assert.Nil(t, err)
	assert.Equal(t, wantTimeline, gotTimeline)
}
