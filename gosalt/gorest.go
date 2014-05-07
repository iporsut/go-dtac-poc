package main

import (
	. "../logger"
	. "./services"
//--close_log	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
//--close_log	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Message struct {
	Msg string
	Err string
}

type SaveSpkdAddWSInput struct {
	USER_CODE              string
	BLPD_INDC              string
	CS_SPKD_PCN__CUST_NUMB string
	CS_SPKD_PCN__SUBR_NUMB string
	CS_SPKD_PCN__PACK_CODE string
	CS_SUBR_PCN__SUBR_TYPE string
	RD_TELP__TELP_TYPE     string
	SAVE_FLAG              string
}

type SaveSpkdAddWSoutput struct {
	Status                       string
	CS_SPKD_PCN__PACK_CODE       string
	CS_PKPL_PCN__PACK_DESC       string
	CS_PACK_TYPE__PACK_TYPE_DESC string
	CS_SPKD_PCN__PACK_STRT_DTTM  string
	CS_SPKD_PCN__PACK_END_DTTM   string
	CS_SPKD_PCN__DISC_CODE       string
	TBL_OCCR                     string
}

type SaveSpkdAddWSFaultoutput struct {
	Status      string
	Faultcode   string
	Faultstring string
}

type Activities struct {
	Contact_Direction                string
	Master_Contact_Header            string
	Parent_Contact_Header            string
	Activity_Number                  string
	Activity_Type                    string
	Contact_Channel                  string
	Contact_Sub_Channel              string
	Contact_Location                 string
	Contact_Status                   string
	Contact_Sub_Status               string
	ContactNum_or_Email              string
	Con_ID                           string
	QConn_ID                         string
	E_Service_Ref_Number             string
	SFA_Ref_Number                   string
	Survey_Score                     string
	Contact_Type                     string
	Contact_Sub_Type                 string
	Contact_Category                 string
	Contact_Sub_Category             string
	Activiry_Created_by_Employee     string
	Activity_Created_by_Position     string
	Activity_Created_by_Division     string
	ACS_Agent                        string
	Account_Number                   string
	Asset_Number                     string
	Account_Name                     string
	Handset_Brand                    string
	Handset_Model                    string
	Telephone_Type                   string
	Account_Type                     string
	Identification_Number            string
	Group_Code                       string
	Business                         string
	Sub_Business                     string
	Product                          string
	Contact_Item_Start_Date_Time     string
	Contact_Item_End_Date_Time       string
	Contact_Item_Status              string
	Contact_Duration                 string
	Parent_Start_Date                string
	Parent_End_Date                  string
	Parent_ActualEnd                 string
	Parent_Duration                  string
	Parent_ActualDuration            string
	Employee_ID                      string
	Parent_Telephone_Type_Contact_No string
	Parent_Contact_Location          string
	Description                      string
	Channel_Group                    string
	Partner_Code                     string
	Partner_Name                     string
	Related_SRTT                     string
	SRTT_Number                      string
	Company_Name                     string
}

type MongoPersister struct {
}

var logger Logger

func main() {
	//num_cores := 8
	//for i := 0; i < num_cores; i++ {
	//	go func() {

	file, err := os.OpenFile("./poc.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", os.Stdout, ":", err)
	}

	multi := io.MultiWriter(file, os.Stdout)

	logger.Init(ioutil.Discard, multi, os.Stdout, os.Stderr) // multi/file

/*	handler := rest.ResourceHandler{
		PreRoutingMiddlewares: []rest.Middleware{
			&rest.CorsMiddleware{
				RejectNonCorsRequests: false,
				//OriginValidator: func(origin string, request *rest.Request) bool {
				//	return origin == "http://10.89.104.16"
				//},
				AllowedMethods:                []string{"GET", "POST", "PUT"},
				AllowedHeaders:                []string{"Accept", "Content-Type", "X-Custom-Header"},
				AccessControlAllowCredentials: true,
				AccessControlMaxAge:           3600,
			},
		},
	}
*/
	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
	}
	handler.SetRoutes(
		&rest.Route{"POST", "/post", SaveSpkdAddWSService},
		&rest.Route{"OPTIONS", "/post", SaveSpkdAddWSService},
	)
	http.ListenAndServe(":11080", &handler)
	//http.ListenAndServeTLS(":9443","/usr/app/go/ssl/pocnccaapp3/self-ssl.pem","/usr/app/go/ssl/pocnccaapp3/self-ssl.key", &handler)	
	//}()
	//}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func (mongoPersister *MongoPersister) InsertPerson(saltOutput *Activities) {
	session, err := mgo.Dial("10.89.104.10")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(saltOutput)

	if err != nil {
		panic(err)
	}
}

