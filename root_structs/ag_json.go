//thanks to https://mholt.github.io/json-to-go/ for generating this terrifying struct
package root_structs

type Ag_json struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ID              int     `json:"id"`
		Slug            string  `json:"slug"`
		Name            string  `json:"name"`
		Released        string  `json:"released"`
		Tba             bool    `json:"tba"`
		BackgroundImage string  `json:"background_image"`
		Rating          float64 `json:"rating"`
		RatingTop       int     `json:"rating_top"`
		Ratings         []struct {
			ID      int     `json:"id"`
			Title   string  `json:"title"`
			Count   int     `json:"count"`
			Percent float64 `json:"percent"`
		} `json:"ratings"`
		RatingsCount     int `json:"ratings_count"`
		ReviewsTextCount int `json:"reviews_text_count"`
		Added            int `json:"added"`
		AddedByStatus    struct {
			Yet     int `json:"yet"`
			Owned   int `json:"owned"`
			Beaten  int `json:"beaten"`
			Toplay  int `json:"toplay"`
			Dropped int `json:"dropped"`
			Playing int `json:"playing"`
		} `json:"added_by_status"`
		Metacritic       int         `json:"metacritic"`
		Playtime         int         `json:"playtime"`
		SuggestionsCount int         `json:"suggestions_count"`
		Updated          string      `json:"updated"`
		UserGame         interface{} `json:"user_game"`
		ReviewsCount     int         `json:"reviews_count"`
		SaturatedColor   string      `json:"saturated_color"`
		DominantColor    string      `json:"dominant_color"`
		Platforms        []struct {
			Platform struct {
				ID              int         `json:"id"`
				Name            string      `json:"name"`
				Slug            string      `json:"slug"`
				Image           interface{} `json:"image"`
				YearEnd         interface{} `json:"year_end"`
				YearStart       interface{} `json:"year_start"`
				GamesCount      int         `json:"games_count"`
				ImageBackground string      `json:"image_background"`
			} `json:"platform"`
			ReleasedAt     string      `json:"released_at"`
			RequirementsEn interface{} `json:"requirements_en"`
			RequirementsRu interface{} `json:"requirements_ru"`
		} `json:"platforms"`
		ParentPlatforms []struct {
			Platform struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"platform"`
		} `json:"parent_platforms"`
		Genres []struct {
			ID              int    `json:"id"`
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			GamesCount      int    `json:"games_count"`
			ImageBackground string `json:"image_background"`
		} `json:"genres"`
		Stores []struct {
			ID    int `json:"id"`
			Store struct {
				ID              int    `json:"id"`
				Name            string `json:"name"`
				Slug            string `json:"slug"`
				Domain          string `json:"domain"`
				GamesCount      int    `json:"games_count"`
				ImageBackground string `json:"image_background"`
			} `json:"store"`
		} `json:"stores"`
		Clip interface{} `json:"clip"`
		Tags []struct {
			ID              int    `json:"id"`
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			Language        string `json:"language"`
			GamesCount      int    `json:"games_count"`
			ImageBackground string `json:"image_background"`
		} `json:"tags"`
		EsrbRating struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"esrb_rating"`
		ShortScreenshots []struct {
			ID    int    `json:"id"`
			Image string `json:"image"`
		} `json:"short_screenshots"`
	} `json:"results"`
	SeoTitle       string `json:"seo_title"`
	SeoDescription string `json:"seo_description"`
	SeoKeywords    string `json:"seo_keywords"`
	SeoH1          string `json:"seo_h1"`
	Noindex        bool   `json:"noindex"`
	Nofollow       bool   `json:"nofollow"`
	Description    string `json:"description"`
	Filters        struct {
		Years []struct {
			From   int    `json:"from"`
			To     int    `json:"to"`
			Filter string `json:"filter"`
			Decade int    `json:"decade"`
			Years  []struct {
				Year     int  `json:"year"`
				Count    int  `json:"count"`
				Nofollow bool `json:"nofollow"`
			} `json:"years"`
			Nofollow bool `json:"nofollow"`
			Count    int  `json:"count"`
		} `json:"years"`
	} `json:"filters"`
	NofollowCollections []string `json:"nofollow_collections"`
}
