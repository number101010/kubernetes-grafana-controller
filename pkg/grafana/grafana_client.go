package grafana

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"kubernetes-grafana-controller/pkg/prometheus"

	"github.com/imroc/req"
	"k8s.io/apimachinery/pkg/util/runtime"
)

const NO_ID = ""

type Interface interface {
	PostDashboard(string, string) (string, error)
	DeleteDashboard(string) error
	GetAllDashboardIds() ([]string, error)

	PostAlertNotification(string, string) (string, error)
	DeleteAlertNotification(string) error
	GetAllAlertNotificationIds() ([]string, error)

	PostDataSource(string, string) (string, error)
	DeleteDataSource(string) error
	GetAllDataSourceIds() ([]string, error)
}

type Client struct {
	address string
}

func init() {
	// cost is required for prom metrics
	req.SetFlags(req.LstdFlags | req.Lcost)
}

func NewClient(address string) *Client {

	client := &Client{
		address: address,
	}

	return client
}

func (client *Client) PostDashboard(dashboardJSON string, uid string) (string, error) {
	dashboardJSON, err := sanitizeObject(dashboardJSON)

	if err != nil {
		return "", err
	}

	if uid != NO_ID {
		dashboardJSON, err = setId(dashboardJSON, "uid", uid)

		if err != nil {
			return "", err
		}
	}

	postJSON := fmt.Sprintf(`{
		"dashboard": %v,
		"folderId": 0,
		"overwrite": true
	}`, dashboardJSON)

	return client.postGrafanaObject(postJSON, "/api/dashboards/db", "uid")
}

func (client *Client) DeleteDashboard(id string) error {
	resp, err := req.Delete(client.address + "/api/dashboards/uid/" + id)
	prometheus.GrafanaDeleteLatencyMilliseconds.WithLabelValues(prometheus.TypeDashboard).Observe(float64(resp.Cost() / time.Millisecond))

	if err != nil {
		return err
	}

	if !responseIsSuccess(resp) {
		return errors.New(resp.Response().Status)
	}

	return nil
}

func (client *Client) GetAllDashboardIds() ([]string, error) {
	var resp *req.Resp
	var err error
	var dashboards []map[string]interface{}

	// Request existing notification channels
	if resp, err = req.Get(client.address + "/api/search"); err != nil {
		return nil, err
	}
	prometheus.GrafanaGetLatencyMilliseconds.WithLabelValues(prometheus.TypeDashboard).Observe(float64(resp.Cost() / time.Millisecond))

	if err = resp.ToJSON(&dashboards); err != nil {
		return nil, err
	}

	var ids []string

	for _, dashboard := range dashboards {
		ids = append(ids, dashboard["uid"].(string))
	}

	return ids, nil
}

func (client *Client) PostAlertNotification(alertNotificationJson string, id string) (string, error) {
	alertNotificationJson, err := sanitizeObject(alertNotificationJson)

	if err != nil {
		return "", err
	}

	if id != NO_ID {
		// alert notification requires the id in the object for unknown reasons
		alertNotificationJson, err = setId(alertNotificationJson, "id", id)

		if err != nil {
			return "", err
		}

		id, err := client.putGrafanaObject(alertNotificationJson, fmt.Sprintf("/api/alert-notifications/%v", id), "id")

		if err != nil {
			runtime.HandleError(err)
			prometheus.GrafanaWastedPutTotal.WithLabelValues(prometheus.TypeAlertNotification).Inc()

			return client.postGrafanaObject(alertNotificationJson, "/api/alert-notifications", "id")
		} else {
			return id, err
		}

	} else {
		return client.postGrafanaObject(alertNotificationJson, "/api/alert-notifications", "id")
	}
}

func (client *Client) DeleteAlertNotification(id string) error {
	resp, err := req.Delete(client.address + "/api/alert-notifications/" + id)
	prometheus.GrafanaDeleteLatencyMilliseconds.WithLabelValues(prometheus.TypeAlertNotification).Observe(float64(resp.Cost() / time.Millisecond))

	if err != nil {
		return err
	}

	if !responseIsSuccess(resp) {
		return errors.New(resp.Response().Status)
	}

	return nil
}

func (client *Client) GetAllAlertNotificationIds() ([]string, error) {
	var resp *req.Resp
	var err error
	var channels []map[string]interface{}

	// Request existing notification channels
	if resp, err = req.Get(client.address + "/api/alert-notifications"); err != nil {
		return nil, err
	}
	prometheus.GrafanaGetLatencyMilliseconds.WithLabelValues(prometheus.TypeAlertNotification).Observe(float64(resp.Cost() / time.Millisecond))

	if err = resp.ToJSON(&channels); err != nil {
		return nil, err
	}

	var ids []string

	for _, channel := range channels {
		ids = append(ids, fmt.Sprintf("%v", channel["id"]))
	}

	return ids, nil
}

