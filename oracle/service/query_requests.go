package service

import "oracle/models/database"

func (d *Service) Requests(requestId, page, limit, status int, order string) ([]database.RandomnessRequest, int64, error) {
	return d.Store.Db.GetPaginatedRequests(page, limit, status, order)
}
