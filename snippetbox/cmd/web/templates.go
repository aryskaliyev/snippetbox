package main

import "lincoln.boris/snippetbox/pkg/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}
