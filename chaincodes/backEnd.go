package main

import (
	"context"
	"fmt"
	"os"

	"github.com/darksidergod/githubfs-test"
	"github.com/google/go-github/github"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/spf13/afero"
	"golang.org/x/oauth2"
)

type mycc struct{}

var m = make(map[string]string)

// Init Implements the Init method
func (clientdid *mycc) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Init executed")
	return shim.Success([]byte("true"))
}

func (clientdid *mycc) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	funcName, args := stub.GetFunctionAndParameters()
	fmt.Println("Function=", funcName)
	fmt.Println(args)
	if funcName == "read" {
		return clientdid.read(stub, args[0], args[1])
	} else if funcName == "write" {
		return clientdid.write(stub, args[0], args[1])
	} else if funcName == "add" {
		return clientdid.add(stub, args[0], args[1])
	}
	return shim.Error(("Bad Function Name = " + funcName + "!"))
}
func (clientdid *mycc) add(stub shim.ChaincodeStubInterface, key string, value string) peer.Response {
	enrollID, _, err := cid.GetAttributeValue(stub, "hf.EnrollmentID")
	if err != nil {
		return shim.Error("error occured while getting the enrollment ID.")
	}
	if enrollID != "healthcare-admin" {
		return shim.Error("You are not allowed to perform this function.")
	}
	m[key] = value
	return shim.Success([]byte("Added path"))
}
func (clientdid *mycc) read(stub shim.ChaincodeStubInterface, patientID string, levelToRead string) peer.Response {
	githubToken := "GITHUB_ACCESS_TOKEN"
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	fs, err := githubfs.NewGithubfs(client, "darksidergod", "recordStorage", "master")
	if err != nil {
		panic(err)
	}
	patientID = m[patientID]
	path := patientID + "/" + levelToRead + "/EHR"
	if err != nil {
		println("error")
	}
	data, _ := afero.ReadFile(fs, path)
	fmt.Println(string(data))
	toPrint := string(data)
	return shim.Success([]byte(toPrint))
}

func (clientdid *mycc) write(stub shim.ChaincodeStubInterface, patientID string, levelToRead string) peer.Response {
	githubToken := "GITHUB_ACCESS_TOKEN"
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	fs, err := githubfs.NewGithubfs(client, "darksidergod", "recordStorage", "master")
	if err != nil {
		panic(err)
	}
	sourcepath := m[patientID]
	path := sourcepath + "/" + levelToRead + "/EHR"
	f, err := fs.OpenFile(path, os.O_APPEND, 0644)
	if err != nil {
		return shim.Error("error")
	}
	f.Write([]byte("Write operation successfull."))
	err = f.Close()
	if err != nil {
		return shim.Error("error")
	}
	return shim.Success([]byte("Success."))
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode.\n")
	m["dataOwner1"] = "patient"
	err := shim.Start(new(mycc))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}

