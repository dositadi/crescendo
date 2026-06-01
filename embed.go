package groupietracker

import (
	"embed"
)

//go:embed migrations src/output.css internal/web/static/auth/*
var embeddedFiles embed.FS

type Embedded struct {
	embedded embed.FS
}

func New() *Embedded {
	return &Embedded{
		embedded: embeddedFiles,
	}
}

func (m *Embedded) Get() embed.FS {
	return m.embedded
}

func (m *Embedded) GetPath(dir, fileName string) string {
	switch dir {
	case "auth":
		return "internal/web/static/auth/" + fileName
	}
	return ""
}
