package integration

import (
	"testing"
	"os"
	"github.com/bestreyer/carfinder/pkg/di"
)

func TestMain(m *testing.M) {
	setupDatabase()
	result := m.Run()
	os.Exit(result)
}

func setupDatabase() {
	d := di.New()

	d.GetDbConn().QueryRow("DELETE FROM driver_location")
	d.GetDbConn().QueryRow("ALTER SEQUENCE driver_location_driver_id_seq RESTART WITH 1")

}
