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
		s_file, strings= minify_remove(s_file,strings,bfr)
	}

	//s_file, strings= minify_remove_white_space(s_file,strings)

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

func minify_remove(s_file string, strings []StringInFile,comments []StringInFile)(string, []StringInFile){
	for _, comment := range comments {
		is_comment_in_string := false
		for _, str := range strings{
			if comment.Begin > str.Begin && comment.End < str.End{
				is_comment_in_string = true
			}
		}
		if !is_comment_in_string{
			s_file = s_file[:comment.Begin] + s_file[comment.End:]
			for _, str := range strings{
				if str.Begin > comment.End{
					str.Begin = str.Begin - comment.End + comment.Begin
					str.End = str.End - comment.End + comment.Begin
				}
			}
			for _, cts := range comments{
				if cts.Begin > comment.End{
					cts.Begin = cts.Begin - comment.End + comment.Begin
					cts.End = cts.End - comment.End + comment.Begin
				}
			}
		}
	}
	return s_file, strings
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

func minify_remove_white_space(s_file string, protected []StringInFile)(string,[]StringInFile){
	f_sfile := "";
	for i, char := range s_file {
		s_char := string(char)
		is_whitespace := false
		is_deletable := false

		switch(s_char){
			case "\t":is_whitespace=!is_whitespace
			case "\n":is_whitespace=!is_whitespace
			case "\r":is_whitespace=!is_whitespace
			case "\r\n":is_whitespace=!is_whitespace
			case " ":is_whitespace=!is_whitespace
		}

		if is_whitespace{
			is_deletable = true
			for _,p := range protected {
				if p.Begin < i && p.End > i{
					is_deletable = false
				}
			}
		}
		if !is_deletable{
			f_sfile += s_char
		}else{
			for _,p := range protected {
				if  i < p.Begin{
					p.Begin--
					p.End--
				}
			}
		}
	}
	return f_sfile, protected
}