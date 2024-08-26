package collector

import (
	"github.com/stmcginnis/gofish"
	redfish "github.com/stmcginnis/gofish/redfish"
	"github.com/prometheus/client_golang/prometheus"
)

var _ prometheus.Collector = &chassisCollector{}

type chassisCollector struct {
	BladeStatus   *prometheus.Desc
	BladeHealth   *prometheus.Desc
	chassisList   []interface{}
	username      string
	password      string
}

func NewChassisCollector(chassisList []interface{}, username string, password string) *chassisCollector {
	return &chassisCollector {
		BladeStatus: prometheus.NewDesc(
			"chassis_blade_status",
			"The blades health status in the chassis",
			[]string{
				"status", 
				"id",
				"chassis",
			}, 
			nil,
		),
		chassisList: chassisList,
		username: username,
		password: password,
	}
}

func (c *chassisCollector) Describe(ch chan<- *prometheus.Desc) {
	ds := []*prometheus.Desc{
		c.BladeStatus,
	}

	for _, d := range ds{
		ch <- d 
	}
}

func (c *chassisCollector) Collect(ch chan<- prometheus.Metric) {
	getBladeStatus(c, ch)
}

func newChassisClient(chassisIP, username, password string) (*gofish.APIClient, error) {
	config := gofish.ClientConfig{
		Endpoint: "https://" + username + ":" + password + "@" + chassisIP + ":5000",
		Username: username,
		Password: password,
		Insecure: true,
	}

	redfishClient, err := gofish.Connect(config)
	if err != nil {
		return nil, err
	}

	return redfishClient, nil
}

func getBladeStatus(c *chassisCollector, ch chan<- prometheus.Metric) {
	for _, chassis := range c.chassisList {
		redfishClient, err := newChassisClient(chassis.(string), c.username, c.password)
		if err != nil {
			return
		}

		defer redfishClient.Logout()
		service := redfishClient.Service

		metrics, err := QueryChassis(service)
		if err != nil {
			return
		}
		for _, metric := range metrics {
			ch <- prometheus.MustNewConstMetric(
				c.BladeStatus, 
				prometheus.GaugeValue, 
				float64(1), 
				string(metric.Status.State), 
				string(metric.ID),
				string(metric.SerialNumber),
				chassis.(string),
			)
		}
	}
}

func QueryChassis(service *gofish.Service)  ([]*redfish.Chassis, error) {
	chassis, err := service.Chassis()
	if err != nil {
		panic(err)
	}

	return chassis, nil
}
