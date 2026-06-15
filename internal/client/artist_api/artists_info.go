package artistapi

func New() *ArtistInfo {
	return &ArtistInfo{}
}

func (a *ArtistInfo) Init() {
	a.mapArtistsInfo()
}

func (a *ArtistInfo) GetByIdKey() map[int]ArtistInfo {
	return byId
}