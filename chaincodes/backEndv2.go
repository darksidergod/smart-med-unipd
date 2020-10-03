package main

import (
	"context"
	"crypto/sha1"
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
type recordHash struct{ levels [10]string }

var m = make(map[string]string)
var h = make(map[string]recordHash)
var index = make(map[string]int)

func getfs(repoName string) afero.Fs {
	githubToken := "GITHUB_ACCESS_TOKEN"
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	fs, err := githubfs.NewGithubfs(client, "darksidergod", repoName, "master")
	if err != nil {
		panic(err)
	}
	return fs
}

func (clientdid *mycc) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Init executed")
	return shim.Success([]byte("true"))
}

func computeHash(s string) []byte {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return bs
}

func (clientdid *mycc) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	funcName, args := stub.GetFunctionAndParameters()
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
	fs := getfs("recordStorage")
	patientID = m[patientID]
	path := patientID + "/" + levelToRead + "/EHR"
	data, _ := afero.ReadFile(fs, path)
	toPrint := string(data)
	return shim.Success([]byte(toPrint))
}

func (clientdid *mycc) readLogs(stub shim.ChaincodeStubInterface, patientID string) peer.Response {

	enrollID, _, err := cid.GetAttributeValue(stub, "hf.EnrollmentID")
	if err != nil {
		return shim.Error("error occured while getting the enrollment ID.")
	}
	if enrollID != "healthcare-admin" {
		return shim.Error("You are not allowed to perform this function.")
	}

	fs := getfs("recordStorage")
	patientID = m[patientID]
	path := patientID + "/" + "logs"

	data, _ := afero.ReadFile(fs, path)
	fmt.Println(string(data))
	toPrint := string(data)
	return shim.Success([]byte(toPrint))
}

func (clientdid *mycc) write(stub shim.ChaincodeStubInterface, patientID string, levelToRead string) peer.Response {
	fs := getfs("recordStorage")
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
	data, _ := afero.ReadFile(fs, path)
	toPrint := string(data)
	hashTree := h[patientID]
	i := index[levelToRead]
	hashTree.levels[i] = string(computeHash(toPrint))

	//adding logs
	logpath := patientID + "/" + "/logs"
	logs, err := fs.OpenFile(logpath, os.O_APPEND, 0644)
	if err != nil {
		return shim.Error("error")
	}
	t, err := stub.GetTxTimestamp()
	timeStamp := t.String()
	enrollID, _, _ := cid.GetAttributeValue(stub, "hf.EnrollmentID")
	logs.Write([]byte(timeStamp + "\n" + enrollID + "write operation on " + levelToRead + "\n"))
	err = logs.Close()
	if err != nil {
		return shim.Error("error")
	}
	//logs part ends
	return shim.Success([]byte("Success."))
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode.\n")
	m["dataOwner1"] = "patient"
	index["lv01"] = 0
	index["lv02"] = 1
	index["lv03"] = 2
	index["lv04"] = 3
	index["lv05"] = 4
	index["lv06"] = 5
	index["lv07"] = 6
	index["lv08"] = 7
	index["lv09"] = 8
	index["lv10"] = 9
	err := shim.Start(new(mycc))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
