package statistics

import (
	"time"
	"testing"
	influx_client "github.com/influxdata/influxdb/client/v2"
)

// Default configuration for tests
var influxHttpConfig = influx_client.HTTPConfig {
	Addr: "http://localhost:8086",
	Username: "rpo",
	Password: "rpo",
}

var influxBatchPointsConfig = influx_client.BatchPointsConfig {
	Precision: "s",
	Database: "cloudtrust_grafana_test",
	RetentionPolicy: "",
	WriteConsistency: "",
}

func TestStatisticsModule(t *testing.T) {

	// Make client
	var influxClient influx_client.Client
	{
		var err error
		influxClient, err = influx_client.NewHTTPClient(influxHttpConfig)
		if err != nil {
			t.Errorf("%s: Cannot create Influx client.\nError: %s", t.Name(), err)
		}
		defer influxClient.Close()
	}

	// Create a new point batch
	var batchPoints influx_client.BatchPoints
	{
		var err error
		batchPoints, err = influx_client.NewBatchPoints(influxBatchPointsConfig)
		if err != nil {
			t.Errorf("%s: Cannot create BatchPoints.\nError: %s", t.Name(), err)
		}
	}

	// Create a point and add to batch
	var tags = map[string]string{"cpu": "cpu-total"}
	var fields = map[string]interface{} {
		"idle":   100.1,
		"system": 3.3,
		"user":   464.6,
	}

	var point *influx_client.Point
	{
		var err error
		point, err = influx_client.NewPoint("cpu_usage", tags, fields, time.Now())
		if err != nil {
			t.Errorf("%s: Cannot create Point.\nError: %s", t.Name(), err)
		}
		batchPoints.AddPoint(point)
	}

	// Write the batch
	var err = influxClient.Write(batchPoints)
	if err != nil {
		t.Errorf("%s: Cannot write batch.\nError: %s", t.Name(), err)
	}
}


func TestStatisticsModule2(t *testing.T) {

	// Make client
	var influxClient influx_client.Client
	{
		var err error
		influxClient, err = influx_client.NewHTTPClient(influxHttpConfig)
		if err != nil {
			t.Errorf("%s: Cannot create Influx client.\nError: %s", t.Name(), err)

		}
		defer influxClient.Close()
	}

	// Create a new point batch
	var batchPoints influx_client.BatchPoints
	{
		var err error
		batchPoints, err = influx_client.NewBatchPoints(influxBatchPointsConfig)
		if err != nil {
			t.Errorf("%s: Cannot create BatchPoints.\nError: %s", t.Name(), err)
		}
	}


	// Create a point and add to batch
	var tags = map[string]string{"event": "login", "realm": "realmId",}
	var fields = map[string]interface{} {
		"connexion":   1,
	}

	var point *influx_client.Point
	{
		var err error
		point, err = influx_client.NewPoint("event_stat", tags, fields, time.Now())
		if err != nil {
			t.Errorf("%s: Cannot create Point.\nError: %s", t.Name(), err)
		}
		batchPoints.AddPoint(point)
	}

	// Write the batch
	var err = influxClient.Write(batchPoints)
	if err != nil {
		t.Errorf("%s: Cannot write batch.\nError: %s", t.Name(), err)
	}
}

/*
Mock for influx
 */
type mockInflux struct {}

func (client *mockInflux) Ping(timeout time.Duration) (time.Duration, string, error) {return 1 * time.Second, "" , nil}
func (client *mockInflux) Write(bp influx_client.BatchPoints) error {return nil}
func (client *mockInflux) Query(q influx_client.Query) (*influx_client.Response, error) {return nil, nil}
func (client *mockInflux) Close() error {return nil}