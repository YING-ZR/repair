package util

func Updatesinventorystatus (number int ,alertquantity int) string{
	var sinventorystatus = ""
	switch {
	case number > alertquantity :
		sinventorystatus = "正常"
	case number == alertquantity :
		sinventorystatus = "临界"
	case number>0 && number<alertquantity :
		sinventorystatus = "警示"
	case number==0:
		sinventorystatus = "缺货"
	default:
		sinventorystatus = "number"
	}
	return sinventorystatus
}

