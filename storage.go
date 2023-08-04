package appwrite

import (
	"encoding/json"
	"strings"
)

// Storage service
type Storage struct {
	Client Client
}

func NewStorage(clt Client) Storage {
	service := Storage{
		Client: clt,
	}

	return service
}

type Bucket struct {
	Id                    string   `json:"$id"`
	Name                  string   `json:"name,omitempty"`
	CreatedAt             string   `json:"$createdAt"`
	UpdatedAt             string   `json:"$updatedAt"`
	Permissions           []string `json:"permissions"`
	FileSecurity          bool     `json:"fileSecurity"`
	Enabled               bool     `json:"enabled"`
	MaximumFileSize       int      `json:"maximumFileSize"`
	AllowedFileExtensions []string `json:"allowedFileExtensions"`
	CompressionType       string   `json:"compression"`
	Encryption            bool     `json:"encryption"`
	Antivirus             bool     `json:"antivirus"`
}

type BucketListResponse struct {
	Total   int      `json:"total"`
	Buckets []Bucket `json:"buckets"`
}

type File struct {
	Id             string   `json:"$id"`
	BucketId       string   `json:"bucketId"`
	CreatedAt      string   `json:"$createdAt"`
	UpdatedAt      string   `json:"$updatedAt"`
	Permissions    []string `json:"permissions"`
	Name           string   `json:"name"`
	Signature      string   `json:"signature"`
	MimeType       string   `json:"mimeType"`
	SizeOriginal   int      `json:"sizeOriginal"`
	ChunksTotal    int      `json:"chunksTotal"`
	ChunksUploaded int      `json:"chunksUploaded"`
}

type FileListResponse struct {
	Total int    `json:"total"`
	Files []File `json:"files"`
}

// ListBuckets get all the bucket in the project. This endpoint response returns a JSON
// object with the list of bucket objects.
func (srv *Storage) ListBuckets(Search string, Limit int, Offset int, OrderType string) (*BucketListResponse, error) {
	path := "/storage/buckets/"

	params := map[string]interface{}{
		"search":    Search,
		"limit":     Limit,
		"offset":    Offset,
		"orderType": OrderType,
	}

	resp, err := srv.Client.CallAPI("GET", path, srv.Client.headers, params)
	if err != nil {
		return nil, err
	}
	var result BucketListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetBucket get bucket by its unique ID. This endpoint response returns a JSON
// object with the bucket metadata.
func (srv *Storage) GetBucket(bucketId string) (*Bucket, error) {
	r := strings.NewReplacer("{bucketId}", bucketId)
	path := r.Replace("/storage/buckets/{bucketId}")

	params := map[string]interface{}{}

	resp, err := srv.Client.CallAPI("GET", path, srv.Client.headers, params)
	if err != nil {
		return nil, err
	}
	var result Bucket
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListFiles get a list of all the user files. You can use the query params to
// filter your results. On admin mode, this endpoint will return a list of all
// of the project files. [Learn more about different API modes](/docs/admin).
func (srv *Storage) ListFiles(bucketId, Search string, Limit int, Offset int, OrderType string) (*FileListResponse, error) {
	r := strings.NewReplacer("{bucketId}", bucketId)
	path := r.Replace("/storage/buckets/{bucketId}/files")

	params := map[string]interface{}{
		"search":    Search,
		"limit":     Limit,
		"offset":    Offset,
		"orderType": OrderType,
	}
	resp, err := srv.Client.CallAPI("GET", path, srv.Client.headers, params)
	if err != nil {
		return nil, err
	}
	var result FileListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateFile create a new file. The user who creates the file will
// automatically be assigned to read and write access unless he has passed
// custom values for read and write arguments.
func (srv *Storage) CreateFile(File string, Read []interface{}, Write []interface{}) (map[string]interface{}, error) {
	path := "/storage/files"

	params := map[string]interface{}{
		"file":  File,
		"read":  Read,
		"write": Write,
	}

	return srv.Client.Call("POST", path, nil, params)
}

// GetFile get file by its unique ID. This endpoint response returns a JSON
// object with the file metadata.
func (srv *Storage) GetFile(bucketId, fileId string) (*File, error) {
	r := strings.NewReplacer("{bucketId}", bucketId, "{fileId}", fileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}")

	params := map[string]interface{}{}

	resp, err := srv.Client.CallAPI("GET", path, srv.Client.headers, params)
	if err != nil {
		return nil, err
	}
	var result File
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateFile update file by its unique ID. Only users with write permissions
// have access to update this resource.
func (srv *Storage) UpdateFile(FileId string, Read []interface{}, Write []interface{}) (map[string]interface{}, error) {
	r := strings.NewReplacer("{fileId}", FileId)
	path := r.Replace("/storage/files/{fileId}")

	params := map[string]interface{}{
		"read":  Read,
		"write": Write,
	}

	return srv.Client.Call("PUT", path, nil, params)
}

// DeleteFile delete a file by its unique ID. Only users with write
// permissions have access to delete this resource.
func (srv *Storage) DeleteFile(FileId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{fileId}", FileId)
	path := r.Replace("/storage/files/{fileId}")

	params := map[string]interface{}{}

	return srv.Client.Call("DELETE", path, nil, params)
}

// GetFileDownload get file content by its unique ID. The endpoint response
// return with a 'Content-Disposition: attachment' header that tells the
// browser to start downloading the file to user downloads directory.
func (srv *Storage) GetFileDownload(FileId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{fileId}", FileId)
	path := r.Replace("/storage/files/{fileId}/download")

	params := map[string]interface{}{}

	return srv.Client.Call("GET", path, nil, params)
}

// GetFilePreview get a file preview image. Currently, this method supports
// preview for image files (jpg, png, and gif), other supported formats, like
// pdf, docs, slides, and spreadsheets, will return the file icon image. You
// can also pass query string arguments for cutting and resizing your preview
// image.
func (srv *Storage) GetFilePreview(FileId string, Width int, Height int, Quality int, Background string, Output string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{fileId}", FileId)
	path := r.Replace("/storage/files/{fileId}/preview")

	params := map[string]interface{}{
		"width":      Width,
		"height":     Height,
		"quality":    Quality,
		"background": Background,
		"output":     Output,
	}

	return srv.Client.Call("GET", path, nil, params)
}

// GetFileView get file content by its unique ID. This endpoint is similar to
// the download method but returns with no  'Content-Disposition: attachment'
// header.
func (srv *Storage) GetFileView(FileId string, As string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{fileId}", FileId)
	path := r.Replace("/storage/files/{fileId}/view")

	params := map[string]interface{}{
		"as": As,
	}

	return srv.Client.Call("GET", path, nil, params)
}
