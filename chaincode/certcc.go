package main

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Certificate struct {
	ID           string `json:"id"`
	StudentID    string `json:"studentID"`
	Course       string `json:"course"`
	Grade        string `json:"grade"`
	IssueDate    string `json:"issueDate"`
	IssuedBy     string `json:"issuedBy"`
	IsRevoked    bool   `json:"isRevoked"`
}

type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Institution struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (s *SmartContract) RegisterStudent(ctx contractapi.TransactionContextInterface, id, name, email string) error {
	// TODO

}

func (s *SmartContract) RegisterInstitution(ctx contractapi.TransactionContextInterface, id, name string) error {
	// TODO

}

func (s *SmartContract) IssueCertificate(ctx contractapi.TransactionContextInterface, certID, studentID, course, grade, institutionID string) error {
	// TODO

}

func (s *SmartContract) RevokeCertificate(ctx contractapi.TransactionContextInterface, certID string) error {
	// TODO

}

func (s *SmartContract) VerifyCertificate(ctx contractapi.TransactionContextInterface, certID string) (bool, error) {
	// TODO

}

func (s *SmartContract) GetStudentCertificates(ctx contractapi.TransactionContextInterface, studentID string) ([]*Certificate, error) {
	// TODO

}

func main() {
	cc, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		panic(err.Error())
	}
	if err := cc.Start(); err != nil {
		panic(err.Error())
	}
}
