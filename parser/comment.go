package parser

import (
	"go/ast"
	"regexp"
)

var genRegex = regexp.MustCompile(`\/\/\s*replica:gen`)

func CheckGen(comment string) bool {
	return genRegex.MatchString(comment)
}

func CheckComments(comments *ast.CommentGroup) bool {
	if comments == nil {
		return false
	}

	for _, comment := range comments.List {
		if CheckGen(comment.Text) {
			return true
		}
	}

	return false
}
