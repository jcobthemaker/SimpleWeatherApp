package urlBuilder

import (
	"net/url"
	"weatherAPI/config"
)


func BuildAPIUrl(params map[string]string) (string,error) {

	config.LoadConfig();
	params["key"] = config.AppConfig.API.Key

	apiBaseUrl := "http://api.weatherapi.com/v1/current.json"

	u,err := url.Parse(apiBaseUrl)

	if err != nil {
		return "", err
	}

	query := u.Query()

	for key, value := range params{
		query.Set(key, value)
	}

	u.RawQuery = query.Encode()

	return u.String(), nil
}