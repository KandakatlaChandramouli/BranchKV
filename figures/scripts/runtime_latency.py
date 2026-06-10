import pandas as pd
import matplotlib.pyplot as plt

df = pd.read_csv(
    "experiments/datasets/runtime_benchmarks.csv"
)

df = df.sort_values(
    "ns_op"
)

plt.figure(figsize=(12,6))

plt.barh(
    df["benchmark"],
    df["ns_op"]
)

plt.xlabel(
    "Latency (ns/op)"
)

plt.ylabel(
    "Operation"
)

plt.title(
    "BranchKV Runtime Benchmark Latency"
)

plt.tight_layout()

plt.savefig(
    "figures/runtime_latency.png",
    dpi=300
)

plt.savefig(
    "figures/runtime_latency.pdf"
)

print("saved")
