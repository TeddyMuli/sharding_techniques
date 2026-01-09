#!/usr/bin/env python3
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

sns.set_theme(style="whitegrid")
plt.rcParams["figure.figsize"] = (12, 6)

def plot_movement():
    df = pd.read_csv("movement.csv")
    
    plt.figure()
    ax = sns.barplot(x="Algorithm", y="PercentMoved", data=df, palette="viridis")
    
    plt.title("Data Movement During Resharding (Lower is Better)", fontsize=16)
    plt.ylabel("% of Keys Moved", fontsize=12)
    plt.xlabel("")
    plt.xticks(rotation=45)
    
    for i in ax.containers:
        ax.bar_label(i, fmt='%.1f%%')
        
    plt.tight_layout()
    plt.savefig("benchmark_movement.png")
    print("Saved benchmark_movement.png")

def plot_skew():
    df = pd.read_csv("distribution.csv")
        
    plt.figure()
    sns.boxplot(x="Algorithm", y="KeyCount", data=df, palette="coolwarm")
    
    plt.title("Load Balance / Hotspot Analysis (Tighter is Better)", fontsize=16)
    plt.ylabel("Keys assigned to Shard", fontsize=12)
    plt.xlabel("")
    plt.xticks(rotation=45)
    
    plt.tight_layout()
    plt.savefig("benchmark_skew.png")
    print("Saved benchmark_skew.png")

def plot_hotspot_heatmap():
    df = pd.read_csv("distribution.csv")

    pivot_df = df.pivot(index="Algorithm", columns="ShardID", values="KeyCount")

    pivot_df = pivot_df.fillna(0).astype(int)

    plt.figure(figsize=(14, 8))
    
    sns.heatmap(pivot_df, annot=True, fmt="d", cmap="YlOrRd", cbar_kws={'label': 'Number of Keys'})

    plt.title("Shard Load Heatmap: Identifying Hotspots", fontsize=16)
    plt.xlabel("Shard ID")
    plt.ylabel("Algorithm")
    
    plt.tight_layout()
    plt.savefig("benchmark_heatmap.png")
    print("Saved benchmark_heatmap.png", flush=True)

if __name__ == "__main__":
    plot_movement()
    plot_skew()
    plot_hotspot_heatmap()
