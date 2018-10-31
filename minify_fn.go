package main
import(
//	"fmt"
)

func minify(file string){
	t_file := getFileType(file)
	s_file := getFileContent(file)
	var strings []StringInFile
	var delimiters []Delimiter

	switch t_file {
		case "html" : 
			//strings =  append( strings, minify_searchForStringHTML(s_file)... )
			delimiters = append(delimiters, Delimiter{"<!--", "-->"})
		case "css":
			strings =  append( strings, minify_searchForString(s_file, []string{"\"", "'", "`"})... )
			delimiters = append(delimiters, Delimiter{"//", newlinetoken})
			delimiters = append(delimiters, Delimiter{"/*", "*/"})
		case "js":
			strings =  append( strings, minify_searchForString(s_file, []string{"\"", "'", "`"})... )
			delimiters = append(delimiters, Delimiter{"//", newlinetoken})
			delimiters = append(delimiters, Delimiter{"/*", "*/"})
	}
	for _, delimiter := range delimiters {
		s_file, strings = minify_filter(s_file,strings,delimiter)
	}
	writeToFile(s_file, file[:len(file)-len(t_file)] + "min." + t_file)
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

func minify_filter(s_file string, protected []StringInFile, delimiters Delimiter)(f_sf string, strings []StringInFile){
	return
}

func minify_extract(strings []string, str string)(rstrings []string, rstr string){
	return
}