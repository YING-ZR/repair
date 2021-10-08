package Table_Struct

type Repair_table struct {
	Rnumber string
	Rtime string
	Rphenomenon string
	Rfault string
	Rproduct string
	Rremarks string
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