import math
import random
import sys

import matplotlib.pyplot as plt
import numpy as np


# 贪心算法 Greedy algorithm
def distance(city1, city2):
    x1, y1 = city1
    x2, y2 = city2
    return math.sqrt((x1 - x2) ** 2 + (y1 - y2) ** 2)


def tsp(cities):
    # 计算所有城市之间的距离矩阵
    n = len(cities)
    distances = [[0.0] * n for i in range(n)]
    for i in range(n):
        for j in range(n):
            distances[i][j] = distance(cities[i], cities[j])

    # 计算一条近似最优的路径
    path = [0] * n
    visited = [False] * n
    visited[0] = True
    for i in range(1, n):
        min_distance = float("inf")
        next_city = None
        for j in range(1, n):
            # 如果该城市未被访问过，并且到上一个城市的距离更短
            if not visited[j] and distances[path[i - 1]][j] < min_distance:
                # 更新最短距离和下一个城市
                next_city = j
                min_distance = distances[path[i - 1]][j]
        # 将下一个城市添加到路径中，并标记为已访问
        path[i] = next_city  # type: ignore
        visited[next_city] = True  # type: ignore

    # 计算路径长度
    total_distance = 0.0
    for i in range(n - 1):
        total_distance += distances[path[i]][path[i + 1]]
    total_distance += distances[path[-1]][path[0]]

    # 返回路径和长度
    return path, total_distance


# 测试
# cities = [[0, 0], [1, 1], [2, 2], [3, 3]]
# path, length = tsp(cities)


def plot_tsp(path, cities):
    x = [city[0] for city in cities]
    y = [city[1] for city in cities]
    plt.figure(figsize=(5, 5))
    # 绘制城市坐标点
    plt.scatter(x, y)
    # 绘制路径
    for i in range(len(path) - 1):
        plt.plot([x[path[i]], x[path[i + 1]]], [y[path[i]], y[path[i + 1]]], "b")
    plt.plot([x[path[-1]], x[path[0]]], [y[path[-1]], y[path[0]]], "b")

    plt.savefig("./tmp/greedy.png")
    # plt.show()





if __name__ == "__main__":
    args = sys.argv
    data = [[int(num) for num in seq.split(",")] for seq in args[1].split()]

    path, length = tsp(data)
    # print("路径：", path)
    # print("长度：", length)
    plot_tsp(path, data)
