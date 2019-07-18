## 针对insert语句的解析处理：
'''

    原始语句： "INSERT INTO A.T (C1 , C2 ) VALUES (1133123, 'ABC')"  

    解析后返回SQLObject 对象
    type SQLObject struct {
        DB            string
        TableList     []string
        SQLQuery      string
        SQLType       int
        SQLCommand    int
        CreateColumns []CreateColumnDefine // including create_table/insert col info
        QueryColumns  []QueryColumnDefine
        Changecoldata []ChangeColumnData
        YoctoSQLBaseListener
    }
    
    type ChangeColumnData struct {
        DataBefore string
        DataAfter  string
        DataType   int
    }
    
    解析后结果
    &{A [T] INSERT INTO A.T (C1 , C2 ) VALUES (1133123, 'ABC') 5 82 [{C1 0   []} {C2 0   []}] [] [{ 1133123 0} { 'ABC' 0}] {<nil>}}
    
    对应关系
    DB => A
    TableList => T
    SQLquery => 语句
    SQLType => 5(DML)
    SQLCommand => 82(inseret)
    CreateColumns => [{C1 0   []} {C2 0   []} // insert语句中columndefine的信息暂时可以忽略
    QueryColumns => []
    Changecoldata => [{ 1133123 0} { 'ABC' 0}] //此处针对增删改语句，包含databefore， dataafter， datatype三个信息，如果是insert语句，则直接从dataafter获取数据即可，其他两个变量暂时忽略

    
'''
