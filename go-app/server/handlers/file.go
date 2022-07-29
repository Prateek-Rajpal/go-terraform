package handlers

import (
	"encoding/json"
	"golang-app/db"
	"golang-app/httptypes"
	"io"

	"io/ioutil"
	"net/http"
	"os"
)

func AddEmployeeHandler(d *db.MyDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 * 1024 * 1024) // 10 mb
		var employees httptypes.Employees
		file, handler, err := r.FormFile("myfile")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		jsonFile, err := os.Create(os.TempDir() + "/" + handler.Filename)
		if err != nil {
			panic(err)
		}
		defer jsonFile.Close()
		// byteValue, _ := ioutil.ReadAll(file)
		// json.Unmarshal(byteValue, x)
		io.Copy(jsonFile, file)
		defer jsonFile.Close()

		newFile, err := os.Open(os.TempDir() + "/" + handler.Filename)
		if err != nil {
			panic(err)
		}

		defer newFile.Close()
		byteValue, _ := ioutil.ReadAll(newFile)
		json.Unmarshal(byteValue, &employees)

		if err := d.AddEmployee(&employees); err != nil {
			panic(err)
		}

		respond(w, map[string]interface{}{"fileName": handler.Filename, "type": handler.Header.Get("Content-Type")}, http.StatusOK)

	}
}

func GetEmployeesHandler(d *db.MyDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		value := d.GetEmployees()

		respond(w, value, http.StatusOK)
	}

}

// ***********************Helper Functions**********************************//

// Response function
func respond(w http.ResponseWriter, data interface{}, status int) {
	// Writing response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, "Could not encode in json", status)

		}
	}
}

// Decoding json body
func decode(w http.ResponseWriter, r *http.Request, v interface{}, status int) error {
	return json.NewDecoder(r.Body).Decode(v)
}
