package beater

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/sukanthm/site24x7beat/config"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func getTime(str string) time.Time {
	var replacer = strings.NewReplacer("-", ",", "T", ",", ":", ",")
	str = replacer.Replace(str)
	s := strings.Split(str[:len(str)-5], ",")
	a, _ := strconv.Atoi(s[0])
	b, _ := strconv.Atoi(s[1])
	c, _ := strconv.Atoi(s[2])
	d, _ := strconv.Atoi(s[3])
	e, _ := strconv.Atoi(s[4])
	f, _ := strconv.Atoi(s[5])

	return time.Date(a, time.Month(b), c, d, e, f, 000000000, time.UTC)
}

type RUM_default struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		AjaxPageViews string `json:"ajax_page_views"`
		PageViews     string `json:"page_views"`
		AppInfo       struct {
			ApdexThreshold  string   `json:"apdex_threshold"`
			ApplicationName string   `json:"application_name"`
			BeaconType      string   `json:"beacon_type"`
			Locations       string   `json:"locations"`
			State           string   `json:"state"`
			Type            string   `json:"type"`
			ApplicationKey  string   `json:"application_key"`
			ApplicationID   string   `json:"application_id"`
			LicenseKey      string   `json:"license_key"`
			Tags            []string `json:"tags"`
		} `json:"app_info"`
		PageViewsLimit string `json:"page_views_limit"`
		AppData        struct {
			FrustratedUsers        string  `json:"frustrated_users"`
			DocumentRenderingTime  float64 `json:"document_rendering_time"`
			PageRenderingTime      float64 `json:"page_rendering_time"`
			TotalCount             float64 `json:"total_count"`
			NetworkTime            float64 `json:"network_time"`
			DocumentProcessingTime float64 `json:"document_processing_time"`
			ErrorCount             string  `json:"error_count"`
			ConnectionTime         float64 `json:"connection_time"`
			DNSTime                float64 `json:"dns_time"`
			MaxRt                  float64 `json:"max_rt"`
			DocumentDownloadTime   float64 `json:"document_download_time"`
			ToleratedUsers         string  `json:"tolerated_users"`
			TotalResponseTime      float64 `json:"total_response_time"`
			SatisfiedUsers         string  `json:"satisfied_users"`
			MinRt                  float64 `json:"min_rt"`
			RedirectionTime        float64 `json:"redirection_time"`
			BackendTime            float64 `json:"backend_time"`
			FrontendTime           float64 `json:"frontend_time"`
			FirstByteTime          float64 `json:"first_byte_time"`
		} `json:"app_data"`
		LastArchievedTime string `json:"last_archieved_time"`
		UniqUsers         string `json:"uniq_users"`
	} `json:"data"`
}

type RUM_respTime struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Num0 struct {
			Responsetimegraph struct {
				Average      []float64       `json:"average"`
				ChartData    [][]interface{} `json:"chart_data"`
				Min          []float64       `json:"min"`
				Max          []float64       `json:"max"`
				Percentile95 []float64       `json:"percentile_95"`
			} `json:"responsetimegraph"`
		} `json:"0"`
		Legends []string `json:"legends"`
		Units   []string `json:"units"`
	} `json:"data"`
}

type RUM_throughput struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Num0 struct {
			Throughputgraph struct {
				Average      []float64       `json:"average"`
				ChartData    [][]interface{} `json:"chart_data"`
				Min          []float64       `json:"min"`
				Max          []float64       `json:"max"`
				Percentile95 []float64       `json:"percentile_95"`
			} `json:"throughputgraph"`
		} `json:"0"`
		Legends []string `json:"legends"`
		Units   []string `json:"units"`
	} `json:"data"`
}

type RUM_webpageDetails struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		WtList []struct {
			AverageResponseTime float64 `json:"average_response_time"`
			TotalResponseTime   float64 `json:"total_response_time"`
			TotalCount          float64 `json:"total_count"`
			Name                string  `json:"name"`
			Throughput          float64 `json:"throughput"`
		} `json:"wt_list"`
	} `json:"data"`
}