func SaveActivities() {
	saveActivities := make([]Activities, 3)
	saveActivities[0] = Activities{Contact_Direction: "Inbound", Master_Contact_Header: "1-9913205841", Parent_Contact_Header: "1-9913205841", Activity_Number: "1-9913205851", Activity_Type: "Inbound", Contact_Channel: "Call", Contact_Sub_Channel: "1678 #0", Contact_Location: "", Contact_Status: "Done", Contact_Sub_Status: "", ContactNum_or_Email: "66950643823", Con_ID: "0176023aa430a843", QConn_ID: "", E_Service_Ref_Number: "", SFA_Ref_Number: "", Survey_Score: "", Contact_Type: "Request Package", Contact_Sub_Type: "Package Prepaid", Contact_Category: "สมัคร แฮปปี้บุฟเฟ่ต์ กลางคืน", Contact_Sub_Category: "N/A", Activiry_Created_by_Employee: "OS6068", Activity_Created_by_Position: "CustCallServ_CMI_Prepaid1_Staff", Activity_Created_by_Division: "CMI CC Prepaid Service 1 Team", ACS_Agent: "SADMIN", Account_Number: "1042146551", Asset_Number: "66950643823", Account_Name: "Unregistered Prepaid", Handset_Brand: "NOKIA", Handset_Model: "LUMIA 520", Telephone_Type: "Prepaid", Account_Type: "Dealer", Identification_Number: "", Group_Code: "", Business: "", Sub_Business: "", Product: "", Contact_Item_Start_Date_Time: "41723.0000231481", Contact_Item_End_Date_Time: "41723.5419560185", Contact_Item_Status: "Done", Contact_Duration: "46823", Parent_Start_Date: "41722.9985069444", Parent_End_Date: "41722.9998726852", Parent_ActualEnd: "41723.5419560185", Parent_Duration: "118", Parent_ActualDuration: "46954", Employee_ID: "OS6068", Parent_Telephone_Type_Contact_No: "Prepaid", Parent_Contact_Location: "Call Center Chiang Mai Department", Description: "", Channel_Group: "", Partner_Code: "", Partner_Name: "", Related_SRTT: "", SRTT_Number: "", Company_Name: "TriNet"}
	saveActivities[1] = Activities{Contact_Direction: "Inbound", Master_Contact_Header: "1-9913296891", Parent_Contact_Header: "1-9913296891", Activity_Number: "1-9913296897", Activity_Type: "Inbound", Contact_Channel: "Walk", Contact_Sub_Channel: "SMS 16789", Contact_Location: "", Contact_Status: "Done", Contact_Sub_Status: "", ContactNum_or_Email: "", Con_ID: "", QConn_ID: "", E_Service_Ref_Number: "", SFA_Ref_Number: "", Survey_Score: "", Contact_Type: "Enquiry VAS Package", Contact_Sub_Type: "Additional Package Happy Internet", Contact_Category: "สอบถามโปรโมชั่นเสริม Internet อื่นๆ", Contact_Sub_Category: "N/A", Activiry_Created_by_Employee: "NATAWATW", Activity_Created_by_Position: "CustCallServ_RGS_OutbServ2_Sup", Activity_Created_by_Division: "RGS CC Outbound Service 2 Team", ACS_Agent: "SADMIN", Account_Number: "606212347", Asset_Number: "66900353982", Account_Name: "SMART VAN PROJECT (FREE SIM PROJECT)", Handset_Brand: "SAMSUNG", Handset_Model: "S5830 GALAXY ACE", Telephone_Type: "Prepaid", Account_Type: "Individual", Identification_Number: "URPP0215", Group_Code: "", Business: "", Sub_Business: "", Product: "", Contact_Item_Start_Date_Time: "41723.000150463", Contact_Item_End_Date_Time: "41723.0001851852", Contact_Item_Status: "Done", Contact_Duration: "3", Parent_Start_Date: "41722.9999884259", Parent_End_Date: "41723.0001851852", Parent_ActualEnd: "41723.0001851852", Parent_Duration: "17", Parent_ActualDuration: "17", Employee_ID: "5927", Parent_Telephone_Type_Contact_No: "Prepaid", Parent_Contact_Location: "Call Center Rangsit Department", Description: "", Channel_Group: "", Partner_Code: "", Partner_Name: "", Related_SRTT: "", SRTT_Number: "", Company_Name: "TriNet"}
	saveActivities[2] = Activities{Contact_Direction: "Inbound", Master_Contact_Header: "1-9913315459", Parent_Contact_Header: "1-9913315459", Activity_Number: "1-9913315465", Activity_Type: "Inbound", Contact_Channel: "Call", Contact_Sub_Channel: "", Contact_Location: "", Contact_Status: "Done", Contact_Sub_Status: "", ContactNum_or_Email: "", Con_ID: "", QConn_ID: "", E_Service_Ref_Number: "", SFA_Ref_Number: "", Survey_Score: "", Contact_Type: "Request VAS", Contact_Sub_Type: "ขอเปิด-ปิดสัญญาณ GPRS/MMS/SMS", Contact_Category: "ขอเปิด-ปิด สัญญาณ GPRS/MMS (โดยลูกค้า)", Contact_Sub_Category: "N/A", Activiry_Created_by_Employee: "OS6063", Activity_Created_by_Position: "CustCallServ_CMI_Prepaid1_Staff", Activity_Created_by_Division: "CMI CC Prepaid Service 1 Team", ACS_Agent: "SADMIN", Account_Number: "602408688", Asset_Number: "66894392439", Account_Name: "นาง ธัญวลัย เปี่ยมดำรงศักดิ์", Handset_Brand: "LG", Handset_Model: "D802 G2", Telephone_Type: "Postpaid", Account_Type: "Individual", Identification_Number: "3669800043020", Group_Code: "", Business: "", Sub_Business: "", Product: "", Contact_Item_Start_Date_Time: "41723.0002083333", Contact_Item_End_Date_Time: "41723.0002314815", Contact_Item_Status: "Done", Contact_Duration: "2", Parent_Start_Date: "41722.9989583333", Parent_End_Date: "41723.0002314815", Parent_ActualEnd: "41723.0002314815", Parent_Duration: "110", Parent_ActualDuration: "110", Employee_ID: "OS6063", Parent_Telephone_Type_Contact_No: "Postpaid", Parent_Contact_Location: "Call Center Chiang Mai Department", Description: "เปิดสัญญาณเน็ตให้ ลค.23.59.37 น.", Channel_Group: "", Partner_Code: "", Partner_Name: "", Related_SRTT: "", SRTT_Number: "", Company_Name: "TriNet"}

	saveActivitie := saveActivities[random(0, 2)]

	mongoPersister := MongoPersister{}

	mongoPersister.InsertPerson(&saveActivitie)
}

