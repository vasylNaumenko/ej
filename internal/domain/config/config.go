package config

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//go:generate go-validator

type (
	Config struct {
		BaseUrl      string    `valid:"required"`
		User         string    `valid:"required"`
		Token        string    `valid:"required"`
		MyAccountTag string    `valid:"required"`
		Reviewers    Assignees `valid:"check,deep"`

		DiscordWebHook string `valid:"required"`
	}

	Assignee struct {
		ID        string `valid:"required"`
		Tag       string `valid:"required"`
		Name      string `valid:"required"`
		DiscordID string `valid:"required"`
	}

	Assignees []Assignee
)

var ErrAssigneeNotFoundByTag = errors.New("assignee is not found by the tag")
var ErrAssigneeListEmpty = errors.New("an assignee list is empty")
var ErrCountTooHigh = errors.New("a count of random assignee is too high")

// GetByTag returns assignee by tag
func (a Assignees) GetByTag(tag string) (Assignee, error) {
	for _, assignee := range a {
		if assignee.Tag == tag {
			return assignee, nil
		}
	}

	return Assignee{}, fmt.Errorf("%w: %s", ErrAssigneeNotFoundByTag, tag)
}

// GetRandom returns a random assignee and shrinks a collection.
func (a *Assignees) GetRandom(count int, excludeIDs ...Assignee) (Assignees, error) {
	total := len(*a)
	if total == 0 || count < 1 {
		return Assignees{}, ErrAssigneeListEmpty
	}

	total -= len(excludeIDs)
	if count > total {
		return Assignees{}, ErrCountTooHigh
	}

	excluded := make(map[string]struct{}, len(excludeIDs))
	for _, x := range excludeIDs {
		excluded[x.ID] = struct{}{}
	}

	// Removing all excluded ID's from the list
	var list Assignees
	for _, assignee := range *a {
		if _, ok := excluded[assignee.ID]; !ok {
			list = append(list, assignee)
		}
	}

	seed := rand.NewSource(time.Now().Unix())
	random := rand.New(seed)

	var assignees Assignees
	for ; count > 0; count-- {
		lenList := len(list)
		index := random.Intn(lenList)
		assignees = append(assignees, list[index])
		list = append(list[:index], list[index+1:]...)
	}

	return assignees, nil
}

// GetSliceOfIDs returns a slice of strings with the assignee's ID's
func (a Assignees) GetSliceOfIDs() []string {
	var res []string

	for _, assignee := range a {
		res = append(res, assignee.ID)
	}

	return res
}
