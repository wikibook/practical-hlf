package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct {
}

type personalInfo struct {
	Id              string `json:"id"`
	Gender          string `json:"gender"`
	RegistrationNum string `json:"registrationNum"`
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "savePersonalInfo" { //create a new marble
		return t.savePersonalInfo(stub, args)
	} else if function == "getPersonalInfo" { //change owner of a specific marble
		return t.getPersonalInfo(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

func (t *SimpleChaincode) savePersonalInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	transMap, err := stub.GetTransient()
	if err != nil {
		return shim.Error(err.Error())
	}

	var pi personalInfo
	err = json.Unmarshal(transMap["personalInfo"], &pi)
	if err != nil {
		return shim.Error("Failed to decode JSON of: " + string(transMap["personalInfo"]))
	}

	fmt.Println("Start picc: savePersonalInfo func")
	if len(pi.Id) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(pi.Gender) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(pi.RegistrationNum) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}


	infoBytes, err := json.Marshal(pi)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutPrivateData("personalInfo", pi.Id, infoBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("KEY: %s \n", pi.RegistrationNum)
	bytes, err := json.Marshal(pi.Id)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(bytes)
}
func (t *SimpleChaincode) getPersonalInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	id := args[0]
	infoBytes, err := stub.GetPrivateData("personalInfo", id)

	if err != nil {
		return shim.Error(err.Error())
	} else if infoBytes == nil {
		jsonResp := "{\"ERROR\":\"personal information does not exist\" }"
		return shim.Error(jsonResp)
	}
	return shim.Success(infoBytes)
}
