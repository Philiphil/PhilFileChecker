package main

type OrderType string;

const(
	OT_delete OrderType = "delete"
	OT_git_push OrderType = "git_push"
	OT_ftp_sync OrderType = "ftp_sync"
	OT_execute	OrderType= "execute"
)

type Order struct{
	Type OrderType
	Params []string
}

type Trigger struct{
	Files []string
	Orders []Order
}
