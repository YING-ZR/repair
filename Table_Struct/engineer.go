package Table_Struct

//技术工程师表
type Engineer struct {
	Ename string //名字
	Eage string //工龄
	Enumber string //工号
	Etelephone string //电话
	Ecost string //人工费
	Eproblem string //维修问题编号
	Estatue string //是否正在维修
}
//CREATE TABLE public.engineer
//(
//ename character(20) COLLATE pg_catalog."default",
//eage character(10) COLLATE pg_catalog."default",
//enumber character(20) COLLATE pg_catalog."default" NOT NULL,
//etelephone character(20) COLLATE pg_catalog."default" NOT NULL,
//ecost character(20) COLLATE pg_catalog."default" NOT NULL,
//eequipment character(20) COLLATE pg_catalog."default" NOT NULL,
//etime character(10) COLLATE pg_catalog."default" NOT NULL,
//CONSTRAINT engineer_pkey PRIMARY KEY (enumber)
//)
