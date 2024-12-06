namespace DefaultNamespace;

public partial class Solution
{
  public bool IsPalindrome(int x)
  {
    var numString = x.ToString();
    var i = 0;
    var j = numString.Length - 1;

    while (i < j)
    {
      if (numString[i] != numString[j])
      {
        return false;
      }

      i++;
      j--;
    }
    
    return true;
  }
}