func SaveSpkdAddWSService(w rest.ResponseWriter, r *rest.Request) {

	input := SaveSpkdAddWSInput{}
	err := r.DecodeJsonPayload(&input)

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		println(err.Error())
		return
	}
	if input.USER_CODE == "" {
		rest.Error(w, "USER_CODE required", 400)
		println("USER_CODE required")
		return
	}
	if input.BLPD_INDC == "" {
		rest.Error(w, "BLPD_INDC required", 400)
		println("BLPD_INDC required")
		return
	}

//--close_log	subr_numb:=input.CS_SPKD_PCN__SUBR_NUMB;

//--close_log	go logger.INFO.Println(subr_numb+" I'm gonna save into mongodb.")
	go SaveActivities()
//--close_log	go logger.INFO.Println(subr_numb+" I've saved.")

	var saveSpkdAddWS = SaveSpkdAddWS{USER_CODE: input.USER_CODE, BLPD_INDC: input.BLPD_INDC, CS_SPKD_PCN__CUST_NUMB: input.CS_SPKD_PCN__CUST_NUMB, CS_SPKD_PCN__SUBR_NUMB: input.CS_SPKD_PCN__SUBR_NUMB, CS_SPKD_PCN__PACK_CODE: input.CS_SPKD_PCN__PACK_CODE, RD_TELP__TELP_TYPE: input.RD_TELP__TELP_TYPE, SAVE_FLAG: input.SAVE_FLAG}

	var bodyElement *SaveSpkdAddWSResponse
	var faultElement *FaultResponse
	var tuxerr error
	var state string

