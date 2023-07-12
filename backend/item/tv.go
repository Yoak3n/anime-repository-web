package item

import "strings"

type TVItem struct {
	FullPath string
	Name     string
	Path     string
}

func NewTVItem(fullPath string) *TVItem {
	full := strings.Split(fullPath, "/")
	l := len(full)
	tvItem := &TVItem{}
	tvItem.FullPath = fullPath
	tvItem.Name = full[l-1]
	//tvItem.Path = strings.Join(full[:l-1], "/")
	tvItem.Path = strings.Replace(fullPath, tvItem.Name, "", 1)
	return tvItem
}
