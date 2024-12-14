package shared

import "time"

func GetBRLocationDefault() (*time.Location, error) {
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return nil, err
	}

	return location, nil
}
