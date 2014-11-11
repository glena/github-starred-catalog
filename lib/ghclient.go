package lib

//https://github.com/google/go-github

import (
	"fmt"
	"github.com/google/go-github/github"
)

func InitGHClient() *GHClient {

	client := github.NewClient(nil)

	gh := &GHClient{
		Client:  client,
		PerPage: 100,
	}

	return gh
}

type GHClient struct {
	Username string
	PerPage  int
	Type     string
	Client   *github.Client
	Catalog  map[string][]github.Repository
}

func (me *GHClient) loadRepos(page int) bool {

	opt := &github.ActivityListStarredOptions{
		ListOptions: github.ListOptions{PerPage: me.PerPage, Page: page},
	}

	repos, _, err := me.Client.Activity.ListStarred(me.Username, opt)

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	}

	fmt.Println("page", page, "count", len(repos), me.PerPage)

	me.processRepos(repos)

	return len(repos) == me.PerPage
}

func (me *GHClient) Load(Username string) {

	me.Catalog = make(map[string][]github.Repository)

	page := 1

	me.Username = Username

	for me.loadRepos(page) {
		page++
	}
}

func (me *GHClient) processRepos(repos []github.Repository) {
	var language string
	for i := range repos {
		repo := repos[i]
		if repo.Language == nil {
			language = "Empty"
		} else {
			language = *repo.Language
		}

		me.Catalog[language] = me.addRepository(me.Catalog[language], repo)
	}
}

func (me *GHClient) addRepository(array []github.Repository, repo github.Repository) []github.Repository {

	newArray := make([]github.Repository, len(array)+1)

	for i := range array {
		newArray[i] = array[i]
	}

	newArray[len(newArray)-1] = repo
	return newArray
}

func (me *GHClient) GetLanguages() []string {

	languages := make([]string, len(me.Catalog))
	position := 0

	for key, _ := range me.Catalog {
		languages[position] = key
		position++
	}

	return languages
}

func (me *GHClient) GetRepositories(language string) []github.Repository {
	return me.Catalog[language]
}
