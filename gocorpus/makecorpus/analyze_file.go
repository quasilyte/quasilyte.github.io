package main

import (
	"go/ast"
	"strconv"
	"strings"
)

type repositoryFileInfo struct {
	isTest         bool
	isAutogen      bool
	isMain         bool
	importsC       bool
	importsUnsafe  bool
	importsReflect bool
}

func analyzeFile(filename string, f *ast.File, src []byte) *repositoryFileInfo {
	info := &repositoryFileInfo{}

	info.isTest = strings.HasSuffix(f.Name.String(), "_test") ||
		strings.HasSuffix(filename, "_test.go")

	info.isMain = f.Name.String() == "main"

	for _, comment := range f.Comments {
		if isAutogenComment(comment) {
			info.isAutogen = true
			break
		}
	}

	for _, imp := range f.Imports {
		path, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			panic(err) // should never happen
		}
		switch path {
		case "C":
			info.importsC = true
		case "unsafe":
			info.importsUnsafe = true
		case "reflect":
			info.importsReflect = true
		}
	}

	return info
}

func isAutogenComment(comment *ast.CommentGroup) bool {
	generated := false
	doNotEdit := false
	for _, c := range comment.List {
		s := strings.ToLower(c.Text)
		if !generated {
			generated = strings.Contains(s, " code generated ") ||
				strings.Contains(s, " generated by ")
		}
		if !doNotEdit {
			doNotEdit = strings.Contains(s, "do not edit") ||
				strings.Contains(s, "don't edit")
		}
		if generated && doNotEdit {
			return true
		}
	}
	return false
}