# TSP_Gin

毕业设计的后端部分，使用Gin框架。

## TSP

TSP(Traveling Salesman Problem,旅行商问题)：一个商品推销员要去若干个城市推销商品，该推销员从一个城市出发，需要经过所有城市后，回到出发地。应如何选择行进路线，以使总的行程最短。

### 遗传算法

1. 初始化编码：设置最大进化代数T_max、选择概率、交叉概率、变异概率、随机生成m个染色体的群体，每个染色体的编码对于一个可行的路径；
2. 适应度函数：对每一个染色体x，其个体适应度函数设置为f(x)=1/d，其中d表示该条路径的总长度；
3. 选择：将旧群体中的染色体以一定概率选择到新群体，每条染色体选中的概率与对应的适应度函数只相对应，本文采用随机遍历选择；
4. 交叉：在交叉概率的控制下，对选择群体中的个体进行两两交叉；
5. 变异：在变异概率的控制下，对单个染色体随机交换两个点的位置；
6. 进化逆转：将选择的染色体随机选择两个位置人r1、r2，将r1:r2的元素翻转为r2:r1，如果翻转后的适应度更高，则替换原染色体，否则不变；
7. 重插：选择的子代与父代结合，形成新的种群，循环操作。

### 贪心算法

每一步都是最优解。

### Dijkstra算法

Dijkstra算法是一种用于解决单源最短路径问题的算法，它的输入是一个加权有向图和一个起始节点，输出是从起始节点到所有其他节点的最短路径。Dijkstra算法的核心是计算最短路径，它并不关注城市的坐标信息，因此在算法实现上不需要使用坐标信息。

## 技术栈

+ gin
+ gorm
