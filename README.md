### Installing
You should copy .env.dist to .env

### Important
https://github.com/bestreyer/carfinder is private repository. I don't forget about:

"Do not make either your solution or this problem statement publicly available by, for
example, using github or bitbucket or by posting this problem to a blog or forum"

### Tech stack

1. Golang
2. Postgre + Postgis
3. Vegeta (https://github.com/tsenart/vegeta) - Load testing
4. Realize (https://github.com/oxequa/realize) - Live reloading
5. Docker

### Development
Set up environment:
`make development_up`

Stop:
`make development_down`

### Start docker-compose with production build
Up:
`make production_up`

Down:
`make production_down`

### Load Testing

`make loadtest`

You can specify rate and duration:
`make -e LOADTEST_RATE=10 -e LOADTEST_DURATION=50s loadtest

Result of load testing for rate=20, duration=60s, 1000000 drivers in database (I know that application has drivers limit=50000,
but for this load I removed this limitation):

```
Requests      [total, rate]            1000, 20.02
Duration      [total, attack, wait]    49.953476071s, 49.950121199s, 3.354872ms
Latencies     [mean, 50, 95, 99, max]  3.743573ms, 3.284406ms, 5.568155ms, 7.567626ms, 41.267073ms
Bytes In      [total, mean]            2000, 2.00
Bytes Out     [total, mean]            83000, 83.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:1000
```

### Tests
I have written only 3 unit tests and a lot of integration tests (`/integration` folder).
My opinion, that unit tests are useful if you write a library or mathematical software or if you have enough time :).
REST API applications should be covered with integration tests firstly. ( It's only my opinion :)

Run tests:

*Please, start integration tests after server is been started, otherwise tests doesn't work, because dependencies are installing*
```
#!/usr/bin/env sh

echo "Installing dependencies"
dep ensure --vendor-only

exec realize start
```

Run tests
```
make development_tests
```

Run only unit tests
```
make development_unit_tests
```

Run only integration tests
```
make development_integration_tests
```


### Work with accuracy

I wanted to implement Kalman filter, but changed my mind.
Because Kalman filter should be implemented on a device, because the device have access to inertial sensors, not just gps data.
Thus, devices should send the already the optimal latitude, longitude.

### Important notes

1. If the driver has sent their coordinates over 10 minutes ago, this driver won't returned in the response even if the driver meet the conditions

2. If a accuracy less than 0.4% (approximately 90 meters error), the application don't update the driver location

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