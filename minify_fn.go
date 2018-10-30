package main
import(
	
)

func minify(file string){
	t_file := getFileType(file)
	s_file := getFileContent(file)
	var strings []StringInFile
	var delimiters []Delimiter

	if t_file != "html"{
		strings =  append( strings, minify_searchForString(s_file, []string{"\"", "'", "`"})... )
	}else{
		//strings =  append( strings, minify_searchForStringHTML(s_file, []string{"\"", "'", "`"})... )
	}
	switch t_file {
	case "css":
		delimiters = append(delimiters, Delimiter{"//", newlinetoken})
	}

	
}

func minify_searchForString(s_file string, delimiters []string)(strings []StringInFile){
	var stack []string
	in_string := false

	for index, char := range s_file {
		is_delimiter := HasElem(delimiters,string(char))
		if is_delimiter {
			if len(stack) > 1  && stack[len(stack)-2] != "\\"{
				if in_string{
					strings[len(strings)-1].End = index
				}else{
					strings = append(strings,StringInFile{})
					strings[len(strings)-1].Begin = index
				}
				in_string = !in_string
			}
		}
		stack = append(stack, string(char))
	}
	return
}

func minify_filterBetween(s_file string, protected []StringInFile, from string, to string)(f_sf string){

	return
}
