import pandas as pd
import matplotlib.pyplot as plt

df = pd.read_csv(
    "experiments/datasets/runtime_benchmarks.csv"
)

df["ops_sec"] = 1e9 / df["ns_op"]

df = df.sort_values(
    "ops_sec",
    ascending=False
)

plt.figure(figsize=(12,6))

plt.bar(
    df["benchmark"],
    df["ops_sec"]
)

plt.xticks(rotation=45)

plt.ylabel(
    "Operations / second"
)

plt.title(
    "BranchKV Estimated Throughput"
)

plt.tight_layout()

plt.savefig(
    "figures/runtime_throughput.png",
    dpi=300
)

plt.savefig(
    "figures/runtime_throughput.pdf"
)

print("saved")
