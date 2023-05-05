import random

import numpy as np
import greedy

if __name__ == "__main__":
    np.random.seed(random.randint(0, 10))
    data = np.random.rand(1000, 2) * 10
    path, length = greedy.tsp(data)
    greedy.plot_tsp(path, data)
