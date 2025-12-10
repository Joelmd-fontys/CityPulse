# frontend/heatmap.py
import json
import os

import matplotlib.pyplot as plt
import numpy as np


def main():
    here = os.path.dirname(os.path.abspath(__file__))
    grid_path = os.path.join(here, "grid.json")

    with open(grid_path, "r", encoding="utf-8") as f:
        grid = json.load(f)

    arr = np.array(grid, dtype=float)

    plt.figure(figsize=(5, 5))
    im = plt.imshow(
        arr,
        cmap="hot",
        origin="lower",
        interpolation="nearest",
        vmin=0.0,
        vmax=1.0,
    )
    plt.colorbar(im, label="Congestion (0–1)")
    plt.title("CityPulse congestion tile (example area)")
    plt.xlabel("X (tile index)")
    plt.ylabel("Y (tile index)")
    plt.tight_layout()
    plt.show()


if __name__ == "__main__":
    main()