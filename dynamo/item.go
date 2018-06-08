package dynamo

import (
	"github.com/rs/xid"
)

// ID type
type ID string

// XID type
type XID xid.ID

// Item type
type Item struct {
	ID      ID       `json:"id"`
	XID     XID      `json:"xid"`
	Title   string   `json:"title"`
	Format  string   `json:"format"`
	Content string   `json:"content"`
	URL     string   `json:"url"`
	Tags    []string `json:"tags"`
	Status  string   `json:"publishStatus"`
}

// Items type
type Items []Item

// Repository interface
type Repository interface {
	Get(id ID) (*Item, error)
	List() (*Items, error)
	Create(item *Item) error
	Update(id ID, item *Item) error
	Delete(id ID) error
	Load(items Items) error
}

// Service interface
type Service interface {
	Get(id ID) (*Item, error)
	List() (*Items, error)
	Create(item *Item) error
	Update(id ID, item *Item) error
	Delete(id ID) error
	Load(items Items) error
}
