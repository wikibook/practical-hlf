# 5.5 프라이빗 데이터 사용하기
프라이빗 데이터를 사용하기 위한 블록체인 네트워크는 4.1절에서의 BYFN 블록체인 네트워크 구성과 동일 합니다.

## collection_config.json
프라이빗 데이터에 대한 접근 정책과 정책의 이름을 정하는 컬렉
션 정의 파일

[컬렉션 정의 내용]
1. 컬렉션 이름: personalInfo
2. 컬렉션 정책: ORG1조직의 구성원 만이 접근 가능
3. 최소 응답 피어수: 0개
4. 전파 대상 최대 피어수: 3개
5. 임시 저장소에 기록되는 최대 블록 수: 2개


## picc.go
개인 정보를 저장 및 조회할 수 있는 스마트 컨트랙트

### 주요 함수
1. savePersonalInfo
프라이빗 데이터에 개인 정보(개인 식별값, 성별, 등록번호) 저장

2. getPersonalInfo
프라이빗 데이터에서 개인 정보(개인 식별값, 성별, 등록번호)를 조회

## 프라이빗 데이터를 사용하는 체인코드 설치
1. 체인코드 개발 후 BYFN 체인코드 디렉토리에 탑재
{fabric-sample 디렉토리 상위 위치}/fabric-samples/chaincode/picc

2.CLI 접속
docker exec -it cli bash

3. 체인코드 install
peer chaincode install -n picc -v 1.0 -p github.com/chaincode/picc

4. 체인코드 intantiate
peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile $ORDERER_CA -C mychannel -n picc -v 1.0 -c
'{"Args":["init"]}' -P "OR('Org1MSP.member','Org2MSP.member')" --collections-config $GOPATH/src/github.com/chaincode/picc/collection_config.json

5. 각각의 피어에 체인코드 설치를 위해 피어 연결 정보 변경 후 3,4 단계 반복
  - ORG1의 피어 연결 정보로 변경
  export CORE_PEER_ADDRESS=peer0.org1.example.com:7051
  export CORE_PEER_LOCALMSPID=Org1MSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
  export PEER0_ORG1_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt

  - ORG2의 피어 연결 정보로 변경
  export CORE_PEER_LOCALMSPID=Org2MSP
  export PEER0_ORG2_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
  export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
  export CORE_PEER_ADDRESS=peer0.org2.example.com:7051

6. 체인코드 초기화
  - 오더러 경로 설정
  export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

## 프라이빗 데이터를 사용하는 체인코드 호출
  - ORG1의 PEER0 연결 정보로 CLI 설정
  export CORE_PEER_ADDRESS=peer0.org1.example.com:7051
  export CORE_PEER_LOCALMSPID=Org1MSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
  export PEER0_ORG1_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt

  - 개인 정보 저장 호출(savePersonalInfo)을 위한 데이터 생성
  export PI=$(echo -n "{\"id\":\"LEE\",\"gender\":\"male\",\"registrationNum\":\"890105\"}" | base64 | tr -d \\n)

  - 개인 정보 저장 호출(savePersonalInfo)
  peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n picc -c '{"Args":["savePersonalInfo"]}' --transient "{\"personalInfo\":\"$PI\"}"

  - 개인 정보 조회 호출(getPersonalInfo)
  peer chaincode query -C mychannel -n picc -c '{"Args":["getPersonalInfo","LEE"]}'