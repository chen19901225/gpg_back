# 备注

## 注意事项

* 只能通过`root`来运行, 不然的话需要输入密码,郁闷

* 主要用于宝塔的`postgres` 备份

## 例子

假设路径是 `bin_path`


* 备份数据库 `test`

    bin_path --database=test

* 备份数据库 `test`, 最大备份数目10

    bin_path --database=test --backup=10

