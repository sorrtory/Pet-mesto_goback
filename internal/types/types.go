package mestoTypes

import (
	"fmt"
	"mesto-goback/internal/common"
)

type ImageLink struct {
	link  string
	valid bool
}

func NewImageLink(link string) (ImageLink, error) {

	if common.IsImage(link) {
		return ImageLink{link, true}, nil
	}
	return ImageLink{link, false}, fmt.Errorf("not valid")
}



