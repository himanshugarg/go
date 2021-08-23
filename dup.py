import sys
from collections import Counter

counts = Counter()
for name in sys.argv[1:]:
  with open(name) as f:
    counts.update([line.strip() for line in f])

for k, v in counts.items():
  if v > 1:
    print(f"{v}\t{k}")
