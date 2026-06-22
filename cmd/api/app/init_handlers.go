package app

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) fileServer() {
	a.router.Handle("/src/output.css", http.FileServerFS(a.embedded.Get()))
}

func (a *App) initHandlers() {
	a.router.Use(middleware.CleanPath)
	a.router.Use(a.midleware.Recover)
	a.fileServer()

	// Get request routes
	a.router.Group(func(r chi.Router) {
		// Auth routes
		r.Get("/", a.handler.Get.Auth.LandingPageHandler)
		r.Get(utils.LOGIN.String(), a.handler.Get.Auth.LoginPageHandler)
		r.Get(utils.REGISTER.String(), a.handler.Get.Auth.SignupHandler)

		// Legal pages
		r.Get(utils.ABOUT.String(), a.handler.Get.LegalPage.AboutPageHandler)
		r.Get(utils.TERMS.String(), a.handler.Get.LegalPage.TermsPageHandler)
		r.Get(utils.PRIVACY.String(), a.handler.Get.LegalPage.PrivacyPageHandler)

		// App pages
		r.With(a.midleware.VerifyAccessToken).Get(utils.HOME.String(), a.handler.Get.HomePage.HomeHandler)
		r.With(a.midleware.VerifyAccessToken).Get(utils.ARTIST_SEARCH.String(), a.handler.Get.HomePage.SearchHandler)
		r.With(a.midleware.VerifyAccessToken).Get(utils.ARTIST_DETAILS.String()+"/{id}", a.handler.Get.DetailPage.DetailPageHandler)
		r.With(a.midleware.VerifyAccessToken).Get(utils.ALL_EVENTS_ROUTES.String()+"/{id}", a.handler.Get.DetailPage.AllEventsPageHandler)
		r.With(a.midleware.VerifyAccessToken).Get(utils.TICKET.String(), a.handler.Get.DetailPage.TicketPageHandler)
		r.With(a.midleware.VerifyAccessToken).Get(utils.Payment.String(), a.handler.Get.DetailPage.RenderPaymentPage)
		r.With(a.midleware.VerifyAccessToken).Get(utils.ALL_EVENTS_ROUTES.String(), a.handler.Get.HomePage.PaidEventsPageHandler)
		r.With(a.midleware.VerifyAccessToken).Get(utils.FAVORITES.String(), a.handler.Get.HomePage.FavoritePageHandler)
		r.With(a.midleware.VerifyAccessToken).Get(utils.RECIEPTS.String(), a.handler.Get.DetailPage.RecieptPageHandler)
		r.With(a.midleware.VerifyAccessToken).Get(utils.SETTINGS.String(), a.handler.Get.Auth.SettingsPageHandler)
	})

	// Post request routes
	a.router.Group(func(r chi.Router) {
		// Auth routes
		r.Post(utils.REGISTER.String(), a.handler.Post.Auth.RegisterHandler)
		r.Post(utils.LOGIN.String(), a.handler.Post.Auth.LoginHandler)
		r.With(a.midleware.VerifyAccessToken).Post(utils.FILTER_SORT_ROUTE.String(), a.handler.Post.HomePage.FilterSortHandler)
		r.With(a.midleware.VerifyAccessToken).Post(utils.UPLOAD_PROFILE.String(), a.handler.Post.Auth.UploadProfilePicture)

		// App post request
		r.With(a.midleware.VerifyAccessToken).Post(utils.FAVORITE.String(), a.handler.Post.HomePage.UpdateFavoriteHandler)
		r.With(a.midleware.VerifyAccessToken).Post(utils.TicketType.String(), a.handler.Post.TicketPage.TicketTypeHandler)
		r.With(a.midleware.VerifyAccessToken).Post(string(utils.TicketQuantity), a.handler.Post.TicketPage.QuantitySelectorHandler)
		r.With(a.midleware.VerifyAccessToken).Post(utils.Payment.String(), a.handler.Post.TicketPage.SoldTicketHandler)
	})
}
