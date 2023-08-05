package appwrite

import (
	"encoding/json"
	"strings"
)

// Database service
type Database struct {
	Client Client
}

func NewDatabase(clt Client) Database {
	service := Database{
		Client: clt,
	}

	return service
}

type DatabaseObject struct {
	Id        string `json:"$id"`
	Name      string `json:"name"`
	CreatedAt string `json:"$createdAt"`
	UpdatedAt string `json:"$updatedAt"`
}

type DatabaseList struct {
	Total     int64            `json:"total"`
	Databases []DatabaseObject `json:"databases"`
}

type AttributeOptions struct {
	Key      string `json:"key"`
	Type     string `json:"type"`
	Status   string `json:"status"`
	Required bool   `json:"required"`
	Array    bool   `json:"array"`
}

type Attribute struct {
	AttributeOptions
}

type Index struct {
	Key        string   `json:"key"`
	Type       string   `json:"type"`
	Status     string   `json:"status"`
	Attributes []string `json:"attributes"`
	Orders     []string `json:"orders"`
}

type Collection struct {
	Id               string      `json:"$id"`
	Name             string      `json:"name"`
	CreatedAt        string      `json:"$createdAt"`
	UpdatedAt        string      `json:"$updatedAt"`
	DatabaseId       string      `json:"databaseId"`
	Permissions      []string    `json:"permissions"`
	Enabled          bool        `json:"enabled"`
	DocumentSecurity bool        `json:"documentSecurity"`
	Attributes       []Attribute `json:"attributes"`
	Indexes          []Index     `json:"indexes"`
}

type CollectionList struct {
	Total       int          `json:"total"`
	Collections []Collection `json:"collections"`
}

func (srv *Database) ListDatabases(Search string, Queries []string) (*DatabaseList, error) {
	path := "/databases"
	params := map[string]interface{}{
		"search":  Search,
		"queries": Queries,
	}

	resp, err := srv.Client.CallAPI("GET", path, srv.Client.headers, params)
	if err != nil {
		return nil, err
	}
	var result DatabaseList
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (srv *Database) GetDatabase(databaseId string) (*DatabaseObject, error) {
	r := strings.NewReplacer("{databaseId}", databaseId)
	path := r.Replace("/databases/{databaseId}")
	params := map[string]interface{}{}

	resp, err := srv.Client.CallAPI("GET", path, srv.Client.headers, params)
	if err != nil {
		return nil, err
	}
	var result DatabaseObject
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListCollections get a list of all the user collections. You can use the
// query params to filter your results. On admin mode, this endpoint will
// return a list of all of the project collections. [Learn more about
// different API modes](/docs/admin).
func (srv *Database) ListCollections(databaseId, Search string, Queries []string) (*CollectionList, error) {
	r := strings.NewReplacer("{databaseId}", databaseId)
	path := r.Replace("/database/{dattableId}/collections")

	params := map[string]interface{}{
		"search":  Search,
		"queries": Queries,
	}

	resp, err := srv.Client.CallAPI("GET", path, srv.Client.headers, params)
	if err != nil {
		return nil, err
	}
	var result CollectionList
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil

}

// CreateCollection create a new Collection.
func (srv *Database) CreateCollection(Name string, Read []interface{}, Write []interface{}, Rules []interface{}) (map[string]interface{}, error) {
	path := "/database/collections"

	params := map[string]interface{}{
		"name":  Name,
		"read":  Read,
		"write": Write,
		"rules": Rules,
	}

	return srv.Client.Call("POST", path, nil, params)
}

// GetCollection get collection by its unique ID. This endpoint response
// returns a JSON object with the collection metadata.
func (srv *Database) GetCollection(databaseId, collectionId string) (*Collection, error) {
	r := strings.NewReplacer("{databaseId}", databaseId, "{collectionId}", collectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}")

	params := map[string]interface{}{}

	resp, err := srv.Client.CallAPI("GET", path, srv.Client.headers, params)
	if err != nil {
		return nil, err
	}
	var result Collection
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateCollection update collection by its unique ID.
func (srv *Database) UpdateCollection(CollectionId string, Name string, Read []interface{}, Write []interface{}, Rules []interface{}) (map[string]interface{}, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}")

	params := map[string]interface{}{
		"name":  Name,
		"read":  Read,
		"write": Write,
		"rules": Rules,
	}

	return srv.Client.Call("PUT", path, nil, params)
}

// DeleteCollection delete a collection by its unique ID. Only users with
// write permissions have access to delete this resource.
func (srv *Database) DeleteCollection(CollectionId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}")

	params := map[string]interface{}{}

	return srv.Client.Call("DELETE", path, nil, params)
}

// ListDocuments get a list of all the user documents. You can use the query
// params to filter your results. On admin mode, this endpoint will return a
// list of all of the project documents. [Learn more about different API
// modes](/docs/admin).
func (srv *Database) ListDocuments(CollectionId string, Filters []interface{}, Offset int, Limit int, OrderField string, OrderType string, OrderCast string, Search string, First int, Last int) (map[string]interface{}, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/documents")

	params := map[string]interface{}{
		"filters":     Filters,
		"offset":      Offset,
		"limit":       Limit,
		"order-field": OrderField,
		"order-type":  OrderType,
		"order-cast":  OrderCast,
		"search":      Search,
		"first":       First,
		"last":        Last,
	}

	return srv.Client.Call("GET", path, nil, params)
}

// CreateDocument create a new Document.
func (srv *Database) CreateDocument(CollectionId string, Data map[string]interface{}, Read []interface{}, Write []interface{}, ParentDocument string, ParentProperty string, ParentPropertyType string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/documents")

	params := map[string]interface{}{
		"data":               Data,
		"read":               Read,
		"write":              Write,
		"parentDocument":     ParentDocument,
		"parentProperty":     ParentProperty,
		"parentPropertyType": ParentPropertyType,
	}

	return srv.Client.Call("POST", path, nil, params)
}

// GetDocument get document by its unique ID. This endpoint response returns a
// JSON object with the document data.
func (srv *Database) GetDocument(CollectionId string, DocumentId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/database/collections/{collectionId}/documents/{documentId}")

	params := map[string]interface{}{}

	return srv.Client.Call("GET", path, nil, params)
}

// UpdateDocument
func (srv *Database) UpdateDocument(CollectionId string, DocumentId string, Data map[string]interface{}, Read []interface{}, Write []interface{}) (map[string]interface{}, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/database/collections/{collectionId}/documents/{documentId}")

	params := map[string]interface{}{
		"data":  Data,
		"read":  Read,
		"write": Write,
	}

	return srv.Client.Call("PATCH", path, nil, params)
}

// DeleteDocument delete document by its unique ID. This endpoint deletes only
// the parent documents, his attributes and relations to other documents.
// Child documents **will not** be deleted.
func (srv *Database) DeleteDocument(CollectionId string, DocumentId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/database/collections/{collectionId}/documents/{documentId}")

	params := map[string]interface{}{}

	return srv.Client.Call("DELETE", path, nil, params)
}
