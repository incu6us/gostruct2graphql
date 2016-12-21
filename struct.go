package main

type Test struct {
	Stringer string
}

type Repository []struct {
	CacheMaxSeconds int64 `json:"cacheMaxSeconds"`
	CurrentTime     int64 `json:"currentTime"`
	Doc             struct {
		TropData struct {
			Two016 []struct {
				Active   bool   `json:"active"`
				Category string `json:"category"`
				Status   string `json:"status"`
				TropID   string `json:"tropId"`
				TropName string `json:"tropName"`
			} `json:"Two016"`
		} `json:"TropData"`
		TropHdr struct {
			TNum int64 `json:"tNum"`
		} `json:"tropHdr"`
	} `json:"doc"`
	GeneratedTime int64  `json:"generatedTime"`
	ID            string `json:"id"`
	Status        int64  `json:"status"`
}
