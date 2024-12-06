namespace DefaultNamespace;

public partial class Solution {
  public int FindShortestSubArray(int[] nums)
  {
    var dict = new Dictionary<int, int>();
    
    foreach (var num in nums)
    {
      dict[num]++;
    }
    
    var degree = dict.Keys.Max();

    
  }
}