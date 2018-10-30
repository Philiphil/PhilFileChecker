package main


func minify(file string){
	t_file := getFileType(file)
	s_file := getFileContent(file)
	strings :=  minify_searchForString(s_file, []string{"\"", "'", "`"})




}

func minify_searchForString(s_file string, delimiters []string)(strings []StringInFile){
	var stack []string
	in_string := false

	for _, char := range s_file {
		is_delimiter := HasElem(delimiters,string(char))
		if is_delimiter {
			if len(stack) > 1  && stack[len(stack)-2] != "\\"{
				if in_string{
					
				}
				in_string = !in_string
			}
		}
		stack = append(stack, string(char))

		
	}
	return
}
