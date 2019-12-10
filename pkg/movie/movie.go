package movie

// Session represents the providing session of the movie
type Session struct {
	Value string `json:"value"`
	Text  string `json:"text"`
}

// Date represents the providing date of the movie
type Date struct {
	Text      string     `json:"text"`
	TimeValue string     `json:"value"`
	Sessions  []*Session `json:"sessions,omitempty"`
}

// Movie represents the movie information
type Movie struct {
	Name  string  `json:"name"`
	ID    string  `json:"id"`
	Dates []*Date `json:"dates,omitempty"`
}
