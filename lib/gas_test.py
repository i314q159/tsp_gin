import random
import numpy as np
import gas

if __name__ == "__main__":
    np.random.seed(random.randint(0, 10))
    data = np.random.rand(40, 2) * 10
    gas.gas(np.array(data))
