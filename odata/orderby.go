package odata

import (
	"fmt"
	"strings"
)

func NewOrderBy(allowed ...string) *OrderBy {
	return &OrderBy{
		Values: map[string]string{},
	}
}

type OrderBy struct {
	Values  map[string]string
	allowed []string
}

func (o *OrderBy) Add(field string, order string) bool {
	if !o.IsAllowed(field) {
		return false
	}

	o.Values[field] = order
	return true
}

func (o *OrderBy) IsAllowed(key string) bool {
	// no documentation on what fields are allowed, so we assume all fields are allowed
	return true

	ok := false
	for _, a := range o.allowed {
		if a == key {
			ok = true
			continue
		}
	}
	return ok
}

func (o *OrderBy) MarshalSchema() string {
	if o == nil {
		return ""
	}
	pairs := []string{}
	for k, v := range o.Values {
		v = strings.ToLower(v)
		pair := fmt.Sprintf("%s %s", k, v)
		pairs = append(pairs, pair)
	}
	return strings.Join(pairs, ",")
}
