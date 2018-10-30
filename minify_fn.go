package main


func minify(file string){
	t_file := getFileType(file)
	s_file := getFileContent(file)
	strings :=  minify_searchForString(s_file, []string{"\"", "'", "`"})




}

func minify_searchForString(s_file string, delimiters []string)(strings []StringInFile){
	var stack []string
	for _, chars := range s_file {
		
	}
	return
}