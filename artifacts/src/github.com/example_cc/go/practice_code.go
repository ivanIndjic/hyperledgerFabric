/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("example_cc0")

type client struct {
	IDClient      string
	AccountNumber string
	Name          string
	LastName      string
	Email         string
	MoneyAmount   float64
	IDCredit      string
}

type bank struct {
	IDBank            string
	Name              string
	EstYear           int
	OriginCountry     string
	BusinessCountries []string
	Clients           []client
}

type transaction struct {
	ID         string
	Date       time.Time
	IDSender   string
	IDReciever string
	Amount     float64
}

type credit struct {
	IDCredit        string
	ApprovalDate    time.Time
	EndDate         time.Time
	RateSize        float64
	Interest        float64
	TotalNumOfRates int
	PaidRates       int
	MoneyAmount     float64
}

type clientHistory struct {
	IDClient string
	History  []float64
}

// Global variables for ID
var nextIDClient int
var nextIDBank string
var nextTransactionID string
var nextCreditID int

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Init initializes ledger
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### Init ###########")

	var duration = time.Duration(10)
	var duration2 = time.Duration(20)
	creditDuration1, _ := time.ParseDuration("120h")
	creditDuration2, _ := time.ParseDuration("420h")
	creditDuration3, _ := time.ParseDuration("520h")
	creditDuration4, _ := time.ParseDuration("620h")
	var client1 = client{"c1", "12345", "Ivan", "Indjic", "ivan123indjic@gmail.com", 500, "-1"}
	var client2 = client{"c2", "22345", "Milan", "Simic", "simke123simic@gmail.com", 5500000, "cr2"}
	var client3 = client{"c3", "12346", "Nikolina", "Tomic", "curka@gmail.com", 500000, "cr3"}
	var transaction1 = transaction{"t1", time.Now(), "2", "1", 5000}
	var transaction2 = transaction{"t2", time.Now().Add(duration), "1", "2", 15000}
	var transaction3 = transaction{"t3", time.Now().Add(duration2), "3", "1", 8000}
	var transaction4 = transaction{"t4", time.Now().Add(duration), "2", "3", 9000}
	var credit1 = credit{"cr1", time.Now(), time.Now().Add(creditDuration1), 1000, 0.7, 100, 22, 100000}
	var credit2 = credit{"cr2", time.Now(), time.Now().Add(creditDuration2), 2000, 0.2, 100, 12, 200000}
	var credit3 = credit{"cr3", time.Now(), time.Now().Add(creditDuration3), 3000, 0.3, 100, 25, 300000}
	var credit4 = credit{"cr4", time.Now(), time.Now().Add(creditDuration4), 4000, 0.1, 100, 55, 400000}
	nextCreditID = 5
	nextTransactionID = "t5"
	nextIDClient = 4
	var countries = make([]string, 1, 20)
	var clients1 = make([]client, 1, 20)
	var clients2 = make([]client, 1, 20)
	var historyValues []float64
	var historyValue2 []float64
	var historyValue3 []float64

	historyValues = append(historyValues, 8000.0, 5000.0)
	historyValue2 = append(historyValue2, 15000.0)
	historyValue3 = append(historyValue3, 9000.0)

	var historyForC1 = clientHistory{"c1h", historyValues}
	var historyForC2 = clientHistory{"c2h", historyValue2}
	var historyForC3 = clientHistory{"c3h", historyValue3}

	clients1 = append(clients1, client1, client2)
	clients2 = append(clients2, client3)
	countries = append(countries, "Senegal")
	countries = append(countries, "Bosina")
	var bank1 = bank{"b1", "Intesa", 1999, "Serbia", countries, clients1}
	var bank2 = bank{"b2", "Unicredit", 2002, "Serbia", countries, clients2}
	nextIDBank = "b3"

	// Write the state to the ledger

	ajson, _ := json.Marshal(historyForC1)
	err := stub.PutState(historyForC1.IDClient, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(historyForC2)
	err = stub.PutState(historyForC2.IDClient, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(historyForC3)
	err = stub.PutState(historyForC3.IDClient, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(client1)
	err = stub.PutState(client1.IDClient, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(client2)
	err = stub.PutState(client2.IDClient, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(client3)
	err = stub.PutState(client3.IDClient, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(transaction1)
	err = stub.PutState(transaction1.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(transaction2)
	err = stub.PutState(transaction2.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(transaction3)
	err = stub.PutState(transaction3.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(transaction4)
	err = stub.PutState(transaction4.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(credit1)
	err = stub.PutState(credit1.IDCredit, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(credit2)
	err = stub.PutState(credit2.IDCredit, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(credit3)
	err = stub.PutState(credit3.IDCredit, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(credit4)
	err = stub.PutState(credit4.IDCredit, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(bank1)
	err = stub.PutState(bank1.IDBank, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(bank2)
	err = stub.PutState(bank2.IDBank, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Invoke routes functions
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// logger.Info("########### example_cc0 Invoke ###########")

	function, args := stub.GetFunctionAndParameters()
	if function == "query" {
		return t.query(stub, args)
	}

	if function == "transfer" {
		return t.transfer(stub, args)
	}
	if function == "credit" {
		return t.credit(stub, args)
	}
	if function == "payRate" {
		return t.payRate(stub, args)
	}

	if function == "addClient" {
		return t.addClient(stub, args)
	}

	// logger.Errorf("Unknown action, check the first argument, must be one of 'delete', 'query'. But got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: "))
}

func (t *SimpleChaincode) addClient(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 6 {
		return shim.Error("Wrong number of arguments! Usage: AccountNum,Name,Last Name,Email, Money Amount, IDCredit defalut[-1]")
	}
	id := "c"+strconv.Itoa(nextIDClient)
	nextIDClient = nextIDClient + 1
	account := args[0]
	name := args[1]
	lastName := args[2]
	email := args[3]
	amount, errConv := strconv.ParseFloat(args[4], 64)
	if errConv != nil {
		return shim.Error("Wrong value for amount! Use float64 values")

	}
	credit := args[5]

	clientBytes, _ := stub.GetState(id)

	if clientBytes != nil {
		return shim.Error("Client with given ID is already in ledger")
	}
	var clientC = client{id, account, name, lastName, email, amount, credit}

	ajson, _ := json.Marshal(clientC)
	err := stub.PutState(clientC.IDClient, ajson)
	if err != nil {
		return shim.Error("Error when adding user")
	}

	return shim.Success(nil)

}

func (t *SimpleChaincode) transfer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var SenderID, RecieverID string
	var allowMinus string
	if len(args) != 4 {
		return shim.Error("Wrong number of arguments! Usage: SenderID,RecieverID,Amount,AllowMinus")
	}
	SenderID = args[0]
	RecieverID = args[1]
	allowMinus = args[3]
	amount, errConv := strconv.ParseFloat(args[2], 64)
	if errConv != nil {
		return shim.Error("Wrong value for amount! Use float64 values")

	}
	var clientSender = client{}
	var clientReciever = client{}

	clientBytes, err := stub.GetState(SenderID)
	if err != nil {
		return shim.Error("Cannot get state for Sender")
	}
	if clientBytes == nil {
		return shim.Error("Wrong value for Sender")
	}

	errUnmarshal := json.Unmarshal(clientBytes, &clientSender)
	if errUnmarshal != nil {
		return shim.Error("Cannot unmarshal client sender")

	}

	recieverClientBytes, errReciever := stub.GetState(RecieverID)
	if errReciever != nil {
		return shim.Error("Cannot get stat for reciever")
	}
	if recieverClientBytes == nil {
		return shim.Error("Wrong value for reciever")
	}

	errUnmarshal = json.Unmarshal(recieverClientBytes, &clientReciever)
	if errUnmarshal != nil {
		return shim.Error("Cannot unmarshal client reciever")

	}

	if amount <= clientSender.MoneyAmount {
		clientSender.MoneyAmount -= amount
		clientReciever.MoneyAmount += amount
		toJSON, errMarshal := json.Marshal(clientSender)
		if errMarshal != nil {
			return shim.Error("Cannot marshal client sender")
		}
		er := stub.PutState(clientSender.IDClient, toJSON)
		if er != nil {
			return shim.Error("Error when adding client sender")
		}
		toJSON, errMarshal = json.Marshal(clientReciever)
		if errMarshal != nil {
			return shim.Error("Cannot marshal client reciever")
		}
		er = stub.PutState(clientReciever.IDClient, toJSON)
		if er != nil {
			return shim.Error("Error when adding client reciever")
		}

		iterator, err := stub.GetHistoryForKey(RecieverID + "h")
		if err != nil {
			return shim.Error(err.Error())
		}
		var CH = clientHistory{}
		var content string
		for iterator.HasNext() {
			queryResponse, errR := iterator.Next()
			if errR != nil {
				shim.Error("Cannot get next value in iterator")
			}
			content += string(queryResponse.Value)

		}
		errUnmarshal = json.Unmarshal([]byte(content), &CH)
		if errUnmarshal != nil {
			return shim.Error("Cannot unmarshal client history")

		}
		CH.History = append(CH.History, amount)
		ajson, _ := json.Marshal(CH)
		errQW := stub.PutState(CH.IDClient, ajson)
		if errQW != nil {
			return shim.Error(err.Error())
		}

		return shim.Success(nil)
	} else {

		valueMinus, errConv := strconv.Atoi(allowMinus)
		if errConv != nil {
			return shim.Error("Cannot cast value")

		}

		if valueMinus == 0 {
			return shim.Error("Not enough money")
		} else {
			iterator, err := stub.GetHistoryForKey(SenderID + "h")
			if err != nil {
				return shim.Error(err.Error())
			}
			var CH = clientHistory{}
			var content string
			var sum float64
			for iterator.HasNext() {
				queryResponse, errR := iterator.Next()
				if errR != nil {
					shim.Error("Cannot get next value in iterator")
				}
				content += string(queryResponse.Value)

			}
			errUnmarshal = json.Unmarshal([]byte(content), &CH)
			if errUnmarshal != nil {
				return shim.Error("Cannot unmarshal client history")

			}

			for i := 0; i < len(CH.History); i++ {
				sum += CH.History[i]
			}
			if len(CH.History) == 0 {
				return shim.Error("Not enough money")
			}
			average := sum / float64(len(CH.History))
			if average >= amount {
				clientSender.MoneyAmount -= amount
				clientReciever.MoneyAmount += amount
				toJSON, errMarshal := json.Marshal(clientSender)
				if errMarshal != nil {
					return shim.Error("Cannot marshal client sender")
				}
				er := stub.PutState(clientSender.IDClient, toJSON)
				if er != nil {
					return shim.Error("Error when adding client sender")
				}
				toJSON, errMarshal = json.Marshal(clientReciever)
				if errMarshal != nil {
					return shim.Error("Cannot marshal client reciever")
				}
				er = stub.PutState(clientReciever.IDClient, toJSON)
				if er != nil {
					return shim.Error("Error when adding client reciever")
				}

				iteratorK, errK := stub.GetHistoryForKey(RecieverID + "h")
				if errK != nil {
					return shim.Error(errK.Error())
				}
				var CH2 = clientHistory{}
				var content2 string
				for iteratorK.HasNext() {
					queryResponse2, errR2 := iteratorK.Next()
					if errR2 != nil {
						shim.Error("Cannot get next value in iterator")
					}
					content2 += string(queryResponse2.Value)

				}
				errUnmarshal2 := json.Unmarshal([]byte(content2), &CH2)
				if errUnmarshal2 != nil {
					return shim.Error("Cannot unmarshal client history")

				}
				CH2.History = append(CH2.History, amount)
				ajson, _ := json.Marshal(CH2)
				errQW := stub.PutState(CH.IDClient, ajson)
				if errQW != nil {
					return shim.Error(err.Error())
				}

				return shim.Success(nil)

			} else {
				return shim.Error("Average is lower than amount")
			}

		}

	}

	return shim.Success(nil)

}

func (t *SimpleChaincode) credit(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Usage: Client ID, Amount, Number of rates")
	}

	ClientID := args[0]
	amount := args[1]
	numOfRates := args[2]
	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return shim.Error(err.Error())
	}
	numOfRatesInt, err2 := strconv.Atoi(numOfRates)
	if err2 != nil {
		return shim.Error(err2.Error())
	}

	var CH = clientHistory{}
	var content string
	var sum float64

	iteratorK, errK := stub.GetHistoryForKey(ClientID + "h")
	if errK != nil {
		return shim.Error(errK.Error())
	}

	for iteratorK.HasNext() {
		queryResponse, errR := iteratorK.Next()
		if errR != nil {
			shim.Error("Cannot get next value in iterator")
		}
		content += string(queryResponse.Value)

	}
	errUnmarshal := json.Unmarshal([]byte(content), &CH)
	if errUnmarshal != nil {
		return shim.Error("Cannot unmarshal client history")

	}

	for i := 0; i < len(CH.History); i++ {
		sum += CH.History[i]
	}
	if len(CH.History) == 0 {
		return shim.Error("Not enough money")
	}
	average := sum / float64(len(CH.History))

	if amountFloat <= average*5 {

		var clientC = client{}

		clientBytes, err := stub.GetState(ClientID)
		if err != nil {
			return shim.Error("Cannot get state for Client")
		}
		if clientBytes == nil {
			return shim.Error("Client doesnt exists")
		}
		errUnmarshal := json.Unmarshal(clientBytes, &clientC)
		if errUnmarshal != nil {
			return shim.Error("Cannot unmarshal client")

		}

		if clientC.IDCredit != "-1" {

			var oldCredit = credit{}

			creditBytes, err := stub.GetState(clientC.IDCredit)
			if err != nil {
				return shim.Error("Cannot get state for Credit")
			}
			if clientBytes == nil {
				return shim.Error("Credit doesnt exists")
			}
			errUnmarshal := json.Unmarshal(creditBytes, &oldCredit)
			if errUnmarshal != nil {
				return shim.Error("Cannot unmarshal credit")

			}

			if oldCredit.PaidRates == oldCredit.TotalNumOfRates {
				rate := (amountFloat / float64(numOfRatesInt)) * 0.4
				hours := numOfRatesInt * 30 * 24
				creditDuration, _ := time.ParseDuration(strconv.Itoa(hours) + "h")
				var newCredit = credit{"cr" + strconv.Itoa(nextCreditID), time.Now(), time.Now().Add(creditDuration), rate, 0.4, numOfRatesInt, 0, amountFloat}
				clientC.IDCredit = "cr" + strconv.Itoa(nextCreditID)
				clientC.MoneyAmount = amountFloat + clientC.MoneyAmount
				nextCreditID += 1
				ajson, _ := json.Marshal(newCredit)
				err := stub.PutState(newCredit.IDCredit, ajson)
				if err != nil {
					return shim.Error(err.Error())
				}
				ajson, _ = json.Marshal(clientC)
				err = stub.PutState(clientC.IDClient, ajson)
				if err != nil {
					return shim.Error(err.Error())
				}

				return shim.Success(nil)

			} else {
				return shim.Error("Cannot get credit! The old one isn't paid")
			}

		} else {
			rate := (amountFloat / float64(numOfRatesInt)) * (1 + 0.4)
			hours := numOfRatesInt * 30 * 24
			creditDuration, _ := time.ParseDuration(strconv.Itoa(hours) + "h")
			var newCredit = credit{"cr" + strconv.Itoa(nextCreditID), time.Now(), time.Now().Add(creditDuration), rate, 0.4, numOfRatesInt, 0, amountFloat}
			clientC.IDCredit = "cr" + strconv.Itoa(nextCreditID)
			clientC.MoneyAmount = amountFloat + clientC.MoneyAmount
			nextCreditID += 1
			ajson, _ := json.Marshal(newCredit)
			err := stub.PutState(newCredit.IDCredit, ajson)
			if err != nil {
				return shim.Error(err.Error())
			}
			ajson, _ = json.Marshal(clientC)
			err = stub.PutState(clientC.IDClient, ajson)
			if err != nil {
				return shim.Error(err.Error())
			}

			return shim.Success(nil)

		}
	} else {
		return shim.Error("Amount is too big")

	}

}

func (t *SimpleChaincode) payRate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Usage: Id Client, Amount")
	}
	var clientC = client{}
	var clientsCredit = credit{}

	clientID := args[0]
	amount, errConv := strconv.ParseFloat(args[1], 64)
	if errConv != nil {
		return shim.Error("Wrong value for amount! Use float64 values")

	}

	clientBytes, err := stub.GetState(clientID)
	if err != nil {
		return shim.Error("Cannot get state for Client")
	}
	if clientBytes == nil {
		return shim.Error("Client doesnt exists")
	}
	errUnmarshal := json.Unmarshal(clientBytes, &clientC)
	if errUnmarshal != nil {
		return shim.Error("Cannot unmarshal client")

	}

	creditBytes, err2 := stub.GetState(clientC.IDCredit)
	if err2 != nil {
		return shim.Error("Cannot get state for Credit")
	}
	if creditBytes == nil {
		return shim.Error("Credit doesnt exists")
	}
	errUnmarshal = json.Unmarshal(creditBytes, &clientsCredit)
	if errUnmarshal != nil {
		return shim.Error("Cannot unmarshal client")

	}

	if amount >= clientsCredit.RateSize {

		if clientsCredit.TotalNumOfRates != clientsCredit.PaidRates+1 {

			clientsCredit.TotalNumOfRates = clientsCredit.TotalNumOfRates - 1
			clientsCredit.PaidRates = clientsCredit.PaidRates + 1

			if amount > clientC.MoneyAmount {
				return shim.Error("You donw have enough money")
			} else {
				clientC.MoneyAmount = clientC.MoneyAmount - amount
			}

		} else {

			if amount > clientC.MoneyAmount {
				return shim.Error("You donw have enough money")
			} else {
				clientC.MoneyAmount = clientC.MoneyAmount - amount
				clientC.IDCredit = "-1"
			}

		}

		ajson, _ := json.Marshal(clientsCredit)
		err := stub.PutState(clientsCredit.IDCredit, ajson)
		if err != nil {
			return shim.Error(err.Error())
		}
		ajson, _ = json.Marshal(clientC)
		err = stub.PutState(clientC.IDClient, ajson)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(nil)

	} else {
		return shim.Error("Incorrect amount")
	}

}

func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the entity to query")
	}
	A := args[0]
	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"ID\":\"" + A + "\",\"Entity:\":\"" + string(Avalbytes) + "\"}"
	logger.Infof("Query Response:%s\n", jsonResp)

	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}
