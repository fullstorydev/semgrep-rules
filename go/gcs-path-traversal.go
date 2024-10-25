package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

const (
	bucketName         = "user-data"
	projectID          = "my-project-id"
	serviceAccountPath = "path/to/service-account.json"
)

type GCSClient struct {
	client *storage.Client
}

func main() {
	http.HandleFunc("/gcs", handleGCSRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleGCSRequest(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	filepath := r.URL.Query().Get("filepath")

	if username == "" || filepath == "" {
		http.Error(w, "Missing username or filepath parameter", http.StatusBadRequest)
		return
	}

	fullPath := fmt.Sprintf("users/%s/%s", username, filepath)

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(serviceAccountPath))
	if err != nil {
		http.Error(w, "Failed to create GCS client: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	gcsClient := &GCSClient{client}

	fileContent, err := gcsClient.readFile(ctx, fullPath)
	if err != nil {
		http.Error(w, "Failed to read file from GCS: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File content read from GCS:\n%s", string(fileContent))

	// Optionally, write something back to GCS
	newContent := fmt.Sprintf("%s\nAdded some content!", string(fileContent))
	gcsClient.writeFile(ctx, fullPath, []byte(newContent))
}

func handleGCSRequest_FP(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	filepath := r.URL.Query().Get("filepath")

	if username == "" || filepath == "" {
		http.Error(w, "Missing username or filepath parameter", http.StatusBadRequest)
		return
	}

	// username = sanitizeInput(username)
	// filepath = sanitizeInput(filepath)

	fullPath := fmt.Sprintf("users/%s/%s", username, filepath)

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(serviceAccountPath))
	if err != nil {
		http.Error(w, "Failed to create GCS client: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	gcsClient := &GCSClient{client}

	fileContent, err := gcsClient.readFile(ctx, fullPath)
	if err != nil {
		http.Error(w, "Failed to read file from GCS: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File content read from GCS:\n%s", string(fileContent))

	// Optionally, write something back to GCS
	newContent := fmt.Sprintf("%s\nAdded some content!", string(fileContent))
	gcsClient.writeFile(ctx, fullPath, []byte(newContent))
}

func (g *GCSClient) readFile(ctx context.Context, filePath string) ([]byte, error) {
	bucket := g.client.Bucket(bucketName)
	// proruleid: gcs-path-traversal
	obj := bucket.Object(filePath)

	reader, err := obj.NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to open file %s: %v", filePath, err)
	}
	defer reader.Close()

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("unable to read file %s: %v", filePath, err)
	}

	return data, nil
}

func (g *GCSClient) writeFile(ctx context.Context, filePath string, data []byte) error {
	bucket := g.client.Bucket(bucketName)
	// ruleid: gcs-path-traversal
	obj := bucket.Object(filePath)

	writer := obj.NewWriter(ctx)
	defer writer.Close()

	if _, err := writer.Write(data); err != nil {
		return fmt.Errorf("unable to write file %s: %v", filePath, err)
	}

	return nil
}

func sanitizeInput(input string) string {
	// Allow only alphanumeric characters, dashes, and underscores
	re := regexp.MustCompile(`[^\w\-\/\.]`)
	sanitized := re.ReplaceAllString(input, "")
	return sanitized
}
