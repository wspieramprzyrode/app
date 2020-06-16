package category

type CategoryID string
type Category struct {
	ID               CategoryID  `json:"id"`
	Name             string      `json:"name" binding:"required"`
	DocstoreRevision interface{} `json:"-"`
}
