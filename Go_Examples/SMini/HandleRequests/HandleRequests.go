package HandleRequests

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	ptime "github.com/yaa110/go-persian-calendar"
	"log"
	"net/http"
)
var db *gorm.DB
var err error

//{"day":"15","new_plan":"true","emergency":"true","reset":"false"}
type NewPlanStruct struct {
	gorm.Model
	RelayID   string
	Day       int `json:"day" xml:"day" form:"day" query:"day"`
	NewPlan   string `json:"new_plan" xml:"new_plan" form:"new_plan" query:"new_plan"`
	Emergency string `json:"emergency" xml:"emergency" form:"emergency" query:"emergency"`
	Reset     string `json:"reset" xml:"reset" form:"reset" query:"reset"`
}
type NewPlanJson struct {
	Day       int `json:"day" xml:"day" form:"day" query:"day"`
	NewPlan   string `json:"new_plan" xml:"new_plan" form:"new_plan" query:"new_plan"`
	Emergency string `json:"emergency" xml:"emergency" form:"emergency" query:"emergency"`
	Reset     string `json:"reset" xml:"reset" form:"reset" query:"reset"`
}
//{"15":[{"startHour":"22","startMinute":"4","endHour":"23","endMinute":"20"}]}
type CheckNewData struct {
	gorm.Model
	RelayID  	string
	Day         string `json:"day" xml:"day" form:"day" query:"day"`
	StartHour   string `json:"startHour" xml:"startHour" form:"startHour" query:"startHour"`
	StartMinute string `json:"startMinute" xml:"startMinute" form:"startMinute" query:"startMinute"`
	EndHour     string `json:"endHour" xml:"endHour" form:"endHour" query:"endHour"`
	EndMinute   string `json:"endMinute" xml:"endMinute" form:"endMinute" query:"endMinute"`
}
type CheckNewDataDay struct {
	StartHour   string `json:"startHour" xml:"startHour" form:"startHour" query:"startHour"`
	StartMinute string `json:"startMinute" xml:"startMinute" form:"startMinute" query:"startMinute"`
	EndHour     string `json:"endHour" xml:"endHour" form:"endHour" query:"endHour"`
	EndMinute   string `json:"endMinute" xml:"endMinute" form:"endMinute" query:"endMinute"`
}

func HandlePlanStatus(w http.ResponseWriter, r * http.Request){
	db, err = gorm.Open("mysql", "root:@/SMini?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("All Users")
		panic(err.Error())
	}
	defer db.Close()

	vars := mux.Vars(r)
	ID := vars["ID"]

	Plan := NewPlanStruct{}
	db.Where("relay_id = ?", ID).Find(&Plan)
	Response := NewPlanJson{
		Day:       Plan.Day,
		NewPlan:   Plan.NewPlan,
		Emergency: Plan.Emergency,
		Reset:     Plan.Reset,
	}

	NewPlanJson, err := json.Marshal(Response)
	if err != nil {
		log.Fatal("Error Marshaling NewPlan",err.Error())
	}

	_,err = w.Write(NewPlanJson)
	if err != nil {
		log.Fatal("Error Marshaling NewPlan",err.Error())
	}
}

func HandleTimeStatus(w http.ResponseWriter, r * http.Request){
	db, err = gorm.Open("mysql", "root:@/SMini?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Check_New_Data")
		panic(err.Error())
	}
	defer db.Close()

	vars := mux.Vars(r)
	ID := vars["ID"]

	gg := make(map[string][]CheckNewDataDay)

	var Plan []CheckNewData
	db.Where("relay_id = ?", ID).Find(&Plan)
	for i := range Plan {
		ff := CheckNewDataDay{
			StartHour:   Plan[i].StartHour,
			StartMinute: Plan[i].StartMinute,
			EndHour:     Plan[i].EndHour,
			EndMinute:   Plan[i].EndMinute,
		}

		gg[Plan[i].Day] = append(gg[Plan[i].Day],ff)
	}

	NewPlanJson, err := json.Marshal(gg)
	if err != nil {
		log.Fatal("Error Marshaling NewPlan",err.Error())
	}

	_,err = w.Write(NewPlanJson)
	if err != nil {
		log.Fatal("Error Marshaling NewPlan",err.Error())
	}
}

