package location

import (
	"database/sql"
	"context"
)

type postgreRepo struct {

	db *sql.DB
}

func NewPostgreRepository(db *sql.DB) (Repository) {
	return &postgreRepo{db: db}
}

func (r *postgreRepo) Update(ctx context.Context, ul *UpdateLocation) (error) {
	r.db.QueryRowContext(ctx, `
		UPDATE driver_location 
		SET 
			latitude = $1, 
			longitude = $2, 
			location = ST_SetSRID(ST_MakePoint($2, $1),4326)::geography,
			updated_at = $3
	  	WHERE driver_id = $4
	`, ul.Latitude, ul.Longitude, ul.UpdatedAt, ul.DriverId)

	return nil
}

func (r *postgreRepo) Create(ctx context.Context, l *Location) (error) {
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO driver_location(latitude, longitude, l, updated_at) 
		VALUES ($1, $2, ST_SetSRID(ST_MakePoint($2, $1),4326)::geography, $3)
		RETURNING driver_id
	`, l.Latitude, l.Longitude, l.UpdatedAt).Scan(&l.DriverId)

	return err
}

func (r *postgreRepo) GetDrivers(ctx context.Context, dal *DriverAroundLocation) ([]LocationWithDistance, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT 
			driver_id as id, latitude, longitude, 
			ST_Distance(location, ST_SetSRID(ST_MakePoint($2, $1), 4326)::geography) as distance 
	  	FROM driver_location 
	  	WHERE ST_DWithin(location, ST_SetSRID(ST_MakePoint($2, $1), 4326)::geography, $3)
 		LIMIT $4;
	`, dal.Latitude, dal.Longitude, dal.Radius, dal.Limit)

	if nil != err {
		return nil, err
	}

	results := []LocationWithDistance{}

	for rows.Next() {
		var r LocationWithDistance

		if err = rows.Scan(&r.Id, &r.Latitude, &r.Longitude, &r.Distance); nil != err {
			return nil, err
		}

		results = append(results, r)
	}

	return results, nil
}