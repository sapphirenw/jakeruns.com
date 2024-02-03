package xtempl

import (
	"fmt"
	"html/template"
	"strings"
	"time"
)

func defaultFuncMap() *template.FuncMap {
	return &template.FuncMap{
		"dict":     dict,
		"truncstr": truncstr,
		"substr":   substr,
		"join":     join,
		"fmtDate":  fmtDate,
		"add":      add,
		"gt":       gt,
	}
}

// allows for arbitrary values to be passed to the template with custom names
func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

// truncates a string `str` to `num` characters
func truncstr(str string, num int) string {
	if len(str) > num {
		return str[:num] + "..."
	}
	return str
}

func substr(str string, num int) string {
	if len(str) > num {
		return str[:num]
	}
	return str
}

func join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}

func fmtDate(str string) string {
	parsedDate, _ := time.Parse("2006-01-02 15:04:05", str)
	return parsedDate.Format("January 2, 2006")
}

func add(a, b int) int {
	return a + b
}

func gt(a, b int) bool {
	return a > b
}