func DeleteTimeStatus(w http.ResponseWriter, r * http.Request){
	db, err = gorm.Open("mysql", "root:@/SMini?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Check_New_Data")
		panic(err.Error())
	}
	defer db.Close()

	vars := mux.Vars(r)
	Status := vars["ID"]
	Day := vars["day"]
	if Status == "all"{
		db.Delete(&CheckNewData{})
		_, _ = w.Write([]byte("All Datas deleted"))
	}  else if Day == "all" {
		db.Where("relay_id = ?",Status).Delete(&CheckNewData{})
		_, _ = w.Write([]byte("All Data From Relay is deleted"))
	} else {
		db.Where("day = ? AND relay_id = ?", Day, Status).Delete(&CheckNewData{})
		_, _ = w.Write([]byte("day Data From Relay is deleted"))
	}
}

func UpdateTimeStatus(w http.ResponseWriter, r * http.Request){
	db, err = gorm.Open("mysql", "root:@/SMini?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Check_New_Data")
		panic(err.Error())
	}
	defer db.Close()

	vars := mux.Vars(r)
	Status := vars["ID"]
	Day := vars["day"]
	StartHour := vars["StartHour"]
	StartMinute := vars["StartMinute"]
	EndHour := vars["EndHour"]
	EndMinute := vars["EndMinute"]
	Update := CheckNewData{
		StartHour:   StartHour,
		StartMinute: StartMinute,
		EndHour:     EndHour,
		EndMinute:   EndMinute,
	}

	db.Model(&CheckNewData{}).Where("relay_id = ? AND day = ?", Status, Day).Updates(Update)
	_, _ = w.Write([]byte("Relay Time Updated Successfully"))


}

func UpdatePlanStatus(w http.ResponseWriter, r * http.Request){
	db, err = gorm.Open("mysql", "root:@/SMini?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Check_New_Data")
		panic(err.Error())
	}
	defer db.Close()

	vars := mux.Vars(r)
	RelayID := vars["ID"]
	Emergency := vars["emergency"]
	Reset := vars["reset"]

	NewPlan := NewPlanStruct{
		Model:     gorm.Model{},
		Emergency: Emergency,
		Reset:     Reset,
	}

	db.Model(&NewPlanStruct{}).Where("relay_id = ?",RelayID).Updates(NewPlan)
	_, _ = w.Write([]byte("Relay Plan Status Updated Successfully"))
}

func CreateTimeStatus(w http.ResponseWriter, r * http.Request){
	db, err = gorm.Open("mysql", "root:@/SMini?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Check_New_Data")
		panic(err.Error())
	}
	defer db.Close()

	vars := mux.Vars(r)
	Status := vars["ID"]
	Day := vars["day"]
	StartHour := vars["StartHour"]
	StartMinute := vars["StartMinute"]
	EndHour := vars["EndHour"]
	EndMinute := vars["EndMinute"]
	Update := CheckNewData{
		RelayID:	 Status,
		Day:		 Day,
		StartHour:   StartHour,
		StartMinute: StartMinute,
		EndHour:     EndHour,
		EndMinute:   EndMinute,
	}
	db.Create(&Update)
	_, _ = w.Write([]byte("Relay New Time Status Created Successfully"))
}

func GetDay(){
	pt := ptime.Now(ptime.Iran())
	var day int

	if pt.Day() != day {
		db, err = gorm.Open("mysql", "root:@/SMini?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		day = pt.Day()
	}
	NewPlan := NewPlanStruct{
		Model:     gorm.Model{},
		Day:       day,
	}
	for i:=1;i<9;i++ {
		db.Model(&NewPlanStruct{}).Where("relay_id = ?",i).Updates(NewPlan)
	}

}


func SqlMigration() {
	db, err = gorm.Open("mysql", "root:@/SMini?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("SQL Migration")
		panic(err.Error())
	}
	defer db.Close()

	db.AutoMigrate(&CheckNewData{})
	db.AutoMigrate(&NewPlanStruct{})
}