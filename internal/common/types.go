package common

import (
	"fmt"
)

type ImageLink struct {
	link  string
	valid bool
}

func NewImageLink(link string) (ImageLink, error) {

	if IsImage(link) {
		return ImageLink{link, true}, nil
	}
	return ImageLink{link, false}, fmt.Errorf("not valid")
}


type CRUD interface {
    set()
}
