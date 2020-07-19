package scrape

import "github.com/gocolly/colly"

type Info struct {
	Name        string `json:"party`
	Party       string `json:"party"`
	YearsActive string `json:"years_active"`
	Bills       string `json:"bills"`
}

func getCandidateInfo(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()

	c.OnHTML("body > div.search-column-main.basic-search-results.nav-on > div:nth-child(1) > div.result-heading > b", func(e *colly.HTMLElement) {
		politicianInfo := new(Info)
		politicianInfo.Name = ""
		politicianInfo.Party = ""
		politicianInfo.YearsActive = ""
		politicianInfo.Bills = ""

		bf := bytes.NewBuffer([]byte{})
		jsonEncoder := json.NewEncoder(bf)
		jsonEncoder.SetEscapeHTML(false)
		jsonEncoder.Encode(politicianInfo)

		w.Header().Set("Content-Type", "application/json")
		w.Write(bf.Bytes())
	}
	c.visit("URL")
}
