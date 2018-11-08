package main
import(
	//"fmt"
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
		bfr = minify_filter(s_file,strings,delimiter)

		//fmt.Println(s_file[bfr[0].Begin : bfr[0].End])
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
					strings[len(strings)-1].End = index+1
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
	f_sf=s_file
	for _, comment := range comments {
		is_comment_in_string := false
		for _, str := range strings{
			if comment.Begin > str.Begin && comment.End < str.End{//comment end in string
				is_comment_in_string = true
			}
		}
		if !is_comment_in_string{
			f_sf = f_sf[:comment.Begin] + f_sf[comment.End:]
		}
	}
	return
}


func minify_filter(s_file string, protected []StringInFile, delimiter Delimiter)(comments []StringInFile){
	is_in_delimiter := false
	s_bfr := ""

	for index, char := range s_file {
		s_bfr += string(char)
		if is_in_delimiter && minify_detect_delimiter(delimiter.End, s_bfr){
			is_in_delimiter = false
			comments[len(comments)-1].End = index + minify_get_delimiter_size(delimiter.End,s_bfr)
		}else if !is_in_delimiter && minify_detect_delimiter(delimiter.Begin, s_bfr){
			is_in_delimiter = true
			comments = append(comments, StringInFile{})
			comments[len(comments)-1].Begin = index- minify_get_delimiter_size(delimiter.Begin,s_bfr)
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
func minify_get_delimiter_size(delimiter string, haystack string)(int){
	if delimiter == newlinetoken{
		if minify_detect_delimiter(string("\n"),haystack) {
			return minify_get_delimiter_size(string("\n"),haystack)
		}
		if minify_detect_delimiter(string("\r\n"),haystack){
			return minify_get_delimiter_size(string("\r\n"),haystack)
		}
	}
	return len(delimiter)-1
}