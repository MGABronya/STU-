# STU

[TOC]



## 添加边

接口： http://106.53.120.252:5504/add/edge 

方法：POST

传入数据格式：

​	Body部分，raw，json格式

​	需要包含左端点（非负数整数）pointa，右端点（非负数整数）pointb，长度（浮点数）length，路名（字符串）name，备注（字符串）remark，如下图

![1662017393466](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662017393466.png)

注意，边将会以添加的顺序编号（从1开始），删除一个边不会对影响其他边的编号。请按所需顺序添加边。

传回数据格式：

​	传回数据在Body的Pretty内，json格式，如下图

![1662017578160](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662017578160.png)

​	如传回失败信息，可能报错为“没有对应点”，“系统错误”或者“数据格式错误”，msg字符串中将会有详细信息

## 添加点

接口： http://106.53.120.252:5504/add/point

方法：POST

传入数据格式：

​	Body部分，raw，json格式

​	需要包含点名（字符串）name，备注（字符串）remark，如下图

![1662017817887](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662017817887.png)

注意，点将会以添加的顺序编号（从1开始），删除一个点不会对影响其他点的编号。请按所需顺序添加点。

传回数据格式：

​	传回数据在Body的Pretty内，json格式，如下图

![1662017983760](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662017983760.png)

如传回失败信息，可能报错为“系统错误”或者“数据格式错误”，msg字符串中将会有详细信息

## 删除边

接口： http://106.53.120.252:5504/delete/edge/:id 

方法：DELETE

传入数据格式：

​	将:id改为需要删除的边的id即可，如http://106.53.120.252:5504/delete/edge/1

​	注意，删除边对其它点的编号不会产生影响。

传回数据格式：

​	传回数据在Body的Pretty内，json格式，如下图

![1662018308938](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662018308938.png)

​	如传回失败信息，可能报错为“该边不存在”，msg字符串中将会有详细信息

## 删除点

接口： http://106.53.120.252:5504/delete/point/:id 

方法：DELETE

传入数据格式：

​	将:id改为需要删除的点的id即可，如http://106.53.120.252:5504/delete/point/1

​	注意，删除该点后，与该点相关的边也会被级联删除。

​	删除点对其它点的编号不会产生影响。

传回数据格式：

​	传回数据在Body的Pretty内，json格式，如下图

![1662018577155](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662018577155.png)

​	如传回失败信息，可能报错为“该点不存在”，msg字符串中将会有详细信息

## 显示所有边的信息

接口： http://106.53.120.252:5504/show/edges 

方法：GET

传入数据格式：

​	无需传入数据

传回数据格式：

​	传回数据在Body的Pretty内，json格式，如下图

![1662020043738](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662020043738.png)

## 显示所有点的信息：

接口： http://106.53.120.252:5504/show/points

方法：GET

传入数据格式：

​	无需传入数据

传回数据格式：

​	传回数据在Body的Pretty内，json格式，如下图

![1662020104493](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662020104493.png)

## 显示某条边的信息

接口： http://106.53.120.252:5504/show/edge/:id

方法：GET

传入数据格式：

​	将/:id替换为需要查看的边的编号即可

传回数据格式：

​	传回数据在Body的Pretty内，json格式，如下图

![1662020249228](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662020249228.png)

如传回失败信息，可能报错为“该边不存在”，msg字符串中将会有详细信息

## 显示某个点的信息

接口： http://106.53.120.252:5504/show/point/:id

方法：GET

传入数据格式：

​	将/:id替换为需要查看的点的编号即可

传回数据格式：

​	传回数据在Body的Pretty内，json格式，如下图

![1662020249228](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662020249228.png)

如传回失败信息，可能报错为“该点不存在”，msg字符串中将会有详细信息

![1662020334955](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662020334955.png)

## 更改某条边的信息

接口： http://106.53.120.252:5504/change/edge/:id

方法：POST

传入数据格式：

​	将:id改为需要更改的某条边的编号即可

​	在Body中的raw以json格式给出长度length（浮点数），名称name（字符串），备注remark（字符串）

![1662020753601](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662020753601.png)

传回数据格式：

传回数据在Body的Pretty内，json格式，如下图

![1662020899387](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662020899387.png)

如传回失败信息，可能报错为“该边不存在”，msg字符串中将会有详细信息

## 更改某个点的信息

接口： http://106.53.120.252:5504/change/point/:id

方法：POST

传入数据格式：

​	将:id改为需要更改的某个点的编号即可

​	在Body中的raw以json格式给出名称name（字符串），备注remark（字符串）

![1662021025146](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662021025146.png)

传回数据格式：

传回数据在Body的Pretty内，json格式，如下图

![1662021051627](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662021051627.png)

如传回失败信息，可能报错为“该点不存在”，msg字符串中将会有详细信息

## 查询路径

接口： http://106.53.120.252:5504/path

方法：POST

在Body中的raw以json格式给出起点标号start（整型），终点end（整型）,必经点location（整形数组），如下

![1662726393014](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662726393014.png)

传回数据格式：

传回数据在Body的Pretty内，json格式，包含code，msg，data

data内包含点数组Points,边数组 Edges,最终路径长度 length如下

![1662726474495](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662726474495.png)



![1662726485093](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\1662726485093.png)