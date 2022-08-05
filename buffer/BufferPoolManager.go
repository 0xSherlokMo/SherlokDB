package buffer

import (
	"github.com/0xSherlokMo/SherlokDB/storage"
)

type BufferPoolManager interface {
	GetPoolSize() int
	FetchPage(id storage.PageId) *storage.Page
	UnpinPage(id storage.PageId, isDirty bool) bool
	FlushPage(id storage.PageId) bool
	FlushAllPage()
	NewPage(id storage.PageId) *storage.Page
	DeletePage(id storage.PageId) bool
}
