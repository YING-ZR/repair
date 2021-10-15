package Table_Struct

//用户资料表
type Customer struct {
	Cname string
	Ctelephone string
	Cpassword string
	Cid string
	Ctel string
	Cemail string
	Caddress string
	Cpostcode string
	Ccompany string
	Cnature string
}
//CREATE TABLE public.customer
//(
//cname character(20) COLLATE pg_catalog."default" NOT NULL,
//ctelephone character(20) COLLATE pg_catalog."default" NOT NULL,
//cpassword character(20) COLLATE pg_catalog."default" NOT NULL,
//cid character(20) COLLATE pg_catalog."default" NOT NULL,
//ctel character(15) COLLATE pg_catalog."default",
//cemail character(40) COLLATE pg_catalog."default" NOT NULL,
//caddress character(50) COLLATE pg_catalog."default" NOT NULL,
//cpostcode character(15) COLLATE pg_catalog."default" NOT NULL,
//ccompany character(30) COLLATE pg_catalog."default",
//cnature character(15) COLLATE pg_catalog."default",
//CONSTRAINT customer_pkey PRIMARY KEY (cid)
//)
