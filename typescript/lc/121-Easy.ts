/*
  Buy and sell stock 1
 */
function maxProfit(prices: number[]): number {
  let profit = 0;
  let buy = prices[0];
  let sell = buy;

  for (let i = 1; i < prices.length - 1; i++) {
    if (prices[i] < buy) {
      buy = prices[i]
      sell = buy
    } else if (prices[i] > sell) {
      sell = prices[i]
    }
    if (sell - buy > profit) {
      profit = sell - buy;
    }
  }
  return profit;
}