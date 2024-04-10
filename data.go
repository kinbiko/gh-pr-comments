package ghprcomments

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

type User struct {
	Login string `json:"login"`
}

type HRef struct {
	Href string `json:"href"`
}

type Links struct {
	Self        HRef `json:"self"`
	HTML        HRef `json:"html"`
	PullRequest HRef `json:"pull_request"`
}

type Reactions struct {
	Num1     int `json:"+1"`
	Num10    int `json:"-1"`
	Laugh    int `json:"laugh"`
	Hooray   int `json:"hooray"`
	Confused int `json:"confused"`
	Heart    int `json:"heart"`
	Rocket   int `json:"rocket"`
	Eyes     int `json:"eyes"`
}

type Side string

const (
	left  Side = "LEFT"
	right Side = "RIGHT"
)

type PRComment struct {
	URL                 string    `json:"url"`
	PullRequestReviewID int       `json:"pull_request_review_id"`
	ID                  int       `json:"id"`
	DiffHunk            string    `json:"diff_hunk"`
	Path                string    `json:"path"`
	CommitID            string    `json:"commit_id"`
	OriginalCommitID    string    `json:"original_commit_id"`
	User                User      `json:"user"`
	Body                string    `json:"body"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	AuthorAssociation   string    `json:"author_association"`
	Links               Links     `json:"_links"`
	Reactions           Reactions `json:"reactions"`

	StartLine         *int `json:"start_line"`
	OriginalStartLine *int `json:"original_start_line"`
	Line              int  `json:"line"`
	OriginalLine      int  `json:"original_line"`

	StartSide *Side `json:"start_side"`
	Side      Side  `json:"side"`

	OriginalPosition int `json:"original_position"`
	Position         int `json:"position"`

	SubjectType string `json:"subject_type"`
	InReplyToID int    `json:"in_reply_to_id,omitempty"`
}

func (c *PRComment) ThreadID() string {
	// This is known to be slightly false, as it's possible to have multiple
	// comment threads on the same line.
	return fmt.Sprintf("%s:%d", c.Path, c.Line)
}

func (c *PRComment) String() string {
	body := strings.ReplaceAll(c.Body, "\n", "\n\t\t\t")
	return fmt.Sprintf("\t@%s:\t%s", color.RedString(c.User.Login), body)
}
