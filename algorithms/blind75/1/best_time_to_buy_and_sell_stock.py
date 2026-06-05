class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        

        max_profit = 0
        mi = float('+inf')
        for i in range(len(prices)):
            if mi > prices[i]:
                mi = prices[i]
            elif prices[i] - mi > max_profit:
                max_profit = prices[i] - mi

        
        return max_profit
