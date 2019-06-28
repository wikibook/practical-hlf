## 5.6 하이퍼레저 패브릭 멀티 호스트 구성
멀티 호스트(3대의 VM) 기반의 1개의 조직(카우치디비를 사용하는 2개의 피어), 1개의 오더러, 1개의 CA, 3개의 kafka, zookeeper로 이루어진 블록체인 네트워크를 구축 하기위한 내용 입니다. 

## 환경
환경 : [Hyperledger fabric] v1.4
      [OS] macOS High Sierra

docker swarm 버전:

## 하이퍼레저 패브릭 설정 파일
cryptogen과 configtxgen을 사용하여 인증서 및 패브릭 설정 트랜잭션들을 생성 하였습니다.
- crypto-config.yaml
1개의 오더러, 1개의 조직(ORG1) 그리고 ORG1 조직에 속하는 2개의 피어 구성 정보를 포함 합니다.

- configtx.yaml
kafka 기반의 오더링 서비스 설정으로 1개의 오더러, 1개의 조직 으로 동작하는 채널을 구성하기 위한 내용을을 포함 합니다.

## 도커 컴포즈 파일
- docker-compose-mq.yaml
3대의 VM에 각각 1개의 kafka, zookeeper 컨테이너를 구동시키기 위한 설정

- docker-compose.yaml
첫 번째 VM에 첫 번쨰 피어와 해당 피어가 월드 스테이트로 사용하는 카우치디비 컨테이너를 구동
두 번쨰 VM에 두 번쨰 피어와 해당 피어가 월드 스테이트로 사용하는 카우치디비 컨테이너를 구동
세 번째 VM에 CA와, 오더러 컨테이너를 구동 하기 위한 설정을 포함 합니다.