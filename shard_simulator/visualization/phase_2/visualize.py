#!/usr/bin/env python3
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

def plot_latency_comparison():
    df = pd.read_csv("latency_results.csv")
    
    df_melted = df.melt(id_vars="Algorithm", 
                        value_vars=["P50_ms", "P99_ms"], 
                        var_name="Percentile", 
                        value_name="Latency_ms")

    plt.figure(figsize=(12, 7))
    sns.barplot(data=df_melted, x="Algorithm", y="Latency_ms", hue="Percentile", palette="muted")
    
    plt.title("Sharding Performance: P50 vs P99 Tail Latency", fontsize=16)
    plt.ylabel("Latency (milliseconds)", fontsize=12)
    plt.xticks(rotation=45)
    plt.grid(axis='y', linestyle='--', alpha=0.7)
    
    plt.tight_layout()
    plt.savefig("latency_comparison.png")
    print("Saved latency_comparison.png", flush=True)

if __name__ == "__main__":
    plot_latency_comparison()
