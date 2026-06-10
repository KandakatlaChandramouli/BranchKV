import pandas as pd
import matplotlib.pyplot as plt

df = pd.read_csv(
    "experiments/results/branch_scaling.csv"
)

plt.figure(figsize=(10,6))

plt.plot(
    df["branches"],
    df["latency"],
    marker="o"
)

plt.xscale("log", base=2)

plt.xlabel("Number of Branches")
plt.ylabel("Latency (ns)")
plt.title("BranchKV Branch Scaling")

plt.grid(True)

plt.tight_layout()

plt.savefig(
    "figures/branch_scaling.png",
    dpi=300
)

print("saved")
