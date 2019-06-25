IO 主流程
写入流程：
	写入memtable -> 分级合并 -> 写入数据文件
读取流程：
	读取memtable -> 分级扫描 -> 读取数据文件


引擎初始化流程：
	buffer