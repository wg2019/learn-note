# InnoDB
*InnoDB是一种通用储存引擎，平衡了高可靠性和高性能。在MySQL中，InnoDB是默认的储存引擎。除非配置了不同的默认储存引擎，否则，所有的表都会使用InnoDB储存引擎。*
## InnoDB的主要优势
* DML操作遵循ACID模型，交易具有提交、回滚和崩溃恢复功能，以保护用户数据。
* 行级锁定和甲骨文风格的一致读取提高了多用户并发性和性能。
* 支持FOREIGN KEY约束，保持数据完整性。使用外键，会检查插入、更新和删除，以确保它们不会导致表之间的不一致。
## InnoDB储存引擎功能
|特点|支持|
|:---:|:---:|
|B树索引|是|
|备份/时间点恢复|是|
|集群数据库支持|否|
|聚类索引|是|
|压缩数据|是|
|数据缓存|是|
|加密数据|是|
|外键|是|
|全文搜索索引|是|
|地理空间数据类型支持|是|
|地理空间索引支持|是|
|哈希索引|否|
|索引缓存|是|
|锁定粒度|行|
|MVCC|是|
|复制支持|是|
|存储限制|64TB|
|T树索引|否|
|事务|是|
|更新数据字典的统计数据|是|
## ACID模型
### 原子性
* AUTOCOMMIT设置。
* COMMIT声明。
* ROLLBACK语句。
### 一致性
* InnoDB双写缓冲区。
* InnoDB崩溃恢复。
### 隔离性
* AUTOCOMMIT设置。
* 事务隔离级别和SET TRANSACTION语句。
* InnoDB锁定粒度。
### 持久性
* InnoDB 双写缓冲区。
* innodb_flush_log_at_trx_commit变量。
* sync_binlog变量。
* innodb_file_per_table变量。
* 存储设备（如磁盘驱动器、SSD或RAID阵列）中的写入缓冲区。
* 存储设备中的电池后备缓存。
* 用于运行MySQL的操作系统，特别是它对fsync()系统调用的支持。
* 不间断电源（UPS）保护运行MySQL服务器和存储MySQL数据的所有计算机服务器和存储设备的电力。
* 您的备份策略，例如备份的频率和类型，以及备份保留期。
* 对于分布式或托管数据应用程序，MySQL服务器硬件所在的数据中心的特殊特征，以及数据中心之间的网络连接。
## 数据页结构
*页是InnoDB管理储存空间的基本单位，一个页的大小默认是16KB。*

```
SHOW GLOBAL STATUS like 'Innodb_page_size';
```
|名称|中文名|占用空间|描述|
|:--:|:--:|:--:|:--:|
|File Header|文件头部|38字节|页的一些通用信息|
|Page Header|页面头部|56字节|数据页专有的一些信息|
|Infimun + Supremum Records|最小记录和最大记录|26字节|两个虚拟的行记录|
|User Records|用户记录|不确定|实际存储的行记录内容|
|Free Space|空闲空间|不确定|页中尚未使用的空间|
|PageDirectory|页面目录|不确定|页中的某些记录的相对位置|
|File Trailer|文件尾部|8字节|校验页是否完整|

## InnoDB行格式
*InnoDB支持两种行格式：Compact和Redundant。Compressed 和 Dynamic是基于Compact改进的紧凑行格式。*

```
CREATE TABLE 表名 (列的信息) ROW_FORMAT=行格式名称
ALTER TABLE 表名 ROW_FORMAT=行格式名称
```

