## 什么是map
map是Go中的内置类型，它将一个值与一个键关联起来。可以使用相应的键检索值。

Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值 Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的，也是引用类型

使用map过程中需要注意的几点：

- map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取 

- map的长度是不固定的，也就是和slice一样，也是一种引用类型
    
- 内置的len函数同样适用于map，返回map拥有的key的数量
    
- map的key可以是所有可比较的类型，如布尔型、整数型、浮点型、复杂型、字符串型……也可以键。
- 与切片相似，映射是引用类型。当将映射分配给一个新变量时，它们都指向相同的内部数据结构。因此，一个的变化会反映另一个
- map元素可比较，非nil的map之间不可比较
- delete(map, key) 函数用于删除集合的元素, 参数为 map 和其对应的 key。删除函数不返回任何值。


