/*
Buy and sell stocks 3
 */

function maxProfit(prices: number[]): number {
  if(prices.length === 0) return 0;
  let firstBuy = -prices[0];
  let secondBuy = -prices[0];
  let firstSell = 0;
  let secondSell = 0;

  for(let i = 0; i < prices.length; i++) {
    firstBuy = Math.max(firstBuy, -prices[i]);
    firstSell = Math.max(firstSell, firstBuy+prices[i]);
    secondBuy = Math.max(secondBuy, firstSell-prices[i]);
    secondSell = Math.max(secondSell, secondBuy+prices[i]);
  }

  return secondSell;
}

console.log(maxProfit([3,3,5,0,0,3,1,4]));