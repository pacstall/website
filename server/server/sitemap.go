package server

import (
	"fmt"
	"net/http"
	"time"

	"pacstall.dev/webserver/types/pac/pacstore"
)

type SitemapEntry struct {
	Location        string
	ChangeFrequency string
	LastUpdated     *string
}

func registerSiteMap() {
	Router().HandleFunc("/sitemap.xml", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/xml")
		w.Write([]byte(generateSiteMapXML()))

	}).Methods("GET")
}

func generateStaticSiteMap() []SitemapEntry {
	return []SitemapEntry{
		{
			Location:        "https://pacstall.dev/",
			ChangeFrequency: "monthly",
		},
		{
			Location:        "https://pacstall.dev/privacy/",
			ChangeFrequency: "yearly",
		},
		{
			Location:        "https://pacstall.dev/packages/",
			ChangeFrequency: "daily",
		},
	}
}

func generateDynamicSiteMap() []SitemapEntry {
	packages := pacstore.GetAll()
	entries := make([]SitemapEntry, len(packages))

	for idx, pkg := range packages {
        lastUpdated := pkg.LastUpdatedAt.Format(time.RFC3339)
		entries[idx] = SitemapEntry{
			Location:        fmt.Sprintf("https://pacstall.dev/packages/%s/", pkg.PackageName),
			ChangeFrequency: "monthly",
			LastUpdated:     &lastUpdated,
		}
	}

	return entries
}

func (entry SitemapEntry) generateSiteMapUrls() string {
	return fmt.Sprintf(`<url>
	<loc>%s</loc>
	<changefreq>%s</changefreq>
    %s
</url>`, entry.Location, entry.ChangeFrequency, func() string {
		if entry.LastUpdated != nil {
			return fmt.Sprintf("<lastmod>%s</lastmod>", *entry.LastUpdated)
		}
		return ""
	}())
}

func generateSiteMapXML() string {
	entries := generateStaticSiteMap()
	entries = append(entries, generateDynamicSiteMap()...)

	urls := ""
	for _, entry := range entries {
		urls += entry.generateSiteMapUrls() + "\n"
	}

	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
	%s
</urlset>
	`, urls)
}
