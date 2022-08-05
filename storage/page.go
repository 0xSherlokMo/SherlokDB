package storage

import (
	"sync"

	"github.com/0xSherlokMo/SherlokDB/common"
)

type PageId = int32

type Page struct {
	data     [common.PAGE_SIZE]byte
	pageId   PageId
	pinCount int
	isDirty  bool
	latch    *sync.RWMutex
}
