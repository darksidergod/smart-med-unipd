package main

/**
 * Demonstrates the use of CID
 **/
import (
	// For printing messages on console
	"crypto/x509"
	"fmt"

	// April 2020, Updated to Fabric 2.0 Shim
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"

	peer "github.com/hyperledger/fabric-protos-go/peer"
)

// mycc Represents our chaincode object
type mycc struct {
}

// Channel Name & name of the chaincode that is to be called.
const Channel = "healthcarechannel"

// Invoke method
func (clientdid *mycc) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Get the function name and parameters
	funcName, args := stub.GetFunctionAndParameters()

	// Just to satisfy the compiler - otherwise it will complain that args declared but not used
	fmt.Println(len(args))

	if funcName == "readattrs" {
		return clientdid.ReadAttributesOfCaller(stub, args[0])
	} else if funcName == "writeperm" {
		return clientdid.ApproveWrite(stub, args[0])
	} else if funcName == "elevateperm" {
		return clientdid.ApproveElevate(stub, args[0], args[1])
	} else if funcName == "read" {
		return clientdid.readRecord(stub, args[0], args[1])
	} else if funcName == "write" {
		return clientdid.writeToRecord(stub, args[0], args[1])
	} else if funcName == "elevate" {
		return clientdid.elevatePriv(stub, args[0], args[1])
	}

	return shim.Error("Wrong function call.")
}

// ReadAttributesOfCaller reads the attributes of the callers cert and return it as JSON
func (clientdid *mycc) ReadAttributesOfCaller(stub shim.ChaincodeStubInterface, dataOwner string) peer.Response {

	// Variable to hold the result
	jsonResult := "{"

	// 1. Get the unique ID of the user
	id, err := cid.GetID(stub)

	if err != nil {
		fmt.Println("Error GetID() =" + err.Error())
		return shim.Error(err.Error())
	}
	// Format and add the attribute to JSON
	jsonResult += SetJSONNV("id", id)

	// 2. Get the MSP ID of the user
	var mspid string
	mspid, err = cid.GetMSPID(stub)

	if err != nil {
		fmt.Println("Error GetMSPID() =" + err.Error())
		return shim.Error(err.Error())
	}
	// Format and add the attribute to JSON
	jsonResult += "," + SetJSONNV("MSPID", mspid)

	// 3. Get the standard attributes added by default
	// "hf.Affiliation" ,"hf.EnrollmentID", "hf.Type"
	affiliation, _, _ := cid.GetAttributeValue(stub, "hf.Affiliation")
	enrollID, _, _ := cid.GetAttributeValue(stub, "hf.EnrollmentID")
	userType, _, _ := cid.GetAttributeValue(stub, "hf.Type")
	// Format and add the attribute to JSON
	jsonResult += "," + SetJSONNV("affiliation", affiliation)
	jsonResult += "," + SetJSONNV("enrollID", enrollID)
	jsonResult += "," + SetJSONNV("userType", userType)

	// 5. Get the attr value for "dataOwner level"
	attrValue, flag, _ := cid.GetAttributeValue(stub, dataOwner)
	if !flag {
		attrValue = "NOT SET"
	}
	// Format and add the attribute to JSON
	jsonResult += "," + SetJSONNV(dataOwner, attrValue)

	// 6. Get the Certificate of the caller - not sending it back in JSON
	var cert *x509.Certificate
	cert, err = cid.GetX509Certificate(stub)
	if err != nil {
		fmt.Println("Error GetX509Certificate() =" + err.Error())
		return shim.Error(err.Error())
	}
	fmt.Println("GetX509Certificate() = " + string(cert.RawSubject))

	// Close the JSON and send it as response
	jsonResult += "}"
	return shim.Success([]byte(jsonResult))
}

// SetJSONNV returns a name value pair in JSON format
func SetJSONNV(attr, value string) string {
	return " \"" + attr + "\":\"" + value + "\""
}

func (clientdid *mycc) ApproveWrite(stub shim.ChaincodeStubInterface, writeUser string) peer.Response {
	attrValue, _, _ := cid.GetAttributeValue(stub, "hf.Affiliation")
	checkString := string(attrValue)
	fmt.Println("Value of affiliation is " + checkString)
	if checkString != "healthcare.patient" {
		return shim.Error("You're not authorized.")
	}
	stub.PutState(writeUser, []byte("Write Auth."))
	response := shim.Success([]byte("Write access granted."))
	return response

}

func (clientdid *mycc) ApproveElevate(stub shim.ChaincodeStubInterface, writeUser string, levelReq string) peer.Response {
	attrValue, _, _ := cid.GetAttributeValue(stub, "hf.Affiliation")
	if attrValue == "healthcare.patient" {
		stub.PutState(writeUser, []byte(levelReq))
		response := shim.Success([]byte("New level added to state, please consult CA admin."))
		return response
	}
	return shim.Error("not authorized to perform this operation")
}

func (clientdid *mycc) readRecord(stub shim.ChaincodeStubInterface, dataOwner string, levelReq string) peer.Response {

	fmt.Println(dataOwner)
	attrValue, _, _ := cid.GetAttributeValue(stub, dataOwner)
	fmt.Println(attrValue)
	fmt.Println(levelReq)
	if attrValue < levelReq {
		return shim.Error("Access denied.")
	}
	args := make([][]byte, 3)
	args[0] = []byte("read")
	args[1] = []byte(dataOwner)
	args[2] = []byte(levelReq)
	targetChaincode := "backEnd"
	response := stub.InvokeChaincode(targetChaincode, args, Channel)
	return response
}

func (clientdid *mycc) writeToRecord(stub shim.ChaincodeStubInterface, dataOwner string, writeLevel string) peer.Response {
	caller, _, err := cid.GetAttributeValue(stub, "hf.EnrollmentID")
	//caller, err := cid.GetID(stub)
	if err != nil {
		return shim.Error("callerID malfunctioned")
	}
	val, err := stub.GetState(caller)
	if err != nil {
		return shim.Error("Error while getting value of state.")
	}
	valueString := string(val)
	if valueString != "Write Auth." {
		return shim.Error("No such request was made by the data owner.\n")
	}
	args := make([][]byte, 3)
	args[0] = []byte("write")
	args[1] = []byte(dataOwner)
	args[2] = []byte(writeLevel)
	targetChaincode := "backEnd"
	response := stub.InvokeChaincode(targetChaincode, args, Channel)
	stub.DelState(caller)
	return response
}

func (clientdid *mycc) elevatePriv(stub shim.ChaincodeStubInterface, dataOwner string, newLevel string) peer.Response {
	caller, err := cid.GetID(stub)
	if err != nil {
		return shim.Error("callerID malfunctioned")
	}
	val, err := stub.GetState(caller)
	if err != nil {
		return shim.Error("No such authorization has been made.\n")
	}
	authMade := string(val)
	if authMade != "Elevate" {
		return shim.Error("Wrong function call made for the authorization.\n")
	}
	args := make([][]byte, 3)
	args[0] = []byte("elevate")
	args[1] = []byte(dataOwner)
	args[2] = []byte(newLevel)
	targetChaincode := "backEnd"
	response := stub.InvokeChaincode(targetChaincode, args, Channel)
	return response
}

// Init Implements the Init method
func (clientdid *mycc) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// Simply print a message
	fmt.Println("Init executed in history")
	// Return success
	return shim.Success(nil)
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode.\n")
	err := shim.Start(new(mycc))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
