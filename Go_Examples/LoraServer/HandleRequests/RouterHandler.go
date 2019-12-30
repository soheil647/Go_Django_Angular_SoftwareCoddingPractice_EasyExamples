package HandleRequests

import (
	"bufio"
	"bytes"
	"crypto/tls"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//Gorm db Format Save
type GetFormat struct {
	gorm.Model
	ModuleID 		string `json:"ModuleID" xml:"ModuleID" form:"ModuleID" query:"ModuleID"`
	Adr             bool `json:"adr" xml:"adr" form:"adr" query:"adr"`
	Acknowledged    bool `json:"acknowledged" xml:"Acknowledged" form:"Acknowledged" query:"Acknowledged"`
	Data            string `json:"data" xml:"data" form:"data" query:"data"`
	FCnt            int `json:"fCnt" xml:"fCnt" form:"fCnt" query:"fCnt"`
	GatewayID       string `json:"gatewayID" xml:"gatewayID" form:"gatewayID" query:"gatewayID"`
	LoraSNR         float32 `json:"loraSNR" xml:"loraSNR" form:"loraSNR" query:"loraSNR"`
	Rssi            int `json:"rssi" xml:"rssi" form:"rssi" query:"rssi"`
	TimeNow         string `json:"timeNow" xml:"timeNow" form:"timeNow" query:"timeNow"`
	Type            string `json:"type" xml:"type" form:"type" query:"type"`
}

//For Parsing Response
type Result struct {
	Result Response `json:"result" xml:"result" form:"result" query:"result"`
}
type Response struct {
	Type        string `json:"type" xml:"type" form:"type" query:"type"`
	PayloadJSON string `json:"payloadJSON" xml:"payloadJSON" form:"payloadJSON" query:"payloadJSON"`
}
type RxInfo struct {
	GatewayID       string   `json:"gatewayID" xml:"gatewayID" form:"gatewayID" query:"gatewayID"`
	LoraSNR         float32   `json:"loraSNR" xml:"loraSNR" form:"loraSNR" query:"loraSNR"`
	Location		map[string]float32 `json:"location" xml:"location" form:"location" query:"location"`
	Name            string   `json:"name" xml:"name" form:"name" query:"name"`
	Rssi            int 	 `json:"rssi" xml:"rssi" form:"rssi" query:"rssi"`
	TimeNow         string   `json:"time" xml:"time" form:"time" query:"time"`
}
type Uplink struct {
	Adr             bool `json:"adr" xml:"adr" form:"adr" query:"adr"`
	Acknowledged    bool `json:"acknowledged" xml:"Acknowledged" form:"Acknowledged" query:"Acknowledged"`
	Data            string `json:"data" xml:"data" form:"data" query:"data"`
	FCnt            int    `json:"fCnt" xml:"fCnt" form:"fCnt" query:"fCnt"`
	RxInfo 			[]RxInfo `json:"rxInfo" xml:"rxInfo" form:"rxInfo" query:"rxInfo"`
	TxInfo 			map[string]int `json:"txInfo" xml:"txInfo" form:"txInfo" query:"txInfo"`
}
type Ack struct {
	Acknowledged    bool `json:"acknowledged" xml:"Acknowledged" form:"Acknowledged" query:"Acknowledged"`
	FCnt            int    `json:"fCnt" xml:"fCnt" form:"fCnt" query:"fCnt"`
}

//For Sending DAta
type SendFormat struct {
	gorm.Model
	Confirmed  bool   `json:"confirmed"`
	Data       string `json:"data"`
	DevEui     string `json:"dev_eui"`
	FCnt       int    `json:"f_cnt"`
	FPort      int    `json:"f_port"`
	JsonObject string `json:"json_object"`
}

//Jwt For Login
type jwt struct {
	Jwt string `json:"jwt"`
}

//For Errors
type HandleGetError struct {
	Error GetError `json:"error" xml:"error" form:"error" query:"error"`
}
type GetError struct {
	gorm.Model
	GrpcCode   int    `json:"grpcCode" xml:"grpcCode" form:"grpcCode" query:"grpcCode"`
	HttpCode   int    `json:"httpCode" xml:"httpCode" form:"httpCode" query:"httpCode"`
	Message    string `json:"message" xml:"message" form:"message" query:"message"`
	HttpStatus string `json:"httpStatus" xml:"httpStatus" form:"httpStatus" query:"httpStatus"`
}

//Ides for each module
type ModuleID struct {
	gorm.Model
	ModuleID 		string `json:"ModuleID"`
	ApplicationID   string `json:"applicationID" xml:"applicationID" form:"applicationID" query:"applicationID"`
	ApplicationName string `json:"applicationName" xml:"applicationName" form:"applicationName" query:"applicationName"`
	DevEUI          string `json:"devEUI" xml:"devEUI" form:"devEUI" query:"devEUI"`
	DeviceName      string `json:"deviceName" xml:"deviceName" form:"deviceName" query:"deviceName"`
	FPort           int    `json:"fPort" xml:"fPort" form:"fPort" query:"fPort"`
	Dr              int `json:"dr" xml:"dr" form:"dr" query:"dr"`
	Frequency       int `json:"frequency" xml:"frequency" form:"frequency" query:"frequency"`
}
type GatewayID struct {
	GatewayID       string   `json:"gatewayID" xml:"gatewayID" form:"gatewayID" query:"gatewayID"`
	Altitude        float32 `json:"altitude" xml:"altitude" form:"altitude" query:"altitude"`
	Latitude        float32 `json:"latitude" xml:"latitude" form:"latitude" query:"latitude"`
	Longitude       float32 `json:"longitude" xml:"longitude" form:"longitude" query:"longitude"`
	Name            string `json:"name" xml:"name" form:"name" query:"name"`
}

func LoginRequest() string {
	LoginInfo, err := json.Marshal(map[string]string{
		"username": "orgadmin1",
		"password": "orgadmin1",
	})
	if err != nil {
		fmt.Println("LoginRequest First")
		log.Fatal(err.Error())
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://io.lpwandata.com:8080/api/internal/login", bytes.NewBuffer(LoginInfo))
	if err != nil {
		log.Println("LoginRequest Second")
		log.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Its Important i need to reLogin")
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("LoginRequest Third")
		log.Fatal(err.Error())
	}
	jwt := jwt{}
	err = json.Unmarshal(body, &jwt)
	if err != nil {
		fmt.Println("OMG Cool")
		log.Fatal(err.Error())
	}
	fmt.Println(jwt.Jwt)
	return jwt.Jwt
}



//Get data for each Module
func GetData(jwt string, moduleID string) string {
	client := &http.Client{}

	//Bypass SSL certificate
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	//Get Request For ModuleID
	req, err := http.NewRequest("GET", "https://io.lpwandata.com:8080/api/devices/"+moduleID+"/events", nil)
	if err != nil {
		log.Fatal("Error Creating Request: ", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Grpc-Metadata-Authorization", " Bearer "+jwt)
	//Do The request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error Sending Get Request: ", err.Error())
	}
	defer resp.Body.Close()

	//Reading Response Body
	reader := bufio.NewReader(resp.Body)
	body, err := reader.ReadBytes('\n')
	if err != nil {
		log.Println("Error Reading Response Body: ", err.Error())
	}
	fmt.Println("Resp Body:  ", string(body))

	return string(body)
}

//Unmarshal data or error
func UnmarshalGetData(ModuleID string,body string, db *gorm.DB) string {
	var SaveGetData GetFormat

	decoder := json.NewDecoder(strings.NewReader(body))
	var decodedRequest Result
	var decodedError HandleGetError

	err := decoder.Decode(&decodedRequest)
	if err != nil {
		log.Println("Getting Error: ",err.Error())
	}
	if decodedRequest.Result.Type != "" {
		switch decodedRequest.Result.Type {
		case "uplink":
			var uplinkObject Uplink
			if err := json.Unmarshal([]byte(decodedRequest.Result.PayloadJSON), &uplinkObject); err != nil {
				log.Println(err.Error())
			}
			SaveGetData = GetFormat{
				Model:        gorm.Model{},
				ModuleID:     ModuleID,
				Adr:          uplinkObject.Adr,
				Acknowledged: uplinkObject.Acknowledged,
				Data:         uplinkObject.Data,
				FCnt:         uplinkObject.FCnt,
				GatewayID:    uplinkObject.RxInfo[0].GatewayID,
				LoraSNR:      uplinkObject.RxInfo[0].LoraSNR,
				Rssi:         uplinkObject.RxInfo[0].Rssi,
				TimeNow:      uplinkObject.RxInfo[0].TimeNow,
				Type:         decodedRequest.Result.Type,
			}
			db.Create(&SaveGetData)
			return uplinkObject.Data
		case "join":
			SaveGetData = GetFormat{
				Model:    gorm.Model{},
				ModuleID: ModuleID,
				Type:     decodedRequest.Result.Type,
			}
			db.Create(&SaveGetData)
			return ""
		case "ack":
			var AckObject Ack
			if err := json.Unmarshal([]byte(decodedRequest.Result.PayloadJSON), &AckObject); err != nil {
				log.Println(err.Error())
			}
			SaveGetData = GetFormat{
				Model:        gorm.Model{},
				ModuleID:     ModuleID,
				Acknowledged: AckObject.Acknowledged,
				FCnt:         AckObject.FCnt,
				Type:         decodedRequest.Result.Type,
			}
			db.Create(&SaveGetData)
			return ""
		}
	}

	err = decoder.Decode(&decodedError)
	if err != nil {
		log.Println(err.Error())
	}
	if decodedError.Error.Message != "" {
		ErrorIS := GetError{
			GrpcCode:   decodedError.Error.GrpcCode,
			HttpCode:   decodedError.Error.HttpCode,
			Message:    decodedError.Error.Message,
			HttpStatus: decodedError.Error.HttpStatus,
		}
		db.Create(&ErrorIS)
	}
	return ""
}


func SendTest(moduleID string, jwt string, data string){

	SendDataTest := "{ \n" +
		"   \"device_queue_item\": {  \n" +
		"     \"confirmed\": true,  \n" +
		"     \"data\": \"" + data + "\",  \n" +
		"     \"dev_eui\": \"323631395d365505\", \n" +
		"     \"f_cnt\": 0,  \n" +
		"     \"f_port\": 10,  \n" +
		"     \"json_object\": \"\"  \n" +
		"   } \n" +
		" }"



	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://io.lpwandata.com:8080/api/devices/"+moduleID+"/queue", bytes.NewBuffer([]byte(SendDataTest)))
	if err != nil {
		log.Println("SendData Second")
		log.Fatal(err.Error())
	}
	req.Header.Set("Content-Type","application/json")
	req.Header.Add("Accept","application/json")
	req.Header.Add("Grpc-Metadata-Authorization",jwt)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("SendData Third")
		log.Fatal(err.Error())
	}
	fmt.Println("MYTest Response: ",resp)
	req.Body.Close()

	myBody, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(myBody))
	}


func HandleGet(jwt string,moduleID string,db *gorm.DB){
	body := GetData(jwt, moduleID)
	data := UnmarshalGetData(moduleID,body, db)
	if moduleID == "323631395d365505" {
		Temperature := UnmarshalGetData(moduleID,body,db)
		fmt.Println("This is Temperature", Temperature)
		if Temperature != "0000" {
		i1, _ := b64.StdEncoding.DecodeString(Temperature)
			fmt.Println("this is i1: ", i1)
			fmt.Println("this is string i1: ", string(i1))

			CmpTemp, err := strconv.Atoi(string(i1))
			if err != nil {
				fmt.Println("MY Atoi Err:", err)
			}
			fmt.Println("This is CmpTemp:", CmpTemp)
			slice := strconv.Itoa(CmpTemp)
			fmt.Println(slice[:2])
			fmt.Println(slice[2:])
			ss, err := strconv.Atoi(slice[2:])
			if err != nil {
				fmt.Println("MY Atoi2 Err:", err)
			}
			fmt.Println("this is ss: ", ss)
			if ss >= 40 {
					SendTest("3938383360377f04", jwt, "MTAx")

				fmt.Println("DAta Send1: ")
			} else {
					SendTest("3938383360377f04", jwt, "MTEx")
				fmt.Println("DAta Send2: ")
			}
		}
	}
	fmt.Println(data)
}



func SqlMigration(url string) *gorm.DB {
	var db *gorm.DB
	var err error

	db, err = gorm.Open("mysql", url)

	if err != nil {
		log.Println("SQL Migration")
		panic(err.Error())
	}
	db.AutoMigrate(&GetFormat{})
	db.AutoMigrate(&ModuleID{})
	db.AutoMigrate(&GatewayID{})
	db.AutoMigrate(&GetError{})

	//Create New Gateway
	/*gateway1 := GatewayID{
		GatewayID: "689e19fffe999470",
		Altitude:  1250,
		Latitude:  35.7024,
		Longitude: 51.4059,
		Name:      "lpwan-gw-0003",
	}
	db.Create(&gateway1)*/
	
	//Create New Module
/*	module1 := ModuleID{
		Model:           gorm.Model{},
		ModuleID:        "3938383360377f04",
		ApplicationID:   "8",
		ApplicationName: "agri",
		DevEUI:          "3938383360377f04",
		DeviceName:      "modem-1",
		FPort:           2,
		Dr:               5,
		Frequency:		 868100000,
	}
	db.Create(&module1)*/
	
	return db
}

