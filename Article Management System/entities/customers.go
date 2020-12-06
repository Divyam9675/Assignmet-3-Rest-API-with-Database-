package entities

import "fmt"

type Customer struct {
	ID          int64  `json:"Id"`
	TITLE       string `json:"title"`
	DESCRIPTION string `json:"description"`
	CONTENT     string `json:"content"`
}

func (custom Customer) ToString() string {

	return fmt.Sprintf("ID: %d\n TITLE: %s\n DESCRIPTION: %s\n CONTENT: %s\n", custom.ID, custom.TITLE, custom.DESCRIPTION, custom.CONTENT)

}
