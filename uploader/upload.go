package uploader

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Participant struct {
	Name     string
	Username string
	Email    string
	Points   int
}

var tempDir = "uploader/temp"

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading file\n")

	// input multipart form data
	r.ParseMultipartForm(10 << 20) // maximum upload of 10 MB file

	// using HTTP request, retrieve data from the "file" path.
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(r.FormFile("file"))
		fmt.Println("Error retrieving data")
		fmt.Println(err)
		return
	}
	defer file.Close()
	// Print uploaded file details
	fmt.Printf("Uploaded File : %+v\n", handler.Filename)
	fmt.Printf("File Size : %+v\n", handler.Size)
	fmt.Printf("MIME Header : %+v\n", handler.Header)

	// write Temporary file using create file path of temp-pdfs
	tempFile, err := ioutil.TempFile(tempDir, "upload-*.csv")
	if err != nil {
		fmt.Println(err)

	}
	defer os.Remove(tempFile.Name())

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	fmt.Fprintf(w, "Created File")
	// dir := path.Join(tempDir, `*.csv`)
	records := readCSV(tempFile.Name())
	// json.NewEncoder(w).Encode(records)
	w.Header().Set("Content-Type", "application/json")
	w.Write(records)
}

func readCSV(path string) []byte {
	csvFile, err := os.Open(path)
	if err != nil {
		log.Println(path)
		log.Fatal("file not found")
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Participant
	var allRecords []Participant

	for _, each := range csvData {
		oneRecord.Name = each[0]
		oneRecord.Username = each[1]
		oneRecord.Email = each[2]
		oneRecord.Points, _ = strconv.Atoi(each[3])
		allRecords = append(allRecords, oneRecord)
	}

	jsondata, err := json.Marshal(allRecords) // convert to JSON

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// check sanity
	fmt.Println(string(jsondata))
	return jsondata
}
