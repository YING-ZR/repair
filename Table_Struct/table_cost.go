package Table_Struct

//维修费用单表
type Table_Cost struct {
	tcost string //总费用
	Tnumber string //维修信息单单号
	tsparepart_cost string //备件费
	tengineer_cost string //人工费
	tproblem_cost string //问题费用
}
//CREATE TABLE public.table_cost
//(
//tcost character(20) COLLATE pg_catalog."default" NOT NULL,
//tnumber character(20) COLLATE pg_catalog."default" NOT NULL,
//tsparepart_cost character(20) COLLATE pg_catalog."default" NOT NULL,
//tengineer_cost character(20) COLLATE pg_catalog."default" NOT NULL,
//tproblem_cost character(20) COLLATE pg_catalog."default" NOT NULL,
//CONSTRAINT table_cost_pkey PRIMARY KEY (tnumber),
//CONSTRAINT table_cost_tnumber_fkey FOREIGN KEY (tnumber)
//REFERENCES public.repair_table (rnumber) MATCH SIMPLE
//ON UPDATE NO ACTION
//ON DELETE NO ACTION
//)
