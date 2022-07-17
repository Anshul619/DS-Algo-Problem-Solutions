package leetCodeProblems.PriorityQueue;

/**
 * LeetCode - https://leetcode.com/problems/stock-price-fluctuation/
 *
 * TimeComplexity - O(1)
 * SpaceComplexity - O(n)
 */

import java.util.Comparator;
import java.util.HashMap;
import java.util.PriorityQueue;

public class StockPrice2034 {

    static class Stock {
        int timestamp;
        int price;

        Stock(int timestamp, int price) {
            this.timestamp = timestamp;
            this.price = price;
        }
    }

    static class PriceAscendingComparator implements Comparator<Stock> {

        public int compare(Stock stock1, Stock stock2) {
            return stock1.price - stock2.price;
        }

    }

    static class PriceDescendingComparator implements Comparator<Stock> {

        public int compare(Stock stock1, Stock stock2) {
            return stock2.price - stock1.price;
        }

    }

    int latestTimeStamp;
    HashMap<Integer, Integer> timeStampMap;

    PriorityQueue<Stock> maxQueue;
    PriorityQueue<Stock> minQueue;

    public StockPrice2034() {

        latestTimeStamp = 0;
        timeStampMap = new HashMap<>();
        maxQueue = new PriorityQueue<>(new PriceDescendingComparator());
        minQueue = new PriorityQueue<>(new PriceAscendingComparator());
    }

    public void update(int timestamp, int price) {

        if (latestTimeStamp < timestamp) {
            latestTimeStamp = timestamp;
        }

        timeStampMap.put(timestamp, price);

        Stock stockObj = new Stock(timestamp, price);
        maxQueue.add(stockObj);
        minQueue.add(stockObj);
    }

    public int current() {
        return timeStampMap.get(latestTimeStamp);
    }

    public int maximum() {

        Stock max = maxQueue.peek();

        while (timeStampMap.get(max.timestamp) != max.price) {
            maxQueue.remove();
            max = maxQueue.peek();
        }

        return max.price;
    }

    public int minimum() {
        Stock min = minQueue.peek();

        while (timeStampMap.get(min.timestamp) != min.price) {
            minQueue.remove();
            min = minQueue.peek();
        }

        return min.price;
    }
}