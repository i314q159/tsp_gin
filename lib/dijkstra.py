import heapq
import math
import random
import sys

import matplotlib.pyplot as plt
import numpy as np


def dijkstra(graph, start):
    n = len(graph)
    dist = [math.inf] * n
    dist[start] = 0
    visited = [False] * n
    pq = [(0, start)]
    while len(pq) > 0:
        _, u = heapq.heappop(pq)
        if visited[u]:
            continue
        visited[u] = True
        for v, d in enumerate(graph[u]):
            if d is None:
                continue
            alt = dist[u] + d
            if alt < dist[v]:
                dist[v] = alt
                heapq.heappush(pq, (alt, v))
    return dist


def plot_tsp(path, cities):
    x = [city[0] for city in cities]
    y = [city[1] for city in cities]
    plt.figure(figsize=(5, 5))
    plt.scatter(x, y)
    for i in range(len(path) - 1):
        plt.plot([x[path[i]], x[path[i + 1]]], [y[path[i]], y[path[i + 1]]], "b")
    plt.plot([x[path[-1]], x[path[0]]], [y[path[-1]], y[path[0]]], "b")

    plt.savefig("./tmp/dijkstra.png")
    # plt.show()


# 把二维坐标转换为图，图上每对城市之间的距离都是它们之间的欧几里得距离。
def generate_graph(cities):
    n = len(cities)
    graph = [[None] * n for _ in range(n)]
    for i in range(n):
        for j in range(i + 1, n):
            xi, yi = cities[i]
            xj, yj = cities[j]
            dist = math.sqrt((xi - xj) ** 2 + (yi - yj) ** 2)  # 计算欧几里得距离
            graph[i][j] = graph[j][i] = dist  # type: ignore
    return graph


def test():
    np.random.seed(random.randint(0, 10))
    data = np.random.rand(20, 2) * 10
    graph = generate_graph(data)
    dist = dijkstra(graph, 0)
    path = [i for i, _ in sorted(enumerate(dist), key=lambda x: x[1])]
    plot_tsp(path, data)


if __name__ == "__main__":
    args = sys.argv
    data = [[int(num) for num in seq.split(",")] for seq in args[1].split()]

    graph = generate_graph(data)
    dist = dijkstra(graph, 0)
    path = [i for i, _ in sorted(enumerate(dist), key=lambda x: x[1])]
    plot_tsp(path, data)