//--close_log	go logger.INFO.Println(subr_numb+" I'm gonna call tuxedo.")
	bodyElement, faultElement, tuxerr, state = CallTux(saveSpkdAddWS)
//--close_log	go logger.INFO.Println(subr_numb+" I've called tuxedo.")

	if tuxerr != nil {
		println(tuxerr.Error())
		w.WriteJson(&SaveSpkdAddWSFaultoutput{
			Status:      "fail",
			Faultcode:   "tux error",
			Faultstring: "tux error",
		})

	}
	if faultElement != nil {
		println("fault error")

		w.WriteJson(&SaveSpkdAddWSFaultoutput{
			Status:      "fail",
			Faultcode:   faultElement.Faultcode,
			Faultstring: faultElement.Detail.SaveSpkdAddWSFault.Errbuf.MESSAGE_TEXT_THA,
		})
//--close_log		go logger.INFO.Println(subr_numb+" TUX_RESULT(fail):")
	}
	if state == "body" {
		println("body state")
	} else {
		println(state)
	}
	if bodyElement == nil {
		println("no body")
	} else {
		var saveSpkdAddWSoutput = SaveSpkdAddWSoutput{
			Status:                       "success",
			CS_SPKD_PCN__PACK_CODE:       bodyElement.CS_SPKD_PCN__PACK_CODE,
			CS_PKPL_PCN__PACK_DESC:       bodyElement.CS_PKPL_PCN__PACK_DESC,
			CS_PACK_TYPE__PACK_TYPE_DESC: bodyElement.CS_PACK_TYPE__PACK_TYPE_DESC,
			CS_SPKD_PCN__PACK_STRT_DTTM:  bodyElement.CS_SPKD_PCN__PACK_STRT_DTTM,
			CS_SPKD_PCN__PACK_END_DTTM:   bodyElement.CS_SPKD_PCN__PACK_END_DTTM,
			CS_SPKD_PCN__DISC_CODE:       bodyElement.CS_SPKD_PCN__DISC_CODE,
			TBL_OCCR:                     bodyElement.TBL_OCCR,
		}
		w.WriteJson(&saveSpkdAddWSoutput)

/*--close_log		b, err := json.Marshal(saveSpkdAddWSoutput)
		if err != nil {
			fmt.Println("error:", err)
		}
		go logger.INFO.Println(subr_numb+" TUX_RESULT(success):", string(b[:]))
--close_log */
	}

}

//func PostParameter(w rest.ResponseWriter, r *rest.Request) {
//	fmt.Printf("HEADER: %s\n", r.Header)
//	input := SaveSpkdAddWSInput{}
//	err := r.DecodeJsonPayload(&input)

//	if err != nil {
//		rest.Error(w, err.Error(), http.StatusInternalServerError)
//		println(err.Error())
//		return
//	}
//	if input.USER_CODE == "" {
//		rest.Error(w, "USER_CODE required", 400)
//		println("USER_CODE required")
//		return
//	}
//	if input.BLPD_INDC == "" {
//		rest.Error(w, "BLPD_INDC required", 400)
//		println("BLPD_INDC required")
//		return
//	}

//	w.WriteJson(&SaveSpkdAddWSoutput{
//		CS_SPKD_PCN__PACK_CODE:       "testa",
//		CS_PKPL_PCN__PACK_DESC:       "testb",
//		CS_PACK_TYPE__PACK_TYPE_DESC: "testc",
//		CS_SPKD_PCN__PACK_STRT_DTTM:  "testd",
//		CS_SPKD_PCN__PACK_END_DTTM:   "teste",
//		CS_SPKD_PCN__DISC_CODE:       "testf",
//		TBL_OCCR:                     "999",
//	})
//}
