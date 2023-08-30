package model

type Si struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type Gu struct {
	Code     string `json:"code"`
	SiCode   string `json:"siCode"`
	SiName   string `json:"siName"`
	FullName string `json:"fullName"`
	Name     string `json:"name"`
	Active   bool   `json:"active"`
}

type Dong struct {
	Code     string `json:"code"`
	SiCode   string `json:"siCode"`
	SiName   string `json:"siName"`
	GuCode   string `json:"guCode"`
	GuName   string `json:"guName"`
	FullName string `json:"fullName"`
	Name     string `json:"name"`
	Active   bool   `json:"active"`
}

type Detail struct {
	Code     string `json:"code"`
	SiCode   string `json:"siCode"`
	SiName   string `json:"siName"`
	GuCode   string `json:"guCode"`
	GuName   string `json:"guName"`
	DongCode string `json:"dongCode"`
	DongName string `json:"dongName"`
	FullName string `json:"fullName"`
	Name     string `json:"name"`
	Active   bool   `json:"active"`
}
