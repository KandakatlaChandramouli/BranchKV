import pandas as pd
import matplotlib.pyplot as plt

df = pd.read_csv(
    "experiments/datasets/runtime_benchmarks.csv"
)

plt.figure(figsize=(12,6))

plt.bar(
    df["benchmark"],
    df["bytes_op"]
)

plt.xticks(rotation=45)

plt.ylabel(
    "Bytes/op"
)

plt.title(
    "BranchKV Memory Cost Per Operation"
)

plt.tight_layout()

plt.savefig(
    "figures/runtime_memory.png",
    dpi=300
)

plt.savefig(
    "figures/runtime_memory.pdf"
)

print("saved")
