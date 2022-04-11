package messageBox

import (
	"errors"
	"strings"
)

func NewWarningBox(message string) string {
	return `
		<div class="alert alert-warning alert-dismissible fade show" role="alert">
			` + message + `
			<button type="button" class="close" data-dismiss="alert" aria-label="Close">
	  			<span aria-hidden="true">&times;</span>
			</button>
		</div>`
}

func NewDangerBox(message string) string {
	return `
		<div class="alert alert-danger alert-dismissible fade show" role="alert">
			` + message + `
			<button type="button" class="close" data-dismiss="alert" aria-label="Close">
			<span aria-hidden="true">&times;</span>
			</button>
		</div>
	`
}

func NewSuccessBox(message string) string {
	return `
		<div class="alert alert-success alert-dismissible fade show" role="alert">
			` + message + `
			<button type="button" class="close" data-dismiss="alert" aria-label="Close">
				<span aria-hidden="true">&times;</span>
			</button>
		</div>
	`
}

func EmptyMessageBox() string {
	return "<div></div>"
}

func NewMessageBox(message string, boxType ...string) (string, error) {
	if len(boxType) == 0 || boxType[0] == "" { // If no value is provided, default to success
		return `
		<div class="alert alert-success alert-dismissible fade show" role="alert">
			` + message + `
			<button type="button" class="close" data-dismiss="alert" aria-label="Close">
				<span aria-hidden="true">&times;</span>
			</button>
		</div>
		`, nil
	}
	messageType := strings.ToLower(boxType[0])
	if messageType != "success" && messageType != "danger" && messageType != "warning" {
		return "", errors.New("Invalid Message Box type")
	}

	return `
	<div class="alert alert-` + messageType + ` alert-dismissible fade show" role="alert">
		` + message + `
		<button type="button" class="close" data-dismiss="alert" aria-label="Close">
			<span aria-hidden="true">&times;</span>
		</button>
	</div>
	`, nil
}
