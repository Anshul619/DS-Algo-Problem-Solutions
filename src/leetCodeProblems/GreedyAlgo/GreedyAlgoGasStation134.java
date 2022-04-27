package leetCodeProblems.GreedyAlgo;

public class GreedyAlgoGasStation134 {
	
	public int canCompleteCircuit(int[] gas, int[] cost) {
        
        int total_gases = Arrays.stream(gas).sum();
        int total_cost = Arrays.stream(cost).sum();
        
        if (total_gases - total_cost < 0) {
            return -1;
        }
        
        // Difference of the TOTAL available gas & TOTAL cost, is greater than 0. 
        // Hence there must exists a complete circuit. 
        // We just need to find the starting index ("ans") for this.
        
        //System.out.println("Array Sum - "+sum);
        
        int sum = 0;
        int start_point = 0;
        
        for (int i=0; i < gas.length;i++) {
            
        	sum += gas[i] - cost[i];
            
            if (sum < 0) {
                start_point = i + 1;
                sum = 0;
            }
            
        }
        
        return start_point;
        
    }

}
