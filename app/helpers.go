package app

import (
	"errors"
	"net/url"
	"strconv"
)

func getPaginationParams(queryParams url.Values) (int, int, error) {
	perPageQuery := queryParams.Get("per_page")
	offsetQuery := queryParams.Get("offset")

	if perPageQuery == "" || offsetQuery == "" {
		return 0, 0, nil
	}

	perPage, err := strconv.Atoi(perPageQuery)
	if err != nil {
		return 0, 0, BadRequestError{err: errors.New("per_page is not a valid int")}
	}
	offset, err := strconv.Atoi(offsetQuery)
	if err != nil {
		return 0, 0, BadRequestError{err: errors.New("offset is not a valid int")}
	}

	if offset != 0 && perPage == 0 {
		return 0, 0, BadRequestError{err: errors.New("Can't have an offset without a per_page")}
	}

	return perPage, offset, nil
}
