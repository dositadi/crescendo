package legal

import (
	"fmt"
	"html/template"

	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
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

	name := "about.html"
	switch page {
	case Terms:
		name = "terms.html"
	case Privacy:
		name = "privacy.html"
	case About:
		name = "about.html"
	}

	prevPage := l.request.FormValue(utils.PAGE_KEY)
	if prevPage == "" {
		prevPage = utils.LOGIN.String()
	}

	data := struct {
		PrevPage, TermsPageUrl, PrivacyPageUrl, RefundPolicyUrl, SupportUrl string
	}{
		PrevPage:       prevPage,
		TermsPageUrl:   fmt.Sprintf("%s?%s=%s", utils.TERMS.String(), utils.PAGE_KEY, l.request.URL.EscapedPath()),
		PrivacyPageUrl: fmt.Sprintf("%s?%s=%s", utils.PRIVACY.String(), utils.PAGE_KEY, l.request.URL.EscapedPath()),
	}

	temp, err := template.New(name).ParseFS(l.embedded.Get(), fs...)
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
