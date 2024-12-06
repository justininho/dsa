/*
Buy and sell stocks 2
 */
function maxProfit(prices: number[]): number {
  let profit = 0;
  let buy = prices[0];
  let sell = buy;

  for (let i = 0; i < prices.length; i++) {
    if (prices[i] < buy) {
      buy = prices[i];
      sell = prices[i];
    } else if (prices[i] > prices[i + 1] || i+1 === prices.length) {
      sell = prices[i]
      profit += sell - buy;
      buy = sell;
    }
  }
  return profit;
}