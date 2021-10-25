package Table_Struct

//分配表
type Distribution struct {
	Dsparepart_name string //备件名称
	Dtable_number string //维修表单号
	Dengineer_number string //技术工程师员工号
	Dtime string //分配技术工程师时间
	Dstatus string //维修状态
	Ddelay_state string //维修延迟程度
}
//CREATE TABLE public.distribution
//(
//dsparepart_name character(1) COLLATE pg_catalog."default" NOT NULL,
//dtable_number character(1) COLLATE pg_catalog."default" NOT NULL,
//dengineer_number character(1) COLLATE pg_catalog."default" NOT NULL,
//dtime character(1) COLLATE pg_catalog."default" NOT NULL,
//dstatus character(1) COLLATE pg_catalog."default" NOT NULL,
//ddelay_state character(1) COLLATE pg_catalog."default" NOT NULL,
//CONSTRAINT distribution_pkey PRIMARY KEY (dtable_number),
//CONSTRAINT distribution_dengineer_number_fkey FOREIGN KEY (dengineer_number)
//REFERENCES public.engineer (enumber) MATCH SIMPLE
//ON UPDATE NO ACTION
//ON DELETE NO ACTION,
//CONSTRAINT distribution_dsparepart_name_fkey FOREIGN KEY (dsparepart_name)
//REFERENCES public.sparepart (sname) MATCH SIMPLE
//ON UPDATE NO ACTION
//ON DELETE NO ACTION,
//CONSTRAINT distribution_dtable_number_fkey FOREIGN KEY (dtable_number)
//REFERENCES public.repair_table (rnumber) MATCH SIMPLE
//ON UPDATE NO ACTION
//ON DELETE NO ACTION
//)
