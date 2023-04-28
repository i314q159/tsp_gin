import random
from math import floor

import matplotlib.pyplot as plt
import numpy as np


# 构建一个类保存遗传算法的初始化参数和函数计算
class GAs(object):
    def __init__(
        self,
        data,
        max_gen=200,
        size_pop=200,
        cross_prob=0.9,
        pmuta_prob=0.01,
        select_prob=0.8,
    ):
        self.max_gen = max_gen  # 最大迭代次数
        self.size_pop = size_pop  # 群体个数
        self.cross_prob = cross_prob  # 交叉概率
        self.pmuta_prob = pmuta_prob  # 变异概率
        self.select_prob = select_prob  # 选择概率

        self.data = data  # 城市的左边数据
        self.num = len(data)  # 城市个数 对应染色体长度

        # 距离矩阵n*n, 第[i,j]个元素表示城市i到j距离
        self.matrix_distance = self.matrix_dis()

        # 通过选择概率确定子代的选择个数
        self.select_num = max(floor(self.size_pop * self.select_prob + 0.5), 2)

        # 父代和子代群体的初始化（不直接用np.zeros是为了保证单个染色体的编码为整数，np.zeros对应的数据类型为浮点型）
        self.chrom = np.array([0] * self.size_pop * self.num).reshape(
            self.size_pop, self.num
        )
        self.sub_sel = np.array([0] * self.select_num * self.num).reshape(
            self.select_num, self.num
        )

        # 存储群体中每个染色体的路径总长度，对应单个染色体的适应度就是其倒数
        self.fitness = np.zeros(self.size_pop)

        # 保存每一步的群体的最优路径和距离
        self.best_fit = []
        self.best_path = []

    # 计算城市间的距离函数
    def matrix_dis(self):
        res = np.zeros((self.num, self.num))
        for i in range(self.num):
            for j in range(i + 1, self.num):
                res[i, j] = np.linalg.norm(self.data[i, :] - self.data[j, :])
                res[j, i] = res[i, j]
        return res

    # 随机产生初始化群体函数
    def rand_chrom(self):
        rand_ch = np.array(range(self.num))
        for i in range(self.size_pop):
            np.random.shuffle(rand_ch)
            self.chrom[i, :] = rand_ch
            self.fitness[i] = self.comp_fit(rand_ch)

    # 计算单个染色体的路径距离值，可利用该函数更新self.fittness
    def comp_fit(self, one_path):
        res = 0
        for i in range(self.num - 1):
            res += self.matrix_distance[one_path[i], one_path[i + 1]]
        res += self.matrix_distance[one_path[-1], one_path[0]]
        return res

    # 路径可视化函数
    def out_path(self, one_path):
        res = str(one_path[0] + 1) + "->"
        for i in range(1, self.num):
            res += str(one_path[i] + 1) + "->"
        res += str(one_path[0] + 1) + "\n"
        print(res)

    # 子代选取，根据选中概率与对应的适应度函数，采用随机遍历选择方法
    def select_sub(self):
        fit = 1.0 / (self.fitness)  # 适应度函数
        cumsum_fit = np.cumsum(fit)
        pick = (
            cumsum_fit[-1]
            / self.select_num
            * (np.random.rand() + np.array(range(self.select_num)))
        )
        i, j = 0, 0
        index = []
        while i < self.size_pop and j < self.select_num:
            if cumsum_fit[i] >= pick[j]:
                index.append(i)
                j += 1
            else:
                i += 1
        self.sub_sel = self.chrom[index, :]

    # 交叉，依概率对子代个体进行交叉操作
    def cross_sub(self):
        if self.select_num % 2 == 0:
            num = range(0, self.select_num, 2)
        else:
            num = range(0, self.select_num - 1, 2)
        for i in num:
            if self.cross_prob >= np.random.rand():
                self.sub_sel[i, :], self.sub_sel[i + 1, :] = self.intercross(
                    self.sub_sel[i, :], self.sub_sel[i + 1, :]
                )

    def intercross(self, ind_a, ind_b):
        r1 = np.random.randint(self.num)
        r2 = np.random.randint(self.num)
        while r2 == r1:
            r2 = np.random.randint(self.num)
        left, right = min(r1, r2), max(r1, r2)
        ind_a1 = ind_a.copy()
        ind_b1 = ind_b.copy()
        for i in range(left, right + 1):
            ind_a2 = ind_a.copy()
            ind_b2 = ind_b.copy()
            ind_a[i] = ind_b1[i]
            ind_b[i] = ind_a1[i]
            x = np.argwhere(ind_a == ind_a[i])
            y = np.argwhere(ind_b == ind_b[i])
            if len(x) == 2:
                ind_a[x[x != i]] = ind_a2[i]
            if len(y) == 2:
                ind_b[y[y != i]] = ind_b2[i]
        return ind_a, ind_b

    # 变异模块
    def mutation_sub(self):
        for i in range(self.select_num):
            if np.random.rand() <= self.cross_prob:
                r1 = np.random.randint(self.num)
                r2 = np.random.randint(self.num)
                while r2 == r1:
                    r2 = np.random.randint(self.num)
                self.sub_sel[i, [r1, r2]] = self.sub_sel[i, [r2, r1]]

    # 进化逆转
    def reverse_sub(self):
        for i in range(self.select_num):
            r1 = np.random.randint(self.num)
            r2 = np.random.randint(self.num)
            while r2 == r1:
                r2 = np.random.randint(self.num)
            left, right = min(r1, r2), max(r1, r2)
            sel = self.sub_sel[i, :].copy()

            sel[left : right + 1] = self.sub_sel[i, left : right + 1][::-1]
            if self.comp_fit(sel) < self.comp_fit(self.sub_sel[i, :]):
                self.sub_sel[i, :] = sel

    # 子代插入父代，得到相同规模的新群体
    def reins(self):
        index = np.argsort(self.fitness)[::-1]
        self.chrom[index[: self.select_num], :] = self.sub_sel


