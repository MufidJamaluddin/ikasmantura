package utils

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

// @author Mufid Jamaluddin
type GetParams struct {
	Ids    []uint `query:"id,omitempty" json:"-" xml:"-" form:"-"`
	Sort   string `query:"_sort,omitempty" json:"-" xml:"-" form:"-"`
	Order  string `query:"_order,omitempty" json:"-" xml:"-" form:"-"`
	Start  uint   `query:"_start,omitempty" json:"-" xml:"-" form:"-"`
	End    uint   `query:"_end,omitempty" json:"-" xml:"-" form:"-"`
	Search string `query:"q,omitempty" json:"-" xml:"-" form:"-"`
}

func (p *GetParams) Filter(db *gorm.DB, searchFields []string) {
	if p.Ids != nil {
		if len(p.Ids) > 0 {
			db.Where("id IN ?", p.Ids)
		}
	}
	if p.Sort != "" {
		if strings.EqualFold(p.Order, "asc") || strings.EqualFold(p.Order, "desc") {
			db.Order(fmt.Sprintf("%s %s", p.Sort, p.Order))
		}
	}
	if p.Start < p.End {
		db.Limit(int(p.End - p.Start)).Offset(int(p.Start))
	}
	if p.Search != "" && searchFields != nil {
		for _, field := range searchFields {
			db.Where(fmt.Sprintf("%s LIKE ?", field), fmt.Sprintf("%s%", p.Search))
		}
	}
}
