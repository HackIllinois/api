package service

import (
	"errors"
	"github.com/HackIllinois/api/services/auth/config"
	"github.com/HackIllinois/api/services/auth/models"
	"github.com/levigross/grequests"
	"strconv"
	"strings"
)

type GitHubOAuthProvider struct {
	token          string
	isVerifiedUser bool
}

func NewGitHubOAuth() *GitHubOAuthProvider {
	return &GitHubOAuthProvider{
		token:          "",
		isVerifiedUser: false,
	}
}

/*
	Returns the url to redirects to for OAuth authorization
*/
func (provider *GitHubOAuthProvider) GetAuthorizationRedirect(redirect_uri string) (string, error) {
	return ConstructSafeURL("https", "github.com", "login/oauth/authorize",
		map[string]string{
			"client_id":    config.GITHUB_CLIENT_ID,
			"scope":        "user:email",
			"redirect_uri": redirect_uri,
		},
	)
}

/*
	Exchanges an OAuth code for an OAuth token
*/
func (provider *GitHubOAuthProvider) Authorize(code string, redirect_uri string) error {
	request, err := grequests.Post("https://github.com/login/oauth/access_token", &grequests.RequestOptions{
		Params: map[string]string{
			"client_id":     config.GITHUB_CLIENT_ID,
			"client_secret": config.GITHUB_CLIENT_SECRET,
			"code":          code,
		},
		Headers: map[string]string{
			"Accept": "application/json",
		},
	})

	if err != nil {
		return err
	}

	var oauth_token models.GithubOauthToken
	err = request.JSON(&oauth_token)

	if err != nil {
		return err
	}

	if oauth_token.Token == "" {
		return errors.New("Invalid oauth code")
	}

	provider.token = oauth_token.Token

	return nil
}

/*
	Retrieves user info from the OAuth provider
*/
func (provider *GitHubOAuthProvider) GetUserInfo() (*models.UserInfo, error) {
	var user_info models.UserInfo

	request, err := grequests.Get("https://api.github.com/user", &grequests.RequestOptions{
		Headers: map[string]string{
			"Authorization": "token " + provider.token,
		},
	})

	if err != nil {
		return nil, err
	}

	var github_user_info models.GithubUserInfo
	err = request.JSON(&github_user_info)

	if err != nil {
		return nil, err
	}

	if github_user_info.ID == 0 {
		return nil, errors.New("Invalid oauth token")
	}

	user_info.Username = github_user_info.Username
	user_info.ID = "github" + strconv.Itoa(github_user_info.ID)

	split_name := strings.SplitAfterN(github_user_info.Name, " ", 2)

	user_info.FirstName = strings.TrimSpace(split_name[0])

	if len(split_name) < 2 {
		user_info.LastName = ""
	} else {
		user_info.LastName = strings.TrimSpace(split_name[1])
	}

	request, err = grequests.Get("https://api.github.com/user/emails", &grequests.RequestOptions{
		Headers: map[string]string{
			"Authorization": "token " + provider.token,
		},
	})

	if err != nil {
		return nil, err
	}

	var emails []models.GithubEmail
	err = request.JSON(&emails)

	if err != nil {
		return nil, err
	}

	for _, email := range emails {
		if email.IsPrimary {
			user_info.Email = email.Email
			provider.isVerifiedUser = email.IsVerified
		}
	}

	return &user_info, nil
}

/*
	Returns true if the user has a verified email
*/
func (provider *GitHubOAuthProvider) IsVerifiedUser() bool {
	return provider.isVerifiedUser
}
