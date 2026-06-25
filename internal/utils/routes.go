package utils

type Route string

func (r Route) String() string {
	return string(r)
}

const (
	LOGIN             Route = Route("/auth/session")
	REGISTER          Route = Route("/auth/registration")
	LOGOUT            Route = Route("/auth/logout")
	DELETE_ACCOUNT    Route = Route("/auth/delete-account")
	SETTINGS          Route = Route("/user/settings")
	UPLOAD_PROFILE    Route = Route("/user/upload-avatar")
	UPDATE_USER_INFO  Route = Route("/user/update-info")
	ABOUT             Route = Route("/crescendo/about")
	TERMS             Route = Route("/crescendo/terms")
	PRIVACY           Route = Route("/crescendo/privacy")
	HOME              Route = Route("/artists")
	FAVORITES         Route = Route("/artists/favorites")
	RECIEPTS          Route = Route("/artists/event/reciept")
	ARTIST_DETAILS    Route = Route("/artists/detail")
	ARTIST_SEARCH     Route = Route("/artists/search") //?query=
	EVENTS            Route = Route("/artists/concert")
	FAVORITE          Route = Route("/artists/favorite")
	FILTER_SORT_ROUTE Route = Route("/artists/filter-sort")
	ALL_EVENTS_ROUTES Route = Route("/artists/all-events")
	TICKET            Route = Route("/artists/events/ticket")
	TicketType        Route = Route("/artists/events/ticket/type")
	TicketQuantity    Route = Route("/artists/events/ticket/quantity")
	Payment           Route = Route("/artists/events/ticket/payment")
)