type website struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		TableData struct {
			Num0 struct {
				RESPONSETIME struct {
					Average                    float64         `json:"average"`
					Nine5Percentile            float64         `json:"95_percentile"`
					Min                        float64         `json:"min"`
					AverageDescription         string          `json:"average_description"`
					Max                        float64         `json:"max"`
					ChartData                  [][]interface{} `json:"chart_data"`
					Title                      string          `json:"title"`
					Unit                       string          `json:"unit"`
					Nine5PercentileDescription string          `json:"95_percentile_description"`
				} `json:"RESPONSETIME"`
			} `json:"0"`
		} `json:"table_data"`
		Info struct {
			FormattedEndTime       string `json:"formatted_end_time"`
			MonitorType            string `json:"monitor_type"`
			ResourceID             string `json:"resource_id"`
			ResourceTypeName       string `json:"resource_type_name"`
			PeriodName             string `json:"period_name"`
			GeneratedTime          string `json:"generated_time"`
			MetricAggregationName  string `json:"metric_aggregation_name"`
			ReportName             string `json:"report_name"`
			EndTime                string `json:"end_time"`
			MetricAggregation      int    `json:"metric_aggregation"`
			StartTime              string `json:"start_time"`
			SegmentType            int    `json:"segment_type"`
			ReportType             int    `json:"report_type"`
			Period                 int    `json:"period"`
			ResourceName           string `json:"resource_name"`
			SegmentTypeName        string `json:"segment_type_name"`
			FormattedStartTime     string `json:"formatted_start_time"`
			FormattedGeneratedTime string `json:"formatted_generated_time"`
			ResourceType           int    `json:"resource_type"`
		} `json:"info"`
		ChartData []struct {
			Num0 struct {
				ResponseTimeReportChart struct {
					Max             []float64       `json:"max"`
					Min             []float64       `json:"min"`
					Nine5Percentile []float64       `json:"95_percentile"`
					Average         []float64       `json:"average"`
					ChartData       [][]interface{} `json:"chart_data"`
				} `json:"ResponseTimeReportChart"`
			} `json:"0,omitempty"`
			Num1 struct {
				ThroughputChart struct {
					Max             []float64       `json:"max"`
					Min             []float64       `json:"min"`
					Nine5Percentile []float64       `json:"95_percentile"`
					Average         []float64       `json:"average"`
					ChartData       [][]interface{} `json:"chart_data"`
				} `json:"ThroughputChart"`
			} `json:"1,omitempty"`
			LocationResponseTimeChart []struct {
				Num6 struct {
					Max             []float64       `json:"max"`
					Label           string          `json:"label"`
					Min             []float64       `json:"min"`
					Nine5Percentile []float64       `json:"95_percentile"`
					Average         []float64       `json:"average"`
					ChartData       [][]interface{} `json:"chart_data"`
				} `json:"6,omitempty"`
				Num15 struct {
					Max             []float64       `json:"max"`
					Label           string          `json:"label"`
					Min             []float64       `json:"min"`
					Nine5Percentile []float64       `json:"95_percentile"`
					Average         []float64       `json:"average"`
					ChartData       [][]interface{} `json:"chart_data"`
				} `json:"15,omitempty"`
				Num1 struct {
					Max             []float64       `json:"max"`
					Label           string          `json:"label"`
					Min             []float64       `json:"min"`
					Nine5Percentile []float64       `json:"95_percentile"`
					Average         []float64       `json:"average"`
					ChartData       [][]interface{} `json:"chart_data"`
				} `json:"1,omitempty"`
			} `json:"LocationResponseTimeChart,omitempty"`
		} `json:"chart_data"`
	} `json:"data"`
}

