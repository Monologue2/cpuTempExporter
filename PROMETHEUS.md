# Prometheus? 
프로메테우스는 모니터링, 알람을 위해 사운드클라우드에서 만든 오픈 소스 시스템 툴 킷입니다.<br>
- key/value 쌍을 통한 시계열을 가진 다중 차원 데이터 모델의 식별
- 다중 차원 데이터를 활용하기 위한 쿼리 언어 PromQL 지원
- 분산 스토리지에 의존하지 않은 자율 운영(Autonomous) 서버
- 시계열 데이터 수집은 HTTP를 통한 Pull 모델로 발생
- 시계열 데이터 Push는 중간 게이트웨이(Intermediary gateway)가 지원합니다.
- 서비스 검색 또는 정적 구성을 통해 검색합니다.
- 그래프 작성 및 대시보드를 지원합니다.(보통 그라파나로 하지만요..)

## Whatr are Metrics?
메트릭(Metric, 직역시 측정 항목)은 일반적으로 `수치를 측정한 것`. 의미합니다.<br>
시계열(Time Series)는 `시간에 따른 변화를 기록하는 것`을 의미합니다.<br>
사용자가 측정하려는 항목은 애플리케이션마다 다릅니다. (웹 서버라면 요청 수를 측정할 겁니다. , 데이터베이스에선 활성화 된 연결 수 또는 활성화된 쿼리 수 가 되겠네요.)<br>


메트릭은 왜 애플리케이션이 어떤 방식으로 동작하는지 아는데에 중요한 역할을 합니다.<br>
웹 애플리케이션을 운영하고, 왜 느린지 밝혀내야한다고 가정해봅시다.<br>
애플리케이션에서 무슨 일이 일어났는지 알기 위해선 정보가 필요합니다.<br>
예시로, 많은 요청이 생기면 애플리케이션이 느려질 수도 있습니다.<br>
만약 당신이 요청 수 Count 메트릭이 있을 경우, 원인을 파악하고 트래픽을 감당하기 위한 서버를 늘릴 수 있을겁니다.<br>

## Architecture
이 다이어그램에서 Prometheus의 전체 구조를 볼 수 있습니다.
<img src="./img/Prometheus Architecture.png">

### Pull Metrics
프로메테우스 서버가 직접 메트릭을 가져오는 방식입니다. <br>
주로 Jobs/Exporters에서 메트릭을 가져옵니다.
- Exporter : 다양한 시스템, 서비스, 애플리케이션에서 메트릭을 수집하여 **프로메테우스가 이해할 수 있는 형식**으로 제공하는 역할을 합니다.
- 단기 작업(Short-lived jobs)이나 애플리케이션에서 Pushgateway로 메트릭을 푸시합니다.(Push) 프로메테우스는 Pushgateway로부터 메트릭을 가져옵니다.(Pull)

### Discover Targets
서비스 디스커버리 메커니즘을 통해 모니터링 대상의 위치를 자동으로 발견합니다.<br>
- `Kubernetes`와 `file_sd`와 같은 서비스 디스커버리 도구를 통해 모니터링할 대상들을 자동으로 발견합니다.