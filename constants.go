package neodb_go_client

const NeoDBUrl = "https://neodb.social"
const TestAccessToken = ""

type ShelfType string

const (
	ShelfTypeWishlist ShelfType = "wishlist"
	ShelfTypeProgress ShelfType = "progress"
	ShelfTypeComplete ShelfType = "complete"
)

type Category string

const (
	CategoryBook        Category = "book"
	CategoryMovie       Category = "movie"
	CategoryTV          Category = "tv"
	CategoryMusic       Category = "music"
	CategoryGame        Category = "game"
	CategoryPodcast     Category = "podcast"
	CategoryFanfic      Category = "fanfic"
	CategoryPerformance Category = "performance"
	CategoryExhibition  Category = "exhibition"
	CategoryCollection  Category = "collection"
)
