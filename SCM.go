
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
)



 
// POC is a high level smart contract that POCs together business artifact based smart contracts
type POC struct {

}

// UserDetails is for storing User Details

type User struct{	
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Country string `json:"country"`
	Address string `json:"address"`
	PhoneNo string `json:"phoneNo"`
	Role string `json:"role"`
	UserId string `json:"userId"`
	Password string `json:"password"`
	NodeId string `json:"nodeId"`	
}

type BookingDetails struct{	
	UserId string `json:"userId"`
	SourcePort string `json:"sourcePort"`
	DestinationPort string `json:"destinationPort"`
	BookingNo string `json:"bookingNo"`
	Consingnee string `json:"consingnee"`
	CargoType string `json:"cargoType"`
	CargoDsec string `json:"cargoDsec"`
	CargoId string `json:"cargoId"`
	ContainerId string `json:"containerId"`
	SpecialReq string `json:"specialReq"`	
	Measurements string `json:"measurements"`	
	Weight string `json:"weight"`	
	Status string `json:"status"`	
	NodeId string `json:"nodeId"`	
}

// Init initializes the smart contracts
func (t *POC) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Check if table already exists
	_, err := stub.GetTable("User")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("User", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "firstName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "email", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "country", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "address", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "phoneNo", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "role", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "userId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "password", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "nodeId", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating User table.")
	}
	
	// Check if table already exists
	_, err = stub.GetTable("BookingDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("BookingDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "userId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "sourcePort", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "destinationPort", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "bookingNo", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "consingnee", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "cargoType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "cargoDsec", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "cargoId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "containerId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "specialReq", Type: shim.ColumnDefinition_STRING, Key: false},		
		&shim.ColumnDefinition{Name: "measurements", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "weight", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "status", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "nodeId", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating Booking Details table.")
	}
	
	stub.PutState("BookingNoincrement", []byte("1"))
	return nil, nil
}

//registerUser to register a user
func (t *POC) registerUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 10 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 10. Got: %d.", len(args))
		}
		
		firstName:=args[0]
		lastName:=args[1]
		email:=args[2]
		country:=args[3]
		address:=args[4]
		phoneNo:=args[5]
		role:=args[6]
		userId:=args[7]
		password:=args[8]
		nodeId:=args[9]
		
		// Insert a row
		ok, err := stub.InsertRow("User", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: email}},
				&shim.Column{Value: &shim.Column_String_{String_: country}},
				&shim.Column{Value: &shim.Column_String_{String_: address}},
				&shim.Column{Value: &shim.Column_String_{String_: phoneNo}},
				&shim.Column{Value: &shim.Column_String_{String_: role}},
				&shim.Column{Value: &shim.Column_String_{String_: userId}},
				&shim.Column{Value: &shim.Column_String_{String_: password}},
				&shim.Column{Value: &shim.Column_String_{String_: nodeId}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}

// generate booking number for shipping item
func (t *POC) createBooking(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 11 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 12. Got: %d.", len(args))
		}
		
		userId:=args[0]
		sourcePort:=args[1]
		destinationPort:=args[2]
		Avalbytes, err := stub.GetState("BookingNoincrement") 
		Aval, _ := strconv.ParseInt(string(Avalbytes), 10, 0) 
		newAval:=int(Aval) + 1 
		newBookingNoincrement:= strconv.Itoa(newAval) 
		stub.PutState("BookingNoincrement", []byte(newBookingNoincrement))
		bookingNo:=string(Avalbytes)
		consingnee:=args[3]
		cargoType:=args[4]
		cargoDsec:=args[5]
		cargoId:=args[6]
		containerId:="NA"
		specialReq:=args[7]
		measurements:=args[8]
		weight:=args[9]
		status:="New"
		nodeId:=args[10]
		
		// Insert a row
		ok, err := stub.InsertRow("BookingDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: userId}},
				&shim.Column{Value: &shim.Column_String_{String_: sourcePort}},
				&shim.Column{Value: &shim.Column_String_{String_: destinationPort}},
				&shim.Column{Value: &shim.Column_String_{String_: bookingNo}},
				&shim.Column{Value: &shim.Column_String_{String_: consingnee}},
				&shim.Column{Value: &shim.Column_String_{String_: cargoType}},
				&shim.Column{Value: &shim.Column_String_{String_: cargoDsec}},
				&shim.Column{Value: &shim.Column_String_{String_: cargoId}},
				&shim.Column{Value: &shim.Column_String_{String_: containerId}},
				&shim.Column{Value: &shim.Column_String_{String_: specialReq}},
				&shim.Column{Value: &shim.Column_String_{String_: measurements}},
				&shim.Column{Value: &shim.Column_String_{String_: weight}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: nodeId}},				
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}	

