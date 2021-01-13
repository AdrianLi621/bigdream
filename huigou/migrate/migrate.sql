##################轮播图表
INSERT INTO hg_carousel ( img_url, background_color ) VALUES( "http://pic.rmb.bdstatic.com/819a044daa66718c2c40a48c1ba971e6.jpeg", "#ff3e96" )
INSERT INTO hg_carousel ( img_url, background_color ) VALUES( "http://img.zcool.cn/community/017a4e58b4eab6a801219c77084373.jpg", "#fff8dc" )
INSERT INTO hg_carousel ( img_url, background_color ) VALUES( "http://images.jaadee.com/images/201702/goods_img/30150_d85aed83521.jpg", "#f5f5dc" )



##################属性表
insert into hg_goods_attrribute(attr_name) values("颜色");
insert into hg_goods_attrribute(attr_name) values("尺码");
insert into hg_goods_attrribute(attr_name) values("材质");
insert into hg_goods_attrribute(attr_name) values("机身颜色");
insert into hg_goods_attrribute(attr_name) values("存储容量");


##################属性值表

insert into hg_attribute_value(attr_value_name,attr_id) values("黑色",1);
insert into hg_attribute_value(attr_value_name,attr_id) values("白色",1);
insert into hg_attribute_value(attr_value_name,attr_id) values("黄色",1);
insert into hg_attribute_value(attr_value_name,attr_id) values("蓝色",1);
insert into hg_attribute_value(attr_value_name,attr_id) values("xl",2);
insert into hg_attribute_value(attr_value_name,attr_id) values("xxl",2);
insert into hg_attribute_value(attr_value_name,attr_id) values("l",2);
insert into hg_attribute_value(attr_value_name,attr_id) values("x",2);
insert into hg_attribute_value(attr_value_name,attr_id) values("丝绵",3);
insert into hg_attribute_value(attr_value_name,attr_id) values("涤纶",3);
insert into hg_attribute_value(attr_value_name,attr_id) values("羊毛",3);
insert into hg_attribute_value(attr_value_name,attr_id) values("亮黑色",4);
insert into hg_attribute_value(attr_value_name,attr_id) values("釉白色",4);
insert into hg_attribute_value(attr_value_name,attr_id) values("秘银色",4);
insert into hg_attribute_value(attr_value_name,attr_id) values("秋日胡杨",4);
insert into hg_attribute_value(attr_value_name,attr_id) values("8+128GB",5);
insert into hg_attribute_value(attr_value_name,attr_id) values("8+256GB",5);
insert into hg_attribute_value(attr_value_name,attr_id) values("12+128GB",5);
insert into hg_attribute_value(attr_value_name,attr_id) values("12+256GB",5);

##################菜单分类表
insert into  hg_goods_class (gc_name,gc_pid,gc_sort,gc_level,gc_pids,gc_img_url)values("女装",0,1,1,"","http://images.jaadee.com/images/201702/goods_img/30150_d85aed83521.jpg");
insert into  hg_goods_class (gc_name,gc_pid,gc_sort,gc_level,gc_pids,gc_img_url)values("上衣",1,1,2,"1","http://images.jaadee.com/images/201702/goods_img/30150_d85aed83521.jpg");
insert into  hg_goods_class (gc_name,gc_pid,gc_sort,gc_level,gc_pids,gc_img_url)values("连衣裙",2,1,3,"1,2","http://images.jaadee.com/images/201702/goods_img/30150_d85aed83521.jpg");
insert into  hg_goods_class (gc_name,gc_pid,gc_sort,gc_level,gc_pids,gc_img_url)values("男装",0,1,1,"","http://images.jaadee.com/images/201702/goods_img/30150_d85aed83521.jpg");
insert into  hg_goods_class (gc_name,gc_pid,gc_sort,gc_level,gc_pids,gc_img_url)values("上衣",4,1,2,"4","http://images.jaadee.com/images/201702/goods_img/30150_d85aed83521.jpg");
insert into  hg_goods_class (gc_name,gc_pid,gc_sort,gc_level,gc_pids,gc_img_url)values("短袖",5,1,3,"4,5","http://images.jaadee.com/images/201702/goods_img/30150_d85aed83521.jpg");
insert into  hg_goods_class (gc_name,gc_pid,gc_sort,gc_level,gc_pids,gc_img_url)values("数码",0,1,1,"","http://images.jaadee.com/images/201702/goods_img/30150_d85aed83521.jpg");
insert into  hg_goods_class (gc_name,gc_pid,gc_sort,gc_level,gc_pids,gc_img_url)values("手机",7,1,2,"7","http://images.jaadee.com/images/201702/goods_img/30150_d85aed83521.jpg");
insert into  hg_goods_class (gc_name,gc_pid,gc_sort,gc_level,gc_pids,gc_img_url)values("华为",8,1,3,"7,8","http://images.jaadee.com/images/201702/goods_img/30150_d85aed83521.jpg");



##################店铺表
INSERT into hg_store(store_name,store_logo,province_id,city_id,area_id,area_detail,address_info,store_state,store_keywords,store_phone,store_open_time,store_close_time)values("腾飞超市","http://pic.rmb.bdstatic.com/819a044daa66718c2c40a48c1ba971e6.jpeg",1,2,3,"刘集乡","河南省南阳市刘集乡",1,"刘集腾飞超市","111111","09:00","18:00");
INSERT into hg_store(store_name,store_logo,province_id,city_id,area_id,area_detail,address_info,store_state,store_keywords,store_phone,store_open_time,store_close_time)values("大润发超市","http://pic.rmb.bdstatic.com/819a044daa66718c2c40a48c1ba971e6.jpeg",1,2,3,"刘集乡","河南省南阳市刘集乡",1,"刘集大润发超市","22222","08:30","22:00");

