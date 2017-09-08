package model

import (
	"time"
)

type Artile struct {
	name    string    // file name
	title   string    // title
	content string    // content
	tag     []string  // tags
	date    time.Time // date
	md5     string    // unique id (md5)
}

// load from content file
func (a *Artile) load(path string) {
	a.name = path


}

// write with template
