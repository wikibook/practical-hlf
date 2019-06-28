# 5.4 체인코드 만들고 사용하기 예제코드
### 책 본문에 자세한 설명이 있기때문에 해당 파일에서는 간단하게만 설명하겠습니다.
1. 해당 코드는 4.1절에서 다룬 first-network 예제에서 사용한 chaincode_example02.go 체인코드를 활용하였습니다.
2. 체인코드의 이름은 사용자마음대로 변경 가능하지만, 체인코드 개발에 집중하는것이 이번 장에 목표임으로, 체인코드명을 그대로하고, 공식홈페이지에서 제공하는
    명령어를 사용하겠습니다.
3. 업로드한 chaincode_example02.go 파일을 모두 복사하여 덮어쓰기해서 활용해도 되고, 기존 체인코드를 열고, 한줄한줄 수정해가면서 따라해도 됩니다.
4. 해당 체인코드의 이름과 경로를 바꾸어 설치 할 경우 책 본문을 학습하고, 그에 맞는 명령어로 변경하여 사용하면 됩니다.



## 초기 셋팅 
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=peer0.org1.example.com:7051
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CHANNEL_NAME=mychannel
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt


## 채널생성
$ peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

## 조인채널
$peer channel join -b mychannel.block

## 체인코드 설치
$peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/chaincode_example02/go/

## 체인코드 초기화
$peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -v 1.0 -c '{"Args":["init"]}'

## 사용자 등록과 토큰발행
peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $TLS_ROOTCERT_FILE -c '{"Args":["makeIdAndVal","LSH","1000"]}'
peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $TLS_ROOTCERT_FILE -c '{"Args":["makeIdAndVal","KSJ","3000"]}'

## 토큰거래
peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $TLS_ROOTCERT_FILE -c '{"Args":["moveVal","KSJ","LSH","1000"]}'

## 전체조회
peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $TLS_ROOTCERT_FILE -c '{"Args":["query"]}'

## 아이디로 조회
peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $TLS_ROOTCERT_FILE -c '{"Args":["queryById","LSH"]}'

## 체인코드 업그레이드 ( 새로운 버전으로 설치 후에 사용)
peer chaincode upgrade -o orderer.example.com:7050 --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -v 1.1 -c '{"Args":["init"]}'
