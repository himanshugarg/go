package popcount

func PopCount(x uint64) byte {
  var count byte = 0
  for i := 1; i < 64; i++ {
    count = count + byte(x&1);
    x = x >> 1
  }

  return count
}
