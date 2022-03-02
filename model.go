package main

type Params struct {
	Keyword        string `json:"keyword"`
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
	PortalType     int    `json:"portal_type"`
	PortalEntrance int    `json:"portal_entrance"`
}

type DataTemp struct {
	Code int `json:"code"`
	Data struct {
		JobPostList []struct {
			Id          string `json:"id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Requirement string `json:"requirement"`
			JobCategory struct {
				Id       string `json:"id"`
				Name     string `json:"name"`
				EnName   string `json:"en_name"`
				I18NName string `json:"i18n_name"`
				Depth    int    `json:"depth"`
				Parent   *struct {
					Id       string `json:"id"`
					Name     string `json:"name"`
					EnName   string `json:"en_name"`
					I18NName string `json:"i18n_name"`
					Depth    int    `json:"depth"`
					Parent   *struct {
						Id       string      `json:"id"`
						Name     string      `json:"name"`
						EnName   string      `json:"en_name"`
						I18NName string      `json:"i18n_name"`
						Depth    int         `json:"depth"`
						Parent   interface{} `json:"parent"`
						Children interface{} `json:"children"`
					} `json:"parent"`
					Children interface{} `json:"children"`
				} `json:"parent"`
				Children interface{} `json:"children"`
			} `json:"job_category"`
			CityInfo struct {
				Code         string      `json:"code"`
				Name         string      `json:"name"`
				EnName       string      `json:"en_name"`
				LocationType interface{} `json:"location_type"`
				I18NName     string      `json:"i18n_name"`
				PyName       interface{} `json:"py_name"`
			} `json:"city_info"`
		} `json:"job_post_list"`
		Count int    `json:"count"`
		Extra string `json:"extra"`
	} `json:"data"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}
