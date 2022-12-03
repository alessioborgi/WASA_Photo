package database

func (db *appdbimpl) CreateFountain(u User) (User, error) {
	res, err := db.c.Exec(`INSERT INTO fountains (id, latitude, longitude, status) VALUES (?, ?, ?, ?)`,
		u.ID, u.Latitude, u.Longitude, u.Status)
	if err != nil {
		return u, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return u, err
	}

	u.ID = uint64(lastInsertID)
	return u, nil
}
