## 工厂方法模式
### 优点
1. 更符合开-闭原则
新增一种产品时，只需要增加相应的具体产品类和相应的工厂子类即可。
2. 符合单一职责原则
每个具体工厂类只负责创建对应的产品。
3. 不使用静态工厂方法
可以形成基于继承的等级结构。
### 缺点
1. 添加新产品时，除了增加新产品类外，
还要提供与之对应的具体工厂类，系统类的个数将成对增加，
在一定程度上增加了系统的复杂度。
同时，有更多的类需要编译和运行，会给系统带来一些额外的开销。
2. 由于考虑到系统的可扩展性，需要引入抽象层，
在客户端代码中均使用抽象层进行定义，增加了系统的抽象性和理解难度，
且在实现时可能需要用到DOM、反射等技术，增加了系统的实现难度。
3. 虽然保证了工厂方法内的对修改关闭，但对于使用工厂方法的类，
如果要更换另外一种产品，仍然需要修改实例化的具体工厂类。
4，一个具体工厂只能创建一种具体产品。