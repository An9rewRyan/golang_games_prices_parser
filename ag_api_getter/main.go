package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"prices_parser/config"
	"prices_parser/root_structs"
	"prices_parser/utils"
	"strconv"

	"github.com/tidwall/gjson"
)

func main() {
	// utils.Create_article_table()
	for i := 1; i < config.Amount_of_pages+1; i++ {
		request := fmt.Sprintf("https://api.rawg.io/api/games?key=%s&page=%d&page_size=%d",
			config.Api_key, i, config.Page_size)
		resp, err := http.Get(request)
		if err != nil {
			fmt.Println(err)
		}
		resp_body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		resp_json_string := string(resp_body)
		var resp_struct root_structs.Ag_json
		json.Unmarshal([]byte(resp_json_string), &resp_struct)
		for i, result := range resp_struct.Results {
			fmt.Println(resp_json_string)
			utils.Write_game_to_db(
				result.Name,
				result.Slug,
				result.Released,
				result.BackgroundImage,
				float32(result.Rating),
				int16(result.Playtime),
				int8(result.Metacritic),
				gjson.Get(resp_json_string, "results."+strconv.Itoa(i)+".platforms.#(platform.id==4).requirements_en.minimum").String(),
				gjson.Get(resp_json_string, "results."+strconv.Itoa(i)+".platforms.#(platform.id==4).requirements_en.recommended").String(),
				gjson.Get(resp_json_string, "results."+strconv.Itoa(i)+".parent_platforms.#.platform.slug").String(),
				gjson.Get(resp_json_string, "results."+strconv.Itoa(i)+".genres.#.slug").String(),
				gjson.Get(resp_json_string, "results."+strconv.Itoa(i)+".tags.#.slug").String(),
				gjson.Get(resp_json_string, "results."+strconv.Itoa(i)+".short_screenshots.#.image").String(),
			)
		}
	}
}