def gas(data):
    path_short = GAs(data)  # 根据位置坐标，生成一个遗传算法类
    path_short.rand_chrom()  # 初始化父类

    # 循环迭代遗传过程
    for i in range(path_short.max_gen):
        path_short.select_sub()  # 选择子代
        path_short.cross_sub()  # 交叉
        path_short.mutation_sub()  # 变异
        path_short.reverse_sub()  # 进化逆转
        path_short.reins()  # 子代插入

        # 重新计算新群体的距离值
        for j in range(path_short.size_pop):
            path_short.fitness[j] = path_short.comp_fit(path_short.chrom[j, :])

        # 每隔三十步显示当前群体的最优路径
        index = path_short.fitness.argmin()
        if (i + 1) % 20 == 0:
            print("第" + str(i + 1) + "步后的最短的路程: " + str(path_short.fitness[index]))
            path_short.out_path(path_short.chrom[index, :])  # 显示每一步的最优路径

        # 存储每一步的最优路径及距离
        path_short.best_fit.append(path_short.fitness[index])
        path_short.best_path.append(path_short.chrom[index, :])

    # 结果图
    fig, ax = plt.subplots()
    x = data[:, 0]
    y = data[:, 1]
    ax.scatter(x, y, linewidths=0.1)
    for i, txt in enumerate(range(1, len(data) + 1)):
        ax.annotate(str(txt), (x[i], y[i]))
    res0 = path_short.chrom[0]
    x0 = x[res0]
    y0 = y[res0]
    for i in range(len(data) - 1):
        plt.quiver(
            x0[i],
            y0[i],
            x0[i + 1] - x0[i],
            y0[i + 1] - y0[i],
            color="blue",
            width=0.005,
            angles="xy",
            scale=1,
            scale_units="xy",
        )
    plt.quiver(
        x0[-1],
        y0[-1],
        x0[0] - x0[-1],
        y0[0] - y0[-1],
        color="blue",
        width=0.005,
        angles="xy",
        scale=1,
        scale_units="xy",
    )

    plt.savefig("./tmp/tsp_gas.png")
    # plt.show()
    return path_short  # 返回遗传算法结果类


def run(data):
    gas(data)


if __name__ == "__main__":
    np.random.seed(random.randint(0, 10))
    data = np.random.rand(20, 2) * 10
    gas(data)
