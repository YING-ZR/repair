package Table_Struct

//用户填写维修信息单表
type Repair_table struct {
	Rnumber string
	Rtime string
	Rphenomenon string//机器故障现象
	Rfault string//故障类型
	Rproduct string//选择产品类型
	Rremarks string//备注
	Rid string
	Rdistribution string//是否分配工程师
	Predicttime string //预测完成时间
	Predictprice string //预测价格
}
//CREATE TABLE public.repair_table
//(
//rnumber character(20) COLLATE pg_catalog."default" NOT NULL,
//rtime character(30) COLLATE pg_catalog."default" NOT NULL,
//rphenomenon character(60) COLLATE pg_catalog."default" NOT NULL,
//rfault character(20) COLLATE pg_catalog."default" NOT NULL,
//rproduct character(20) COLLATE pg_catalog."default" NOT NULL,
//rremarks character(200) COLLATE pg_catalog."default",
//rid character(20) COLLATE pg_catalog."default" NOT NULL,
//CONSTRAINT repair_table_pkey PRIMARY KEY (rnumber),
//CONSTRAINT repair_table_rid_fkey FOREIGN KEY (rid)
//REFERENCES public.customer (cid) MATCH SIMPLE
//ON UPDATE NO ACTION
//ON DELETE NO ACTION
//NOT VALID
//)