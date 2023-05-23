import random

import numpy as np
import dijkstra


if __name__ == "__main__":
    np.random.seed(random.randint(0, 10))
    data = np.random.rand(70, 2) * 10
    graph = dijkstra.generate_graph(data)
    dist = dijkstra.dijkstra(graph, 0)
    path = [i for i, _ in sorted(enumerate(dist), key=lambda x: x[1])]
    dijkstra.plot_tsp(path, data)
