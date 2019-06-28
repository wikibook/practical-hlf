## 7.2절 프로메테우스와 그라파나를 이용한 메트릭 모니터링
환경 : [Hyperledger fabric] v1.4
      [prometheus] v2.9.1
      [OS] macOS High Sierra

메트릭 모니터링을 위해서는 하이퍼레저 패브릭 환경 설정과 프로메테우스 환경 설정이 필요합니다.
하이퍼레저 패브릭에서 프로메테우스 연동을 1.4버전에서만 지원하기때문에 1.4버전으로만 진행해주시기 바랍니다.
프로메테우스의 경우에는 버전에 상관없이 yml파일을 업데이트 해주면 부팅시에 해당 설정으로 동작합니다.

## 하이퍼레저 패브릭 설정 파일
해당 실습에서는 First-network를 기준으로 블록체인 네트워크를 구성하였습니다.
아래 세가지 파일을 해당 위치에 업데이트합니다.

- core.yaml
위치 : ~/fabric-samples-release-1.4/config

- orderer.yaml
위치 : ~/fabric-samples-release-1.4/config

- docker-compose-base.yaml
위치 : ~/fabric-samples-release-1.4/first-network/base


## 프로메테우스 설정 파일
- prometheus.yml
위치 : ~/prometheus-2.9.1.darwin-amd64

해당 실습에서는 블록체인 네트워크가 동작한 뒤 프로메테우스를 도커로 구동시킵니다.
프로메테우스를 설치하게되면 해당 폴더에 기존 prometheus.yml파일이 존재하는데 이 파일을 업데이트합니다(prometheus-2.9.1.darwin-amd64는 예시입니다).


## 메트릭 모니터링 구동 커맨드
[블록체인 네트워크 구동]
cd ~/fabric-samples-release-1.4/first-network
./byfn.sh -m down
./byfn.sh -m generate
./byfn.sh -m up


[프로메테우스 구동]
 docker run -d --name prometheus-server -p 9090:9090 \
 --restart always \
 -v /Users/name/fabric-samples-release-1.4/prometheus-2.9.1.darwin-amd64/prometheus.yml:/prometheus.yml \
 prom/prometheus \
 --config.file=/prometheus.yml

 위 명령어중 v옵션에 해당하는 경로는 작성자의 local pc를 기준으로 작성하였습니다. 
 ':'을 기준으로 왼쪽 부분에 해당하는 경로는 각자 사용자의 pc에 해당하는 경로로 수정하여 사용하시기 바랍니다.


 [블록체인 네트워크와 프로메테우스 연동]
 docker network connect net_byfn 6c7edd355328

 위 명령어중 '6c7edd355328' 는 프로메테우스의 실습 당시의 도커 컨테이너 아이디입니다. 
 해당 부분은 사용자가 설치한 컨테이너의 아이디로 수정하여 사용하시기 바랍니다.
