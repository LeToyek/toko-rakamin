package utils

import (
	"regexp"
	"strings"
)

func GenerateSlug(title string) string {

	slug := strings.ToLower(title)

	reg := regexp.MustCompile(`[^a-z0-9]+`)
	slug = reg.ReplaceAllString(slug, " ")

	slug = strings.TrimSpace(slug)

	slug = strings.ReplaceAll(slug, " ", "-")

	return slug
}
