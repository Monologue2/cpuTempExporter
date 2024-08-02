package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// CPU 온도 Collector 정의
// Collector Interface를 따라 Describe, Collect Method를 작성한다.
type cpuTempCollector struct {
	// prometheus.Desc은 이름, 도움말, 텍스트, 라벨을 포함하여 측정 항목에 대한 메타 데이터를 제공하는 설명자
	tempDesc *prometheus.Desc
}

// 새 CPU 온도 Collector 생성
func newCpuTempCollector() *cpuTempCollector {
	return &cpuTempCollector{
		// func NewDesc(fqName, help string, variableLabels []string, constLabels Labels) *Desc
		// HELP cpu_temperature_celsius Current CPU temperature in Celsius
		// 이름은 cpu_temperature_celsius, HELP 는 Current CPU temperature in Celsius
		tempDesc: prometheus.NewDesc("cpu_temperature_celsius", "Current CPU temperature in Celsius", nil, nil),
	}
}

// 수집기에서 메트릭을 수집하는 메서드 구현
// func (m *MetricVec) Describe(ch chan<- *Desc)
// Describe를 보내는 역할
func (collector *cpuTempCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.tempDesc
}

// Collect, 실제 지표를 수집하는 역할
func (collector *cpuTempCollector) Collect(ch chan<- prometheus.Metric) {
	// Float64, Err 반환받음
	temp, err := readCPUTemperature()
	if err != nil {
		log.Printf("Error reading CPU temperature: %v", err)
		return
	}
	// Descriptor와 Gauge Metric Type, temp 전달
	// func MustNewConstMetric(desc *Desc, valueType ValueType, value float64, labelValues ...string) Metric
	ch <- prometheus.MustNewConstMetric(collector.tempDesc, prometheus.GaugeValue, temp)
}

// CPU 온도 읽기
func readCPUTemperature() (float64, error) {
	// os 라이브러리에서 로그 파일을 읽는다.
	data, err := os.ReadFile("/sys/class/thermal/thermal_zone0/temp") // 해당 파일을 cat 할 시 39000
	if err != nil {
		return 0, err
	}

	// Data를 Trim한다. 문자열 좌, 우의 공백,을 제거한다.
	tempStr := strings.TrimSpace(string(data))
	// String에서 Float64로 변환
	tempMilliCelsius, err := strconv.ParseFloat(tempStr, 64)
	if err != nil {
		return 0, err
	}

	// 섭씨로 변환 후 반환
	return tempMilliCelsius / 1000.0, nil
}

func main() {
	collector := newCpuTempCollector()
	// func NewRegistry() *Registry
	// NewRegistry creates a new vanilla Registry without any Collectors pre-registered.
	reg := prometheus.NewRegistry()
	// func (r *Registry) MustRegister(cs ...Collector)
	// 여러 개의 Collector 등록 가능하다.
	reg.MustRegister(collector)

	// http Handle 함수에 promhttp 라이브러리의 Handler를 전달하여 /metrics 엔드포인트 expose
	// PromHandler 에는 Register와, HandlerOption이 포함된다.
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	fmt.Println("Starting server on :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