func (client *Client) PostDataSource(dataSourceJson string, id string) (string, error) {

	dataSourceJson, err := sanitizeObject(dataSourceJson)

	if err != nil {
		return "", err
	}

	if id != NO_ID {
		id, err := client.putGrafanaObject(dataSourceJson, fmt.Sprintf("/api/datasources/%v", id), "id")

		if err != nil {
			runtime.HandleError(err)
			prometheus.GrafanaWastedPutTotal.WithLabelValues(prometheus.TypeDataSource).Inc()

			return client.postGrafanaObject(dataSourceJson, "/api/datasources", "id")
		} else {
			return id, err
		}

	} else {
		return client.postGrafanaObject(dataSourceJson, "/api/datasources", "id")
	}
}

func (client *Client) DeleteDataSource(id string) error {
	resp, err := req.Delete(client.address + "/api/datasources/" + id)
	prometheus.GrafanaDeleteLatencyMilliseconds.WithLabelValues(prometheus.TypeDataSource).Observe(float64(resp.Cost() / time.Millisecond))

	if err != nil {
		return err
	}

	if !responseIsSuccess(resp) {
		return errors.New(resp.Response().Status)
	}

	return nil
}

func (client *Client) GetAllDataSourceIds() ([]string, error) {
	var resp *req.Resp
	var err error
	var datasources []map[string]interface{}

	// Request existing notification channels
	if resp, err = req.Get(client.address + "/api/datasources"); err != nil {
		return nil, err
	}
	prometheus.GrafanaGetLatencyMilliseconds.WithLabelValues(prometheus.TypeDataSource).Observe(float64(resp.Cost() / time.Millisecond))

	if err = resp.ToJSON(&datasources); err != nil {
		return nil, err
	}

	var ids []string

	for _, datasource := range datasources {
		ids = append(ids, fmt.Sprintf("%v", datasource["id"]))
	}

	return ids, nil
}

//
// shared
//

func (client *Client) postGrafanaObject(postJSON string, path string, idField string) (string, error) {
	var responseBody map[string]interface{}

	header := req.Header{
		"Content-Type": "application/json",
	}

	resp, err := req.Post(client.address+path, header, postJSON)

	if err != nil {
		return "", err
	}

	if resp == nil {
		return "", errors.New("Error and response are nil")
	}

	if !responseIsSuccess(resp) {
		return "", errors.New(resp.Response().Status)
	}

	err = resp.ToJSON(&responseBody)

	if err != nil {
		return "", err
	}

	id, ok := responseBody[idField]

	if !ok {
		return "", fmt.Errorf("Response Body did not have field %s", idField)
	}

	// is there a better way to generically convert to string?
	idString := fmt.Sprintf("%v", id)

	return idString, nil
}

func (client *Client) putGrafanaObject(putJSON string, path string, idField string) (string, error) {
	var responseBody map[string]interface{}

	header := req.Header{
		"Content-Type": "application/json",
	}

	resp, err := req.Put(client.address+path, header, putJSON)

	if err != nil {
		return "", err
	}

	if resp == nil {
		return "", errors.New("Error and response are nil")
	}

	if !responseIsSuccess(resp) {
		body, _ := resp.ToString()
		return "", errors.New(resp.Response().Status + ": " + body)
	}

	err = resp.ToJSON(&responseBody)

	if err != nil {
		return "", err
	}

	id, ok := responseBody[idField]

	if !ok {
		return "", fmt.Errorf("Response Body did not have field %s", idField)
	}

	// is there a better way to generically convert to string?
	idString := fmt.Sprintf("%v", id)

	return idString, nil
}

func responseIsSuccess(resp *req.Resp) bool {
	return resp.Response().StatusCode < 300 && resp.Response().StatusCode >= 200
}

func sanitizeObject(obj string) (string, error) {
	var jsonObject map[string]interface{}

	err := json.Unmarshal([]byte(obj), &jsonObject)
	if err != nil {
		return "", err
	}

	delete(jsonObject, "id")
	delete(jsonObject, "version")

	sanitizedBytes, err := json.Marshal(jsonObject)
	if err != nil {
		return "", err
	}

	return string(sanitizedBytes), nil
}

func setId(obj string, idField string, idValue string) (string, error) {
	var jsonObject map[string]interface{}

	err := json.Unmarshal([]byte(obj), &jsonObject)
	if err != nil {
		return "", err
	}

	intValue, err := strconv.Atoi(idValue)
	if err == nil {
		jsonObject[idField] = intValue
	} else {
		jsonObject[idField] = idValue
	}

	bytes, err := json.Marshal(jsonObject)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
