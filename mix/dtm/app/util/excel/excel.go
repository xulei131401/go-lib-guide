package excel

/*func DownloadExcel(w http.ResponseWriter, req *http.Request, filename string, file *xlsx.File){
	w.Header().Add("Content-Disposition", "attachment")
	//w.Header().Add("Content-Type", "application/vnd.ms-excel")
	w.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	var buffer bytes.Buffer
	file.Write(&buffer)
	r := bytes.NewReader(buffer.Bytes())

	http.ServeContent(w, req, filename, time.Now(), r)
}*/