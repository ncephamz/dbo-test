INSERT INTO public.stores (id,"level","name",created_at,updated_at,deleted_at) VALUES
	 (1,'DISTRIBUTOR','Cimaung Store','2024-09-16 15:56:17.894873',NULL,NULL);

INSERT INTO public.products (id,store_id,product_name,image,uom,description,created_at,updated_at,deleted_at) VALUES
	 (1,1,'Baja Ringan 0.7',NULL,'Lente','Baja ringan','2024-09-16 16:39:43.896987',NULL,NULL);

INSERT INTO public.warehouses (id,status,"name",created_at,updated_at,deleted_at) VALUES
	 (1,'ACTIVE','Cimaung Warehouse','2024-09-16 16:01:02.986545',NULL,NULL);

INSERT INTO public.warehouses_coverages (id,warehouse_id,province,city,district,sub_district,zipcode,delivery_fee,service_fee,created_at,updated_at,deleted_at,tax) VALUES
	 (1,1,'JAWA BARAT','KABUPATEN BANDUNG','CIMAUNG','CIMAUNG','40375',200000.000,50000.000,'2024-09-16 17:24:35.030463',NULL,NULL,5);

INSERT INTO public.stores_warehouses (id,store_id,warehouse_coverage_id,product_id,qty,price,created_at,updated_at,deleted_at) VALUES
	 (1,1,1,1,10,70000.000,'2024-09-16 16:40:15.016269',NULL,NULL);



