package CodingInterviewQuestions;

/**
 * Asked in Amex-Online-Test, 29thJune2022
 *
 */
import java.util.*;

class TransactionsAmount {

    static class CardPaymentCountHelper {
        int count;
        int amount;

        CardPaymentCountHelper(int count, int amount) {
            this.count = count;
            this.amount = amount;
        }
    }

    private String getMonth(String date) {

        String[] splitArray = date.split("-");
        return splitArray[1];
    }

    public int solution(int[] A, String[] D) {
        // write your code in Java SE 8

        int amount = 0;
        int freeCardPaymentMonthsCount = 0;

        HashMap<String,CardPaymentCountHelper> monthMap = new HashMap<>();

        for(int i=0; i < A.length; i++) {
            amount += A[i];

            String month = getMonth(D[i]);

            if (A[i] < 0) {

                //System.out.println(month);
                if (monthMap.containsKey(month)) {

                    CardPaymentCountHelper countObj = monthMap.get(month);
                    countObj.count++;
                    countObj.amount += -A[i];
                    monthMap.put(month, countObj);
                }
                else {
                    CardPaymentCountHelper countObj = new CardPaymentCountHelper(1, -A[i]);
                    monthMap.put(month, countObj);
                }
            }

        }

        //System.out.println(monthMap);
        for(CardPaymentCountHelper value: monthMap.values()) {

            //System.out.println("Amount ->" + value.amount);
            if (value.count >= 3 && value.amount >=100) {
                freeCardPaymentMonthsCount++;
            }

        }

        int cardPaymentFees = 5*(12-freeCardPaymentMonthsCount);
        //System.out.println(cardPaymentFees);

        return amount-cardPaymentFees;
    }
}


