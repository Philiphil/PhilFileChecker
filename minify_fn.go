package main
import(
	"fmt"
)

func minify(file string){
	t_file := getFileType(file)
	s_file := getFileContent(file)
	var strings []StringInFile
	var delimiters []Delimiter
	var bfr []StringInFile

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
		s_file, bfr = minify_filter(s_file,strings,delimiter)
		fmt.Println(bfr)
		s_file, strings= minify_remove(s_file,strings,bfr)

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

func minify_remove(s_file string, strings []StringInFile,comments []StringInFile)(f_sf string, f_strings []StringInFile){
	for _, comment := range comments {
		is_comment_in_string := false
		for _, str := range strings{
			if comment.Begin > str.Begin && comment.End < str.End{//comment end in string
				is_comment_in_string = true
			}
		}
		if !is_comment_in_string{

		}
	}
	return
}


func minify_filter(s_file string, protected []StringInFile, delimiter Delimiter)(f_sf string, comments []StringInFile){
	is_in_delimiter := false
	s_bfr := ""

	for index, char := range s_file {
		s_bfr += string(char)
		if is_in_delimiter && minify_detect_delimiter(delimiter.End, s_bfr){
			is_in_delimiter = false
			comments[len(comments)-1].End = index - len(delimiter.End)
		}else if !is_in_delimiter && minify_detect_delimiter(delimiter.Begin, s_bfr){
			is_in_delimiter = true
			comments = append(comments, StringInFile{})
			comments[len(comments)-1].Begin = index- len(delimiter.Begin)
		}
	}
	return
}

func minify_detect_delimiter(delimiter string, haystack string)(bool){
	if delimiter == newlinetoken{
		return minify_detect_delimiter(string("\n"),haystack) || minify_detect_delimiter(string("\r\n"),haystack)
	}
	return len(haystack) >= len(delimiter) && haystack[len(haystack)-len(delimiter):len(haystack)] == delimiter
}