type Site24x7beat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Site24x7beat{
		done:   make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Site24x7beat) Run(b *beat.Beat) error {

	logp.Info("Site24x7beat is running! Hit CTRL-C to stop it.")
	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	counter := 0

	Period_yml, _ := strconv.Atoi(strings.Replace(bt.config.Period.String(), "s", "", -1))
	var Period float64 = 3600 / float64(Period_yml)
	var Period_counter float64 = 1

	total := make([][]interface{}, 0)
	total_time := make([][][]time.Time, 0)

	for _, item := range bt.config.Input {

		enabled := item["enabled"]

		if strings.ToLower(enabled.(string)) == "true" {

			type1 := item["type"]

			if strings.ToLower(type1.(string)) == "website" {
				zohoToken, ok3 := item["zohoToken"]
				id, ok4 := item["id"]
				granularity, ok5 := item["granularity"]
				period, ok6 := item["period"]
				respTimeData, ok7 := item["respTimeData"]
				throughputData, ok8 := item["throughputData"]
				locationData, ok9 := item["locationData"]
				if ok3 && ok4 && ok5 && ok6 && ok7 && ok8 && ok9 {

					total = append(total, []interface{}{zohoToken, type1, id, granularity, period, respTimeData, throughputData, locationData})
					total_time = append(total_time, [][]time.Time{})
					for i := 0; i < reflect.ValueOf(id).Len(); i++ {
						total_time[len(total_time)-1] = append(total_time[len(total_time)-1], []time.Time{time.Time{}, time.Time{}, time.Time{}, time.Time{}, time.Time{}, time.Time{}, time.Time{}, time.Time{}, time.Time{}, time.Time{}})
					}

				} else {
					logp.Info("Please ensure every input has all the required config data (refer the yml file)")
					//return "Please ensure every input has all the required config data (refer the yml file)"
				}
			}
			if strings.ToLower(type1.(string)) == "rum" {
				zohoToken, ok3 := item["zohoToken"]
				id, ok4 := item["id"]
				timeWindow, ok5 := item["timeWindow"]
				respTimeData, ok6 := item["respTimeData"]
				throughputData, ok7 := item["throughputData"]
				webpageDetails, ok8 := item["webpageDetails"]

				if ok3 && ok4 && ok5 && ok6 && ok7 && ok8 {

					total = append(total, []interface{}{zohoToken, type1, id, timeWindow, respTimeData, throughputData, webpageDetails})
					total_time = append(total_time, [][]time.Time{})
					for i := 0; i < reflect.ValueOf(id).Len(); i++ {
						total_time[len(total_time)-1] = append(total_time[len(total_time)-1], []time.Time{time.Time{}, time.Time{}, time.Time{}, time.Time{}})
					}

				} else {
					logp.Info("Please ensure every input has all the required config data (refer the yml file)")
					//return "Please ensure every input has all the required config data (refer the yml file)"
				}
			}
		}
	}

	//fmt.Println(total)
	//fmt.Println(total_time)

	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		for total_counter, tot := range total {

			if strings.ToLower(tot[1].(string)) == "website" {
				switch reflect.TypeOf(tot[2]).Kind() {
				case reflect.Slice:
					s := reflect.ValueOf(tot[2])

					for i := 0; i < s.Len(); i++ {

						client := &http.Client{}
						req, err := http.NewRequest("GET", "https://www.site24x7.com/api/reports/performance/"+s.Index(i).Interface().(string)+"?granularity="+tot[3].(string)+"&period="+tot[4].(string), nil)
						req.Header.Add("Authorization", "Zoho-authtoken"+" "+tot[0].(string))

						resp, err := client.Do(req)
						if err != nil {
							return err
						}
						defer resp.Body.Close()
						var websitedata website
						if resp.StatusCode != http.StatusOK {
							bodyBytes, err2 := ioutil.ReadAll(resp.Body)
							if err2 == nil {
								logp.Info("HTTP ERROR " + string(resp.StatusCode) + " " + strings.Trim(string(bodyBytes), "\n"))
							} else {
								logp.Info(err2.Error())
								return err2
							}
							return nil
						}
						if resp.StatusCode == http.StatusOK {
							bodyBytes, err2 := ioutil.ReadAll(resp.Body)

							if err2 != nil {
								return err
							}

							x := string(bodyBytes)
							r, _ := regexp.Compile(`{\"0\":{\"ThroughputChart\"`)

							y := r.FindStringIndex(x)
							if len(y) != 0 {
								x = x[:y[0]+2] + `1` + x[y[0]+3:]
							}
							bodyBytes = []byte(x)

							json.Unmarshal(bodyBytes, &websitedata)

						}

						if strings.ToLower(tot[5].(string)) == "true" {
							for d := range websitedata.Data.ChartData[0].Num0.ResponseTimeReportChart.ChartData {
								event_time := beat.Event{
									Timestamp: time.Now(),
									Fields: common.MapStr{
										"mytimestamp":    websitedata.Data.ChartData[0].Num0.ResponseTimeReportChart.ChartData[d][0],
										"Type":           "Site24x7_WebsiteAggResponseTime",
										"MonitorID":      s.Index(i).Interface().(string),
										"WebsiteName":    websitedata.Data.Info.ResourceName,
										"Counter":        counter,
										"DNSTime":        websitedata.Data.ChartData[0].Num0.ResponseTimeReportChart.ChartData[d][1],
										"ConnectionTime": websitedata.Data.ChartData[0].Num0.ResponseTimeReportChart.ChartData[d][2],
										"DownloadTime":   websitedata.Data.ChartData[0].Num0.ResponseTimeReportChart.ChartData[d][5],
										"FirstByteTime":  websitedata.Data.ChartData[0].Num0.ResponseTimeReportChart.ChartData[d][4],
										"SSLTime":        websitedata.Data.ChartData[0].Num0.ResponseTimeReportChart.ChartData[d][3],
									},
								}

								if total_time[total_counter][i][0].Sub(getTime(websitedata.Data.ChartData[0].Num0.ResponseTimeReportChart.ChartData[d][0].(string))) < 0 {
									bt.client.Publish(event_time)
									logp.Info("Website AggResponseTime sent")
									if total_time[total_counter][i][1].Sub(getTime(websitedata.Data.ChartData[0].Num0.ResponseTimeReportChart.ChartData[d][0].(string))) < 0 {
										total_time[total_counter][i][1] = getTime(websitedata.Data.ChartData[0].Num0.ResponseTimeReportChart.ChartData[d][0].(string))
									}
								} else {
									logp.Info("Website AggResponseTime duplicate not sent ")
								}

							}
							total_time[total_counter][i][0] = total_time[total_counter][i][1]
						}

						if strings.ToLower(tot[6].(string)) == "true" {
							for f := range websitedata.Data.ChartData[1].Num1.ThroughputChart.ChartData {
								event_tp := beat.Event{
									Timestamp: time.Now(),
									Fields: common.MapStr{
										"mytimestamp": websitedata.Data.ChartData[1].Num1.ThroughputChart.ChartData[f][0],
										"Type":        "Site24x7_Website_Throughput",
										"MonitorID":   s.Index(i).Interface().(string),
										"WebsiteName": websitedata.Data.Info.ResourceName,
										"Counter":     counter,
										"Throughput":  websitedata.Data.ChartData[1].Num1.ThroughputChart.ChartData[f][1],
									},
								}

								if total_time[total_counter][i][2].Sub(getTime(websitedata.Data.ChartData[1].Num1.ThroughputChart.ChartData[f][0].(string))) < 0 {
									bt.client.Publish(event_tp)
									logp.Info("Website Throughput sent")
									if total_time[total_counter][i][3].Sub(getTime(websitedata.Data.ChartData[1].Num1.ThroughputChart.ChartData[f][0].(string))) < 0 {
										total_time[total_counter][i][3] = getTime(websitedata.Data.ChartData[1].Num1.ThroughputChart.ChartData[f][0].(string))
									}
								} else {
									logp.Info("Website Throughput duplicate not sent ")
								}

							}
							total_time[total_counter][i][2] = total_time[total_counter][i][3]
						}

						if strings.ToLower(tot[7].(string)) == "true" {
							for e := range websitedata.Data.ChartData[2].LocationResponseTimeChart[0].Num6.ChartData {
								event_loc1 := beat.Event{
									Timestamp: time.Now(),
									Fields: common.MapStr{
										"mytimestamp":  websitedata.Data.ChartData[2].LocationResponseTimeChart[0].Num6.ChartData[e][0],
										"Type":         "Site24x7_Website_LocResponseTime",
										"MonitorID":    s.Index(i).Interface().(string),
										"WebsiteName":  websitedata.Data.Info.ResourceName,
										"Counter":      counter,
										"Location":     websitedata.Data.ChartData[2].LocationResponseTimeChart[0].Num6.Label,
										"ResponseTime": websitedata.Data.ChartData[2].LocationResponseTimeChart[0].Num6.ChartData[e][1],
									},
								}

								if total_time[total_counter][i][4].Sub(getTime(websitedata.Data.ChartData[2].LocationResponseTimeChart[0].Num6.ChartData[e][0].(string))) < 0 {
									bt.client.Publish(event_loc1)
									logp.Info("Website Loc1 sent")
									if total_time[total_counter][i][5].Sub(getTime(websitedata.Data.ChartData[2].LocationResponseTimeChart[0].Num6.ChartData[e][0].(string))) < 0 {
										total_time[total_counter][i][5] = getTime(websitedata.Data.ChartData[2].LocationResponseTimeChart[0].Num6.ChartData[e][0].(string))
									}
								} else {
									logp.Info("Website Loc1 duplicate not sent ")
								}

							}
							total_time[total_counter][i][4] = total_time[total_counter][i][5]

							for g := range websitedata.Data.ChartData[2].LocationResponseTimeChart[1].Num15.ChartData {
								event_loc2 := beat.Event{
									Timestamp: time.Now(),
									Fields: common.MapStr{
										"mytimestamp":  websitedata.Data.ChartData[2].LocationResponseTimeChart[1].Num15.ChartData[g][0],
										"Type":         "Site24x7_Website_LocResponseTime",
										"MonitorID":    s.Index(i).Interface().(string),
										"WebsiteName":  websitedata.Data.Info.ResourceName,
										"Counter":      counter,
										"Location":     websitedata.Data.ChartData[2].LocationResponseTimeChart[1].Num15.Label,
										"ResponseTime": websitedata.Data.ChartData[2].LocationResponseTimeChart[1].Num15.ChartData[g][1],
									},
								}

								if total_time[total_counter][i][6].Sub(getTime(websitedata.Data.ChartData[2].LocationResponseTimeChart[1].Num15.ChartData[g][0].(string))) < 0 {
									bt.client.Publish(event_loc2)
									logp.Info("Website Loc2 sent")
									if total_time[total_counter][i][7].Sub(getTime(websitedata.Data.ChartData[2].LocationResponseTimeChart[1].Num15.ChartData[g][0].(string))) < 0 {
										total_time[total_counter][i][7] = getTime(websitedata.Data.ChartData[2].LocationResponseTimeChart[1].Num15.ChartData[g][0].(string))
									}
								} else {
									logp.Info("Website Loc2 duplicate not sent")
								}

							}
							total_time[total_counter][i][6] = total_time[total_counter][i][7]

							for h := range websitedata.Data.ChartData[2].LocationResponseTimeChart[2].Num1.ChartData {
								event_loc3 := beat.Event{
									Timestamp: time.Now(),
									Fields: common.MapStr{
										"mytimestamp":  websitedata.Data.ChartData[2].LocationResponseTimeChart[2].Num1.ChartData[h][0],
										"Type":         "Site24x7_Website_LocResponseTime",
										"MonitorID":    s.Index(i).Interface().(string),
										"WebsiteName":  websitedata.Data.Info.ResourceName,
										"Counter":      counter,
										"Location":     websitedata.Data.ChartData[2].LocationResponseTimeChart[2].Num1.Label,
										"ResponseTime": websitedata.Data.ChartData[2].LocationResponseTimeChart[2].Num1.ChartData[h][1],
									},
								}

								if total_time[total_counter][i][8].Sub(getTime(websitedata.Data.ChartData[2].LocationResponseTimeChart[2].Num1.ChartData[h][0].(string))) < 0 {
									bt.client.Publish(event_loc3)
									logp.Info("Website Loc3 sent")
									if total_time[total_counter][i][9].Sub(getTime(websitedata.Data.ChartData[2].LocationResponseTimeChart[2].Num1.ChartData[h][0].(string))) < 0 {
										total_time[total_counter][i][9] = getTime(websitedata.Data.ChartData[2].LocationResponseTimeChart[2].Num1.ChartData[h][0].(string))
									}
								} else {
									logp.Info("Website Loc3 duplicate not sent")
								}

							}
							total_time[total_counter][i][8] = total_time[total_counter][i][9]
						}
					}

				}
			}

			if strings.ToLower(tot[1].(string)) == "rum" {
				switch reflect.TypeOf(tot[2]).Kind() {
				case reflect.Slice:
					s := reflect.ValueOf(tot[2])

					for i := 0; i < s.Len(); i++ {

						//get rum name, pageViews, apdex
						var ApplicationName string
						client := &http.Client{}
						req, err := http.NewRequest("GET", "https://www.site24x7.com/api/rum/web/view/"+s.Index(i).Interface().(string)+"/"+tot[3].(string), nil)
						req.Header.Add("Authorization", "Zoho-authtoken"+" "+tot[0].(string))

						resp, err := client.Do(req)
						if err != nil {
							return err
						}
						defer resp.Body.Close()
						var RUM0 RUM_default
						if resp.StatusCode != http.StatusOK {
							bodyBytes, err2 := ioutil.ReadAll(resp.Body)
							if err2 == nil {
								logp.Info("HTTP ERROR " + string(resp.StatusCode) + " " + strings.Trim(string(bodyBytes), "\n"))
							} else {
								logp.Info(err2.Error())
								return err2
							}
							return nil
						}
						if resp.StatusCode == http.StatusOK {
							bodyBytes, err2 := ioutil.ReadAll(resp.Body)
							if err2 != nil {
								return err
							}

							json.Unmarshal(bodyBytes, &RUM0)
							ApplicationName = RUM0.Data.AppInfo.ApplicationName
							satisfied, _ := strconv.ParseFloat(RUM0.Data.AppData.SatisfiedUsers, 64)
							tolerated, _ := strconv.ParseFloat(RUM0.Data.AppData.ToleratedUsers, 64)

							event_time := beat.Event{
								Timestamp: time.Now(),
								Fields: common.MapStr{
									"mytimestamp":            time.Now(),
									"ApdexThreshold":         RUM0.Data.AppInfo.ApdexThreshold,
									"ApplicationName":        RUM0.Data.AppInfo.ApplicationName,
									"Type":                   "Site24x7_RUM_UserExperience",
									"PageViewsLimit":         RUM0.Data.PageViewsLimit,
									"DocumentRenderingTime":  RUM0.Data.AppData.DocumentRenderingTime,
									"PageRenderingTime":      RUM0.Data.AppData.PageRenderingTime,
									"NetworkTime":            RUM0.Data.AppData.NetworkTime,
									"DocumentProcessingTime": RUM0.Data.AppData.DocumentProcessingTime,
									"ConnectionTime":         RUM0.Data.AppData.ConnectionTime,
									"DNSTime":                RUM0.Data.AppData.DNSTime,
									"MaxResponseTime":        RUM0.Data.AppData.MaxRt,
									"DocumentDownloadTime":   RUM0.Data.AppData.DocumentDownloadTime,
									"TotalResponseTime":      RUM0.Data.AppData.TotalResponseTime,
									"MinResponseTime":        RUM0.Data.AppData.MinRt,
									"RedirectionTime":        RUM0.Data.AppData.RedirectionTime,
									"BackendTime":            RUM0.Data.AppData.BackendTime,
									"FirstByteTime":          RUM0.Data.AppData.FirstByteTime,
									"FrontendTime":           RUM0.Data.AppData.FrontendTime,
									"LastArchievedTime":      RUM0.Data.LastArchievedTime,
									"APDEX":                  (satisfied + tolerated*0.5) / RUM0.Data.AppData.TotalCount,
									"ApplicationID":          s.Index(i).Interface().(string),
								},
							}
							bt.client.Publish(event_time)
							logp.Info("Rum default data sent")
						}
						//get rum name, pageViews, apdex

						if true {
							//get rum name, pageViews, apdex V2

							client := &http.Client{}
							req, err := http.NewRequest("GET", "https://www.site24x7.com/api/rum/web/view/"+s.Index(i).Interface().(string)+"/"+tot[3].(string), nil)
							if err != nil {
								return err
							}
							req.Header.Add("Authorization", "Zoho-authtoken"+" "+tot[0].(string))

							if Period_counter == 1 {

								resp, err := client.Do(req)
								if err != nil {
									return err
								}
								defer resp.Body.Close()
								var RUM0 RUM_default
								if resp.StatusCode != http.StatusOK {
									bodyBytes, err2 := ioutil.ReadAll(resp.Body)
									if err2 == nil {
										logp.Info("HTTP ERROR " + string(resp.StatusCode) + " " + strings.Trim(string(bodyBytes), "\n"))
									} else {
										logp.Info(err2.Error())
										return err2
									}
									return nil
								}
								if resp.StatusCode == http.StatusOK {
									bodyBytes, err2 := ioutil.ReadAll(resp.Body)
									if err2 != nil {
										return err
									}

									json.Unmarshal(bodyBytes, &RUM0)
									ApplicationName = RUM0.Data.AppInfo.ApplicationName
									satisfied, _ := strconv.ParseFloat(RUM0.Data.AppData.SatisfiedUsers, 64)
									tolerated, _ := strconv.ParseFloat(RUM0.Data.AppData.ToleratedUsers, 64)

									event_time := beat.Event{
										Timestamp: time.Now(),
										Fields: common.MapStr{
											"mytimestamp":     time.Now(),
											"Type":            "Site24x7_RUM_UserExperience_Hourly",
											"AjaxPageViews":   RUM0.Data.AjaxPageViews,
											"PageViews":       RUM0.Data.PageViews,
											"FrustratedUsers": RUM0.Data.AppData.FrustratedUsers,
											"TotalUsers":      RUM0.Data.AppData.TotalCount,
											"ErrorCount":      RUM0.Data.AppData.ErrorCount,
											"ToleratedUsers":  RUM0.Data.AppData.ToleratedUsers,
											"SatisfiedUsers":  RUM0.Data.AppData.SatisfiedUsers,
											"UniqueUsers":     RUM0.Data.UniqUsers,
											"APDEX_aggregate": (satisfied + tolerated*0.5) / RUM0.Data.AppData.TotalCount,
											"ApplicationName": RUM0.Data.AppInfo.ApplicationName,
											"ApplicationID":   s.Index(i).Interface().(string),
										},
									}
									bt.client.Publish(event_time)
									logp.Info("Rum default data V2 sent")
								}
							}
							if Period_counter >= Period {
								Period_counter = 1
							} else {
								Period_counter += 1
							}

						}
						//get rum name, pageViews, apdex V2

						if strings.ToLower(tot[4].(string)) == "true" {

							client := &http.Client{}
							req, err := http.NewRequest("GET", "https://www.site24x7.com/api/rum/web/view/"+s.Index(i).Interface().(string)+"/graph/responsetime/"+tot[3].(string), nil)
							req.Header.Add("Authorization", "Zoho-authtoken"+" "+tot[0].(string))

							resp, err := client.Do(req)
							if err != nil {
								return err
							}
							defer resp.Body.Close()
							var RUM1 RUM_respTime
							if resp.StatusCode != http.StatusOK {
								bodyBytes, err2 := ioutil.ReadAll(resp.Body)
								if err2 == nil {
									logp.Info("HTTP ERROR " + string(resp.StatusCode) + " " + strings.Trim(string(bodyBytes), "\n"))
								} else {
									logp.Info(err2.Error())
									return err2
								}
								return nil
							}

							if resp.StatusCode == http.StatusOK {
								bodyBytes, err2 := ioutil.ReadAll(resp.Body)

								if err2 != nil {
									return err
								}

								json.Unmarshal(bodyBytes, &RUM1)

								for _, d := range RUM1.Data.Num0.Responsetimegraph.ChartData {

									event_time := beat.Event{
										Timestamp: time.Now(),
										Fields: common.MapStr{
											"mytimestamp":           d[0],
											"Type":                  "Site24x7_RUM_Responsetime",
											"NetworkTime":           d[1],
											"ServerTime":            d[2],
											"DocumentRenderingTime": d[3],
											"FirstByteTime":         d[4],
											"OverallResponseTime":   d[5],
											"ApplicationID":         s.Index(i).Interface().(string),
											"ApplicationName":       ApplicationName,
										},
									}

									if total_time[total_counter][i][0].Sub(getTime(d[0].(string))) < 0 {
										bt.client.Publish(event_time)
										logp.Info("Rum response time sent")
										if total_time[total_counter][i][1].Sub(getTime(d[0].(string))) < 0 {
											total_time[total_counter][i][1] = getTime(d[0].(string))
										}
									} else {
										logp.Info("Rum response time duplicate not sent")
									}

								}
								total_time[total_counter][i][0] = total_time[total_counter][i][1]

							}
						}

						if strings.ToLower(tot[5].(string)) == "true" {

							client := &http.Client{}
							req, err := http.NewRequest("GET", "https://www.site24x7.com/api/rum/web/view/"+s.Index(i).Interface().(string)+"/graph/throughput/"+tot[3].(string), nil)
							req.Header.Add("Authorization", "Zoho-authtoken"+" "+tot[0].(string))

							resp, err := client.Do(req)
							if err != nil {
								return err
							}
							defer resp.Body.Close()
							var RUM1 RUM_throughput
							if resp.StatusCode != http.StatusOK {
								bodyBytes, err2 := ioutil.ReadAll(resp.Body)
								if err2 == nil {
									logp.Info("HTTP ERROR " + string(resp.StatusCode) + " " + strings.Trim(string(bodyBytes), "\n"))
								} else {
									logp.Info(err2.Error())
									return err2
								}
								return nil
							}

							if resp.StatusCode == http.StatusOK {
								bodyBytes, err2 := ioutil.ReadAll(resp.Body)

								if err2 != nil {
									return err
								}

								json.Unmarshal(bodyBytes, &RUM1)

								for _, d := range RUM1.Data.Num0.Throughputgraph.ChartData {

									event_time := beat.Event{
										Timestamp: time.Now(),
										Fields: common.MapStr{
											"mytimestamp":     d[0],
											"Type":            "Site24x7_RUM_Thrughput",
											"Throughput":      d[1],
											"ApplicationID":   s.Index(i).Interface().(string),
											"ApplicationName": ApplicationName,
										},
									}

									if total_time[total_counter][i][2].Sub(getTime(d[0].(string))) < 0 {
										bt.client.Publish(event_time)
										logp.Info("Rum throughput sent")
										if total_time[total_counter][i][3].Sub(getTime(d[0].(string))) < 0 {
											total_time[total_counter][i][3] = getTime(d[0].(string))
										}
									} else {
										logp.Info("Rum throughput duplicate not sent")
									}

								}
								total_time[total_counter][i][2] = total_time[total_counter][i][3]

							}
						}

						if strings.ToLower(tot[6].(string)) == "true" {

							client := &http.Client{}
							req, err := http.NewRequest("GET", "https://www.site24x7.com/api/rum/web/view/"+s.Index(i).Interface().(string)+"/wt/list/avgRT/"+tot[3].(string), nil)
							req.Header.Add("Authorization", "Zoho-authtoken"+" "+tot[0].(string))

							resp, err := client.Do(req)
							if err != nil {
								return err
							}
							defer resp.Body.Close()
							var RUM1 RUM_webpageDetails
							if resp.StatusCode != http.StatusOK {
								bodyBytes, err2 := ioutil.ReadAll(resp.Body)
								if err2 == nil {
									logp.Info("HTTP ERROR " + string(resp.StatusCode) + " " + strings.Trim(string(bodyBytes), "\n"))
								} else {
									logp.Info(err2.Error())
									return err2
								}
								return nil
							}

							if resp.StatusCode == http.StatusOK {
								bodyBytes, err2 := ioutil.ReadAll(resp.Body)

								if err2 != nil {
									return err
								}

								json.Unmarshal(bodyBytes, &RUM1)

								for _, d := range RUM1.Data.WtList {

									event_time := beat.Event{
										Timestamp: time.Now(),
										Fields: common.MapStr{
											"mytimestamp":         time.Now(),
											"AverageResponseTime": d.AverageResponseTime,
											"TotalResponseTime":   d.TotalResponseTime,
											"TotalCount":          d.TotalCount,
											"Type":                "Site24x7_RUM_WebPageDetails",
											"WebPageName":         d.Name,
											"Throughput":          d.Throughput,
											"ApplicationID":       s.Index(i).Interface().(string),
											"ApplicationName":     ApplicationName,
										},
									}
									bt.client.Publish(event_time)
									logp.Info("Rum web page metrics sent")
								}

							}
						}

					}

				}

			}

		}
		counter++
	}
}

//}    to fool notepad++

func (bt *Site24x7beat) Stop() {
	bt.client.Close()
	close(bt.done)
}

