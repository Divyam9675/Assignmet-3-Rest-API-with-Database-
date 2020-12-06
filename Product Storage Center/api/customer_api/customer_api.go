 package customer_api
 
 import (
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	"config"
	"strconv"
	"Models"
	"entities"
	
 )

 func FindAll(response http.ResponseWriter, request *http.Request){

	db, err := config.GetDB()

	if err!= nil{
		responseWithError(response, http.StatusBadRequest, err.Error())
	}else{
		customerModel := Models.CustomerModel{
			Db: db,
		}
		customer, err2 := customerModel.FindAll()
		if err2 != nil{
			responseWithError(response, http.StatusBadRequest, err.Error())
		}else{
			responseWithJson(response, http.StatusOK, customer) 
		}

	}

 }


 func Search(response http.ResponseWriter, request *http.Request){

	vars := mux.Vars(request)

	keyword := vars["keyword"]

	db, err := config.GetDB()

	if err!= nil{
		responseWithError(response, http.StatusBadRequest, err.Error())
	}else{
		customerModel := Models.CustomerModel{
			Db: db,
		}
		customer, err2 := customerModel.Search(keyword)
		if err2 != nil{
			responseWithError(response, http.StatusBadRequest, err.Error())
		}else{
			responseWithJson(response, http.StatusOK, customer) 
		}

	}

 }



func Update(response http.ResponseWriter, request *http.Request){

	var custum entities.Customer
	err := json.NewDecoder(request.Body).Decode(&custum)

	db, err := config.GetDB()

	if err!= nil{
		responseWithError(response, http.StatusBadRequest, err.Error())
	}else{
		customerModel := Models.CustomerModel{
			Db: db,
		}
		_, err2 := customerModel.Update(&custum)
		if err2 != nil{
			responseWithError(response, http.StatusBadRequest, err.Error())
		}else{
			responseWithJson(response, http.StatusOK, custum) 
		}

	}

 }


 func Delete(response http.ResponseWriter, request *http.Request){

	vars := mux.Vars(request)
	sid  := vars["keyword"]
	
	id, _ := strconv.ParseInt(sid, 10, 64) 

	var custum entities.Customer
	err := json.NewDecoder(request.Body).Decode(&custum)

	db, err := config.GetDB()

	if err!= nil{
		responseWithError(response, http.StatusBadRequest, err.Error())
	}else{
		customerModel := Models.CustomerModel{
			Db: db,
		}
		RowsAffected ,err2 := customerModel.Delete(id)
		if err2 != nil{
			responseWithError(response, http.StatusBadRequest, err2.Error())
		}else{
			responseWithJson(response, http.StatusOK, map[string]int64{
				"RowsAffected": RowsAffected,
			}) 
		}

	}

 }

 func Create(response http.ResponseWriter, request *http.Request){

	vars := mux.Vars(request)

	keyword := vars["keyword"]

	id, _ := strconv.ParseInt(keyword, 10, 64) 


	var custum entities.Customer
	err := json.NewDecoder(request.Body).Decode(&custum)

	db, err := config.GetDB()

	if err!= nil{
		responseWithError(response, http.StatusBadRequest, err.Error())
	}else{
		customerModel := Models.CustomerModel{
			Db: db,
		}
		err2 := customerModel.Create(&custum, id)
		if err2 != nil{
			responseWithError(response, http.StatusBadRequest, err.Error())
		}else{
			responseWithJson(response, http.StatusOK, custum) 
		}

	}

 }






 func responseWithError(w http.ResponseWriter, code int, msg string){

	responseWithJson(w, code, map[string]string{"error": msg})

 }


 func responseWithJson(w http.ResponseWriter, code int, payload interface{}){

	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	//json.NewEncoder(w).Encode(response)

 }