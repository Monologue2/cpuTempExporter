# [Prometheus 용어 사전](https://prometheus.io/docs/introduction/glossary/#client-library)<br>
**Prometheus Client Library**: Go, Java, Python, Ruby 등의 언어로 작성된 라이브러리, 코드를 직접 계측하고 다른 시스템에서 메트릭을 가져오는 사용자 정의 Collecter를 작성하고 Exporter를 만드는걸 쉽게 할 수 있습니다.<br>

**Exporter**: Metric을 얻으려는 Application과 함께 실행되는 바이너리(응용 프로그램)입니다. Prometheus Metrics를 외부에 노출시키며 일반적으로 Prometheus Format이 아닌 채 외부에 노출된 데이터를(File 형태의 Log, Network 인터페이스로 전달된 Json 등) Prometheus Format으로 변환합니다.<br>

**Collector**: Exporter의 일부로서, Metrics 집합을 나타냅니다. 직접 계측의 일부인 경우 단일 Metric일 수 있으며, 다른 시스템에서 pulling 하는 방식인 경우 다수의 Metrics일 수 있습니다.<br>

**Metric** : 일반적으로 `수치를 측정한 것`. 의미합니다. `측정값` 정도로 받아들이면 되겠습니다.



## [Instrumenting a Go Application for Prometheus](https://prometheus.io/docs/guides/go-application/): Prometheus용 Go Application 구현하기
> Prometheus용 Go Application 구현하기<br>
> 개인 Application을 계측하는 경우 Prometheus Client Libaray로 코드를 계측하는 방법에 대한 규칙을 따라야한다.<br>

