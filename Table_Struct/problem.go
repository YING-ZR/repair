package Table_Struct

//维修问题表
type Problem struct {
	Pname string //问题名称
	Pnumber string //问题编号
	Pcost string //问题预计价格
	Pproduct string //问题产品
	Ptime string //预计维修时间
}
//CREATE TABLE public.problem
//(
//pname character(20) COLLATE pg_catalog."default" NOT NULL,
//pnumber character(20) COLLATE pg_catalog."default" NOT NULL,
//pcost character(20) COLLATE pg_catalog."default" NOT NULL,
//pproduct character(20) COLLATE pg_catalog."default" NOT NULL,
//CONSTRAINT problem_pkey PRIMARY KEY (pnumber)
//)
