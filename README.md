### Load Testing

`make loadtest`

### Development

Set up environment:
`make development_up`

Stop:
`make development_down`

### How to store latitude and longitude
`driver_location` table has three columns, which to store location: `latitude`, `longitude`, `location`

`location` column is geography type column. `latitude` and `longitude` columns are double precision type column
I store both location and pair (latitude, longitude), because ST_AsText function, which converts from geometry to latitude, latitude is slow function

```
CREATE TABLE driver_location
(
  driver_id serial PRIMARY KEY,
  location geography(POINT),
  latitude double precision,
  longitude double precision,
  updated_at timestamp without time zone
);

CREATE INDEX driver_location_position_gix ON public.driver_location USING GIST (location);
```