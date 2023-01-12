package common

import (
	"io/ioutil"

	"github.com/juunys/go-webcrawler/entity"
	"gopkg.in/yaml.v3"
)

type App struct {
	Source []*entity.FeedYml
}

func NewApp() (*App, error) {
	source, err := ReadYaml()
	if err != nil {
		return nil, err
	}

	return &App{
		Source: source,
	}, nil
}

func ReadYaml() ([]*entity.FeedYml, error) {
	var feedsSource []*entity.FeedYml
	source, err := ioutil.ReadFile("../feed.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(source, &feedsSource)
	if err != nil {
		return nil, err
	}

	return feedsSource, nil
}
