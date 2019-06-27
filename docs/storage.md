IO 主流程
写入流程：
	写入redo -> 写入memtable |（async）-> 分级合并 -> 写入数据文件
读取流程：
	读取memtable -> 分级扫描 -> 读取数据文件


引擎初始化流程：
	buffer
	
##buffer 管理
1.table cache 刷新	
	
##table
表类型：内存缓存表 内存有序表 磁盘分级表（C0~Cn）

##memtable 结构
[{db.table:[{key:value}]}