//get all booking details for specified cargo status
func (t *POC) viewBookingDetailsByCargoStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting cargo status to query")
	}

	status := args[0]
	
	var columns []shim.Column

	rows, err := stub.GetRows("BookingDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
			
	res2E:= []*BookingDetails{}	
	
	for row := range rows {		
		newApp:= new(BookingDetails)
		newApp.UserId = row.Columns[0].GetString_()
		newApp.SourcePort = row.Columns[1].GetString_()
		newApp.DestinationPort = row.Columns[2].GetString_()
		newApp.BookingNo = row.Columns[3].GetString_()
		newApp.Consingnee = row.Columns[4].GetString_()
		newApp.CargoType = row.Columns[5].GetString_()
		newApp.CargoDsec = row.Columns[6].GetString_()
		newApp.CargoId = row.Columns[7].GetString_()
		newApp.ContainerId = row.Columns[8].GetString_()
		newApp.SpecialReq = row.Columns[9].GetString_()
		newApp.Measurements = row.Columns[10].GetString_()
		newApp.Weight = row.Columns[11].GetString_()
		newApp.Status = row.Columns[12].GetString_()
		newApp.NodeId = row.Columns[13].GetString_()
		
		if newApp.Status == status{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

//get all booking details for specified cargo id
func (t *POC) viewBookingDetailsByCargoId(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting cargo status to query")
	}

	cargoId := args[0]
	
	var columns []shim.Column

	rows, err := stub.GetRows("BookingDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
			
	res2E:= []*BookingDetails{}	
	
	for row := range rows {		
		newApp:= new(BookingDetails)
		newApp.UserId = row.Columns[0].GetString_()
		newApp.SourcePort = row.Columns[1].GetString_()
		newApp.DestinationPort = row.Columns[2].GetString_()
		newApp.BookingNo = row.Columns[3].GetString_()
		newApp.Consingnee = row.Columns[4].GetString_()
		newApp.CargoType = row.Columns[5].GetString_()
		newApp.CargoDsec = row.Columns[6].GetString_()
		newApp.CargoId = row.Columns[7].GetString_()
		newApp.ContainerId = row.Columns[8].GetString_()
		newApp.SpecialReq = row.Columns[9].GetString_()
		newApp.Measurements = row.Columns[10].GetString_()
		newApp.Weight = row.Columns[11].GetString_()
		newApp.Status = row.Columns[12].GetString_()
		newApp.NodeId = row.Columns[13].GetString_()
		
		if newApp.CargoId == cargoId{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

//get all booking details for specified container id
func (t *POC) viewBookingDetailsByContainerId(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting cargo status to query")
	}

	containerId := args[0]
	
	var columns []shim.Column

	rows, err := stub.GetRows("BookingDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
			
	res2E:= []*BookingDetails{}	
	
	for row := range rows {		
		newApp:= new(BookingDetails)
		newApp.UserId = row.Columns[0].GetString_()
		newApp.SourcePort = row.Columns[1].GetString_()
		newApp.DestinationPort = row.Columns[2].GetString_()
		newApp.BookingNo = row.Columns[3].GetString_()
		newApp.Consingnee = row.Columns[4].GetString_()
		newApp.CargoType = row.Columns[5].GetString_()
		newApp.CargoDsec = row.Columns[6].GetString_()
		newApp.CargoId = row.Columns[7].GetString_()
		newApp.ContainerId = row.Columns[8].GetString_()
		newApp.SpecialReq = row.Columns[9].GetString_()
		newApp.Measurements = row.Columns[10].GetString_()
		newApp.Weight = row.Columns[11].GetString_()
		newApp.Status = row.Columns[12].GetString_()
		newApp.NodeId = row.Columns[13].GetString_()
		
		if newApp.ContainerId == containerId{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

// 
func (t *POC) exportHaulage(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4.")
	}
	newContainerId := args[0]
	cargoIds := args[1]
	newStatus := args[2]
	newNodeId := args[3]

	var cargoIdList []string
	json.Unmarshal([]byte(cargoIds), &cargoIdList)

	for row1 := range cargoIdList {
		cargoId := cargoIdList[row1]
		
		

		var columns []shim.Column

		rows, err := stub.GetRows("BookingDetails", columns)
		if err != nil {
			return nil, fmt.Errorf("Failed to retrieve row")
		}

		for row := range rows {	
		
			tempCargoid := row.Columns[7].GetString_()
			if tempCargoid==cargoId{
			
				bookingNo:=row.Columns[3].GetString_()
				
				// Get the row pertaining to this bookingNo
				var columns1 []shim.Column
				col1 := shim.Column{Value: &shim.Column_String_{String_: bookingNo}}
				columns1 = append(columns1, col1)

				row, err := stub.GetRow("BookingDetails", columns1)
				if err != nil {
					return nil, fmt.Errorf("Error: Failed retrieving data with bookingNo %s. Error %s", bookingNo, err.Error())
				}

				// GetRows returns empty message if key does not exist
				if len(row.Columns) == 0 {
					return nil, nil
				}
			
				//End- Check that the currentStatus to newStatus transition is accurate
				// Delete the row pertaining to this applicationId
				err = stub.DeleteRow(
					"BookingDetails",
					columns1,
				)
				if err != nil {
					return nil, errors.New("Failed deleting row while updating cargo status.")
				}
				userId := row.Columns[0].GetString_()
				sourcePort := row.Columns[1].GetString_()
				destinationPort := row.Columns[2].GetString_()
				bookingNo = row.Columns[3].GetString_()
				consingnee := row.Columns[4].GetString_()
				cargoType := row.Columns[5].GetString_()
				cargoDsec := row.Columns[6].GetString_()
				cargoId = row.Columns[7].GetString_()
				containerId := newContainerId
				specialReq := row.Columns[9].GetString_()
				measurements := row.Columns[10].GetString_()
				weight := row.Columns[11].GetString_()
				status := newStatus
				nodeId := newNodeId
		
					// Insert a row
				ok, err := stub.InsertRow("BookingDetails", shim.Row{
					Columns: []*shim.Column{
						&shim.Column{Value: &shim.Column_String_{String_: userId}},
						&shim.Column{Value: &shim.Column_String_{String_: sourcePort}},
						&shim.Column{Value: &shim.Column_String_{String_: destinationPort}},
						&shim.Column{Value: &shim.Column_String_{String_: bookingNo}},
						&shim.Column{Value: &shim.Column_String_{String_: consingnee}},
						&shim.Column{Value: &shim.Column_String_{String_: cargoType}},
						&shim.Column{Value: &shim.Column_String_{String_: cargoDsec}},
						&shim.Column{Value: &shim.Column_String_{String_: cargoId}},
						&shim.Column{Value: &shim.Column_String_{String_: containerId}},
						&shim.Column{Value: &shim.Column_String_{String_: specialReq}},
						&shim.Column{Value: &shim.Column_String_{String_: measurements}},
						&shim.Column{Value: &shim.Column_String_{String_: weight}},
						&shim.Column{Value: &shim.Column_String_{String_: status}},
						&shim.Column{Value: &shim.Column_String_{String_: nodeId}},				
					}})

				if err != nil {
					return nil, err 
				}
				if !ok && err == nil {
					return nil, errors.New("Failed to insert row while updating cargo status.")
				}
			}
		}
	}
	return nil, nil

}

//update cargo status by cargo id
func (t *POC) updateCargoStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3.")
	}
	newStatus := args[0]
	cargoId := args[1]
	newNodeId := args[2]
	
	var columns []shim.Column

	rows, err := stub.GetRows("BookingDetails", columns)
	if err != nil {
			return nil, fmt.Errorf("Failed to retrieve row")
	}

	for row := range rows {	
		
		tempCargoid := row.Columns[7].GetString_()
		if tempCargoid==cargoId{
			
			bookingNo:=row.Columns[3].GetString_()
	
			// Get the row pertaining to this bookingNo
			var columns1 []shim.Column
			col1 := shim.Column{Value: &shim.Column_String_{String_: bookingNo}}
			columns1 = append(columns1, col1)

			row, err := stub.GetRow("BookingDetails", columns1)
			if err != nil {
				return nil, fmt.Errorf("Error: Failed retrieving data with bookingNo %s. Error %s", bookingNo, err.Error())
			}

			// GetRows returns empty message if key does not exist
			if len(row.Columns) == 0 {
				return nil, nil
			}
			
			//End- Check that the currentStatus to newStatus transition is accurate
			// Delete the row pertaining to this applicationId
			err = stub.DeleteRow(
				"BookingDetails",
				columns1,
			)
			if err != nil {
				return nil, errors.New("Failed deleting row while updating cargo status.")
			}
			userId := row.Columns[0].GetString_()
			sourcePort := row.Columns[1].GetString_()
			destinationPort := row.Columns[2].GetString_()
			bookingNo = row.Columns[3].GetString_()
			consingnee := row.Columns[4].GetString_()
			cargoType := row.Columns[5].GetString_()
			cargoDsec := row.Columns[6].GetString_()
			cargoId = row.Columns[7].GetString_()
			containerId := row.Columns[8].GetString_()
			specialReq := row.Columns[9].GetString_()
			measurements := row.Columns[10].GetString_()
			weight := row.Columns[11].GetString_()
			status := newStatus
			nodeId := newNodeId
			
			// Insert a row
			ok, err := stub.InsertRow("BookingDetails", shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: userId}},
					&shim.Column{Value: &shim.Column_String_{String_: sourcePort}},
					&shim.Column{Value: &shim.Column_String_{String_: destinationPort}},
					&shim.Column{Value: &shim.Column_String_{String_: bookingNo}},
					&shim.Column{Value: &shim.Column_String_{String_: consingnee}},
					&shim.Column{Value: &shim.Column_String_{String_: cargoType}},
					&shim.Column{Value: &shim.Column_String_{String_: cargoDsec}},
					&shim.Column{Value: &shim.Column_String_{String_: cargoId}},
					&shim.Column{Value: &shim.Column_String_{String_: containerId}},
					&shim.Column{Value: &shim.Column_String_{String_: specialReq}},
					&shim.Column{Value: &shim.Column_String_{String_: measurements}},
					&shim.Column{Value: &shim.Column_String_{String_: weight}},
					&shim.Column{Value: &shim.Column_String_{String_: status}},
					&shim.Column{Value: &shim.Column_String_{String_: nodeId}},				
				}})

			if err != nil {
				return nil, err 
			}
			if !ok && err == nil {
				return nil, errors.New("Failed to insert row while updating caro status.")
			}
		}
	}
		return nil, nil

}
// Invoke invokes the chaincode
func (t *POC) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "registerUser" {
		t := POC{}
		return t.registerUser(stub, args)	
	} else if function == "createBooking" { 
		t := POC{}
		return t.createBooking(stub, args)
	} else if function == "exportHaulage" { 
		t := POC{}
		return t.exportHaulage(stub, args)
	} else if function == "updateCargoStatus" { 
		t := POC{}
		return t.updateCargoStatus(stub, args)
	}
	return nil, errors.New("Invalid invoke function name.")

}

// query queries the chaincode
func (t *POC) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "viewBookingDetailsByCargoStatus" {
		t := POC{}
		return t.viewBookingDetailsByCargoStatus(stub, args)		
	} else if function == "viewBookingDetailsByCargoId" { 
		t := POC{}
		return t.viewBookingDetailsByCargoId(stub, args)
	}else if function == "viewBookingDetailsByContainerId" { 
		t := POC{}
		return t.viewBookingDetailsByContainerId(stub, args)
	}	
	return nil, errors.New("Invalid query function name.")
}

func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(POC))
	if err != nil {
		fmt.Printf("Error starting POC: %s", err)
	}
} 

