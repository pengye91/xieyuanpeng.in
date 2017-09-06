# Python 函数式编程学习3

### nonlocal变量

考虑如下的函数：

```python
In [14]: def make_averager():
    ...:     count = 0
    ...:     total = 0
    ...:     
    ...:     def averager(new_value):
    ...:         count += 1
    ...:         total += new_value
    ...:         return total / count
    ...:     
    ...:     return averager
    ...: 

In [15]: avg = make_averager()

In [16]: avg(10)
---------------------------------------------------------------------------
UnboundLocalError                         Traceback (most recent call last)
<ipython-input-16-2b3d43cb065d> in <module>()
----> 1 avg(10)

<ipython-input-14-fe9bdbf03cae> in averager(new_value)
      4 
      5     def averager(new_value):
----> 6         count += 1
      7         total += new_value
      8         return total / count

UnboundLocalError: local variable 'count' referenced before assignment

```

函数定义的问题在于：`count += 1`对自由变量赋值了，`total`也是同样的问题。

但是上一节中的实例函数中： `series.append(new_value）`就没有这个问题，这是因为**series是可变类型list**。但是对于不可变类型来说，自由变量只能读取，不能更新，如果尝试重新绑定，如`count = count + 1`，会隐式创建局部变量`count`，这样一来，`count`就不是自由变量了，因此不会保存在闭包中。

**解决: nonlocal声明**

`nonlocal`关键字的作用是把变量标记为自由变量。

如果nonlocal声明的变量被赋予新值，闭包中保存的绑定会更新。

正确的代码如下：

```python
In [17]: def make_averager():
    ...:     count = 0
    ...:     total = 0
    ...:     
    ...:     def averager(new_value):
    ...:         nonlocal count, total
    ...:         count += 1
    ...:         total += new_value
    ...:         return total / count
    ...:     
    ...:     return averager
    ...: 
    ...: 

In [18]: avg = make_averager()

In [19]: avg(10)
Out[19]: 10.0

In [20]: avg(20)
Out[20]: 15.0

In [21]: avg(200)
Out[21]: 76.66666666666667
```

*注意：nonlocal只在python3中可用, [python2需要用hack的方法](http://www.python.org/dev/peps/pep-3104/)*