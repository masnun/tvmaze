package tvmaze

type Result struct {
	Score float64
	Show  Show
}

type Show struct {
	ID       int64
	URL      string
	Name     string
	Type     string
	Language string
	Genres   []string
	Image    Image
	Summary  string
	Status   string
}

type Image struct {
	Original string
	Medium   string
}
