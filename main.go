package main

import (
	"fmt"
	"janic0/gemserv/actions"
	"janic0/gemserv/generators"
	"janic0/gemserv/utils"
	"janic0/gemserv/views"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/cvhariharan/gemini-server"
	"github.com/google/uuid"
)

func main() {
	gemini.HandleFunc("/", func(w *gemini.Response, r *gemini.Request) {
		start := time.Now()
		defer fmt.Println(start.String(), "\t", w.Body.RemoteAddr().String(), "\t", "\t", r.URL.String(), "\t", time.Since(start).Microseconds(), "ms")
		adminUrl := fmt.Sprintf("/%s/admin", utils.AdminSession)
		if r.URL.Path == "/admin" {
			actions.Login(w, r)
		} else if strings.HasPrefix(r.URL.Path, adminUrl) || strings.HasPrefix(r.URL.Path, fmt.Sprintf("/%s/admin", utils.LastAdminSession)) {
			utils.LastAdminSession = utils.AdminSession
			utils.AdminSession = uuid.NewString()
			path := strings.Trim(r.URL.Path[len(adminUrl):], "/")
			if path == "" {
				views.AdminDashboard(w, r)
			} else if path == "add-page" {
				actions.AddPage(w, r)
			} else if strings.HasPrefix(path, "edit-page") {
				relativePath := path[len("edit-page"):]
				if relativePath == "" {
					relativePath = "/"
				}
				views.EditPage(w, r, relativePath)
			} else if strings.HasPrefix(path, "page/") {
				pageSuffixes := map[string]func(w *gemini.Response, r *gemini.Request, path string){
					"/enable":    actions.EnablePage,
					"/disable":   actions.DisablePage,
					"/move":      actions.MovePage,
					"/duplicate": actions.DuplicatePage,
					"/delete":    actions.DeletePage,
					"/add-block": actions.AddBlock,
				}
				for suffix, handler := range pageSuffixes {
					if strings.HasSuffix(path, suffix) {
						page := path[len("page") : len(path)-len(suffix)]
						handler(w, r, page)
						return
					}
				}
				// Edit block
				blockSuffixes := map[string]func(w *gemini.Response, r *gemini.Request, path string, block int64){
					"/edit":         actions.EditBlock,
					"/insert-above": actions.AddBlockAbove,
					"/remove":       actions.DelteBlock,
				}
				for suffix, handler := range blockSuffixes {
					if strings.HasSuffix(path, suffix) {
						segments := strings.Split(path, "/")
						if segments[len(segments)-3] != "blocks" {
							return
						}
						pagePath := "/" + strings.Join(segments[1:len(segments)-3], "/")
						blockIndexStr := segments[len(segments)-2]
						blockIndex, err := strconv.Atoi(blockIndexStr)
						if err != nil {
							w.SetStatus(gemini.StatusTemporaryFailure, "Invalid URL parameter")
						}
						handler(w, r, pagePath, int64(blockIndex))
						return
					}
				}
			} else {
				w.SetStatus(51, "Page not found.")
				w.SendStatus()
			}
		} else {
			truncatedPath := strings.TrimRight(r.URL.Path, "/")
			content, err := generators.GetPage(truncatedPath)
			if err != nil {
				w.SetStatus(51, "Page not found.")
				w.SendStatus()
				return
			}
			w.SetStatus(gemini.StatusSuccess, "text/gemini")
			w.Write(content)
		}
	})

	log.Fatal(gemini.ListenAndServeTLS(":1965", "cert.crt", "cert.key"))
}
