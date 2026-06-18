package legal

import (
	"html/template"

	"github.com/dositadi/groupie-tracker/internal/helper"
)

type LegalPage string

func (l LegalPage) String() string {
	return l.String()
}

const (
	Privacy LegalPage = LegalPage("privacy")
	Terms   LegalPage = LegalPage("terms")
	About   LegalPage = LegalPage("about")

	sourceL = "legal.RenderLegalPages()"
)

func (l *LegalPages) RenderLegalPages(page LegalPage) error {
	fs := []string{
		"internal/web/static/pages/about.html",
		"internal/web/static/pages/privacy.html",
		"internal/web/static/pages/terms.html",
	}

	data := struct{}{}

	temp, err := template.New("about.html").ParseFS(l.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Template creation error", err)
		l.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceL,
		})
		return e
	}

	switch page {
	case Privacy:
		err = temp.Execute(l.responseWriter, data)
	case About:
		err = temp.Execute(l.responseWriter, data)
	case Terms:
		err = temp.Execute(l.responseWriter, data)
	}

	if err != nil {
		e := helper.WrapError("Template execution error", err)
		l.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceL,
		})
		return e
	}
	return nil
}
