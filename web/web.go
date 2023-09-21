package web

import (
	"embed"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed dist/*
var embeddedFiles embed.FS

func Static(r *gin.Engine) {
	// Serve the Svelte app from the embedded filesystem
	r.NoRoute(func(c *gin.Context) {
		// Get the requested path
		requestedPath := c.Request.URL.Path

		// If the path ends with a slash, assume it's a directory and append "index.html"
		if strings.HasSuffix(requestedPath, "/") {
			requestedPath = path.Join(requestedPath, "index.html")
		}

		// Construct the full path in the embedded "dist" directory
		fullPath := path.Join("dist", requestedPath)

		// Open the file from the embedded filesystem
		fileContents, err := embeddedFiles.ReadFile(fullPath)
		if err != nil {
			// File not found, handle it as needed (e.g., return 404)
			c.String(http.StatusNotFound, "File not found")
			return
		}

		// Set the Content-Type header based on the file extension
		contentType := http.DetectContentType(fileContents)

		// Set specific Content-Type for stylesheets and images
		switch {
		case strings.HasSuffix(requestedPath, ".js"):
			contentType = "application/javascript"
		case strings.HasSuffix(requestedPath, ".css"):
			contentType = "text/css"
		case strings.HasSuffix(requestedPath, ".svg"):
			contentType = "image/svg+xml"
		case strings.HasSuffix(requestedPath, ".jpg"), strings.HasSuffix(requestedPath, ".jpeg"):
			contentType = "image/jpeg"
		}

		// Serve the file with its content type
		c.Data(http.StatusOK, contentType, fileContents)
	})
}
