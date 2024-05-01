package server

import (
	"fmt"
	"net/http"
)

type SitemapEntry struct {
	Location        string
	ChangeFrequency string
}

func (s *WebserverService) registerSiteMap() {
	s.router.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/xml")
		w.Write([]byte(s.generateSiteMapXML()))
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

func (s *WebserverService) generateDynamicSiteMap() []SitemapEntry {
	packages := s.packageCacheService.GetAll()
	entries := make([]SitemapEntry, len(packages))

	for idx, pkg := range packages {
		entries[idx] = SitemapEntry{
			Location:        fmt.Sprintf("https://pacstall.dev/packages/%s/", pkg.PackageName),
			ChangeFrequency: "monthly",
		}
	}

	return entries
}

func (entry SitemapEntry) generateSiteMapUrls() string {
	return fmt.Sprintf(`
	<url>
		<loc>%s</loc>
		<changefreq>%s</changefreq>
	</url>`, entry.Location, entry.ChangeFrequency)
}

func (s *WebserverService) generateSiteMapXML() string {
	entries := generateStaticSiteMap()
	entries = append(entries, s.generateDynamicSiteMap()...)

	urls := ""
	for _, entry := range entries {
		urls += entry.generateSiteMapUrls() + "\n"
	}

	return fmt.Sprintf(`
	<?xml version="1.0" encoding="UTF-8"?>
	<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
		%s	
	</urlset> 
	`, urls)
}