### COMPACT行格式
 * **变长字段长度列表**：VARCHAR(M)、VARBINARY(M)、TEXT类型，BLOB类型，这些数据类型修饰列称为变长字段，变长字段中存储多少字节的数据不是固定的，所以我们在存储真实数据的时候需要顺便把这些数据占用的字节数也存起来。在Compact行格式中，把所有变长字段的真实数据占用的字节长度都存放在记录的开头部位，从而形成一个变长字段长度列表。
 * **NULL值列表**：Compact行格式会把可以为NULL的列统一管理起来，存一个标记为在NULL值列表中，如果表中没有允许存储 NULL 的列，则 NULL值列表也不存在了。
    * 二进制位的值为1时，代表该列的值为NULL。
    * 二进制位的值为0时，代表该列的值不为NULL。
 * **记录头信息**：除了变长字段长度列表、NULL值列表之外，还有一个用于描述记录的记录头信息，它是由固定的5个字节组成。5个字节也就是40个二进制位，不同的位代表不同的意思。
    |名称|大小（单位：bit）|描述|
    |:--:|:--:|:--:|
    |预留位1|1|没有使用|
    |预留位2|1|没有使用|
    |delete_mask|1|标记该记录是否被删除|
    |min_rec_mask|1|B+树的每层非叶子节点中的最小记录都会添加该标记|
    |n_owned|4|表示当前记录拥有的记录数|
    |heap_no|13|表示当前记录在记录堆的位置信息|
    |record_type|3|表示当前记录的类型，0表示普通记录，1表示B+树非叶子节点记录，2表示最小记录，3表示最大记录|
    |next_record|16|表示下一条记录的相对位置|
### 记录的真实数据
*记录的真实数据除了我们自己定义的列的数据以外，还会有三个隐藏列：*
|列名|是否必须|真实名称|占用空间|描述|
|:--:|:--:|:--:|:--:|:--:|
|row_id|否|DB_ROW_ID|6字节|行ID，唯一标识一条记录|
|transaction_id|是|DB_TRX_ID|6字节|事务ID|
|roll_pointer|是|DB_ROLL_PTR|7字节|回滚指针|
### 行溢出数据
**VARCHAR(M)类型的列最多可以占用65535个字节。其中的M代表该类型最多存储的字符数量，如果我们使用ascii字符集的话，一个字符就代表一个字节，我们看看VARCHAR(65535)是否可用：**
```
mysql> CREATE TABLE varchar_size_demo(
    ->     c VARCHAR(65535)
    -> ) CHARSET=ascii ROW_FORMAT=Compact;
ERROR 1118 (42000): Row size too large. The maximum row size for the used table type, not counting BLOBs, is 65535. This includes storage overhead, check the manual. You have to change some columns to TEXT or BLOBs
```
报错信息表达的意思是：MySQL对一条记录占用的最大存储空间是有限制的，除BLOB或者TEXT类型的列之外，其他所有的列（不包括隐藏列和记录头信息）占用的字节长度加起来不能超过65535个字节。这个65535个字节除了列本身的数据之外，还包括一些其他的数据，比如说我们为了存储一个VARCHAR(M)类型的列，其实需要占用3部分存储空间：
1. 真实数据
2. 变长字段真实数据的长度
3. NULL值标识
如果该VARCHAR类型的列没有NOT NULL属性，那最多只能存储65532个字节的数据，因为变长字段的长度占用2个字节，NULL值标识需要占用1个字节。

```
mysql> CREATE TABLE varchar_size_demo(
    ->      c VARCHAR(65532)
    -> ) CHARSET=ascii ROW_FORMAT=Compact;
Query OK, 0 rows affected (0.02 sec)
```
```
CREATE TABLE varchar_size_demo(
    c VARCHAR(65533) not null
) CHARSET=ascii ROW_FORMAT=Compact;
Query OK, 0 rows affected (0.02 sec)
```
### 记录中的数据太多产生的溢出
一个页的大小一般是16KB，也就是16384字节，而一个VARCHAR(M)类型的列就最多可以存储65533个字节，这样就可能出现一个页存放不了一条记录。
在Compact和Reduntant行格式中，对于占用存储空间非常大的列，在记录的真实数据处只会存储该列的一部分数据，把剩余的数据分散存储在几个其他的页中，然后记录的真实数据处用20个字节存储指向这些页的地址（当然这20个字节中还包括这些分散在其他页面中的数据的占用的字节数），从而可以找到剩余数据所在的页。
### Dynamic和Compressed行格式
这两种行格式类似于COMPACT行格式，只不过在处理行溢出数据时有点儿分歧，它们不会在记录的真实数据处存储一部分数据，而是把所有的数据都存储到其他页面中，只在记录的真实数据处存储其他页面的地址。另外，Compressed行格式会采用压缩算法对页面进行压缩。