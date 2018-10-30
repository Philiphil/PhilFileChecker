package main

type OrderType string;

const(
	OT_delete OrderType = "delete"
	OT_git_push OrderType = "git_push"
	OT_ftp_sync OrderType = "ftp_sync"
	OT_execute	OrderType= "execute"
	OT_minify OrderType = "minify"
	OT_merge_files OrderType = "merge_files"
)

type Order struct{
	Type OrderType
	Params []string
}

type Trigger struct{
	Files []string
	Orders []Order
}


//MINIFY
const newlinetoken string = "\\~~#~#&&/NEWLINETOKEN/19*!ou;Hze"
type StringInFile struct{
	Begin int
	End int
}

type Delimiter struct{
	Begin string
	End string
}