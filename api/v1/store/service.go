package store

import (
	"github.com/labstack/echo/v4"
	. "key-value-store/api/responses"
	"net/http"
)

var m map[string]string

// Key-Value Set
// @Description add key-value pair to store
// @ID key-value-set
// @Accept json
// @Produce json
// @Param request body SetKeyValueRequest true "request body"
// @Success 200 {object} Done
// @Router /store [POST]
func set(c echo.Context) (error error) {
	req := new(SetKeyValueRequest)
	if error = c.Bind(&req); error != nil {
		return Err(c, "Request Model Error", http.StatusBadRequest)
	}

	if len(m) == 0 {
		m = make(map[string]string)
	}

	m[req.Key] = req.Value

	return Ok(c, "key-value set successfully")
}

// Key-Value List
// @Description list key-value pair to store
// @ID key-value-get
// @Accept json
// @Produce json
// @Param key query string true "find value by key"
// @Success 200 {object} map[string]string
// @Router /store [GET]
func get(c echo.Context) (error error) {
	key := c.QueryParam("key")

	if key != "" {
		if val, ok := m[key]; ok {
			pair := make(map[string]string)
			pair[key] = val

			return Ok(c, pair)
		}
	}

	return Ok(c, nil)
}

// Flush Key-Value List
// @Description flush key-value list on store
// @ID flush-key-value-list
// @Accept json
// @Produce json
// @Success 200 {object} Done
// @Router /store/flush [PUT]
func flush(c echo.Context) (error error) {
	m = make(map[string]string)

	return Ok(c, "key-value flush successfully")
}
